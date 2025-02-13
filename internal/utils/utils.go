package utils

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func GetHostnameNumber() string {
	cmd := exec.Command("cat", "/etc/hostname")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	hostname := strings.TrimSpace(string(output))

	re := regexp.MustCompile(`\d+`)
	number := re.FindString(hostname)

	if number == "" {
		fmt.Printf("число не найдено в hostname: %s", hostname)
		return ""
	}

	return number
}
