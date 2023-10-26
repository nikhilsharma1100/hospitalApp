package doctor

type GetDoctorByIdRequest struct {
	ID string `json:"id" uri:"id"`
}

type CreateDoctorRequest struct {
	Name      string `json:"name"`
	ContactNo string `json:"contact_no"`
}

type UpdateDoctorRequestBody struct {
	ContactNo string `json:"contact_no" binding:"required"`
}

type UpdateDoctorRequest struct {
	ContactNo string `json:"contact_no" binding:"required"`
	ID        string `json:"id"`
}

type UpdateDoctorRequestUri struct {
	ID string `json:"id" uri:"id"`
}

type GetPatientByDoctorIdRequest struct {
	ID string `json:"id" uri:"id"`
}
