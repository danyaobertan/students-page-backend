package models

import (
	"context"
	"time"
)

// GetProfessor returns one professor and error, if any
func (m *DBModel) GetProfessor(id int) (*Professors, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select professor_id, department_id, name, surname, patronymic, degree, id_code, birth_date, gender,
                        phone_number, email, residence_postal_code, residence_address  from professors where professor_id = $1 `

	row := m.DB.QueryRowContext(ctx, query, id)

	var professor Professors

	err := row.Scan(
		&professor.ProfessorId,
		&professor.DepartmentId,
		&professor.Name,
		&professor.Surname,
		&professor.Patronymic,
		&professor.Degree,
		&professor.IdCode,
		&professor.BirthDate,
		&professor.Gender,
		&professor.PhoneNumber,
		&professor.Email,
		&professor.ResidencePostalCode,
		&professor.ResidenceAddress,
	)
	if err != nil {
		return nil, err
	}

	// get groups, if any
	groupsQuery := `select group_id,group_name,professor_id from groups Where professor_id = $1 order by group_id`
	groupsRows, err := m.DB.QueryContext(ctx, groupsQuery, professor.ProfessorId)
	if err != nil {
		return nil, err
	}

	var groups []*Groups

	for groupsRows.Next() {
		var group Groups
		err := groupsRows.Scan(
			&group.GroupId,
			&group.GroupName,
			&group.ProfessorId,
		)
		if err != nil {
			return nil, err
		}
		professor.Groups = append(groups, &group)

	}
	groupsRows.Close()

	return &professor, nil
}

// All returns all departments and error, if any
func (m *DBModel) AllProfessors() ([]*Professors, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select professor_id, department_id, name, surname, patronymic, degree, id_code, birth_date, gender,
                        phone_number, email, residence_postal_code, residence_address from professors order by professor_id`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var professors []*Professors

	for rows.Next() {
		var professor Professors
		err := rows.Scan(
			&professor.ProfessorId,
			&professor.DepartmentId,
			&professor.Name,
			&professor.Surname,
			&professor.Patronymic,
			&professor.Degree,
			&professor.IdCode,
			&professor.BirthDate,
			&professor.Gender,
			&professor.PhoneNumber,
			&professor.Email,
			&professor.ResidencePostalCode,
			&professor.ResidenceAddress,
		)
		if err != nil {
			return nil, err
		}
		// get groups, if any
		groupsQuery := `select group_id,group_name,professor_id from groups Where professor_id = $1 order by group_id`
		groupsRows, err := m.DB.QueryContext(ctx, groupsQuery, professor.ProfessorId)
		if err != nil {
			return nil, err
		}

		var groups []*Groups

		for groupsRows.Next() {
			var group Groups
			err := groupsRows.Scan(
				&group.GroupId,
				&group.GroupName,
				&group.ProfessorId,
			)
			if err != nil {
				return nil, err
			}
			professor.Groups = append(groups, &group)

		}
		groupsRows.Close()

		professors = append(professors, &professor)
	}
	return professors, err
}
