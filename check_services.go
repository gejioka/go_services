package main

import (
    "fmt"
    "os/exec"
    "strings"
)

func isServiceRunning(serviceName string) (bool, error) {
    cmd := exec.Command("systemctl", "is-active", serviceName)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return false, err
    }
    isActive := strings.TrimSpace(string(output))
    return isActive == "active", nil
}

func main() {
    serviceName := "gateway" // TODO: Add a list with all services
    running, err := isServiceRunning(serviceName)
    if err != nil {
        fmt.Printf("Error checking service status: %v\n", err)
        return
    }
    if running {
        fmt.Printf("Service %s is running.\n", serviceName)
    } else {
        fmt.Printf("Service %s is not running.\n", serviceName)
    }
}
