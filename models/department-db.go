package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

// GetDepartment returns one department and error, if any
func (m *DBModel) GetDepartment(id int) (*Departments, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select department_id, department_name from departments where department_id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var department Departments

	err := row.Scan(
		&department.DepartmentId,
		&department.DepartmentName,
	)
	if err != nil {
		return nil, err
	}

	// get professors, if any
	query = `select professor_id,name,surname,patronymic from professors Where department_id = $1`

	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	professors := make(map[int]string)
	for rows.Next() {
		var p Professors
		err := rows.Scan(
			&p.ProfessorId,
			&p.Name,
			&p.Surname,
			&p.Patronymic,
		)
		if err != nil {
			return nil, err
		}
		professors[p.ProfessorId] = p.Name + " " + p.Surname + " " + p.Patronymic
	}
	department.Professors = professors

	return &department, nil
}

// All returns all departments and error, if any
func (m *DBModel) AllDepartments() ([]*Departments, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select department_id, department_name from departments order by department_id`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []*Departments

	for rows.Next() {
		var department Departments
		err := rows.Scan(
			&department.DepartmentId,
			&department.DepartmentName,
		)
		if err != nil {
			return nil, err
		}
		// get professors, if any
		professorsQuery := `select professor_id,name,surname,patronymic from professors Where department_id = $1 order by professor_id`

		professorsRows, _ := m.DB.QueryContext(ctx, professorsQuery, department.DepartmentId)

		professors := make(map[int]string)
		for professorsRows.Next() {
			var p Professors
			err := professorsRows.Scan(
				&p.ProfessorId,
				&p.Name,
				&p.Surname,
				&p.Patronymic,
			)
			if err != nil {
				return nil, err
			}
			professors[p.ProfessorId] = p.Name + " " + p.Surname + " " + p.Patronymic

		}
		professorsRows.Close()

		department.Professors = professors
		departments = append(departments, &department)
	}
	return departments, err
}
