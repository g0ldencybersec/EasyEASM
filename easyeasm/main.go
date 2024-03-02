package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sethlaw/EasyEASM/pkg/active"
	"github.com/sethlaw/EasyEASM/pkg/configparser"
	"github.com/sethlaw/EasyEASM/pkg/passive"
	"github.com/sethlaw/EasyEASM/pkg/utils"
)

func main() {
	// install required tools
	utils.InstallTools()

	// print a banner
	banner := "\x1b[36m****************\n\nEASY EASM\n\n***************\x1b[0m\n"
	fmt.Println(banner)

	// parse the configuration file
	cfg := configparser.ParseConfig()

	// check for previous run file
	var prevRun bool
	if _, err := os.Stat("EasyEASM.csv"); err == nil {
		fmt.Println("Found data from previous run!")
		prevRun = true
		e := os.Rename("EasyEASM.csv", "old_EasyEASM.csv")
		if e != nil {
			panic(e)
		}
		var domains = utils.getActiveDomains("./easyeasm.db")
		fmt.Println("Active domains from previous run: ", len(domains))
	} else {
		fmt.Println("No previous run data found")
		utils.setupDB()
		prevRun = false
	}

	// check the run type specified in the config and perform actions accordingly
	if strings.ToLower(cfg.RunConfig.RunType) == "fast" {
		// fast run: passive enumeration only

		// create a PassiveRunner instance
		Runner := passive.PassiveRunner{
			SeedDomains: cfg.RunConfig.Domains,
		}

		// run passive enumeration and get the results
		passiveResults := Runner.RunPassiveEnum()

		// remove duplicate subdomains
		Runner.Subdomains = utils.RemoveDuplicates(passiveResults)
		Runner.Results = len(Runner.Subdomains)

		fmt.Printf("\x1b[31mFound %d subdomains\n\n\x1b[0m", Runner.Results)
		fmt.Println(Runner.Subdomains)
		fmt.Println("Checking which domains are live and generating assets csv...")

		// run Httpx to check live domains
		Runner.RunHttpx()

		// notify about new domains if prevRun is true
		if prevRun && strings.Contains(cfg.RunConfig.SlackWebhook, "https") {
			utils.NotifyNewDomainsSlack(Runner.Subdomains, cfg.RunConfig.SlackWebhook)
			os.Remove("old_EasyEASM.csv")
			os.Remove("old_tracked.csv")
		} else if prevRun && strings.Contains(cfg.RunConfig.DiscordWebhook, "https") {
			utils.NotifyNewDomainsDiscord(Runner.Subdomains, cfg.RunConfig.DiscordWebhook)
			os.Remove("old_EasyEASM.csv")
			os.Remove("old_tracked.csv")
		}
	} else if strings.ToLower(cfg.RunConfig.RunType) == "complete" {
		// complete run: passive and active enumeration

		// passive enumeration
		PassiveRunner := passive.PassiveRunner{
			SeedDomains: cfg.RunConfig.Domains,
		}
		passiveResults := PassiveRunner.RunPassiveEnum()

		// remove duplicate subdomains
		PassiveRunner.Subdomains = utils.RemoveDuplicates(passiveResults)
		PassiveRunner.Results = len(PassiveRunner.Subdomains)

		// active enumeration
		ActiveRunner := active.ActiveRunner{
			SeedDomains: cfg.RunConfig.Domains,
		}
		activeResults := ActiveRunner.RunActiveEnum(cfg.RunConfig.ActiveWordlist, cfg.RunConfig.ActiveThreads)
		activeResults = append(activeResults, passiveResults...)

		ActiveRunner.Subdomains = utils.RemoveDuplicates(activeResults)

		// permutation scan
		permutationResults := ActiveRunner.RunPermutationScan(cfg.RunConfig.ActiveThreads)
		ActiveRunner.Subdomains = append(ActiveRunner.Subdomains, permutationResults...)
		ActiveRunner.Subdomains = utils.RemoveDuplicates(ActiveRunner.Subdomains)
		ActiveRunner.Results = len(ActiveRunner.Subdomains)

		// httpx scan
		fmt.Printf("Found %d subdomains: ", ActiveRunner.Results)
		fmt.Println(ActiveRunner.Subdomains)
		if !prevRun {
			for _, domain := range ActiveRunner.Subdomains {
				utils.insertDomain(domain, "./easyeasm.db")
			}
		}
		fmt.Println("Checking which domains are live and generating assets csv...")
		ActiveRunner.RunHttpx()

		// notify about new domains if prevRun is true
		if prevRun && strings.Contains(cfg.RunConfig.SlackWebhook, "https") {
			utils.NotifyNewDomainsSlack(ActiveRunner.Subdomains, cfg.RunConfig.SlackWebhook)
			os.Remove("old_EasyEASM.csv")
			os.Remove("old_tracked.csv")
		} else if prevRun && strings.Contains(cfg.RunConfig.DiscordWebhook, "https") {
			utils.NotifyNewDomainsDiscord(ActiveRunner.Subdomains, cfg.RunConfig.DiscordWebhook)
			os.Remove("old_EasyEASM.csv")
			os.Remove("old_tracked.csv")
		}
	} else {
		// invalid run mode specified
		panic("Please pick a valid run mode and add it to your config.yml file! You can set runType to either 'fast' or 'complete'")
	}
}
