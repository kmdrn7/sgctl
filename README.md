# SGCTL
### The missing ctl/cli tools for Sendgrid :D

## Usage
### Export your sendgrid api key
`export SENDGRID_API_KEY='your_sendgrid_api_key'`

### Download dependencies
`go mod download`

### Setup _mails.yaml_ from configs directory
`vim configs/mails.yaml`

### Run in Dryrun mode
`go run main.go --dryrun`

### Run actual command
`go run main.go`

## Note
there are some example files in `example/files` directory to play with