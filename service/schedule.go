package service

import (
	"log"
	"time"

	"github.com/ropel12/email/entities"
	"gorm.io/gorm"
)

func InsertSchedule(data entities.Data, db *gorm.DB) {
	res := []entities.Payment{}
	err := db.Where("school_id = ? AND type = 'interval' AND deleted_at IS NULL", data.Schoolid).Find(&res).Error
	if err != nil {
		log.Printf("[ERROR]WHEN GETTING PAYMENT DATA, Error: %v", err)
	} else {
		for i := 1; i <= 36; i++ {
			total := 0
			for _, val := range res {
				if i%val.Interval == 0 {
					total += int(val.Price)
				}
			}
			t := time.Date(time.Now().Year(), time.May, 7, 0, 0, 0, 0, time.UTC)
			t = t.AddDate(0, i, 0)
			if err := db.Create(&entities.BillingSchedule{StudentName: data.Name, StudentEmail: data.Email, SchoolName: data.School, Date: t.Format("2006-01-02"), Total: total}).Error; err != nil {
				log.Printf("[ERROR]WHEN CREATING SCHEDULE, Err: %v", err)
			}
		}
		log.Println("[INFO]SUCCESS CREATING SCHEDULE")
	}

}
