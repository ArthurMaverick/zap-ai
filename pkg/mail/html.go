package mail

import (
	"bytes"
	"log/slog"
	"text/template"
)

type BodyRequest struct {
	To    string
	Token string
}

func ParseHtml(fileName string, data map[string]string) string {
	html, errParse := template.ParseFiles("template/" + fileName + ".html")
	if errParse != nil {
		defer slog.Error(errParse.Error())
	}

	body := BodyRequest{
		To:    data["to"],
		Token: data["token"],
	}

	buf := new(bytes.Buffer)
	errExecute := html.Execute(buf, body)
	if errExecute != nil {
		defer slog.Error(errExecute.Error())
	}

	return buf.String()
}
