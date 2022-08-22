package correio

import (
	"errors"

	"github.com/GeoinovaDev/lower-resultys/convert/decode"
	"github.com/GeoinovaDev/lower-resultys/exec"
	"github.com/GeoinovaDev/lower-resultys/net/request"
	"github.com/GeoinovaDev/lower-resultys/str"
)

// Client struct
type Client struct {
	IP     string
	Apikey string
	Domain string
}

type protocol struct {
	Status string `json:"status"`
	Email  Email  `json:"data"`
}

// New ...
func New(ip string, apikey string, domain string) *Client {
	return &Client{
		IP:     ip,
		Apikey: apikey,
		Domain: domain,
	}
}

// CreateAndSend ...
func (c *Client) CreateAndSend(email *Email) {
	err := c.Create(email)
	if err != nil {
		return
	}

	c.Send(email)
}

// Create ...
func (c *Client) Create(email *Email) (erro error) {

	exec.Try(func() {
		resp, err := request.New(c.createURL("/email/insert", "")).PostJSON(email)
		if err != nil {
			erro = errors.New("erro ao conectar ao servidor de email")
			return
		}

		protocol := protocol{}
		decode.JSON(resp, &protocol)

		email.ID = protocol.Email.ID
		erro = nil
	}).Catch(func(message string) {
		erro = errors.New(message)
	})

	return
}

// Send ...
func (c *Client) Send(email *Email) {
	request.New(c.createURL("/email/send", email.ID)).Get()
}

func (c *Client) createURL(path string, id string) string {
	return str.Format("http://{0}{1}?apikey={2}&domain={3}&emailID={4}", c.IP, path, c.Apikey, c.Domain, id)
}
