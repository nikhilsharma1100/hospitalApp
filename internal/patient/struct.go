package patient

type CreatePatientRequest struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ContactNo string `json:"contact_no"`
	Address   string `json:"address"`
	DoctorID  string `json:"doctor_id"`
}

type UpdatePatientRequestUri struct {
	ID string `json:"id" uri:"id"`
}

type UpdatePatientRequest struct {
	DoctorID  string `json:"doctor_id"`
	ContactNo string `json:"contact_no"`
	Address   string `json:"address"`
}
