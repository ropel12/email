package entities

import "gorm.io/gorm"

type (
	Data struct {
		Invoice       string `json:"invoice"`
		Total         int    `json:"total"`
		Name          string `json:"name"`
		Email         string `json:"email"`
		PaymentCode   string `json:"payment_code"`
		PaymentMethod string `json:"payment_method"`
		Expire        string `json:"expire"`
		School        string `json:"school"`
		Test          string `json:"test"`
		Type          string `json:"type"`
		Schoolid      int    `json:"school_id"`
		Reason        string `json:"reason"`
		Uid           string `json:"user_id"`
	}
	ReqAddQuiz struct {
		SchoolID int    `json:"school_id"`
		Question string `json:"question"`
		Option1  string `json:"option1"`
		Option2  string `json:"option2"`
		Option3  string `json:"option3"`
		Option4  string `json:"option4"`
		Answer   int    `json:"answer"`
	}
	Payment struct {
		gorm.Model
		SchoolID    uint
		Description string `gorm:"type:varchar(255);not null"`
		Image       string `gorm:"type:varchar(70);not null"`
		Type        string `gorm:"type:varchar(15);not null"`
		Price       int
		Interval    int
	}
	BillingSchedule struct {
		ID           uint `gorm:"primaryKey;not null;autoIncrement"`
		StudentName  string
		StudentEmail string
		SchoolName   string
		DeletedAt    gorm.DeletedAt `gorm:"index"`
		Date         string         `gorm:"type:timestamp;not null"`
		Total        int
	}
	ReqDataQuiz struct {
		PubLink    string       `json:"pub_link"`
		Prevlink   string       `json:"prev_link"`
		ResultLink string       `json:"result_link"`
		Data       []ReqAddQuiz `json:"data"`
	}
)
