package main

import (
	"bufio"
	"fmt"
	"os"
	"password_validator/loader"
)

func readStdin() {
	// load up the common password text
	p, err := loader.LoadCommon(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// on scanning through stdin, pass it to the validator
		res, err := p.IsValid(scanner.Bytes())
		if err != nil {
			// print out any errors as they occur and continue.
			fmt.Printf("%s -> %s \n", res, err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: password_validator [weak_password_list.txt]")
	} else {
		readStdin()
	}
}
