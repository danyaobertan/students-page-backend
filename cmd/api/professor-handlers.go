package main

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app *application) getOneProfessor(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("id is", id)

	professor, err := app.models.DB.GetProfessor(id)

	err = app.writeJSON(w, http.StatusOK, professor, "professor")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllProfessors(w http.ResponseWriter, r *http.Request) {
	professors, err := app.models.DB.AllProfessors()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, professors, "professors")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) deleteProfessor(w http.ResponseWriter, r *http.Request) {

}

func (app *application) insertProfessor(w http.ResponseWriter, r *http.Request) {

}

func (app *application) updateProfessor(w http.ResponseWriter, r *http.Request) {

}

func (app *application) searchProfessor(w http.ResponseWriter, r *http.Request) {

}
