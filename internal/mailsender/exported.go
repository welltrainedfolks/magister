// Copyright (c) 2018, Well Trained Folks and contributors.
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject
// to the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
// CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package mailsender

import (
	// stdlib
	"crypto/tls"

	// local
	"github.com/welltrainedfolks/magister/internal/config"
	"github.com/welltrainedfolks/magister/internal/templater"

	// other
	"github.com/jordan-wright/email"
	"github.com/rs/zerolog/log"
)

func Initialize() {
	log.Info().Msg("Initializing mails sender...")
}

func SendMail(templateName string, data map[string]string) {
	go func() {
		_, found := data["mail.to"]
		if !found {
			log.Error().Msgf("Tried to send email on empty address! Aborting! Passed data was: %+v", data)
			return
		}

		var shouldAuth bool

		if config.Config.MailSender.User != "" {
			shouldAuth = true
		}

		if shouldAuth {
			//sendMailWithAuth(to, templateName, data)
			log.Error().Msg("Mail sender cannot authorize at mail servers! Will not send any email!")
		} else {
			sendMailWithoutAuth(templateName, data)
		}
	}()
}

func sendMailWithoutAuth(templateName string, data map[string]string) {
	log.Info().Msgf("Sending mail to '%s' without authorization on mail server...", data["mail.to"])

	tpl := templater.GetTextTemplate(templateName, data)

	mail := email.NewEmail()
	mail.Headers.Set("Content-Transfer-Encoding", "quoted-printable")
	mail.From = config.Config.MailSender.From
	// ToDo: multiple recipients.
	mail.To = []string{data["mail.to"]}
	mail.Subject = data["mail.subject"]
	mail.Text = []byte(tpl)
	err := mail.SendWithTLS(config.Config.MailSender.Host, nil, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Error().Msgf("Failed to send mail: %s", err.Error())
	}
}
