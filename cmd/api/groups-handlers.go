package main

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app *application) getOneGroup(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("id is", id)

	group, err := app.models.DB.GetGroup(id)

	err = app.writeJSON(w, http.StatusOK, group, "group_professor_students")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getAllGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := app.models.DB.AllGroups()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, groups, "groups")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) deleteGroups(w http.ResponseWriter, r *http.Request) {

}

func (app *application) insertGroups(w http.ResponseWriter, r *http.Request) {

}

func (app *application) updateGroups(w http.ResponseWriter, r *http.Request) {

}

func (app *application) searchGroups(w http.ResponseWriter, r *http.Request) {

}
