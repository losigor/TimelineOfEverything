package handlers

import (
	"net/http"
	"regexp"

	. "user-service/storage"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Status  string
	Message string
}

func RegisterUser(c echo.Context) error {

	var newUser User

	if err := c.Bind(&newUser); err != nil {
		return BadRequest(c, "Could not bind user")
	}

	// Validation of username
	if !validateUsername(newUser.Username) {
		return BadRequest(c, "Bad username")
	}

	// Validation of email
	if !validateEmail(newUser.Email) {
		return BadRequest(c, "Bad Email")
	}

	// Validation of password
	if !validatePassword(newUser.Password) {
		return BadRequest(c, "Password must contains numbers, lowercase and upercase letters")
	}

	// Check if username is unique
	var count int64
	DB.Model(&User{}).Where("username = ?", newUser.Username).Count(&count)
	if count > 0 {
		return BadRequest(c, "User with this username is alredy exists")
	}

	// Check if email is unique
	DB.Model(&User{}).Where("email = ?", newUser.Email).Count(&count)
	if count > 0 {
		return BadRequest(c, "User with this Email is already exists")
	}

	// Hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return BadRequest(c, "Could not hash password")
	}
	newUser.Password = string(hashedPassword)

	// Creating new user in the DB table
	if err := DB.Create(&newUser).Error; err != nil {
		return BadRequest(c, "Could not create new account(DB error)")
	}

	return Success(c, "Account was successfully created")
}

func LoginUser(c echo.Context) error {

	return nil
}

func DeleteUser(c echo.Context) error {

	return nil
}

func ChangeStatus(c echo.Context) error {

	return nil
}

func ChangeName(c echo.Context) error {

	return nil
}

func ChangePassword(c echo.Context) error {

	return nil
}

func Success(c echo.Context, s string) error {
	return c.JSON(http.StatusOK, Response{
		Status:  "OK",
		Message: s,
	})
}

func BadRequest(c echo.Context, s string) error {
	return c.JSON(http.StatusBadRequest, Response{
		Status:  "Error",
		Message: s,
	})
}

func validatePassword(password string) bool {

	// Check if password is long enough
	if len(password) < 8 {
		return false
	}

	// Regular expression to check has password numbers, lowercase and uppercase letters
	hasLower := regexp.MustCompile(`[a-z]`) // Lowercase
	hasUpper := regexp.MustCompile(`[A-Z]`) // Uppercase
	hasDigit := regexp.MustCompile(`[0-9]`) // Numbers

	// hasSpecial := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`)
	// Add if you want password to has special symbols

	return hasLower.MatchString(password) &&
		hasUpper.MatchString(password) &&
		hasDigit.MatchString(password)
	// && hasSpecial.MatchString(password)
}

func validateUsername(username string) bool {

	// Check username's lenght
	if len(username) < 3 {
		return false
	}

	// Regular expression to check if username contains only acceptable symbols
	re := regexp.MustCompile(`^[a-zA-Z0-9_]*$`)
	return re.MatchString(username)
}

func validateEmail(email string) bool {

	// Regular expression to check Email form
	re := regexp.MustCompile(`([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9_-]+)`)
	return re.MatchString(email)
}
