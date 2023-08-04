package main

import (
	"fmt"

	"github.com/g0ldencybersec/EasyEASM/pkg/configparser"
	"github.com/g0ldencybersec/EasyEASM/pkg/passive"
	"github.com/g0ldencybersec/EasyEASM/pkg/utils"
)

func main() {
	cfg := configparser.ParseConfig()
	Runner := passive.Runner{
		SeedDomains: cfg.RunConfig.Domains,
	}

	results := Runner.Run()

	Runner.Subdomains = utils.RemoveDuplicates(results)
	Runner.Results = len(Runner.Subdomains)

	fmt.Printf("Found %d subdomains\n\n", Runner.Results)
	fmt.Println(Runner.Subdomains)

}
