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

type Movie struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Year        int          `json:"year"`
	ReleaseDate time.Time    `json:"release_date"`
	Runtime     int          `json:"runtime"`
	Rating      int          `json:"rating"`
	MPAARating  string       `json:"mpaa_rating"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	MovieGenre  []MovieGenre `json:"-"`
}

type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type MovieGenre struct {
	ID        int       `json:"id"`
	MovieID   int       `json:"movie_id"`
	GenreID   int       `json:"genre_id"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Departments struct {
	DepartmentId   int            `json:"department_id"`
	DepartmentName string         `json:"department_name"`
	Professors     map[int]string `json:"professors"`
}

type Professors struct {
	ProfessorId         int            `json:"professor_id"`
	DepartmentId        int            `json:"department_id"`
	Name                string         `json:"name"`
	Surname             string         `json:"surname"`
	Patronymic          string         `json:"patronymic"`
	Degree              string         `json:"degree"`
	IdCode              string         `json:"id_code"`
	BirthDate           time.Time      `json:"birth_date"`
	Gender              string         `json:"gender"`
	PhoneNumber         string         `json:"phone_number"`
	Email               string         `json:"email"`
	ResidencePostalCode string         `json:"residence_postal_code"`
	ResidenceAddress    string         `json:"residence_address"`
	Groups              map[int]string `json:"groups"`
}

type Groups struct {
	GroupId   int    `json:"group_id"`
	GroupName string `json:"group_name"`
}
