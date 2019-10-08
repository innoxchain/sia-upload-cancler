package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	cmdName := "/var/sia/siac"
	cmdArgs := []string{"renter", "uploads"}

	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(cmdReader)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for scanner.Scan() {			
			line := scanner.Text()
			if line == "No files are uploading." {
				fmt.Println(line)
				break
			}
			// Skip first line of output which might look like "Uploading 98 files:"
			if !strings.Contains(line, " files:") {
				s := strings.Split(line[15:], " ")				
				fmt.Printf("Deleting %s\n", s[0])
				delCmdName := "/var/sia/siac"
				delCmdArgs := []string{"renter", "delete", s[0]}
				delCmd := exec.Command(delCmdName, delCmdArgs...)
				err := delCmd.Start()

				if err != nil {
					log.Fatalf("Error deleting file in Sia: %s\n", err)
				}

				err = delCmd.Wait()
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error waiting for delCmd", err)
					os.Exit(1)
				}
			}
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}

	// Wait for the scanner to finish processing the ouput from the "siac renter uploads command"
	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		os.Exit(1)
	}
}
