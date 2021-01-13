package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//instruction := os.Args[1]
	file_name := os.Args[3]
	if os.Args[1] == "compile" {
		if os.Args[2] == "--log" { // will generated html
			bin_source, err := ioutil.ReadFile(file_name)
			if err != nil {
				fmt.Print("Couldn't open file")
				os.Exit(1)
			}
			source := string(bin_source)
			fmt.Print(source)
		}
		if os.Args[2] == "--make" { // will put generated html in a file
			err := ioutil.WriteFile("generated-files/index.html", []byte("Hello"), 0755)
			if err != nil {
				fmt.Printf("Unable to write file: %v", err)
			}
		}
	}
}
