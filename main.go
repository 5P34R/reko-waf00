package main

import (
	"flag"
	"fmt"
	"go-wafw00f/tool"
	"go-wafw00f/util"
	"go-wafw00f/waf"
)

func main() {
	util.PrintBanner()
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
		if len(resList) > 0 {
			fmt.Printf("%s => %s\n", url, resList)
		}
	}
}
