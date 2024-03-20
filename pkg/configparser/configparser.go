package configparser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/g0ldencybersec/EasyEASM/pkg/utils"
	"gopkg.in/yaml.v3"
)

type Config struct {
	RunConfig struct {
		Domains        []string `yaml:"domains"`
		SlackWebhook   string   `yaml:"slack"`
		DiscordWebhook string   `yaml:"discord"`
		RunType        string   `yaml:"runType"`
		ActiveWordlist string   `yaml:"activeWordList"`
		ActiveThreads  int      `yaml:"activeThreads"`
	} `yaml:"runConfig"`
}

func ParseConfig(flags string) Config {
	// Read file data
	data, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Initialize configuration
	var config Config

	// Unmarshal YAML data into Config struct
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	//runtime config modification if flag -i is provided
	if flags == "interactive" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Do you want to change anything in the config?")
		opt, _ := utils.GetInput("Press \"y\" to make changes or any characther to keep running\n", reader)
		if opt == "y" {
			config = PromptConfigChange(config)
		}
	}

	return config
}

func PromptConfigChange(config Config) (cfg Config) {
	//runtime changes to the config file if flag -i is provided
	fmt.Println("Choose an option or press any other character to run without anymore changes")
	reader := bufio.NewReader(os.Stdin)

	opt, _ := utils.GetInput("1.Domains - 2.Slack - 3.Discord - 4.Run Type 5.N of Threads\n", reader)
	switch opt {

	//add domains to the list
	case "1":
		opt, _ = utils.GetInput("Write the domain you would like to add\n", reader)

		//check if the domain is in a valid format (doesnt ensure that the domain exists)
		if utils.ValidDomain(opt) {
			config.RunConfig.Domains = append(config.RunConfig.Domains, opt)

			yamlData, err := yaml.Marshal(config)
			if err != nil {
				log.Fatalf("error marshalling YAML: %v", err)
			}

			// Write modified data back to YAML file
			err = os.WriteFile("config.yml", yamlData, 0644)
			if err != nil {
				log.Fatalf("error writing YAML file: %v", err)
			}

			fmt.Println("Domain added successfully")
			PromptConfigChange(config)
		} else {
			fmt.Printf("Invalid Domain format\n\n")
			PromptConfigChange(config)
		}

	//add slack webhook at runtime
	case "2":
		opt, _ = utils.GetInput("Insert the Slack Webhook, end by pressing \"Enter\"\n", reader)
		config.RunConfig.SlackWebhook = opt

		//marshal back the data to yml
		yamlData, err := yaml.Marshal(config)
		if err != nil {
			log.Fatalf("error marshalling YAML: %v", err)
		}

		// Write modified data back to YAML file
		err = os.WriteFile("config.yml", yamlData, 0644)
		if err != nil {
			log.Fatalf("error writing YAML file: %v", err)
		}

		fmt.Println("Slack Webhook added successfully")
		PromptConfigChange(config)

	//add discord webhook at runtime
	case "3":
		opt, _ = utils.GetInput("Insert the Discord Webhook, end by pressing \"Enter\"\n", reader)
		config.RunConfig.DiscordWebhook = opt

		//marshal back the data to yml
		yamlData, err := yaml.Marshal(config)
		if err != nil {
			log.Fatalf("error marshalling YAML: %v", err)
		}

		// Write modified data back to YAML file
		err = os.WriteFile("config.yml", yamlData, 0644)
		if err != nil {
			log.Fatalf("error writing YAML file: %v", err)
		}

		fmt.Println("Slack Webhook added successfully")
		PromptConfigChange(config)

	//change the configuration type
	case "4":
		opt, _ = utils.GetInput("Insert the run type (fast or complete). End by pressing \"Enter\"\n", reader)

		//check if the type is setted correctly
		if opt == "fast" || opt == "complete" {
			config.RunConfig.RunType = opt
			yamlData, err := yaml.Marshal(config)
			if err != nil {
				log.Fatalf("error marshalling YAML: %v", err)
			}

			// Write modified data back to YAML file
			err = os.WriteFile("config.yml", yamlData, 0644)
			if err != nil {
				log.Fatalf("error writing YAML file: %v", err)
			}

			fmt.Println("Config type setted correctly")
			PromptConfigChange(config)
		} else {
			//restart the config if the type was invalid
			PromptConfigChange(config)
		}

	//change the number of threads
	case "5":
		opt, _ = utils.GetInput("Insert the number of threads you want to run. End by pressing \"Enter\"\n", reader)

		//check if the value inserted is a number
		num, err := strconv.Atoi(opt)
		if err != nil {
			log.Fatalf("error converting thread number: %v", err)
		}

		//set the number back in the config file
		config.RunConfig.ActiveThreads = num
		yamlData, err := yaml.Marshal(config)
		if err != nil {
			log.Fatalf("error marshalling YAML: %v", err)
		}

		// Write modified data back to YAML file
		err = os.WriteFile("config.yml", yamlData, 0644)
		if err != nil {
			log.Fatalf("error writing YAML file: %v", err)
		}

		fmt.Println("Thread number setted correctly")
		PromptConfigChange(config)

	default:
		return config
	}
	return config
}
