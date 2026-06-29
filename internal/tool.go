package internal

import (
	"fmt"
	"os"
	"os/exec"
)

func RunRecon(input string) {
	if input == "" {
		fmt.Println("Missing Input :(")
		return
	}

	var outputFileSubdo string = fmt.Sprintf("recon_subdo_%s.txt", input)
	var outputFilePorts string = fmt.Sprintf("recon_ports_%s.txt", input)
	var outptFileCrawling string = fmt.Sprintf("recon_crawl_%s.txt", input)
	var outputFileUro string = fmt.Sprintf("recon_uro_%s.txt", input)

	var inSub string = input
	var inNaa string = input
	var inKat string = input
	var inScp string = input

	var cmdShell string = fmt.Sprintf("subfinder -nc -silent -d %s -all | dnsx -nc -silent | httpx -silent -td -sc -fr -title -o %s && subfinder -d %s -all -silent | dnsx -silent | naabu -tp 1000 -silent -o %s && katana -u %s -kf all -jc -silent -cs %s -o %s | uro -o %s", inSub, outputFileSubdo, inNaa, outputFilePorts, inKat, inScp, outptFileCrawling, outputFileUro)

	cmd := exec.Command("cmd", "/C", cmdShell)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
		return
	}
}
