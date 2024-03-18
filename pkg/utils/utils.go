package utils

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func RemoveDuplicates(slice []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, val := range slice {
		if _, ok := seen[val]; !ok {
			seen[val] = true
			result = append(result, val)
		}
	}
	return result
}

func NotifyNewDomainsSlack(newDomains []string, slackWebhook string) {
	// Open the CSV file
	inputFile, err := os.Open("old_EasyEASM.csv")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	// Create a new CSV reader
	reader := csv.NewReader(inputFile)

	// The index of the column you want to extract
	columnToExtract := 3

	// Slice to hold the values from the specified column
	var oldDomains []string

	// Iterate through the records, extracting the value from the specified column
	for {
		record, err := reader.Read()
		if err != nil {
			if err == csv.ErrTrailingComma {
				// Skip records with trailing commas
				continue
			} else if err.Error() == "EOF" {
				// End of file
				break
			} else {
				// Some other error
				panic(err)
			}
		}

		// Append the value from the specified column if the index is within bounds
		if columnToExtract < len(record) {
			oldDomains = append(oldDomains, record[columnToExtract])
		}
	}

	NewDomainsToAlert := difference(newDomains, oldDomains)
	OldDomainsToAlert := difference(oldDomains, newDomains)

	sendToSlack(slackWebhook, fmt.Sprintf("New live domains found: %v", NewDomainsToAlert))
	sendToSlack(slackWebhook, fmt.Sprintf("Domains that were not to be now longer live: %v", OldDomainsToAlert))
}

func NotifyNewDomainsDiscord(newDomains []string, discordWebhook string) {
	// Open the CSV file
	inputFile, err := os.Open("old_EasyEASM.csv")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	// Create a new CSV reader
	reader := csv.NewReader(inputFile)

	// The index of the column you want to extract
	columnToExtract := 3

	// Slice to hold the values from the specified column
	var oldDomains []string

	// Iterate through the records, extracting the value from the specified column
	for {
		record, err := reader.Read()
		if err != nil {
			if err == csv.ErrTrailingComma {
				// Skip records with trailing commas
				continue
			} else if err.Error() == "EOF" {
				// End of file
				break
			} else {
				// Some other error
				panic(err)
			}
		}

		// Append the value from the specified column if the index is within bounds
		if columnToExtract < len(record) {
			oldDomains = append(oldDomains, record[columnToExtract])
		}
	}

	NewDomainsToAlert := difference(newDomains, oldDomains)
	OldDomainsToAlert := difference(oldDomains, newDomains)

	sendToDiscord(discordWebhook, fmt.Sprintf("New live domains found: %v", NewDomainsToAlert))
	sendToDiscord(discordWebhook, fmt.Sprintf("Domains that were not to be now longer live: %v", OldDomainsToAlert))
}

func difference(slice1, slice2 []string) []string {
	// Create a map to hold the elements of slice2 for easy lookup
	lookupMap := make(map[string]bool)
	for _, item := range slice2 {
		lookupMap[item] = true
	}

	// Iterate through slice1 and add elements that are not in slice2
	var result []string
	for _, item := range slice1 {
		if !lookupMap[item] {
			result = append(result, item)
		}
	}

	return result
}

func sendToSlack(webhookURL string, message string) {
	// Create JSON payload
	payload := map[string]string{
		"text": message,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error creating JSON:", err)
		return
	}

	// Send HTTP POST request
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error sending to Slack:", err)
		return
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error response from Slack:", resp.Status)
	}
}

func sendToDiscord(webhookURL string, message string) {
	// Create JSON payload
	payload := map[string]string{
		"content": message,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error creating JSON:", err)
		return
	}

	// Send HTTP POST request
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error sending to Discord:", err)
		return
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusNoContent {
		fmt.Println("Error response from Discord:", resp.Status)
	}
}

func InstallTools() {
	for name, path := range map[string]string{
		"alterx":    "github.com/projectdiscovery/alterx/cmd/alterx@latest",
		"amass":     "github.com/owasp-amass/amass/v3/...@master",
		"dnsx":      "github.com/projectdiscovery/dnsx/cmd/dnsx@latest",
		"httpx":     "github.com/projectdiscovery/httpx/cmd/httpx@latest",
		"oam_subs":  "github.com/owasp-amass/oam-tools/cmd/oam_subs@master",
		"subfinder": "github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest",
		"nuclei":    "github.com/projectdiscovery/nuclei/v3/cmd/nuclei@latest",
	} {
		if !checkTool(name) {
			installGoTool(name, path)
		}
	}

	fmt.Println("All needed tools installed!")
}

func checkTool(name string) bool {
	_, err := exec.LookPath(name)
	if err != nil {
		fmt.Printf("%s is not installed\n", name)
		return false
	}

	return true
}

func installGoTool(name string, path string) {
	// Replace this with the package you want to install
	packagePath := path

	cmd := exec.Command("go", "install", packagePath)
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("An error occurred while installing the package: %s\n%s", err, cmdOutput)
	}

	log.Printf("Successfully installed the package: %s", packagePath)
}

func GetInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input), err
}

func CheckJq() {
	//check if jq is installed, if not, abort the scan
	cmd := exec.Command("jq", "--version")
	err := cmd.Run()
	if err != nil {
		print("Jq is not installed, nuclei scan can't be run.\n\n")
		panic(err)
	} else {
		return
	}
}
