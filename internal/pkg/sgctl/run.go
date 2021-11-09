package sgctl

func Run() {
	// 1. parse yaml config
	mails := ReadConfig()

	// 2. sending email using sendgrid api
	send(mails)
}
