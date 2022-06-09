package models

import (
	"context"
	"time"
)

type GroupStudentProfessor struct {
	Group                         Groups     `json:"group"`
	Student                       Students   `json:"student"`
	Professor                     Professors `json:"professor"`
	StudentGroupMonitorName       string     `json:"student_group_monitor_name"`
	StudentGroupMonitorSurname    string     `json:"student_group_monitor_surname"`
	StudentGroupMonitorPatronymic string     `json:"student_group_monitor_patronymic"`
}

func (m *DBModel) GetGroup(id int) ([]*GroupStudentProfessor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	query := `Select g.group_id as group_id,
       COALESCE(g.group_name, '')            as group_name,
       COALESCE(s.student_id, 0)             as student_id,
       COALESCE(s.name, '')                  as student_name,
       COALESCE(s.surname, '')               as student_surname,
       COALESCE(s.patronymic, '')            as student_patronymic,
       COALESCE(s.phone_number, '')          as student_phone_number,
       COALESCE(s.email, '')                 as student_pemail,
       COALESCE(ss.student_group_monitor, 0) as student_group_monitor_id,
       COALESCE(ss.name, '')                 as student_group_monitor_name,
       COALESCE(ss.surname, '')              as student_group_monitor_surname,
       COALESCE(ss.patronymic, '')           as student_group_monitor_patronymic,
       COALESCE(p.professor_id, 0)           as professor_id,
       COALESCE(p.name, '')                  as professor_name,
       COALESCE(p.surname, '')               as professor_surname,
       COALESCE(p.patronymic, '')            as professor_patronymic,
       COALESCE(p.degree, '')                as professor_degree
		From groups g
         Left Outer JOIN students s on g.group_id = s.group_id
         LEFT Outer JOIN professors p on p.professor_id = g.professor_id
         LEFT JOIN students ss on s.student_group_monitor = ss.student_id
		where g.group_id = $1`

	row, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	var group []*GroupStudentProfessor

	for row.Next() {
		var groupRow GroupStudentProfessor
		err := row.Scan(
			&groupRow.Group.GroupId,
			&groupRow.Group.GroupName,
			&groupRow.Student.StudentId,
			&groupRow.Student.Name,
			&groupRow.Student.Surname,
			&groupRow.Student.Patronymic,
			&groupRow.Student.PhoneNumber,
			&groupRow.Student.Email,
			&groupRow.Student.StudentGroupMonitor,
			&groupRow.StudentGroupMonitorName,
			&groupRow.StudentGroupMonitorSurname,
			&groupRow.StudentGroupMonitorPatronymic,
			&groupRow.Professor.ProfessorId,
			&groupRow.Professor.Name,
			&groupRow.Professor.Surname,
			&groupRow.Professor.Patronymic,
			&groupRow.Professor.Degree,
		)
		if err != nil {
			return nil, err
		}
		group = append(group, &groupRow)

	}
	return group, nil
}

func (m *DBModel) AllGroups() ([]*Groups, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `Select g.group_id , g.group_name From groups g`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []*Groups

	for rows.Next() {
		var group Groups
		err := rows.Scan(
			&group.GroupId,
			&group.GroupName,
		)
		if err != nil {
			return nil, err

		}
		groups = append(groups, &group)

	}
	rows.Close()
	return groups, err
}
