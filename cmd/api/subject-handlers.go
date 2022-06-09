package main

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app *application) getOneSubject(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("id is", id)

	group, err := app.models.DB.GetOneSubject(id)

	err = app.writeJSON(w, http.StatusOK, group, "subject")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllSubjects(w http.ResponseWriter, r *http.Request) {
	groups, err := app.models.DB.GetAllSubjects()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, groups, "subjects")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) deleteSubjects(w http.ResponseWriter, r *http.Request) {

}

func (app *application) insertSubjects(w http.ResponseWriter, r *http.Request) {

}

func (app *application) updateSubjects(w http.ResponseWriter, r *http.Request) {

}

func (app *application) searchSubjects(w http.ResponseWriter, r *http.Request) {

}
