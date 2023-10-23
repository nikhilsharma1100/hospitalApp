package doctor

import "hospitalApp/internal/patient"

type GetDoctorByNameRequest struct {
	Name string `json:"name" uri:"name"`
}

type GetDoctorByIdRequest struct {
	ID string `json:"id" uri:"id"`
}

type CreateDoctorRequest struct {
	Name      string `json:"name"`
	ContactNo string `json:"contact_no"`
}

type UpdateDoctorRequest struct {
	ContactNo string `json:"contact_no" binding:"required"`
}

type UpdateDoctorRequestUri struct {
	ID string `json:"id" uri:"id"`
}

type UpdatePatientRequest struct {
	ID      string          `json:"id"`
	Patient patient.Patient `json:"patient"`
}

// Add response structs
