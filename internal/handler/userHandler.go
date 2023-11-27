package handler

import (
	"encoding/json"
	"net/http"

	"jeevan/jobportal/internal/middleware"
	"jeevan/jobportal/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// Reset the password it takes email,otp,new password and confirmation password
func (h *Handler) ResetPassword(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	var UserResetData models.ResetPassword
	err := json.NewDecoder(c.Request.Body).Decode(&UserResetData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid email,otp,new password and confirmation password",
		})
		return
	}
	err = h.Service.ResetPassword(ctx, UserResetData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "passwod reset successfully")
}

// forgot password function takes email and dob and send otp to their email to reset the password
func (h *Handler) ForgotPasswod(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	var userEmailAndPassword models.ForgotPasswod

	err := json.NewDecoder(c.Request.Body).Decode(&userEmailAndPassword)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid email and dob",
		})
		return
	}
	err = h.Service.CheckUserDataAndSendOtp(ctx, userEmailAndPassword)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "otp sent successfully")

}

// login function takes email and passoword to login
func (h *Handler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	var userData models.NewUser
	err := json.NewDecoder(c.Request.Body).Decode(&userData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid email and password",
		})
		return
	}
	token, err := h.Service.UserLogin(ctx, userData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// User Signup function
func (h *Handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIDKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	var userData models.NewUser
	err := json.NewDecoder(c.Request.Body).Decode(&userData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid username, email and password",
		})
		return
	}
	validate := validator.New()
	err = validate.Struct(userData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid username, email, password and dob",
		})
		return
	}
	userDetails, err := h.Service.UserSignup(ctx, userData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, userDetails)
}
