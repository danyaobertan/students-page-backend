package main

import (
	"backend/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/graphql-go/graphql"
)

var students []*models.Students

// graphql schema definition
var fields = graphql.Fields{
	"student": &graphql.Field{
		Type:        studentType,
		Description: "Get student by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id, ok := p.Args["id"].(int)
			if ok {
				for _, student := range students {
					if student.StudentId == id {
						return student, nil
					}
				}
			}
			return nil, nil
		},
	},
	"list": &graphql.Field{
		Type:        graphql.NewList(studentType),
		Description: "Get all students",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return students, nil
		},
	},
	"search": &graphql.Field{
		Type:        graphql.NewList(studentType),
		Description: "Search students by surname",
		Args: graphql.FieldConfigArgument{
			"surnameContains": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var theList []*models.Students
			search, ok := params.Args["surnameContains"].(string)
			if ok {
				for _, currentStudent := range students {
					if strings.Contains(currentStudent.Surname, search) {
						log.Println("Found one")
						theList = append(theList, currentStudent)
					}
				}
			}
			return theList, nil
		},
	},
}

var studentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Student",
		Fields: graphql.Fields{
			"student_id": &graphql.Field{
				Type: graphql.Int,
			},
			"student_group_monitor": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"surname": &graphql.Field{
				Type: graphql.String,
			},
			"patronymic": &graphql.Field{
				Type: graphql.String,
			},
			"bachelors_enrollment_date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"gender": &graphql.Field{
				Type: graphql.String,
			},
			"group_id": &graphql.Field{
				Type: graphql.Int,
			},
			"tuition": &graphql.Field{
				Type: graphql.String,
			},
			"id_code": &graphql.Field{
				Type: graphql.Int,
			},
			"phone_number": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func (app *application) studentsGraphQL(w http.ResponseWriter, r *http.Request) {
	students, _ = app.models.DB.AllStudents()

	q, _ := io.ReadAll(r.Body)
	query := string(q)

	log.Println(query)

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		app.errorJSON(w, errors.New("failed to create schema"))
		log.Println(err)
		return
	}

	params := graphql.Params{Schema: schema, RequestString: query}
	resp := graphql.Do(params)
	if len(resp.Errors) > 0 {
		app.errorJSON(w, errors.New(fmt.Sprintf("failed: %+v", resp.Errors)))
	}

	j, _ := json.MarshalIndent(resp, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
