package models

import (
	"context"
	"time"
)

// GetProfessor returns one professor and error, if any
func (m *DBModel) GetStudent(id int) (*Students, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select student_id,student_group_monitor, name, patronymic, surname,bachelors_enrollment_date, gender,group_id,tuition,id_code,phone_number,email from students where student_id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var student Students

	err := row.Scan(
		&student.StudentId,
		&student.StudentGroupMonitor,
		&student.Name,
		&student.Patronymic,
		&student.Surname,
		&student.BachelorsEnrollmentDate,
		&student.Gender,
		&student.GroupId,
		&student.Tuition,
		&student.IdCode,
		&student.PhoneNumber,
		&student.Email,
	)
	if err != nil {
		return nil, err
	}

	// get students, if any
	//query = `select group_id ,group_name from groups Where professor_id = $1`
	//
	//rows, _ := m.DB.QueryContext(ctx, query, id)
	//defer rows.Close()
	//
	//groups := make(map[int]string)
	//for rows.Next() {
	//	var g Groups
	//	err := rows.Scan(
	//		&g.GroupId,
	//		&g.GroupName,
	//	)
	//	if err != nil {
	//		return nil, err
	//	}
	//	groups[g.GroupId] = g.GroupName + " " + string(g.GroupId)
	//}
	//professor.Groups = groups

	return &student, nil
	//return nil, nil

}

// All returns all students and error, if any
func (m *DBModel) AllStudents() ([]*Students, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select student_id,student_group_monitor, name, patronymic, surname,bachelors_enrollment_date, gender,group_id,tuition,id_code,phone_number,email from students order by student_id`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []*Students

	for rows.Next() {
		var student Students
		err := rows.Scan(
			&student.StudentId,
			&student.StudentGroupMonitor,
			&student.Name,
			&student.Patronymic,
			&student.Surname,
			&student.BachelorsEnrollmentDate,
			&student.Gender,
			&student.GroupId,
			&student.Tuition,
			&student.IdCode,
			&student.PhoneNumber,
			&student.Email,
		)
		if err != nil {
			return nil, err
		}

		//// get groups, if any
		//groupsQuery := `select group_id,group_name from groups Where professor_id = $1 order by group_id`
		//
		//groupsRows, _ := m.DB.QueryContext(ctx, groupsQuery, professor.ProfessorId)
		//
		//groups := make(map[int]string)
		//for groupsRows.Next() {
		//	var g Groups
		//	err := groupsRows.Scan(
		//		&g.GroupId,
		//		&g.GroupName,
		//	)
		//	if err != nil {
		//		return nil, err
		//	}
		//	groups[g.GroupId] = string(g.GroupId) + " " + g.GroupName
		//
		//}
		//groupsRows.Close()
		//
		//professor.Groups = groups
		//professors = append(professors, &professor)
		students = append(students, &student)
	}
	return students, err
}
