package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

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
}

func (app *application) deleteStudent(w http.ResponseWriter, r *http.Request) {

}

//insert and update are combined in edit
//func (app *application) insertStudent(w http.ResponseWriter, r *http.Request) {
//
//}
//
//func (app *application) updateStudent(w http.ResponseWriter, r *http.Request) {
//
//}

func (app *application) editStudent(w http.ResponseWriter, r *http.Request) {
	type jsonResp struct {
		OK bool `json:"ok"`
	}

	ok := jsonResp{OK: true}

	err := app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) searchStudent(w http.ResponseWriter, r *http.Request) {

}
