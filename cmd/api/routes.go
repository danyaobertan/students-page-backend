package main

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (app *application) routes() http.Handler {
	router := httprouter.New()
	secure := alice.New(app.checkToken)

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodPost, "/v1/graphql", app.studentsGraphQL)

	router.HandlerFunc(http.MethodPost, "/v1/signin", app.Signin)

	router.HandlerFunc(http.MethodGet, "/v1/department/:id", app.getOneDepartment)
	router.HandlerFunc(http.MethodGet, "/v1/departments", app.getAllDepartments)

	router.HandlerFunc(http.MethodGet, "/v1/professor/:id", app.getOneProfessor)
	router.HandlerFunc(http.MethodGet, "/v1/professors", app.getAllProfessors)

	router.HandlerFunc(http.MethodGet, "/v1/students-points", app.getAllStudentsPoints)
	router.HandlerFunc(http.MethodGet, "/v1/student-points/:id", app.getOneStudentPoints)

	router.HandlerFunc(http.MethodGet, "/v1/student-relatives/:id", app.getOneStudentRelative)

	router.HandlerFunc(http.MethodGet, "/v1/student/:id", app.getOneStudent)
	router.HandlerFunc(http.MethodGet, "/v1/students", app.getAllStudents)

	router.HandlerFunc(http.MethodGet, "/v1/group/:id", app.getOneGroup)
	router.HandlerFunc(http.MethodGet, "/v1/groups", app.getAllGroups)

	router.HandlerFunc(http.MethodGet, "/v1/subject/:id", app.getOneSubject)
	router.HandlerFunc(http.MethodGet, "/v1/subject", app.getAllSubjects)

	router.HandlerFunc(http.MethodGet, "/v1/send-mail-group-subject-points", app.sendMailGroupSubjectPoints)

	router.POST("/v1/admin/editstudent", app.wrap(secure.ThenFunc(app.editStudent)))
	router.POST("/v1/admin/deletestudent/:id", app.wrap(secure.ThenFunc(app.deleteStudent)))

	return app.enableCORS(router)
}
