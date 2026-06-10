package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func RunRecon(input string) {
	cmd := exec.Command("subfinder", "-nc", "-silent", "-d", input, "-all", "&&", "findomain", "-t", input)

	cmd.Stdout = os.Stdout
	cmd.Stdout = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Subdomain Recon Running...")
}
