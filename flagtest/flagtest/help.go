package flagtest

var (
	usage = `usage: flagtest [flags]
	do some flag test
	flagtest -h| --help`

	flags = `
	--ipaddr 127.0.0.1
	--port 8080
	--confg-file config.ini`
)
