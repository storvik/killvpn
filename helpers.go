package main

import (
	"fmt"
	"os/exec"
)

func executeUFW(cmd ...string) {
	if verbose {
		fmt.Printf("Running UFW: ")
		for _, i := range cmd {
			fmt.Printf("%s ", i)
		}
		fmt.Printf("\n")
	}
	out, err := exec.Command("/usr/sbin/ufw", cmd...).Output()
	if verbose {
		fmt.Printf("Output: %s", out)
	}
	if err != nil {
		panic(err)
	}

}

func killApp(app string) {
	if verbose {
		fmt.Printf("Killing app: %s\n", app)
	}
	_, _ = exec.Command("killall", app).Output()
}

func printInfo(format string, a ...interface{}) {
	if verbose {
		fmt.Printf(format, a...)
	}
}
