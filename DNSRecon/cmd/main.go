// Author: M-Faheem-Khan
// File: DNSRecon/cmd/main.go
// Description: Lists A, AAAA, CNAME, NameServer, TXT dns records

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Verifying domain name is given as an arg
	if len(os.Args) != 2 {
		fmt.Println("Lists A, AAAA, CNAME, NameServer, TXT dns records")
		fmt.Printf("Usage: %s domain\n", os.Args[0])
		os.Exit(0)
	}

	// Getting the domain name
	domain := os.Args[1]
	coloredOutput := fmt.Sprintf("\x1b[%dm%s%s\x1b[0m", 33, "Domain: ", domain)
	fmt.Println(coloredOutput)

	// Getting DNS Records
	getIP(domain)
	getMXRecords(domain)
	getCNAME(domain)
	getNameServers(domain)
	getTxtRecords(domain)
}

// List A/AAAA records of a given domain
func getIP(domain string) {
	ipRecords, _ := net.LookupIP(domain)
	coloredOutput := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 33, "=== A/AAAA Records ===")
	fmt.Println(coloredOutput)
	for _, ip := range ipRecords {
		fmt.Println(ip)
	}
}

// List MX records of a given domain
func getMXRecords(domain string) {
	mxRecords, _ := net.LookupMX(domain)
	coloredOutput := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 33, "=== MX Records ===")
	fmt.Println(coloredOutput)
	for _, record := range mxRecords {
		fmt.Println(record.Host, record.Pref)
	}
}

// List CNAME record of a given domain
func getCNAME(domain string) {
	cname, _ := net.LookupCNAME(domain)
	coloredOutput := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 33, "=== CNAME Record ===")
	fmt.Println(coloredOutput)
	fmt.Println(cname)
}

// List Name Server records of a given domain
func getNameServers(domain string) {
	nameservers, _ := net.LookupNS(domain)
	coloredOutput := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 33, "=== Name Server Records ===")
	fmt.Println(coloredOutput)
	for _, nameserver := range nameservers {
		fmt.Println(nameserver.Host)
	}
}

// List TXT records of a given domain
func getTxtRecords(domain string) {
	txtRecords, _ := net.LookupTXT(domain)
	coloredOutput := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 33, "=== TXT Records ===")
	fmt.Println(coloredOutput)
	for _, record := range txtRecords {
		fmt.Println(record)
	}
}

// EOF
