package dto

type DTORegisterAuth struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type DTOLoginAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
