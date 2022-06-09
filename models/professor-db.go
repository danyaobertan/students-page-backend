package models

import (
	"context"
	"time"
)

// GetProfessor returns one professor and error, if any
func (m *DBModel) GetProfessor(id int) (*Professors, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	query := `select p.professor_id, p.department_id,
       COALESCE(p.name,''),
       COALESCE(p.surname,''),
       COALESCE(p.patronymic,''),
       COALESCE(p.degree,''),
       COALESCE(p.id_code,''),
       COALESCE(p.birth_date,'0001-01-01'),
       COALESCE(p.gender,''),
       COALESCE(p.phone_number,''),
       COALESCE(p.email,''),
       COALESCE(p.residence_postal_code,''),
       COALESCE(p.residence_address,''),
       COALESCE(d.department_name,'')

		from professors p
Left join departments d on d.department_id = p.department_id
where p.professor_id = $1 `

	//--        ,
	//--        COALESCE(s.subject_id,1),
	//--        COALESCE(s.name,'')
	//-- Left join subjects s on s.professor_id = p.professor_id

	row := m.DB.QueryRowContext(ctx, query, id)

	var professor Professors

	err := row.Scan(
		//COALESCE(ss.name, '')                 as student_group_monitor_name,

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
		&professor.Departments.DepartmentName,
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
		groups = append(groups, &group)

	}
	professor.Groups = groups

	groupsRows.Close()

	subjectsQuery := `select subject_id,subject_name from subjects Where professor_id = $1 order by subject_id`
	subjectRows, err := m.DB.QueryContext(ctx, subjectsQuery, professor.ProfessorId)
	if err != nil {
		return nil, err
	}

	var subjects []*Subjects

	for subjectRows.Next() {
		var subject Subjects
		err := subjectRows.Scan(
			&subject.SubjectId,
			&subject.SubjectName,
		)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, &subject)
	}
	professor.Subjects = subjects

	subjectRows.Close()

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
