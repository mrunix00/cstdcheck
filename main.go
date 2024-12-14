package main

import (
	"fmt"
	"os"
)

func main() {
	ccPath := getCCPath()

	if _, err := os.Stat(ccPath); err != nil {
		fmt.Printf("[-] Error: %s\n", err.Error())
		os.Exit(1)
	}

	if !isExecutable(ccPath) {
		fmt.Printf("[-] Error: %s is not an executable!\n", ccPath)
		os.Exit(1)
	}

	fmt.Printf("[+] Provided path to CC: %s\n", ccPath)
	fmt.Println("[*] Downloading gcc torture tests...")
	err := FetchGccTorture()
	if err != nil {
		fmt.Printf("[-] Error: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("[+] Downloaded gcc torture tests")

	fmt.Println("[*] Running gcc torture tests...")
	err, results := TestGccTorture(ccPath)
	if err != nil {
		fmt.Printf("[-] Error: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("[+] Gcc torture tests results:")
	fmt.Printf("[+] %d/%d passed\n", results.passed, results.all)
	fmt.Printf("[+] Success rate: %d%%\n", (results.passed*100)/results.all)
	
	fmt.Println("[*] Running IEEE tests...")
	err, results = TestIeeeTorture(ccPath)
	if err != nil {
		fmt.Printf("[-] Error: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("[+] IEEE tests results:")
	fmt.Printf("[+] %d/%d passed\n", results.passed, results.all)
	fmt.Printf("[+] Success rate: %d%%\n", (results.passed*100)/results.all)
}
