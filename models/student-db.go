package models

import (
	"context"
	"math"
	"net/url"
	"strconv"
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

func (m *DBModel) InsertStudent(student Students) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `insert into students (name, surname, patronymic, bachelors_enrollment_date, gender,group_id, tuition, phone_number, email) 
				values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := m.DB.ExecContext(ctx, stmt,
		student.Name,
		student.Surname,
		student.Patronymic,
		student.BachelorsEnrollmentDate,
		student.Gender,
		student.GroupId,
		student.Tuition,
		student.PhoneNumber,
		student.Email,
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) UpdateStudent(student Students) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `update students set name = $1, surname=$2, patronymic=$3, bachelors_enrollment_date=$4, gender=$5,
                    group_id=$6, tuition=$7, phone_number=$8, email=$9 where student_id = $10`
	//fmt.Println(student)
	//fmt.Println(`/n`)
	//fmt.Println("/n")
	//fmt.Println(stmt)
	_, err := m.DB.ExecContext(ctx, stmt,
		student.Name,
		student.Surname,
		student.Patronymic,
		student.BachelorsEnrollmentDate,
		student.Gender,
		student.GroupId,
		student.Tuition,
		student.PhoneNumber,
		student.Email,
		student.StudentId,
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) DeleteStudent(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := "delete from students where student_id = $1"

	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}

// All returns all students and error, if any
func (m *DBModel) SearchStudents(f url.Values) ([]*Students, error, int, int, float64) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	whereClause := ``
	orderClause := ``
	paginationClause := ``
	page := 1
	perPage := 100
	//moreThanOneWhereCase := false
	for key, value := range f {
		if key == "name" && len(whereClause) == 0 {
			if len(value) == 1 {
				whereClause += ` where name like '%` + value[0] + `%'`
				//moreThanOneWhereCase = true
			} else {
				whereClause += ` and surname like '%` + value[0] + `%'`
			}
		}
		if key == "surname" {
			if len(value) == 1 && len(whereClause) == 0 {
				whereClause += ` where surname like '%` + value[0] + `%'`
				//moreThanOneWhereCase = true
			} else {
				whereClause += ` and surname like '%` + value[0] + `%'`
			}
		}
		if key == "patronymic" {
			if len(value) == 1 && len(whereClause) == 0 {
				whereClause += ` where patronymic like '%` + value[0] + `%'`
				//moreThanOneWhereCase = true
			} else {
				whereClause += ` and patronymic like '%` + value[0] + `%'`
			}
		}

		if key == `sb` && value[0] == `id` {
			orderClause += ` order by student_id`
		} else if key == "sb" && value[0] == `name` {
			orderClause += ` order by name`
		} else if key == "sb" && value[0] == `surname` {
			orderClause += ` order by surname`
		} else if key == "sb" && value[0] == `patronymic` {
			orderClause += ` order by patronymic`
		}

		if len(orderClause) != 0 && key == `so` && value[0] == `desc` {
			orderClause += ` desc`
		}

		if key == `page` {
			page, _ = strconv.Atoi(value[0])
		}
	}
	query := `select student_id,student_group_monitor, name, patronymic, surname,bachelors_enrollment_date, gender,group_id,tuition,id_code,phone_number,email from students` + whereClause + orderClause

	var total int
	err := m.DB.QueryRow(`select count(*) from students` + whereClause).Scan(&total)
	if err != nil {
		return nil, err, 0, 0, 0
	}
	paginationClause = ` LIMIT ` + strconv.Itoa(perPage) + `OFFSET ` + strconv.Itoa((page-1)*perPage)
	query += paginationClause

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err, 0, 0, 0
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
			return nil, err, 0, 0, 0
		}

		students = append(students, &student)
	}
	lastPage := math.Ceil(float64(total / perPage))
	return students, err, total, page, lastPage
}
