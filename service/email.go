package service

import (
	"bytes"
	"crypto/tls"
	"encoding/base32"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/ropel12/email/config"
	"github.com/ropel12/email/entities"

	gomail "gopkg.in/mail.v2"
)

const (
	FrontEndURL           = "https://go-event.online/verify/"
	FrontEndURL3          = "https://go-event.online/updateverif/"
	FrontEndURL2          = "https://education-hub-fe-3q5c.vercel.app/pwd-new/"
	FrontEndURLDetailCost = "https://education-hub-fe-3q5c.vercel.app/student/transactions/"

	EmailHost = "smtp.gmail.com"
)

type subjectBody struct {
	subject string
	body    bytes.Buffer
}

func SendEmailPendingPayment(sdata config.SenderConfig, rdata entities.Data) {
	log.Printf("[INFO] Sending email pending payment from %s", sdata.Email)

	sb := subjectBody{
		subject: fmt.Sprintf("Tagihan %s", rdata.Invoice),
		body:    bytes.Buffer{},
	}

	t, err := getTemplate("pending.html")
	if err != nil {
		log.Printf("[ERROR] Failed to get template: %s", err)
		return
	}

	err = t.Execute(&sb.body, struct {
		TWT     string
		IG      string
		FB      string
		URL     string
		Email   string
		Telpon  string
		Name    string
		Slogan  string
		Cusname string
		Pcode   string
		Invoice string
		Total   int
		Pmethod string
		Expire  string
	}{
		URL:     FrontEndURL,
		TWT:     sdata.Twitter,
		FB:      sdata.Facebook,
		IG:      sdata.Instagram,
		Email:   sdata.Email,
		Telpon:  sdata.Phone,
		Cusname: rdata.Name,
		Invoice: rdata.Invoice,
		Slogan:  sdata.Slogan,
		Pmethod: rdata.PaymentMethod,
		Name:    sdata.Name,
		Total:   rdata.Total,
		Expire:  rdata.Expire,
		Pcode:   rdata.PaymentCode,
	})

	if err != nil {
		log.Fatalf("[ERROR] Failed to execute template: %v", err)
	}

	if err := sendEmail(sdata.Email, sdata.Password, rdata.Email, sb); err != nil {
		log.Printf("[ERROR] Failed to send email: %s", err)
		return
	}

	log.Printf("[INFO] Email pending payment sent to %s", rdata.Email)

}
func SendEmailSuccessPayment(sdata config.SenderConfig, rdata entities.Data) {
	log.Printf("[INFO] Sending email success payment from %s", sdata.Email)

	sb := subjectBody{
		subject: fmt.Sprintf("Berhasil %s", rdata.Invoice),
		body:    bytes.Buffer{},
	}

	t, err := getTemplate("success.html")
	if err != nil {
		log.Printf("[ERROR] Failed to get template: %s", err)
		return
	}

	err = t.Execute(&sb.body, struct {
		TWT     string
		IG      string
		FB      string
		URL     string
		Email   string
		Telpon  string
		Name    string
		Slogan  string
		Cusname string
		Invoice string
	}{
		URL:     FrontEndURL,
		TWT:     sdata.Twitter,
		FB:      sdata.Facebook,
		IG:      sdata.Instagram,
		Email:   sdata.Email,
		Telpon:  sdata.Phone,
		Cusname: rdata.Name,
		Invoice: rdata.Invoice,
		Slogan:  sdata.Slogan,
		Name:    sdata.Name,
	})

	if err != nil {
		log.Fatalf("[ERROR] Failed to execute template: %v", err)
	}

	if err := sendEmail(sdata.Email, sdata.Password, rdata.Email, sb); err != nil {
		log.Printf("[ERROR] Failed to send email: %s", err)
		return
	}

	log.Printf("[INFO] Email success payment sent to %s", rdata.Email)
}

func SendEmailCancelPayment(sdata config.SenderConfig, rdata entities.Data) {
	log.Printf("[INFO] Sending email cancel payment from %s", sdata.Email)

	sb := subjectBody{
		subject: fmt.Sprintf("Pembatalan %s", rdata.Invoice),
		body:    bytes.Buffer{},
	}

	t, err := getTemplate("cancel.html")
	if err != nil {
		log.Printf("[ERROR] Failed to get template: %s", err)
		return
	}

	err = t.Execute(&sb.body, struct {
		TWT     string
		IG      string
		FB      string
		URL     string
		Email   string
		Telpon  string
		Name    string
		Slogan  string
		Cusname string
		Invoice string
	}{
		URL:     FrontEndURL,
		TWT:     sdata.Twitter,
		FB:      sdata.Facebook,
		IG:      sdata.Instagram,
		Email:   sdata.Email,
		Telpon:  sdata.Phone,
		Cusname: rdata.Name,
		Invoice: rdata.Invoice,
		Slogan:  sdata.Slogan,
		Name:    sdata.Name,
	})

	if err != nil {
		log.Fatalf("[ERROR] Failed to execute template: %v", err)
	}

	if err := sendEmail(sdata.Email, sdata.Password, rdata.Email, sb); err != nil {
		log.Printf("[ERROR] Failed to send email: %s", err)
		return
	}

	log.Printf("[INFO] Email cancel payment sent to %s", rdata.Email)

}
func SendEmailRefundPayment(sdata config.SenderConfig, rdata entities.Data) {
	log.Printf("[INFO] Sending email refund payment from %s", sdata.Email)

	sb := subjectBody{
		subject: fmt.Sprintf("Pengembalian %s", rdata.Invoice),
		body:    bytes.Buffer{},
	}

	t, err := getTemplate("refund.html")
	if err != nil {
		log.Printf("[ERROR] Failed to get template: %s", err)
		return
	}

	err = t.Execute(&sb.body, struct {
		TWT     string
		IG      string
		FB      string
		URL     string
		Email   string
		Telpon  string
		Name    string
		Slogan  string
		Cusname string
		Invoice string
	}{
		URL:     FrontEndURL,
		TWT:     sdata.Twitter,
		FB:      sdata.Facebook,
		IG:      sdata.Instagram,
		Email:   sdata.Email,
		Telpon:  sdata.Phone,
		Cusname: rdata.Name,
		Invoice: rdata.Invoice,
		Slogan:  sdata.Slogan,
		Name:    sdata.Name,
	})

	if err != nil {
		log.Fatalf("[ERROR] Failed to execute template: %v", err)
	}

	if err := sendEmail(sdata.Email, sdata.Password, rdata.Email, sb); err != nil {
		log.Printf("[ERROR] Failed to send email: %s", err)
		return
	}

	log.Printf("[INFO] Email refund payment sent to %s", rdata.Email)

}
func SendTest(sdata config.SenderConfig, rdata entities.Data) {
	log.Printf("[INFO] Sending email SendTest from %s", sdata.Email)

	sb := subjectBody{
		subject: fmt.Sprintf("Test Pendaftaran %s", rdata.School),
		body:    bytes.Buffer{},
	}

	t, err := getTemplate("test.html")
	if err != nil {
		log.Printf("[ERROR] Failed to get template: %s", err)
		return
	}

	err = t.Execute(&sb.body, struct {
		TWT     string
		IG      string
		FB      string
		URL     string
		Email   string
		Telpon  string
		Name    string
		Slogan  string
		Cusname string
		School  string
	}{
		URL:     rdata.Test,
		TWT:     sdata.Twitter,
		FB:      sdata.Facebook,
		IG:      sdata.Instagram,
		Email:   sdata.Email,
		Telpon:  sdata.Phone,
		Cusname: rdata.Name,
		Slogan:  sdata.Slogan,
		Name:    sdata.Name,
		School:  rdata.School,
	})

	if err != nil {
		log.Fatalf("[ERROR] Failed to execute template: %v", err)
	}

	if err := sendEmail(sdata.Email, sdata.Password, rdata.Email, sb); err != nil {
		log.Printf("[ERROR] Failed to send email: %s", err)
		return
	}

	log.Printf("[INFO] Email SendTest sent to %s", rdata.Email)

}
func SendEmailVerification(sdata config.SenderConfig, hashedemail string) {
	log.Printf("[INFO] Sending email verification from %s", sdata.Email)

	sb := subjectBody{
		subject: fmt.Sprintf("Verifikasi Email"),
		body:    bytes.Buffer{},
	}
	email, err := base32.StdEncoding.DecodeString(hashedemail)
	if err != nil {
		log.Printf("[ERROR] Failed to decode email receiver: %s", err)
		return
	}
	t, err := getTemplate("verify.html")
	if err != nil {
		log.Printf("[ERROR] Failed to get template: %s", err)
		return
	}

	err = t.Execute(&sb.body, struct {
		URL string
	}{
		URL: FrontEndURL + hashedemail,
	})

	if err != nil {
		log.Fatalf("[ERROR] Failed to execute template: %v", err)
	}

	if err := sendEmail(sdata.Email, sdata.Password, string(email), sb); err != nil {
		log.Printf("[ERROR] Failed to send email: %s", err)
		return
	}

	log.Printf("[INFO] Email verification sent to %s", string(email))

}
func SendEmailResetPassword(sdata config.SenderConfig, hashedemail string) {
	log.Printf("[INFO] Sending Reset Password from %s", sdata.Email)

	sb := subjectBody{
		subject: fmt.Sprintf("Reset Password"),
		body:    bytes.Buffer{},
	}
	email, err := base32.StdEncoding.DecodeString(hashedemail)
	if err != nil {
		log.Printf("[ERROR] Failed to decode email receiver: %s", err)
		return
	}
	t, err := getTemplate("reset.html")
	if err != nil {
		log.Printf("[ERROR] Failed to get template: %s", err)
		return
	}

	err = t.Execute(&sb.body, struct {
		URL string
	}{
		URL: FrontEndURL2 + hashedemail,
	})

	if err != nil {
		log.Fatalf("[ERROR] Failed to execute template: %v", err)
	}

	if err := sendEmail(sdata.Email, sdata.Password, string(email), sb); err != nil {
		log.Printf("[ERROR] Failed to send email: %s", err)
		return
	}

	log.Printf("[INFO] Email Reset Password sent to %s", string(email))

}
func SendEmailChangeEmail(sdata config.SenderConfig, hashedemail string) {
	log.Printf("[INFO] Sending Change Email from %s", sdata.Email)

	sb := subjectBody{
		subject: fmt.Sprintf("Change Email"),
		body:    bytes.Buffer{},
	}
	email, err := base32.StdEncoding.DecodeString(hashedemail)
	if err != nil {
		log.Printf("[ERROR] Failed to decode email receiver: %s", err)
		return
	}
	t, err := getTemplate("changeemail.html")
	if err != nil {
		log.Printf("[ERROR] Failed to get template: %s", err)
		return
	}

	err = t.Execute(&sb.body, struct {
		URL string
	}{
		URL: FrontEndURL3 + hashedemail,
	})

	if err != nil {
		log.Fatalf("[ERROR] Failed to execute template: %v", err)
	}

	if err := sendEmail(sdata.Email, sdata.Password, string(email), sb); err != nil {
		log.Printf("[ERROR] Failed to send email: %s", err)
		return
	}

	log.Printf("[INFO] Email Change Email sent to %s", string(email))

}

func SendDetailCost(sdata config.SenderConfig, rdata entities.Data) {
	log.Printf("[INFO] Sending Change Email from %s", sdata.Email)

	sb := subjectBody{
		subject: fmt.Sprintf("Detail Biaya %s", rdata.School),
		body:    bytes.Buffer{},
	}

	t, err := getTemplate("info.html")
	if err != nil {
		log.Printf("[ERROR] Failed to get template: %s", err)
		return
	}

	err = t.Execute(&sb.body, struct {
		TWT     string
		IG      string
		FB      string
		URL     string
		Email   string
		Telpon  string
		Name    string
		Slogan  string
		Cusname string
		Type    string
		School  string
	}{
		URL:     fmt.Sprintf("%s%d", FrontEndURLDetailCost, rdata.Schoolid),
		TWT:     sdata.Twitter,
		FB:      sdata.Facebook,
		IG:      sdata.Instagram,
		Email:   sdata.Email,
		Telpon:  sdata.Phone,
		Cusname: rdata.Name,
		Slogan:  sdata.Slogan,
		Name:    sdata.Name,
		Type:    rdata.Type,
		School:  rdata.School,
	})

	if err != nil {
		log.Fatalf("[ERROR] Failed to execute template: %v", err)
	}

	if err := sendEmail(sdata.Email, sdata.Password, rdata.Email, sb); err != nil {
		log.Printf("[ERROR] Failed to send email: %s", err)
		return
	}

	log.Printf("[INFO] Email Change Email sent to %s", rdata.Email)
}
func SendFinishRegister(sdata config.SenderConfig, rdata entities.Data) {
	log.Printf("[INFO] Sending Finish Registration Email from %s", sdata.Email)

	sb := subjectBody{
		subject: fmt.Sprintf("Penerimaan Sekolah %s", rdata.School),
		body:    bytes.Buffer{},
	}

	t, err := getTemplate("finish.html")
	if err != nil {
		log.Printf("[ERROR] Failed to get template: %s", err)
		return
	}

	err = t.Execute(&sb.body, struct {
		TWT     string
		IG      string
		FB      string
		Email   string
		Telpon  string
		Name    string
		Slogan  string
		Cusname string
		School  string
	}{
		TWT:     sdata.Twitter,
		FB:      sdata.Facebook,
		IG:      sdata.Instagram,
		Email:   sdata.Email,
		Telpon:  sdata.Phone,
		Cusname: rdata.Name,
		Slogan:  sdata.Slogan,
		Name:    sdata.Name,
		School:  rdata.School,
	})

	if err != nil {
		log.Fatalf("[ERROR] Failed to execute template: %v", err)
	}

	if err := sendEmail(sdata.Email, sdata.Password, rdata.Email, sb); err != nil {
		log.Printf("[ERROR] Failed to send email: %s", err)
		return
	}

	log.Printf("[INFO] Finish Registration Email sent to %s", rdata.Email)
}

func SendFailRegistration(sdata config.SenderConfig, rdata entities.Data) {
	log.Printf("[INFO] Sending Finish Registration Email from %s", sdata.Email)

	sb := subjectBody{
		subject: fmt.Sprintf("Pendaftaran Gagal %s", rdata.School),
		body:    bytes.Buffer{},
	}

	t, err := getTemplate("failed.html")
	if err != nil {
		log.Printf("[ERROR] Failed to get template: %s", err)
		return
	}

	err = t.Execute(&sb.body, struct {
		TWT     string
		IG      string
		FB      string
		Email   string
		Telpon  string
		Name    string
		Slogan  string
		Cusname string
		School  string
		Reason  string
	}{
		TWT:     sdata.Twitter,
		FB:      sdata.Facebook,
		IG:      sdata.Instagram,
		Email:   sdata.Email,
		Telpon:  sdata.Phone,
		Cusname: rdata.Name,
		Slogan:  sdata.Slogan,
		Name:    sdata.Name,
		School:  rdata.School,
		Reason:  rdata.Reason,
	})

	if err != nil {
		log.Fatalf("[ERROR] Failed to execute template: %v", err)
	}

	if err := sendEmail(sdata.Email, sdata.Password, rdata.Email, sb); err != nil {
		log.Printf("[ERROR] Failed to send email: %s", err)
		return
	}

	log.Printf("[INFO] Fail Registration Email sent to %s", rdata.Email)
}
func SendMonthlyBilling(sdata config.SenderConfig, rdata entities.Data) {
	log.Printf("[INFO] Sending Monthly Billing Email from %s", sdata.Email)

	sb := subjectBody{
		subject: fmt.Sprintf("Tagihan Bulanan - %s", rdata.School),
		body:    bytes.Buffer{},
	}

	t, err := getTemplate("billing.html")
	if err != nil {
		log.Printf("[ERROR] Failed to get template: %s", err)
		return
	}

	err = t.Execute(&sb.body, struct {
		TWT     string
		IG      string
		FB      string
		Email   string
		Telpon  string
		Name    string
		Slogan  string
		Cusname string
		School  string
		Total   int
	}{
		TWT:     sdata.Twitter,
		FB:      sdata.Facebook,
		IG:      sdata.Instagram,
		Email:   sdata.Email,
		Telpon:  sdata.Phone,
		Cusname: rdata.Name,
		Slogan:  sdata.Slogan,
		Name:    sdata.Name,
		School:  rdata.School,
		Total:   rdata.Total,
	})

	if err != nil {
		log.Fatalf("[ERROR] Failed to execute template: %v", err)
	}

	if err := sendEmail(sdata.Email, sdata.Password, rdata.Email, sb); err != nil {
		log.Printf("[ERROR] Failed to send email: %s", err)
		return
	}

	log.Printf("[INFO] Failed Monthly Billing Email sent to %s", rdata.Email)
}
func getTemplate(htmlFile string) (t *template.Template, err error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	wd = wd + "/template/"

	t, err = template.ParseFiles(wd + htmlFile)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func sendEmail(emailSender, passSender, emailReceiver string, sb subjectBody) error {
	m := gomail.NewMessage()

	m.SetHeader("From", emailSender)
	m.SetHeader("To", emailReceiver)
	m.SetHeader("Subject", sb.subject)
	m.SetBody("text/html", string(sb.body.Bytes()))

	d := gomail.NewDialer(EmailHost, 587, emailSender, passSender)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
