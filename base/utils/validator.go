package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

// Validation errors
type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationError struct {
	Errors []FieldError `json:"errors"`
}

func (v *ValidationError) Error() string {
	var sb strings.Builder
	for _, err := range v.Errors {
		sb.WriteString(fmt.Sprintf("%s: %s\n", err.Field, err.Message))
	}
	return sb.String()
}

func NewValidationError(err error) *ValidationError {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]FieldError, len(ve))
		for i, fe := range ve {
			out[i] = FieldError{
				Field:   fe.Field(),
				Message: msgForTag(fe.Field(), fe.Tag(), fe.Param()),
			}
		}
		return &ValidationError{Errors: out}
	}
	return nil
}

func msgForTag(field string, tag string, param string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, param)
	case "max":
		return fmt.Sprintf("%s must be at less than %s characters", field, param)
	case "datetime":
		return fmt.Sprintf("%s Must be in %s format", field, param)
	case "future_date":
		return fmt.Sprintf("%s must be in the future", field)
	case "time_format":
		return fmt.Sprintf("%s (Time) must be in HH:MM format", field)
	case "regexp":
		return fmt.Sprintf("%s must match %s", field, param)
	case "oneof":
		return fmt.Sprintf("%s must be one of %s", field, param)
	case "numeric":
		return fmt.Sprintf("%s must be numeric", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "url":
		return fmt.Sprintf("%s must be a valid URL", field)
	case "uuid":
		return fmt.Sprintf("%s must be a valid UUID", field)
	case "image":
		return fmt.Sprintf("%s must be a valid image", field)
	case "file_size":
		return fmt.Sprintf("%s must be less than %s MB", field, param)
	case "file_type":
		return fmt.Sprintf("%s must be one of %s", field, param)
	case "file_extension":
		return fmt.Sprintf("%s must be one of %s", field, param)
	case "file_mime_type":
		return fmt.Sprintf("%s must be one of %s", field, param)
	case "file_max_size":
		return fmt.Sprintf("%s must be less than %s MB", field, param)
	case "file_max_width":
		return fmt.Sprintf("%s must be less than %s px", field, param)
	case "file_max_height":
		return fmt.Sprintf("%s must be less than %s px", field, param)
	case "file_min_width":
		return fmt.Sprintf("%s must be less than %s px", field, param)
	case "file_min_height":
		return fmt.Sprintf("%s must be less than %s px", field, param)
	case "password_strength":
		return "Password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character"
	case "before_today":
		return fmt.Sprintf("%s must be before today", field)
	case "after_today":
		return fmt.Sprintf("%s must be after today", field)
	case "not_weekend":
		return fmt.Sprintf("%s cannot be on a weekend", field)
	case "min_age":
		return fmt.Sprintf("%s Must be at least %s years old", field, param)
	default:
		return "Invalid value"
	}
}

// Custom validators (time_format)
func ValidateTimeFormat(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(`^(0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$`, fl.Field().String())
	return matched
}

// Custom validators (password_strength)
func ValidatePasswordStrength(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 {
		return false
	}
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[\W_]`).MatchString(password)
	return hasUppercase && hasLowercase && hasNumber && hasSpecial
}

func ValidMobileNumber(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(`^\+?[0-9]{10,15}$`, fl.Field().String())
	return matched
}

// Custom validators (future_date)
func ValidateFutureDate(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	return date.After(time.Now())
}

// validatePastDate checks if the date is in the past
func validatePastDate(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	return date.Before(time.Now())
}

// validateBeforeToday checks if the date is before today (excluding today)
func validateBeforeToday(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}

	// Get today's date at midnight
	today := time.Now().Truncate(24 * time.Hour)
	return date.Before(today)
}

// validateAfterToday checks if the date is after today (excluding today)
func validateAfterToday(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}

	// Get today's date at midnight
	today := time.Now().Truncate(24 * time.Hour)
	return date.After(today)
}

// validateNotWeekend checks if the date is not a weekend (Saturday or Sunday)
func validateNotWeekend(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}

	// Sunday = 0, Saturday = 6
	weekday := date.Weekday()
	return weekday != time.Sunday && weekday != time.Saturday
}

// validateMinAge checks if the person is at least X years old
func validateMinAge(fl validator.FieldLevel) bool {
	date, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}

	// Default minimum age is 18
	minAge := 18
	if param := fl.Param(); param != "" {
		if age, err := strconv.Atoi(param); err == nil {
			minAge = age
		}
	}

	// Calculate age
	now := time.Now()
	age := now.Year() - date.Year()

	// Adjust age if birthday hasn't occurred yet this year
	if now.YearDay() < date.YearDay() {
		age--
	}

	return age >= minAge
}
