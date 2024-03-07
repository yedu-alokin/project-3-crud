package validations

type CreateStudentValidator struct {
	Name   string `json:"name"`
	Email  string `json:"email"  binding:"required"`
	Class  string `json:"class"`
	School string `json:"school"`
}
