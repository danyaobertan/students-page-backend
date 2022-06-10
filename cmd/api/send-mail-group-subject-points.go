package main

import (
	fitmailer "backend/cmd/service"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) sendMailGroupSubjectPoints(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Println(query)
	fmt.Printf("%T\n", query)
	sendMail := false
	if len(query) == 0 {
		errq := errors.New("no query params")
		app.errorJSON(w, errq)
		return
	} else {
		groupSubjectPoints, err := app.models.DB.GetMailGroupSubjectPoints(query)
		for key, value := range query {
			if key == "sendmail" {
				sendMail, _ = strconv.ParseBool(value[0])
			}
		}
		if sendMail {
			go fitmailer.SendMail(groupSubjectPoints)
		}
		if err != nil {
			app.errorJSON(w, err)
			return
		}
		err = app.writeJSON(w, http.StatusOK, groupSubjectPoints, "sendMailGroupSubjectPoints")
		if err != nil {
			app.errorJSON(w, err)
			return
		}
	}

}
