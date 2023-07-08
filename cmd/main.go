package main

import (
	"fmt"

	"github.com/g0ldencybersec/EasyEASM/pkg/configparser"
	"github.com/g0ldencybersec/EasyEASM/pkg/passive"
	"github.com/g0ldencybersec/EasyEASM/pkg/utils"
)

func main() {
	var domains []string
	cfg := configparser.ParseConfig()
	Runner := passive.Runner{
		Seed_domain: cfg.RunConfig.Seed_domain,
	}

	results := Runner.Run()
	for {
		res, ok := <-results
		if !ok {
			fmt.Println("Channel Closed")
			break
		} else {
			domains = append(domains, res)
		}
	}
	Runner.Data = utils.RemoveDuplicates(domains)
	Runner.Results = len(Runner.Data)

	fmt.Printf("Found %d subdomains for %s\n\n", Runner.Results, Runner.Seed_domain)
	fmt.Println(Runner.Data)

}
