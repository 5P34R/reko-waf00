package main

import (
	"fmt"
	"go-wafw00f/tool"
	"go-wafw00f/waf"
)

func main() {
	domains := tool.SubfinderRun("highradius.com")
	urls := tool.HttpxRun(domains)
	waf.ResolveWafLib()

	for _, url := range urls {

		resList := waf.DetectWaf(url)
		fmt.Printf("%s => %s\n", url, resList)
	}
}
