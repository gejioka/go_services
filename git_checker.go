package main

import (
    "fmt"
    "os"
    "os/exec"
    "time"
    "strings"
)

func main() {
	// Create a ticker that ticks once per day at 09:00
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	// Start a loop to check for new commits once per day
	for {
		select {
		case <-ticker.C:
			// Get the current time
			// now := time.Now()

			// Check if it's 09:00
			// if now.Hour() == 9 && now.Minute() == 0 {
			// Fetch remote changes
			fetchCmd := exec.Command("git", "fetch", "origin")
			fetchCmd.Stdout = os.Stdout
			fetchCmd.Stderr = os.Stderr
			if err := fetchCmd.Run(); err != nil {
				fmt.Println("Error fetching remote changes:", err)
				continue
			}

			// Get local and remote commit hashes
			localHashCmd := exec.Command("git", "rev-parse", "HEAD")
			remoteHashCmd := exec.Command("git", "rev-parse", "origin/master")

			localHashOutput, err := localHashCmd.Output()
			if err != nil {
				fmt.Println("Error getting local commit hash:", err)
				continue
			}
			remoteHashOutput, err := remoteHashCmd.Output()
			if err != nil {
				fmt.Println("Error getting remote commit hash:", err)
				continue
			}

			localHash := strings.TrimSpace(string(localHashOutput))
			remoteHash := strings.TrimSpace(string(remoteHashOutput))

			fmt.Printf("Local Hash is: %s\nRemote Hash is: %s\n",localHash,remoteHash)

			// Check if there are new commits
			if localHash != remoteHash {
				// Pull changes if there are new commits
				pullCmd := exec.Command("git", "pull")
				pullCmd.Stdout = os.Stdout
				pullCmd.Stderr = os.Stderr
				if err := pullCmd.Run(); err != nil {
					fmt.Println("Error pulling changes:", err)
				} else {
					fmt.Println("Git pull successful")
				}
			} else {
				fmt.Println("No new commits.")
			}
            // }
        }
    }
}
