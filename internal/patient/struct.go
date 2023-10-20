package patient

type CreatePatientRequest struct {
	PatientId string `json:"patient_id"`
	Name      string `json:"name"`
	ContactNo string `json:"contact_no"`
	Address   string `json:"address"`
}

type UpdatePatientRequestUri struct {
	Id string `json:"id" uri:"id"`
}

type UpdatePatientRequestUriName struct {
	Name string `json:"name" uri:"name"`
}

type UpdatePatientRequest struct {
	DoctorId  string `json:"doctor_id"`
	ContactNo string `json:"contact_no"`
	Address   string `json:"address"`
}
