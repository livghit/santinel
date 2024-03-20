package watchers

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"regexp"
	"sync"
)

type NetworkMetric struct {
	device string
	ip     string
}

var (
	commandOutput []byte
	outputMutex   sync.Mutex
)

func executeCommand() {
	cmd := exec.Command("nmap", "-sL", "192.168.178.0-255")
	output, err := cmd.Output()
	if err != nil {
		log.Printf("Error executing command: %v", err)
		return
	}

	outputMutex.Lock()
	defer outputMutex.Unlock()
	commandOutput = output
}

func NetworkHandler(w http.ResponseWriter, r *http.Request) {
	outputMutex.Lock()
	defer outputMutex.Unlock()

	networkMetrics := []NetworkMetric{}

	// If commandOutput is empty, start a Goroutine to execute the command
	if len(commandOutput) == 0 {
		go executeCommand()
		fmt.Fprint(w, "Executing command, please wait...")
		return
	}

	commandOutputString := string(commandOutput)

	// Find matches using regular expression
	nmapRegex := regexp.MustCompile(`Nmap scan report for\s+([\w.-]+)(?:\.fritz\.box)?\s+\((\d+\.\d+\.\d+\.\d+)\)`)
	matches := nmapRegex.FindAllStringSubmatch(commandOutputString, -1)

	// Extracting hostname and IP address
	for _, match := range matches {
		hostname := match[1]
		ip := match[2]
		metric := NetworkMetric{device: hostname, ip: ip}
		networkMetrics = append(networkMetrics, metric)
		fmt.Fprintf(w, "<p>%s is online with IP Address: %s </p>", hostname, ip)
	}
}
