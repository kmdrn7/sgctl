package sgctl

func Dryrun() {
	// 1. parse yaml config
	mails := ReadConfig()

	// 2. sending email using sendgrid api (dryrun)
	send_dryrun(mails)
}
