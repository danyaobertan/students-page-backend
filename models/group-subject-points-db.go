package models

import (
	"context"
	"net/url"
	"time"
)

type GroupSubjectPoints struct {
	StudentId           int    `json:"student_id"`
	StudentName         string `json:"student_name"`
	StudentSurname      string `json:"student_surname"`
	StudentPatronymic   string `json:"student_patronymic"`
	StudentMail         string `json:"student_mail"`
	SubjectId           string `json:"subject_id"`
	SubjectName         string `json:"subject_name"`
	ExamDate            string `json:"exam_date"`
	ExamType            string `json:"exam_type"`
	SubjectCredits      string `json:"subject_credits"`
	TotalScore          string `json:"total_score"`
	ProfessorId         string `json:"professor_id"`
	ProfessorName       string `json:"professor_name"`
	ProfessorSurname    string `json:"professor_surname"`
	ProfessorPatronymic string `json:"professor_patronymic"`
}

func (m *DBModel) GetMailGroupSubjectPoints(f url.Values) ([]*GroupSubjectPoints, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	groupId := ``
	subjectId := ``
	for key, value := range f {
		if key == "group" {
			groupId = value[0]
		}
		if key == "subject" {
			subjectId = value[0]
		}
	}
	query := `Select s.student_id,
       s.name,
       s.surname,
       s.patronymic,
       s.email,
       ss.subject_id,
       sb.subject_name,
       sb.exam_date,
       sb.exam_type,
       sb.subject_credits,
       ss.total_score,
       p.professor_id,
       p.name,
       p.surname,
       p.patronymic
FROM studentssubjects ss
         inner join students s on ss.student_id = s.student_id
         inner join subjects sb on ss.subject_id = sb.subject_id
         inner join professors p on p.professor_id = sb.professor_id
where s.group_id = $1
  and ss.subject_id = $2`

	rows, err := m.DB.QueryContext(ctx, query, groupId, subjectId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var group_subject_points []*GroupSubjectPoints

	for rows.Next() {
		var group_subject_point GroupSubjectPoints
		err := rows.Scan(
			&group_subject_point.StudentId,
			&group_subject_point.StudentName,
			&group_subject_point.StudentSurname,
			&group_subject_point.StudentPatronymic,
			&group_subject_point.StudentMail,
			&group_subject_point.SubjectId,
			&group_subject_point.SubjectName,
			&group_subject_point.ExamDate,
			&group_subject_point.ExamType,
			&group_subject_point.SubjectCredits,
			&group_subject_point.TotalScore,
			&group_subject_point.ProfessorId,
			&group_subject_point.ProfessorName,
			&group_subject_point.ProfessorSurname,
			&group_subject_point.ProfessorPatronymic,
		)
		if err != nil {
			return nil, err
		}

		group_subject_points = append(group_subject_points, &group_subject_point)
	}
	return group_subject_points, err

}
