package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneDepartment(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("id is", id)

	department, err := app.models.DB.GetDepartment(id)

	err = app.writeJSON(w, http.StatusOK, department, "department")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllDepartments(w http.ResponseWriter, r *http.Request) {
	departments, err := app.models.DB.AllDepartments()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, departments, "departments")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) deleteDepartment(w http.ResponseWriter, r *http.Request) {

}

func (app *application) insertDepartment(w http.ResponseWriter, r *http.Request) {

}

func (app *application) updateDepartment(w http.ResponseWriter, r *http.Request) {

}

func (app *application) searchDepartment(w http.ResponseWriter, r *http.Request) {

}
