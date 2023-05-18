package service

import (
	"log"

	"github.com/ropel12/email/entities"
	"github.com/ropel12/email/helper"
)

func AddQuiz(data entities.ReqDataQuiz, auth string) {
	for _, val := range data.Data {
		helper.AddQuiz(val, auth, data.Prevlink, data.PubLink, data.ResultLink)
	}
	helper.SaveQuestion(auth)
	helper.SetPublish(auth, data.PubLink)
	helper.SetRequiredEmail(auth, data.ResultLink)
	helper.Publish(auth, data.PubLink)
	log.Println("Success To Create A New QUIZ")
}
