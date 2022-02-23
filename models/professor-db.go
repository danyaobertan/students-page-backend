package models

import (
	"context"
	"time"
)

// GetDepartment returns one department and error, if any
func (m *DBModel) GetProfessor(id int) (*Departments, error) {
	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//defer cancel()
	//
	//query := `select department_id, department_name from departments where department_id = $1`
	//
	//row := m.DB.QueryRowContext(ctx, query, id)
	//
	//var department Departments
	//
	//err := row.Scan(
	//	&department.DepartmentId,
	//	&department.DepartmentName,
	//)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// get professors, if any
	//query = `select professor_id,name,surname,patronymic from professors Where department_id = $1`
	//
	//rows, _ := m.DB.QueryContext(ctx, query, id)
	//defer rows.Close()
	//
	//professors := make(map[int]string)
	//for rows.Next() {
	//	var p Professors
	//	err := rows.Scan(
	//		&p.ProfessorId,
	//		&p.Name,
	//		&p.Surname,
	//		&p.Patronymic,
	//	)
	//	if err != nil {
	//		return nil, err
	//	}
	//	professors[p.ProfessorId] = p.Name + " " + p.Surname + " " + p.Patronymic
	//}
	//department.Professors = professors
	//
	//return &department, nil
	return nil, nil
}

// All returns all departments and error, if any
func (m *DBModel) AllProfessors() ([]*Professors, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select professor_id, name, patronymic, surname from professors order by professor_id`

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
			&professor.Name,
			&professor.Surname,
			&professor.Patronymic,
		)
		if err != nil {
			return nil, err
		}
		// get groups, if any
		groupsQuery := `select group_id,group_name from groups Where professor_id = $1 order by group_id`

		groupsRows, _ := m.DB.QueryContext(ctx, groupsQuery, professor.ProfessorId)

		groups := make(map[int]string)
		for groupsRows.Next() {
			var g Groups
			err := groupsRows.Scan(
				&g.GroupId,
				&g.GroupName,
			)
			if err != nil {
				return nil, err
			}
			groups[g.GroupId] = string(g.GroupId) + " " + g.GroupName

		}
		groupsRows.Close()

		professor.Groups = groups
		professors = append(professors, &professor)
	}
	return professors, err
}
