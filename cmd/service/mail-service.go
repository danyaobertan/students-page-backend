package fitmailer

import (
	"backend/models"
	"fmt"
	"log"
	"time"
)

func SendMail(gsp []*models.GroupSubjectPoints) {

	//var mailsArr []string
	//for  key,value := range gsp {
	//
	//}
	config := MailerConfig{
		Host:     "smtp.mailtrap.io",
		Port:     587,
		Username: "a6aedfe3fc5537",
		Password: "144c4ba5079f8c",
		Timeout:  5 * time.Second,
		Sender:   "FITAutoMail@mail.com",
	}

	sender := New(config)

	err := sender.Send("jondoe@gmail.com", "letter.html", nil)
	if err != nil {
		log.Fatal(err)
	}

	for key, value := range gsp {
		err := sender.Send(value.StudentMail, "letter.html", value)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("mail #", key)
	}

	fmt.Println("Email sent!")
}
