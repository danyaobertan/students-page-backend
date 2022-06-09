package models

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"strconv"
	"time"
)

type StudentProfessorGroup struct {
	StudentId                     int       `json:"student_id"`
	StudentGroupMonitor           int       `json:"student_group_monitor"`
	GroupId                       int       `json:"group_id"`
	Surname                       string    `json:"surname"`
	Name                          string    `json:"name"`
	Patronymic                    string    `json:"patronymic"`
	BachelorsEnrollmentDocumentId string    `json:"bachelors_enrollment_document_id"`
	BachelorsEnrollmentDate       time.Time `json:"bachelors_enrollment_date"`
	MastersEnrollmentDocumentId   string    `json:"masters_enrollment_document_id"`
	MastersEnrollmentDate         time.Time `json:"masters_enrollment_date"`
	BachelorsExpulsionDocumentId  string    `json:"bachelors_expulsion_document_id"`
	BachelorsExpulsionDate        time.Time `json:"bachelors_expulsion_date"`
	MastersExpulsionDocumentId    string    `json:"masters_expulsion_document_id"`
	MastersExpulsionDate          time.Time `json:"masters_expulsion_date"`
	Tuition                       string    `json:"tuition"`
	IdCode                        string    `json:"id_code"`
	BirthDate                     time.Time `json:"birth_date"`
	Gender                        string    `json:"gender,omitempty"`
	ResidencePostalCode           string    `json:"residence_postal_code"`
	ResidenceAddress              string    `json:"residence_address"`
	CampusPostalCode              string    `json:"campus_postal_code"`
	CampusAddress                 string    `json:"campus_address"`
	PhoneNumber                   string    `json:"phone_number"`
	Email                         string    `json:"email"`
	GroupName                     string    `json:"group_name"`
	ProfessorID                   string    `json:"professor_id"`
	ProfessorName                 string    `json:"professor_name"`
	ProfessorSurname              string    `json:"professor_surname"`
	ProfessorPatronymic           string    `json:"professor_patronymic"`
	StudentGroupMonitorName       string    `json:"student_group_monitor_name"`
	StudentGroupMonitorSurname    string    `json:"student_group_monitor_surname"`
	StudentGroupMonitorPatronymic string    `json:"student_group_monitor_patronymic"`
	StudentRelativeID             int       `json:"relative_id"`
	StudentRelativeName           string    `json:"student_relative_name"`
	StudentRelativeSurname        string    `json:"student_relative_surname"`
	StudentRelativePatronymic     string    `json:"student_relative_patronymic"`
}

// GetProfessor returns one professor and error, if any
func (m *DBModel) GetStudent(id int) (*StudentProfessorGroup, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	//query := `select student_id,student_group_monitor, name, patronymic, surname,COALESCE(bachelors_enrollment_date,'0001-01-01'), gender,group_id,tuition,id_code,phone_number,email from students where student_id = $1`
	query := `select s.student_id,
       s.student_group_monitor,
       COALESCE(s.group_id, 0)                             as group_id,
       COALESCE(s.name, '')                                as name,
       COALESCE(s.surname, '')                             as surname,
       COALESCE(s.patronymic, '')                          as patronymic,
       COALESCE(s.bachelors_enrollment_document_id, '')    as bachelors_enrollment_document_id,
       COALESCE(s.bachelors_enrollment_date, '0001-01-01') as bachelors_enrollment_date,
       COALESCE(s.masters_enrollment_document_id, '')      as masters_enrollment_document_id,
       COALESCE(s.masters_enrollment_date, '0001-01-01')   as masters_enrollment_date,
       COALESCE(s.bachelors_expulsion_document_id, '')     as bachelors_expulsion_document_id,
       COALESCE(s.bachelors_expulsion_date, '0001-01-01')  as bachelors_expulsion_date,
       COALESCE(s.masters_expulsion_document_id, '')       as masters_expulsion_document_id,
       COALESCE(s.masters_expulsion_date, '0001-01-01')    as masters_expulsion_date,
       COALESCE(s.id_code, '')                             as id_code,
       COALESCE(s.tuition, '')                             as tuition,
       COALESCE(s.birth_date, '0001-01-01')                as birth_date,
       COALESCE(s.gender, '')                              as gender,
       COALESCE(s.residence_address, '')                   as residence_address,
       COALESCE(s.residence_postal_code, '')               as residence_postal_code,
       COALESCE(s.campus_address, '')                      as campus_address,
       COALESCE(s.campus_postal_code, '')                  as campus_postal_code,
       COALESCE(s.phone_number, '')                        as phone_number,
       COALESCE(s.email, '')                               as email,
       COALESCE(g.group_name, '')                          as group_name,
       COALESCE(p.professor_id, 0) 						 as professor_id,
       COALESCE(p.name, '')                                as professor_name,
       COALESCE(p.surname, '')                             as professor_surname,
       COALESCE(p.patronymic, '')                          as professor_patronymic,
       COALESCE(ss.name, '')                 as student_group_monitor_name,
       COALESCE(ss.surname, '')              as student_group_monitor_surname,
       COALESCE(ss.patronymic, '')           as student_group_monitor_patronymic,
       COALESCE(sr.relative_id,0)           as  relative_id,
	   COALESCE(sr.name, '')                 as student_relative_name,
       COALESCE(sr.surname, '')              as student_relative_surname,
       COALESCE(sr.patronymic, '')           as student_relative_patronymic
		from students s 
         Left outer join groups g on g.group_id = s.group_id
         Left outer join professors p on g.professor_id = p.professor_id 
         Left outer join students ss on s.student_group_monitor = ss.student_id
		 Left outer join studentsrelatives sr on sr.student_id = s.student_id
		where s.student_id = $1`
	row := m.DB.QueryRowContext(ctx, query, id)

	var student StudentProfessorGroup

	err := row.Scan(
		&student.StudentId,
		&student.StudentGroupMonitor,
		&student.GroupId,
		&student.Name,
		&student.Surname,
		&student.Patronymic,
		&student.BachelorsEnrollmentDocumentId,
		&student.BachelorsEnrollmentDate,
		&student.MastersEnrollmentDocumentId,
		&student.MastersEnrollmentDate,
		&student.BachelorsExpulsionDocumentId,
		&student.BachelorsExpulsionDate,
		&student.MastersExpulsionDocumentId,
		&student.MastersExpulsionDate,
		&student.IdCode,
		&student.Tuition,
		&student.BirthDate,
		&student.Gender,
		&student.ResidenceAddress,
		&student.ResidencePostalCode,
		&student.CampusAddress,
		&student.CampusPostalCode,
		&student.PhoneNumber,
		&student.Email,
		&student.GroupName,
		&student.ProfessorID,
		&student.ProfessorName,
		&student.ProfessorSurname,
		&student.ProfessorPatronymic,
		&student.StudentGroupMonitorName,
		&student.StudentGroupMonitorSurname,
		&student.StudentGroupMonitorPatronymic,
		&student.StudentRelativeID,
		&student.StudentRelativeName,
		&student.StudentRelativeSurname,
		&student.StudentRelativePatronymic,
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

	//query := `select student_id,student_group_monitor, name, patronymic, surname,COALESCE(bachelors_enrollment_date,'0001-01-01'), gender,group_id,tuition,id_code,phone_number,email from students order by student_id`
	query := `select  student_id,student_group_monitor,   
      COALESCE(group_id,0) as group_id,
	  COALESCE(name,'') as name,
      COALESCE(surname,'') as surname,
      COALESCE(patronymic,'') as patronymic,
      COALESCE(bachelors_enrollment_document_id,'') as bachelors_enrollment_document_id,
      COALESCE(bachelors_enrollment_date,'0001-01-01') as bachelors_enrollment_date,
      COALESCE(masters_enrollment_document_id,'') as masters_enrollment_document_id,
      COALESCE(masters_enrollment_date,'0001-01-01') as masters_enrollment_date,
      COALESCE(bachelors_expulsion_document_id,'') as bachelors_expulsion_document_id,
      COALESCE(bachelors_expulsion_date,'0001-01-01') as bachelors_expulsion_date,
      COALESCE(masters_expulsion_document_id,'') as masters_expulsion_document_id,
      COALESCE(masters_expulsion_date,'0001-01-01') as masters_expulsion_date,
      COALESCE(id_code,'') as id_code,
      COALESCE(tuition,'') as tuition,
      COALESCE(birth_date,'0001-01-01') as birth_date,
      COALESCE(gender,'') as gender,
      COALESCE(residence_address,'') as residence_address,
      COALESCE(residence_postal_code,'') as residence_postal_code,
      COALESCE(campus_address,'') as campus_address,
      COALESCE(campus_postal_code,'') as campus_postal_code,
      COALESCE(phone_number,'') as phone_number,
      COALESCE(email,'') as email
      from students order by student_id`

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
			&student.GroupId,
			&student.Name,
			&student.Surname,
			&student.Patronymic,
			&student.BachelorsEnrollmentDocumentId,
			&student.BachelorsEnrollmentDate,
			&student.MastersEnrollmentDocumentId,
			&student.MastersEnrollmentDate,
			&student.BachelorsExpulsionDocumentId,
			&student.BachelorsExpulsionDate,
			&student.MastersExpulsionDocumentId,
			&student.MastersExpulsionDate,
			&student.IdCode,
			&student.Tuition,
			&student.BirthDate,
			&student.Gender,
			&student.ResidenceAddress,
			&student.ResidencePostalCode,
			&student.CampusAddress,
			&student.CampusPostalCode,
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

func (m *DBModel) InsertStudent(student StudentProfessorGroup) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if student.StudentPayload {

	}
	stmt := `insert into students (name, surname, patronymic, gender,group_id, tuition, phone_number, email,
					  COALESCE(id_code,''),
                      COALESCE(residence_postal_code,''),
                      COALESCE(residence_address,''),
                      COALESCE(campus_postal_code,''),
                      COALESCE(campus_address,''),
                      COALESCE(bachelors_enrollment_document_id,''),
                      COALESCE(masters_enrollment_document_id,''),
                      COALESCE(bachelors_expulsion_document_id,''),
                      COALESCE(masters_expulsion_document_id,''),
					  COALESCE(bachelors_enrollment_date,'0001-01-01'),
                      COALESCE(masters_enrollment_date,'0001-01-01'),
                      COALESCE(bachelors_enrollment_date,'0001-01-01'),
                      COALESCE(masters_expulsion_date,'0001-01-01')
) 
				values ($1, $2, $3, $4, $5, $6, $7, $8 ,$9 ,$10 ,$11, $12, $13, $14, $15, $16, $17, $18, $19 ,$20 ,$21)`

	_, err := m.DB.ExecContext(ctx, stmt,
		student.Name,
		student.Surname,
		student.Patronymic,
		student.Gender,
		student.GroupId,
		student.Tuition,
		student.PhoneNumber,
		student.Email,
		student.IdCode,
		student.ResidencePostalCode,
		student.ResidenceAddress,
		student.CampusPostalCode,
		student.CampusAddress,
		student.BachelorsEnrollmentDocumentId,
		student.MastersEnrollmentDocumentId,
		student.BachelorsExpulsionDocumentId,
		student.MastersExpulsionDocumentId,
		student.BachelorsEnrollmentDate,
		student.MastersEnrollmentDate,
		student.BachelorsExpulsionDate,
		student.MastersExpulsionDate,
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) UpdateStudent(student StudentProfessorGroup) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `update students set 
                    name = $1,
                    surname=$2, 
                    patronymic=$3, 
                    gender=$4,
                    group_id=$5,
                    tuition=$6,
                    phone_number=$7,
                    email=$8 ,
                    id_code=$9 ,
                    residence_postal_code=$10 ,
                    residence_address=$11 ,
                    campus_postal_code=$12 ,
                    campus_address=$13 ,
                    bachelors_enrollment_document_id=$14 ,
                    masters_enrollment_document_id=$15 ,
                    bachelors_expulsion_document_id=$16 ,
                    masters_expulsion_document_id=$17 ,
                    bachelors_enrollment_date=$18 ,
                    masters_enrollment_date=$19 ,
                    bachelors_expulsion_date=$20 ,
                    masters_expulsion_date=$21 
where student_id = $22`
	//fmt.Println(student)
	//fmt.Println(`/n`)
	//fmt.Println("/n")
	//fmt.Println(stmt)
	_, err := m.DB.ExecContext(ctx, stmt,
		student.Name,
		student.Surname,
		student.Patronymic,
		student.Gender,
		student.GroupId,
		student.Tuition,
		student.PhoneNumber,
		student.Email,
		student.IdCode,
		student.ResidencePostalCode,
		student.ResidenceAddress,
		student.CampusPostalCode,
		student.CampusAddress,
		student.BachelorsEnrollmentDocumentId,
		student.MastersEnrollmentDocumentId,
		student.BachelorsExpulsionDocumentId,
		student.MastersExpulsionDocumentId,
		student.BachelorsEnrollmentDate,
		student.MastersEnrollmentDate,
		student.BachelorsExpulsionDate,
		student.MastersExpulsionDate,
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
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
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
				whereClause += ` where patronymic like '(%` + value[0] + `%'`
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
	//query := `select student_id,student_group_monitor, name, patronymic, surname,COALESCE(bachelors_enrollment_date,'0001-01-01') as bachelors_enrollment_date ,COALESCE(masters_enrollment_date,'0001-01-01')as masters_enrollment_date, gender,group_id,tuition,id_code,phone_number,email from students` + whereClause + orderClause
	query := `select  student_id,student_group_monitor,   
      COALESCE(group_id,0) as group_id,
	  COALESCE(name,'') as name,
      COALESCE(surname,'') as surname,
      COALESCE(patronymic,'') as patronymic,
      COALESCE(bachelors_enrollment_document_id,'') as bachelors_enrollment_document_id,
      COALESCE(bachelors_enrollment_date,'0001-01-01') as bachelors_enrollment_date,
      COALESCE(masters_enrollment_document_id,'') as masters_enrollment_document_id,
      COALESCE(masters_enrollment_date,'0001-01-01') as masters_enrollment_date,
      COALESCE(bachelors_expulsion_document_id,'') as bachelors_expulsion_document_id,
      COALESCE(bachelors_expulsion_date,'0001-01-01') as bachelors_expulsion_date,
      COALESCE(masters_expulsion_document_id,'') as masters_expulsion_document_id,
      COALESCE(masters_expulsion_date,'0001-01-01') as masters_expulsion_date,
      COALESCE(id_code,'') as id_code,
      COALESCE(tuition,'') as tuition,
      COALESCE(birth_date,'0001-01-01') as birth_date,
      COALESCE(gender,'') as gender,
      COALESCE(residence_address,'') as residence_address,
      COALESCE(residence_postal_code,'') as residence_postal_code,
      COALESCE(campus_address,'') as campus_address,
      COALESCE(campus_postal_code,'') as campus_postal_code,
      COALESCE(phone_number,'') as phone_number,
      COALESCE(email,'') as email
      from students` + whereClause + orderClause
	shitRow := `select count(*) from (` + query + `)src`
	fmt.Println(shitRow)
	strTotal := ``
	total := 0
	//var total int
	err := m.DB.QueryRow(`select count(*) from (` + query + `)src`).Scan(&strTotal)
	total, err = strconv.Atoi(strTotal)
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
			&student.GroupId,
			&student.Name,
			&student.Surname,
			&student.Patronymic,
			&student.BachelorsEnrollmentDocumentId,
			&student.BachelorsEnrollmentDate,
			&student.MastersEnrollmentDocumentId,
			&student.MastersEnrollmentDate,
			&student.BachelorsExpulsionDocumentId,
			&student.BachelorsExpulsionDate,
			&student.MastersExpulsionDocumentId,
			&student.MastersExpulsionDate,
			&student.IdCode,
			&student.Tuition,
			&student.BirthDate,
			&student.Gender,
			&student.ResidenceAddress,
			&student.ResidencePostalCode,
			&student.CampusAddress,
			&student.CampusPostalCode,
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
