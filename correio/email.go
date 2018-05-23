package correio

// Email struct
type Email struct {
	ID      string `json:"id"`
	From    string `json:"email_from"`
	To      string `json:"email_to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}
