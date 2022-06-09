package main

import (
	"backend/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

type jsonResp struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (app *application) getOneStudent(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("id is", id)

	student, err := app.models.DB.GetStudent(id)

	err = app.writeJSON(w, http.StatusOK, student, "student")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllStudents(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println(query)
	fmt.Printf("%T\n", query)
	if len(query) == 0 {
		students, err := app.models.DB.AllStudents()
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		err = app.writeJSON(w, http.StatusOK, students, "students")
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	} else {
		students, err, total, page, lastPage := app.models.DB.SearchStudents(query)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		fmt.Println(total, page, lastPage)
		err = app.writeJSON(w, http.StatusOK, students, "students")
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	}

}

func (app *application) deleteStudent(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.models.DB.DeleteStudent(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	ok := jsonResp{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

type StudentPayload struct {
	StudentId                     string `json:"id"`
	Surname                       string `json:"surname"`
	Name                          string `json:"name"`
	Patronymic                    string `json:"patronymic"`
	Gender                        string `json:"gender"`
	GroupId                       string `json:"group_id"`
	Tuition                       string `json:"tuition"`
	PhoneNumber                   string `json:"phone_number"`
	Email                         string `json:"email"`
	IdCode                        string `json:"id_code"`
	ResidencePostalCode           string `json:"residence_postal_code"`
	ResidenceAddress              string `json:"residence_address"`
	CampusPostalCode              string `json:"campus_postal_code"`
	CampusAddress                 string `json:"campus_address"`
	BachelorsEnrollmentDocumentId string `json:"bachelors_enrollment_document_id"`
	BachelorsEnrollmentDate       string `json:"bachelors_enrollment_date"`
	MastersEnrollmentDocumentId   string `json:"masters_enrollment_document_id"`
	MastersEnrollmentDate         string `json:"masters_enrollment_date"`
	BachelorsExpulsionDocumentId  string `json:"bachelors_expulsion_document_id"`
	BachelorsExpulsionDate        string `json:"bachelors_expulsion_date"`
	MastersExpulsionDocumentId    string `json:"masters_expulsion_document_id"`
	MastersExpulsionDate          string `json:"masters_expulsion_date"`
}

func (app *application) editStudent(w http.ResponseWriter, r *http.Request) {

	//var payload StudentPayload
	var payload StudentPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var student models.StudentProfessorGroup

	if payload.StudentId != "0" {
		id, _ := strconv.Atoi(payload.StudentId)
		s, _ := app.models.DB.GetStudent(id)
		student = *s
	}

	student.StudentId, _ = strconv.Atoi(payload.StudentId)
	student.Surname = payload.Surname
	student.Name = payload.Name
	student.Patronymic = payload.Patronymic
	student.Gender = payload.Gender
	student.GroupId, _ = strconv.Atoi(payload.GroupId)
	student.Tuition = payload.Tuition
	student.PhoneNumber = payload.PhoneNumber
	student.Email = payload.Email
	student.IdCode = payload.IdCode
	student.ResidencePostalCode = payload.ResidencePostalCode
	student.ResidenceAddress = payload.ResidenceAddress
	student.CampusPostalCode = payload.CampusPostalCode
	student.CampusAddress = payload.CampusAddress
	student.BachelorsEnrollmentDocumentId = payload.BachelorsEnrollmentDocumentId
	student.MastersEnrollmentDocumentId = payload.MastersEnrollmentDocumentId
	student.BachelorsExpulsionDocumentId = payload.BachelorsExpulsionDocumentId
	student.MastersExpulsionDocumentId = payload.MastersExpulsionDocumentId
	student.BachelorsEnrollmentDate, _ = time.Parse("2006-01-02", payload.BachelorsEnrollmentDate)
	student.MastersEnrollmentDate, _ = time.Parse("2006-01-02", payload.MastersEnrollmentDate)
	student.BachelorsExpulsionDate, _ = time.Parse("2006-01-02", payload.BachelorsExpulsionDate)
	student.MastersExpulsionDate, _ = time.Parse("2006-01-02", payload.MastersExpulsionDate)

	if student.StudentId == 0 {
		err = app.models.DB.InsertStudent(student)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	} else {
		err = app.models.DB.UpdateStudent(student)
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	}

	ok := jsonResp{OK: true}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) searchStudent(w http.ResponseWriter, r *http.Request) {

}
