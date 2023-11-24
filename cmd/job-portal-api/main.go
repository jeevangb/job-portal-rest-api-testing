package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"jeevan/jobportal/config"
	"jeevan/jobportal/internal/auth"
	"jeevan/jobportal/internal/cache"
	"jeevan/jobportal/internal/database"
	"jeevan/jobportal/internal/handler"
	"jeevan/jobportal/internal/repository"
	"jeevan/jobportal/internal/service"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func main() {
	err := StartApp()
	if err != nil {
		log.Panic().Err(err).Send()
	}
	log.Info().Msg("Hello this is our app")

}

func StartApp() error {
	cfg := config.GetConfig()

	fmt.Println(cfg.DatabaseConfig)

	log.Info().Interface("cfg", cfg).Msg("config")
	// initializing the authentication support
	log.Info().Msg("main started : initializing the authentication support")

	//reading the private key file
	fmt.Println()

	fmt.Println("PKKKKKKKKKKK===,", cfg.PrivatePublicKey)
	decodedPVKBytes, err := base64.StdEncoding.DecodeString(cfg.PrivatePublicKey.PrivateKey)
	if err != nil {
		fmt.Println("Error decoding base64:", err)
		return err
	}

	decodedPKBytes, err := base64.StdEncoding.DecodeString(cfg.PrivatePublicKey.PublicKey)
	if err != nil {
		fmt.Println("Error decoding base64:", err)
		return err
	}

	// Converting the decoded bytes to a string
	// decodedPKString := string(decodedPVKBytes)

	privatePEM := []byte(decodedPVKBytes)
	publicPEM := []byte(decodedPKBytes)

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	fmt.Println("[[[[[[[[[[[[[[]]]]]]]]]]]]]]", privatePEM)
	if err != nil {
		return fmt.Errorf("error in parsing auth private key : %w", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPEM)
	if err != nil {
		return fmt.Errorf("error in parsing auth public key : %w", err)
	}
	a, err := auth.NewAuth(privateKey, publicKey)
	if err != nil {
		return fmt.Errorf("error in constructing auth %w", err)
	}

	// ***********************************************************************************
	// start the database

	log.Info().Msg("main started : initializing the data")

	db, err := database.Connection(cfg.DatabaseConfig)
	if err != nil {
		return fmt.Errorf("error in opening the database connection : %w", err)
	}

	pg, err := db.DB()
	if err != nil {
		return fmt.Errorf("error in getting the database instance")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("database is not connected: %w", err)
	}
	////////////////////////////////////////////
	///redis connection
	rdb := database.ResdisConnection(cfg.RedisConfig)

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("redis is not connected: %w", err)
	}
	// service.NewService(rdb)
	rdbinter := cache.NewRdbLayer(rdb)

	//****************************************************************************
	// initialize the repository layer
	repo, err := repository.NewRepository(db)
	if err != nil {
		return err
	}

	svc, err := service.NewService(repo, a, rdbinter)
	if err != nil {
		return err
	}
	//*******************************************************************************
	// initializing the http server
	api := http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.AppConfig.AppHost, cfg.AppConfig.Port),
		ReadTimeout:  time.Duration(cfg.AppConfig.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.AppConfig.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.AppConfig.IdleTimeout) * time.Second,
		Handler:      handler.SetApi(a, svc),
	}

	// channel to store any errors while setting up the service
	serverErrors := make(chan error, 1)

	go func() {
		log.Info().Str("Port", api.Addr).Msg("main started : api is listening")
		serverErrors <- api.ListenAndServe()
	}()

	//shutdown channel intercepts ctrl+c signals
	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt)

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error : %w", err)

	case sig := <-shutdown:
		log.Info().Msgf("main: Start shutdown %s", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := api.Shutdown(ctx)
		if err != nil {
			err := api.Close()
			return fmt.Errorf("could not stop server gracefully : %w", err)
		}
	}
	return nil

}
