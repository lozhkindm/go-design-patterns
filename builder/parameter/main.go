package main

import (
	"log"
	"strings"
)

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		log.Fatal("email should contain @")
	}
	b.email.from = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	if !strings.Contains(to, "@") {
		log.Fatal("email should contain @")
	}
	b.email.to = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

type build func(builder *EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}
	action(&builder)
	sendEmail(&builder.email)
}

func sendEmail(email *email) {
	log.Printf(
		"sending email from %q to %q with subject %q and body %q\n",
		email.from,
		email.to,
		email.subject,
		email.body,
	)
}

func main() {
	SendEmail(func(builder *EmailBuilder) {
		builder.
			From("foo@bar.com").
			To("bar@baz.com").
			Subject("Meeting").
			Body("Expecting you at the meeting today.")
	})
}
