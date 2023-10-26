package patient

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
)

func (c *Core) ValidateCreateRequest(input CreatePatientRequest) error {
	validationErr := validation.ValidateStruct(&input,
		validation.Field(&input.ContactNo, validation.Match(regexp.MustCompile("\\d{10}$")), validation.Length(10, 10)),
		validation.Field(&input.DoctorID, validation.Required),
	)

	if validationErr != nil {
		return validationErr
	}

	return nil
}
