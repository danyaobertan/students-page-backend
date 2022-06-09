package models

import (
	"database/sql"
	"time"
)

// Models is the wrapper for database
type Models struct {
	DB DBModel
}

// NewModels returns models with db pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type Departments struct {
	DepartmentId   int            `json:"department_id"`
	DepartmentName string         `json:"department_name"`
	Professors     map[int]string `json:"professors"`
}

type Professors struct {
	ProfessorId         int         `json:"professor_id"`
	DepartmentId        int         `json:"department_id"`
	Name                string      `json:"name"`
	Surname             string      `json:"surname"`
	Patronymic          string      `json:"patronymic"`
	Degree              string      `json:"degree"`
	IdCode              string      `json:"id_code"`
	BirthDate           time.Time   `json:"birth_date"`
	Gender              string      `json:"gender"`
	PhoneNumber         string      `json:"phone_number"`
	Email               string      `json:"email"`
	ResidencePostalCode string      `json:"residence_postal_code"`
	ResidenceAddress    string      `json:"residence_address"`
	Groups              []*Groups   `json:"groups"`
	Subjects            []*Subjects `json:"subjects"`
	Departments         Departments `json:"departments"`
}

type Groups struct {
	GroupId     int    `json:"group_id"`
	ProfessorId int    `json:"professor_id"`
	GroupName   string `json:"group_name"`
}

type Students struct {
	StudentId                     int       `json:"student_id"`
	StudentGroupMonitor           int       `json:"student_group_monitor"`
	GroupId                       int       `json:"group_id"`
	Surname                       string    `json:"surname"`
	Name                          string    `json:"name"`
	Patronymic                    string    `json:"patronymic"`
	BachelorsEnrollmentDocumentId string    `json:"bachelors_enrollment_document_id"`
	BachelorsEnrollmentDate       time.Time `json:"bachelors_enrollment_date"`
	MastersEnrollmentDocumentId   string    `json:"masters_enrollment_document_id"`
	MastersEnrollmentDate         time.Time `json:"masters_enrollment_date"`
	BachelorsExpulsionDocumentId  string    `json:"bachelors_expulsion_document_id"`
	BachelorsExpulsionDate        time.Time `json:"bachelors_expulsion_date"`
	MastersExpulsionDocumentId    string    `json:"masters_expulsion_document_id"`
	MastersExpulsionDate          time.Time `json:"masters_expulsion_date"`
	Tuition                       string    `json:"tuition"`
	IdCode                        string    `json:"id_code"`
	BirthDate                     time.Time `json:"birth_date"`
	Gender                        string    `json:"gender,omitempty"`
	ResidencePostalCode           string    `json:"residence_postal_code"`
	ResidenceAddress              string    `json:"residence_address"`
	CampusPostalCode              string    `json:"campus_postal_code"`
	CampusAddress                 string    `json:"campus_address"`
	PhoneNumber                   string    `json:"phone_number"`
	Email                         string    `json:"email"`
}

type StudentsRelatives struct {
	RelativeId        int       `json:"relative_id"`
	StudentId         int       `json:"student_id"`
	Surname           string    `json:"surname"`
	Name              string    `json:"name"`
	Patronymic        string    `json:"patronymic"`
	RelationToStudent string    `json:"relation_to_student"`
	BirthDate         time.Time `json:"birth_date"`
	PhoneNumber       string    `json:"phone_number"`
	Email             string    `json:"email"`
}

type Subjects struct {
	SubjectId      int       `json:"subject_id"`
	ProfessorId    int       `json:"professor_id"`
	SubjectName    string    `json:"subject_name"`
	SubjectCredits float32   `json:"subject_credits"`
	ExamType       string    `json:"exam_type"`
	ExamDate       time.Time `json:"exam_date"`
}

type StudentsSubjects struct {
	StudentsSubjectsId int    `json:"students_subjects_id"`
	StudentId          int    `json:"student_id"`
	SubjectId          string `json:"subject_id"`
	ExamScore          int    `json:"exam_score"`
	TotalScore         string `json:"total_score"`
}

type User struct {
	ID       int
	Email    string
	Password string
}
