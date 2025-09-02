package utils

import (
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()

	// Register custom validators
	_ = Validate.RegisterValidation("future_date", ValidateFutureDate)
	_ = Validate.RegisterValidation("past_date", validatePastDate)
	_ = Validate.RegisterValidation("before_today", validateBeforeToday)
	_ = Validate.RegisterValidation("after_today", validateAfterToday)
	_ = Validate.RegisterValidation("min_age", validateMinAge)
	_ = Validate.RegisterValidation("not_weekend", validateNotWeekend)
	_ = Validate.RegisterValidation("time_format", ValidateTimeFormat)
	_ = Validate.RegisterValidation("password_strength", ValidatePasswordStrength)
	_ = Validate.RegisterValidation("valid_mobile_number", ValidMobileNumber)
}
