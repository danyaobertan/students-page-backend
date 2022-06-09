package models

import (
	"context"
	"time"
)

type SubjectProfessor struct {
	Subject   Subjects   `json:"subject"`
	Professor Professors `json:"professor"`
}

func (m *DBModel) GetOneSubject(id int) (*SubjectProfessor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	query := `select s.subject_id,
       COALESCE(s.subject_name, '')        as subject_name,
       COALESCE(s.subject_credits, 0)      as subject_credits,
       COALESCE(s.exam_type, '')           as exam_type,
       COALESCE(s.exam_date, '0001-01-01') as exam_date,
       COALESCE(p.professor_id, 0)         as professor_id,
       COALESCE(p.name, '')                as professor_name,
       COALESCE(p.surname, '')             as professor_surname,
       COALESCE(p.patronymic, '')          as professor_patronymic,
       COALESCE(p.degree, '')              as professor_degree
from subjects s
         left join professors p on p.professor_id = s.professor_id
where s.subject_id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var group SubjectProfessor

	err := row.Scan(
		&group.Subject.SubjectId,
		&group.Subject.SubjectName,
		&group.Subject.SubjectCredits,
		&group.Subject.ExamType,
		&group.Subject.ExamDate,
		&group.Professor.ProfessorId,
		&group.Professor.Name,
		&group.Professor.Surname,
		&group.Professor.Patronymic,
		&group.Professor.Degree,
	)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (m *DBModel) GetAllSubjects() ([]*GroupStudentProfessor, error) {
	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//defer cancel()
	//
	//query := `select professor_id, department_id, name, surname, patronymic, degree, id_code, birth_date, gender,
	//                    phone_number, email, residence_postal_code, residence_address from professors order by professor_id`
	//
	//rows, err := m.DB.QueryContext(ctx, query)
	//if err != nil {
	//	return nil, err
	//}
	//defer rows.Close()
	//
	//var professors []*Professors
	//
	//for rows.Next() {
	//	var professor Professors
	//	err := rows.Scan(
	//		&professor.ProfessorId,
	//		&professor.DepartmentId,
	//		&professor.Name,
	//		&professor.Surname,
	//		&professor.Patronymic,
	//		&professor.Degree,
	//		&professor.IdCode,
	//		&professor.BirthDate,
	//		&professor.Gender,
	//		&professor.PhoneNumber,
	//		&professor.Email,
	//		&professor.ResidencePostalCode,
	//		&professor.ResidenceAddress,
	//	)
	//	if err != nil {
	//		return nil, err
	//	}
	//	// get groups, if any
	//	groupsQuery := `select group_id,group_name,professor_id from groups Where professor_id = $1 order by group_id`
	//	groupsRows, err := m.DB.QueryContext(ctx, groupsQuery, professor.ProfessorId)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	var groups []*Groups
	//
	//	for groupsRows.Next() {
	//		var group Groups
	//		err := groupsRows.Scan(
	//			&group.GroupId,
	//			&group.GroupName,
	//			&group.ProfessorId,
	//		)
	//		if err != nil {
	//			return nil, err
	//		}
	//		professor.Groups = append(groups, &group)
	//
	//	}
	//	groupsRows.Close()
	//
	//	professors = append(professors, &professor)
	//}
	//return professors, err
	return nil, nil
}
