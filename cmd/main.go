package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/g0ldencybersec/EasyEASM/pkg/configparser"
	"github.com/g0ldencybersec/EasyEASM/pkg/passive"
	"github.com/g0ldencybersec/EasyEASM/pkg/utils"
)

func main() {
	// Parse the configuraton file
	cfg := configparser.ParseConfig()

	// Check for previous run file
	var prevRun bool
	if _, err := os.Stat("EasyEASM.csv"); err == nil {
		fmt.Println("Found data from previous run!")
		prevRun = true
		e := os.Rename("EasyEASM.csv", "old_EasyEASM.csv")
		if e != nil {
			panic(e)
		}
	} else {
		fmt.Println("No previous run data found")
		prevRun = false
		fmt.Printf("File does not exist\n")
	}

	// Fast run. This is passive enumeration only
	if strings.ToLower(cfg.RunConfig.RunType) == "fast" {
		Runner := passive.PassiveRunner{
			SeedDomains: cfg.RunConfig.Domains,
		}
		passiveResults := Runner.RunPassiveEnum()

		Runner.Subdomains = utils.RemoveDuplicates(passiveResults)
		Runner.Results = len(Runner.Subdomains)

		fmt.Printf("Found %d subdomains\n\n", Runner.Results)
		fmt.Println(Runner.Subdomains)
		fmt.Println("Checking which domains are live and generating assets csv...")
		Runner.RunHttpx()
		if prevRun && strings.Contains(cfg.RunConfig.SlackWebhook, "https") {
			utils.NotifyNewDomainsSlack(Runner.Subdomains, cfg.RunConfig.SlackWebhook)
			os.Remove("old_EasyEASM.csv")
		} else if prevRun && strings.Contains(cfg.RunConfig.DiscordWebhook, "https") {
			utils.NotifyNewDomainsDiscord(Runner.Subdomains, cfg.RunConfig.DiscordWebhook)
			os.Remove("old_EasyEASM.csv")
		}
	} else if strings.ToLower(cfg.RunConfig.RunType) == "complete" {
		fmt.Println("Run complete recon")
	} else {
		panic("Please pick a valid run mode and add it to your config.yml file! You can set runType to either 'fast' or 'complete'")
	}

}
