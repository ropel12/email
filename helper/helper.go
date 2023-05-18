package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/ropel12/email/entities"
	"golang.org/x/net/http2"
)

type Option struct {
	OptionID        string `json:"optionId"`
	OptionText      string `json:"optionText"`
	OptionMatch     string `json:"optionMatch"`
	Points          int    `json:"points"`
	OptionPointsSet int    `json:"optionPointsSet"`
	OptionFileName  string `json:"optionFileName"`
	OptionImageID   string `json:"optionImageId"`
	OptionRow       string `json:"optionRow"`
	OptionChar      string `json:"optionChar"`
}

func AddQuiz(data entities.ReqAddQuiz, auth, prevlink, publink, reslink string) {
	point1 := 0
	point2 := 0
	point3 := 0
	point4 := 0
	switch data.Answer {
	case 1:
		point1 = 1
	case 2:
		point2 = 2
	case 3:
		point3 = 3
	case 4:
		point4 = 4
	}
	options := []Option{}
	for i := 0; i < 4; i++ {
		point := 0
		option := ""
		switch i {
		case 0:
			point = point1
			option = data.Option1
		case 1:
			point = point2
			option = data.Option2
		case 2:
			point = point3
			option = data.Option3
		case 3:
			point = point4
			option = data.Option4
		}
		optionn := Option{
			OptionID:        "00000000-0000-0000-0000-000000000000",
			OptionText:      option,
			OptionMatch:     "",
			Points:          point,
			OptionPointsSet: 0,
			OptionFileName:  "",
			OptionImageID:   "00000000-0000-0000-0000-000000000000",
			OptionRow:       "0",
			OptionChar:      "",
		}
		options = append(options, optionn)
	}
	optionsencoded, _ := json.Marshal(options)
	values := url.Values{}
	values.Set("surveyPageId", prevlink)
	values.Set("questionId", "-1")
	values.Set("questionText", "<p>"+data.Question+"</p>")
	values.Set("questionType", "radioButton")
	values.Set("options", string(optionsencoded))
	values.Set("questionRows", "1")
	values.Set("newQuestionPosition", "ABOVE")
	values.Set("availablePoints", "1")
	values.Set("questionDescription", "")
	values.Set("showQuestionDescription", "false")
	values.Set("questionRequired", "false")
	values.Set("randomizeOptions", "false")
	values.Set("questionBankCategories", "")
	values.Set("useOptionPointsSet", "false")
	values.Set("optionDisplaySize", "L")
	values.Set("codingLanguages", "")
	values.Set("questionIncorrectDescription", "")
	values.Set("showIncorrectQuestionDescription", "false")
	values.Set("showImmediateQuestionDescription", "false")
	values.Set("addAboveQuestionId", "")
	values.Set("questionTimeLimit", "")
	values.Set("isNewMedia", "false")
	values.Set("mediaSource", "")
	values.Set("mediaType", "")
	values.Set("mediaId", "")
	values.Set("mediaUrl", "")
	values.Set("mediaAttributionName", "")
	values.Set("mediaAttributionUrl", "")
	values.Set("mediaAltDescription", "")
	values.Set("mediaDownloadLocation", "")

	t := &http2.Transport{}
	client := http.Client{
		Transport: t,
	}
	req, err := http.NewRequest("POST", "https://www.flexiquiz.com/Create/SaveQuestion", strings.NewReader(values.Encode()))
	if err != nil {
		log.Printf("[ERROR]WHEN SETTING TO PUBLIC, Err: %v", err)
	}
	cookie := &http.Cookie{Name: ".ASPXAUTH", Value: auth}
	req.AddCookie(cookie)
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR]WHEN SETTING TO PUBLIC, Err: %v", err)
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	log.Println("[ERROR]ADD A NEW QUESTION", err)
}

func SetPublish(auth, publink string) {
	values := url.Values{}
	values.Set("publishedSurveyId", publink)
	values.Set("surveyType", "completeSurveyOnce")
	t := &http2.Transport{}
	client := http.Client{
		Transport: t,
	}
	req, err := http.NewRequest("POST", "https://www.flexiquiz.com/Publish/SetPublishedSurveyType", strings.NewReader(values.Encode()))
	if err != nil {
		log.Printf("[ERROR]WHEN SETTING TO PUBLIC, Err: %v", err)
	}
	cookie := &http.Cookie{Name: ".ASPXAUTH", Value: auth}
	req.AddCookie(cookie)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR]WHEN SETTING TO PUBLIC, Err: %v", err)
	}
	fmt.Println(resp.StatusCode)

	defer resp.Body.Close()
}
func SetRequiredEmail(auth, testlink string) {
	values := url.Values{}
	values.Set("SurveyId", testlink)
	values.Set("registrationPublishType", "quiz_link")
	t := &http2.Transport{}
	client := http.Client{
		Transport: t,
	}
	req, err := http.NewRequest("POST", "https://www.flexiquiz.com/Configure/RegistrationItems", strings.NewReader(values.Encode()))
	if err != nil {
		log.Printf("[ERROR]WHEN SETTING TO PUBLIC, Err : %v", err)
	}
	cookie := &http.Cookie{Name: ".ASPXAUTH", Value: auth}
	req.AddCookie(cookie)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR]WHEN SETTING TO PUBLIC, Err : %v", err)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("[ERROR]WHEN SETTING TO PUBLIC, Err : %v", err)
	}
	res := ""
	doc.Find("input[type=checkbox]").Not("[checked]").Each(func(i int, s *goquery.Selection) {
		id, _ := s.Attr("id")
		if strings.Contains(id, "registrationItemRequired") {
			res = id
			return
		}
	})
	res = strings.ReplaceAll(res, "registrationItemRequired_", "")
	resp.Body.Close()
	values2 := url.Values{}
	values2.Set("registrationItemId", res)
	values2.Set("required", "true")
	fmt.Println(values2.Encode())
	req2, err := http.NewRequest("POST", "https://www.flexiquiz.com/Configure/SaveRegistrationItemRequired", strings.NewReader(values2.Encode()))
	if err != nil {
		log.Printf("[ERROR]WHEN SETTING TO PUBLIC, Err : %v", err)
	}
	cookie2 := &http.Cookie{Name: ".ASPXAUTH", Value: auth}
	req2.AddCookie(cookie2)
	req2.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	respp, err := client.Do(req2)
	respp.Body.Close()
	if err != nil {
		log.Printf("[ERROR]WHEN SETTING TO PUBLIC, Err : %v", err)
	}

}
func Publish(auth, publink string) {
	values := url.Values{}
	values.Set("publishedSurveyId", publink)
	client := http.Client{}
	req, err := http.NewRequest("POST", "https://www.flexiquiz.com/Publish/PublishSurvey", strings.NewReader(values.Encode()))
	if err != nil {
		panic(err)
	}
	cookie := &http.Cookie{Name: ".ASPXAUTH", Value: auth}
	req.AddCookie(cookie)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)

	defer resp.Body.Close()
}
