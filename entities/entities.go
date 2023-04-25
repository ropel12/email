package entities

type (
	Data struct {
		Invoice       string `json:"invoice"`
		Total         int    `json:"total"`
		Name          string `json:"name"`
		Email         string `json:"email"`
		PaymentCode   string `json:"payment_code"`
		PaymentMethod string `json:"payment_method"`
		Expire        string `json:"expire"`
	}
)
