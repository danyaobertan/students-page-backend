package main

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app *application) getOneStudentRelative(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("id is", id)

	student, err := app.models.DB.GetStudentRelative(id)

	err = app.writeJSON(w, http.StatusOK, student, "student")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

//func (app *application) getAllStudentsRelatives(w http.ResponseWriter, r *http.Request) {
//	query := r.URL.Query()
//	fmt.Println(query)
//	fmt.Printf("%T\n", query)
//	if len(query) == 0 {
//		students, err := app.models.DB.AllStudentsRelatives()
//		if err != nil {
//			app.errorJSON(w, err)
//			return
//		}
//
//		err = app.writeJSON(w, http.StatusOK, students, "students")
//		if err != nil {
//			app.errorJSON(w, err)
//			return
//		}
//	} else {
//		students, err, total, page, lastPage := app.models.DB.SearchStudents(query)
//		if err != nil {
//			app.errorJSON(w, err)
//			return
//		}
//		fmt.Println(total, page, lastPage)
//		err = app.writeJSON(w, http.StatusOK, students, "students")
//		if err != nil {
//			app.errorJSON(w, err)
//			return
//		}
//	}
//
//
//}

//func (app *application) deleteStudent(w http.ResponseWriter, r *http.Request) {
//	params := httprouter.ParamsFromContext(r.Context())
//
//	id, err := strconv.Atoi(params.ByName("id"))
//	if err != nil {
//		app.errorJSON(w, err)
//		return
//	}
//
//	err = app.models.DB.DeleteStudent(id)
//	if err != nil {
//		app.errorJSON(w, err)
//		return
//	}
//
//	ok := jsonResp{
//		OK: true,
//	}
//
//	err = app.writeJSON(w, http.StatusOK, ok, "response")
//	if err != nil {
//		app.errorJSON(w, err)
//		return
//	}
//}
//
//type StudentPayload struct {
//	StudentId               string `json:"id"`
//	Surname                 string `json:"surname"`
//	Name                    string `json:"name"`
//	Patronymic              string `json:"patronymic"`
//	BachelorsEnrollmentDate string `json:"bachelors_enrollment_date"`
//	Gender                  string `json:"gender"`
//	GroupId                 string `json:"group_id"`
//	Tuition                 string `json:"tuition"`
//	PhoneNumber             string `json:"phone_number"`
//	Email                   string `json:"email"`
//}

//func (app *application) editStudent(w http.ResponseWriter, r *http.Request) {
//
//	//var payload StudentPayload
//	var payload StudentPayload
//
//	err := json.NewDecoder(r.Body).Decode(&payload)
//	if err != nil {
//		app.errorJSON(w, err)
//		return
//	}
//
//	var student models.Students
//
//	if payload.StudentId != "0" {
//		id, _ := strconv.Atoi(payload.StudentId)
//		s, _ := app.models.DB.GetStudent(id)
//		student = *s
//	}
//
//	student.StudentId, _ = strconv.Atoi(payload.StudentId)
//	student.Surname = payload.Surname
//	student.Name = payload.Name
//	student.Patronymic = payload.Patronymic
//	student.BachelorsEnrollmentDate, _ = time.Parse("2006-01-02", payload.BachelorsEnrollmentDate)
//	student.Gender = payload.Gender
//	student.GroupId, _ = strconv.Atoi(payload.GroupId)
//	student.Tuition = payload.Tuition
//	student.PhoneNumber = payload.PhoneNumber
//	student.Email = payload.Email
//
//	if student.StudentId == 0 {
//		err = app.models.DB.InsertStudent(student)
//		if err != nil {
//			app.errorJSON(w, err)
//			return
//		}
//	} else {
//		err = app.models.DB.UpdateStudent(student)
//		if err != nil {
//			app.errorJSON(w, err)
//			return
//		}
//	}
//
//	ok := jsonResp{OK: true}
//
//	err = app.writeJSON(w, http.StatusOK, ok, "response")
//	if err != nil {
//		app.errorJSON(w, err)
//		return
//	}
//}

//func (app *application) searchStudent(w http.ResponseWriter, r *http.Request) {
//
//}
