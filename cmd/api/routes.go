package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	//router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.getOneMovie)
	//router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMovies)

	router.HandlerFunc(http.MethodGet, "/v1/department/:id", app.getOneDepartment)
	router.HandlerFunc(http.MethodGet, "/v1/departments", app.getAllDepartments)

	router.HandlerFunc(http.MethodGet, "/v1/professor/:id", app.getOneProfessor)
	router.HandlerFunc(http.MethodGet, "/v1/professors", app.getAllProfessors)

	router.HandlerFunc(http.MethodGet, "/v1/student/:id", app.getOneStudent)
	router.HandlerFunc(http.MethodGet, "/v1/students", app.getAllStudents)

	return app.enableCORS(router)
}
