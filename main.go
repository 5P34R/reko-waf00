package main

import (
	"flag"
	"fmt"
	"go-wafw00f/tool"
	"go-wafw00f/waf"
)

func main() {

	urlPtr := flag.String("domain", "", "Enter the domain name")
	flag.Parse()

	if *urlPtr == "" {
		fmt.Printf("%s", "Enter a valid domain")
		return
	}

	domains := tool.SubfinderRun(*urlPtr)
	urls := tool.HttpxRun(domains)
	waf.ResolveWafLib()

	for _, url := range urls {

		resList := waf.DetectWaf(url)
		fmt.Printf("%s => %s\n", url, resList)
	}
}
