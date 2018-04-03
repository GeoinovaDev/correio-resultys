package client

import (
	"errors"
	"git.resultys.com.br/framework/lower/exec"
	"git.resultys.com.br/framework/lower/library/convert"
	"git.resultys.com.br/framework/lower/net"
	"git.resultys.com.br/framework/lower/net/request"
	"git.resultys.com.br/framework/lower/str"
	"strconv"
)

type Correio struct {
	Url    string
	Apikey string
	Domain string
}

type Email struct {
	Id      int    `json:"id"`
	From    string `json:"email_from"`
	To      string `json:"email_to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func (e *Email) LoadFromMap(m map[string]interface{}) {
	if val, ok := m["id"]; ok {
		id := val.(string)
		e.Id, _ = strconv.Atoi(id)
	}

	if val, ok := m["email_from"]; ok {
		e.From = val.(string)
	}

	if val, ok := m["email_to"]; ok {
		e.To = val.(string)
	}

	if val, ok := m["subject"]; ok {
		e.Subject = val.(string)
	}

	if val, ok := m["body"]; ok {
		e.Body = val.(string)
	}
}

func (p *Correio) SendEmail(email *Email) {
	err := p.Insert(email)
	if err != nil {
		return
	}

	p.Send(email)
}

func (p *Correio) Insert(email *Email) (erro error) {

	exec.Try(func() {
		protocol := net.Protocol{}
		resp, err := request.PostJson(p.createUrl("/email/insert", 0), email)
		if err != nil {
			erro = errors.New("erro ao conectar ao servidor de email")
			return
		}

		convert.StringToJson(resp, &protocol)
		email.LoadFromMap(protocol.Data.(map[string]interface{}))

		erro = nil
	}).Catch(func(message string) {
		erro = errors.New(message)
	})

	return
}

func (p *Correio) Send(email *Email) {
	request.Get(p.createUrl("/email/send", email.Id))
}

func (p *Correio) createUrl(path string, id int) string {
	return str.Format("{0}{1}?apikey={2}&domain={3}&emailId={4}", p.Url, path, p.Apikey, p.Domain, strconv.Itoa(id))
}
