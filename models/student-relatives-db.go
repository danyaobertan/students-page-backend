package models

import (
	"context"
	"time"
)

func (m *DBModel) GetStudentRelative(id int) (*StudentsRelatives, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	query := `select relative_id, student_id, name, surname, patronymic, relation_to_student, birthdate, phone_number, email 
			from studentsrelatives where relative_id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var sr StudentsRelatives

	err := row.Scan(
		&sr.RelativeId,
		&sr.StudentId,
		&sr.Name,
		&sr.Surname,
		&sr.Patronymic,
		&sr.RelationToStudent,
		&sr.BirthDate,
		&sr.PhoneNumber,
		&sr.Email,
	)
	if err != nil {
		return nil, err
	}

	return &sr, nil

}

func (m *DBModel) AllStudentRelatives(id int) (*StudentProfessorGroup, error) {
	return nil, nil
}
