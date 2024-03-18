package nuclei

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/g0ldencybersec/EasyEASM/pkg/utils"
)

func RunNuclei(domains []string, flags string) {
	//run the nuclei tool
	writeTempFile(domains)

	//run the interactive mode if flag is provided
	if flags == "interactive" {
		reader := bufio.NewReader(os.Stdin)
		std, _ := utils.GetInput("Press y if you want to insert template directory, or any other character to run standard\n", reader)
		switch std {
		case "y":
			reader = bufio.NewReader(os.Stdin)
			opt, _ := utils.GetInput("Insert the template directory or press enter to run standard list\n", reader)

			if _, err := os.Stat(opt); os.IsNotExist(err) {
				fmt.Println("DIRECTORY DOES NOT EXISTS -> Running standard config...")
				//run the nuclei std command
				cmd := exec.Command("nuclei", "-l", "tempNuclei.txt", "-silent", "-o", "temp.json", "-j", "-exclude-severity", "info", "-exclude-severity", "unknown")
				err := cmd.Run()
				if err != nil {
					panic(err)
				}
			} else {
				//run the nuclei cmd on the selected list of templetes
				fmt.Println("Running found template lists...")
				cmd := exec.Command("nuclei", "-l", "tempNuclei.txt", "-silent", "-t", opt, "-o", "temp.json", "-j", "-exclude-severity", "info", "-exclude-severity", "unknown")
				err := cmd.Run()
				if err != nil {
					panic(err)
				}
			}

		default:
			//run the nuclei std command
			cmd := exec.Command("nuclei", "-l", "tempNuclei.txt", "-silent", "-o", "temp.json", "-j", "-exclude-severity", "info", "-exclude-severity", "unknown")
			err := cmd.Run()
			if err != nil {
				panic(err)
			}
		}
	} else {
		//run the standard scan with nuclei, on the provided targets
		cmd := exec.Command("nuclei", "-l", "tempNuclei.txt", "-silent", "-o", "temp.json", "-j", "-exclude-severity", "info", "-exclude-severity", "unknown")
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
	processJson()
	fmt.Printf("Nuclei scan completed!\n")
	os.Remove("temp.json")
	os.Remove("tempNuclei.txt")
}

func writeTempFile(list []string) {
	//create the temporary file and read the domains provided
	file, err := os.OpenFile("tempNuclei.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	datawriter := bufio.NewWriter(file)

	for _, data := range list {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	file.Close()
}

func processJson() {
	// Check if jq is installed
	utils.CheckJq()

	//process the json output to make lines unique using jq
	cmd := exec.Command("jq", "-c", "--unbuffered", "-r", "del(.timestamp) | del(.\"curl-command\") | @json", "temp.json")

	// Create a file to hold the output
	outputFile, err := os.Create("EasyEASM.json")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Set the output of the command to the created file
	cmd.Stdout = outputFile

	// Run the command
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
