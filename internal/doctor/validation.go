package doctor

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
)

func (c *Core) ValidateCreateRequest(input CreateDoctorRequest) error {
	validationErr := validation.ValidateStruct(&input,
		validation.Field(&input.ContactNo, validation.Match(regexp.MustCompile("\\d{10}$")), validation.Length(10, 10)),
	)
	if validationErr != nil {
		return validationErr
	}

	return nil
}
