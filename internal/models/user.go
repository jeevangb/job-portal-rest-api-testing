package models

import "gorm.io/gorm"

type NewUser struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Dob      string `json:"dob" validate:"required"`
}

type User struct {
	gorm.Model
	Username     string `json:"username" gorm:"unique"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash string `json:"-"`
	Dob          string `json:"dob" validate:"required"`
}

type ForgotPasswod struct {
	Email string `json:"email"`
	Dob   string `json:"dob"`
}

type ResetPassword struct {
	Email           string `json:"email" validate:"required"`
	NewPassword     string `json:"newpassword" validate:"required"`
	ConfirmPassword string `json:"confirmpassword" validate:"required"`
	Otp             string `json:"otp" validate:"required"`
}
