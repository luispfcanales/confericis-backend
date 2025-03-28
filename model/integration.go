package model

type ReniecPerson struct {
	DNI              string `json:"DNI"`
	LastNamePaternal string `json:"AP_PAT"`
	LastNameMaternal string `json:"AP_MAT"`
	Names            string `json:"NOMBRES"`
	BirthDate        string `json:"FECHA_NAC"`
	Address          string `json:"DIRECCION"`
	Gender           string `json:"SEXO"`
	MaritalStatus    string `json:"EST_CIVIL"`
}

type UnamadStudent struct {
	UserName        string  `json:"userName"`
	DNI             string  `json:"dni"`
	Name            string  `json:"name"`
	PaternalSurname string  `json:"paternalSurname"`
	MaternalSurname string  `json:"maternalSurname"`
	Email           string  `json:"email"`
	PersonalEmail   *string `json:"personalEmail"`
	CareerName      string  `json:"carrerName"`
	FacultyName     string  `json:"facultyName"`
}

type UnamadTeacher struct {
	UserName            string  `json:"userName"`
	DNI                 string  `json:"dni"`
	Name                string  `json:"name"`
	PaternalSurname     string  `json:"paternalSurname"`
	MaternalSurname     string  `json:"maternalSurname"`
	Email               string  `json:"email"`
	PersonalEmail       *string `json:"personalEmail"`
	AcademicDepartament string  `json:"academicDepartament"`
	FacultyName         string  `json:"facultyName"`
}
