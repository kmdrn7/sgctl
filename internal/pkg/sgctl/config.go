package sgctl

import (
	"errors"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	log "github.com/sirupsen/logrus"
)

type Mail struct {
	Subject string `koanf:"subject"`
	From    struct {
		Name  string `koanf:"name"`
		Email string `koanf:"email"`
	}
	To []struct {
		Name  string `koanf:"name"`
		Email string `koanf:"email"`
	}
	Content struct {
		Type  string `koanf:"type"`
		Value string `koanf:"value"`
	}
	Attachments []struct {
		Path string `koanf:"path"`
	}
}

type MailsConfig struct {
	SendgridApiKey string `koanf:"sendgrid_api_key"`
	Mails          []Mail `koanf:"mails"`
}

func ReadConfig() *MailsConfig {
	k := koanf.New(".")

	// Load sendgrid configs from environment variables
	if err := k.Load(env.Provider("SENDGRID_", ".", func(s string) string {
		return strings.ToLower(s)
	}), nil); err != nil {
		panic(err)
	}

	// Load SGCTL configs from yaml configuration file
	if err := k.Load(file.Provider("configs/mails.yaml"), yaml.Parser()); err != nil {
		log.Panicln("Error parsing configuration, ", err)
	}

	// Unmarshall configs from koanf
	mailsConfig := &MailsConfig{}
	if err := k.Unmarshal("", mailsConfig); err != nil {
		log.Panicln("Error unmarshall configuration, ", err)
	}

	// TODO: Add error checking on emails config
	var err error
	if mailsConfig.SendgridApiKey == "" {
		err = errors.New("missing sendgrid api key")
	}
	for _, mail := range mailsConfig.Mails {
		if mail.From.Email == "" {
			err = errors.New("error email configuration, missing from.email value")
		}
		if mail.Subject == "" {
			err = errors.New("error email configuration, missing subject value")
		}
		if mail.Content.Type == "" {
			err = errors.New("error email configuration, missing content.type value")
		}
		if mail.Content.Value == "" {
			err = errors.New("error email configuration, missing content.value value")
		}
		for _, to := range mail.To {
			if to.Email == "" {
				err = errors.New("error email configuration, missing to.email value")
			}
		}
	}

	if err != nil {
		log.Panicln(err)
	}

	return mailsConfig
}
