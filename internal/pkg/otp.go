package pkg

import (
	"crypto/rand"
	"errors"
	"fmt"
	"net/smtp"

	"github.com/rs/zerolog/log"
)

func GenCaptchaCode() (string, error) {
	codes := make([]byte, 6)
	if _, err := rand.Read(codes); err != nil {
		return "", err
	}

	for i := 0; i < 6; i++ {
		codes[i] = uint8(48 + (codes[i] % 10))
	}

	return string(codes), nil
}

func GenerateOneTimePassword(gmail string) (string, error) {
	// Sender's email address and password
	from := "jeevangb3@gmail.com"
	password := "flwo idka rcwf mcvf"

	// Recipient's email address
	to := "jeevanindian01@gmail.com"

	// SMTP server details
	smtpServer := "smtp.gmail.com"
	smtpPort := 587

	otp, err := GenCaptchaCode()
	if err != nil {
		log.Error().Err(err).Msg("failed to create random number")
		return "", errors.New("could not generate otp")
	}

	// Message content
	message := []byte("your one time password " + otp)

	// Authentication information
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// SMTP connection
	smtpAddr := fmt.Sprintf("%s:%d", smtpServer, smtpPort)
	err = smtp.SendMail(smtpAddr, auth, from, []string{to}, message)
	if err != nil {
		log.Error().Err(err).Msg("failure in sending otp to email")
		return "", errors.New("failed to send otp")
	}

	fmt.Println("Email sent successfully!")
	return otp, nil
}
