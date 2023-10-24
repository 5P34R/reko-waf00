package tool

import (
	"log"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	"github.com/projectdiscovery/httpx/runner"
)

func HttpxRun(domains []string) []string {
	var subdomain []string
	gologger.DefaultLogger.SetMaxLevel(levels.LevelVerbose) // increase the verbosity (optional)

	options := runner.Options{
		Methods:         "GET",
		InputTargetHost: domains,
		OnResult: func(r runner.Result) {

			subdomain = append(subdomain, r.URL)
		},
	}

	if err := options.ValidateOptions(); err != nil {
		log.Fatal(err)
	}

	httpxRunner, err := runner.New(&options)
	if err != nil {
		log.Fatal(err)
	}
	defer httpxRunner.Close()

	httpxRunner.RunEnumeration()

	return subdomain
}
