package doctor

import "hospitalApp/internal/patient"

type GetDoctorByNameRequest struct {
	Name string `json:"name" uri:"name"`
}

type GetDoctorByIdRequest struct {
	ID string `json:"id" uri:"id"`
}

type CreateDoctorRequest struct {
	DoctorId  string `json:"doctor_id"`
	Name      string `json:"name"`
	ContactNo string `json:"contact_no"`
}

type UpdateDoctorRequest struct {
	DoctorId  string `json:"doctor_id"`
	Name      string `json:"name"`
	ContactNo string `json:"contact_no"`
}

type UpdatePatientRequest struct {
	DoctorId string          `json:"doctor_id"`
	Patient  patient.Patient `json:"patient"`
}
