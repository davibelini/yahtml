package main

import (
	"fmt"
	"io/ioutil"
	"os"
	lex "github.com/davibelini/yahtml/mod/lexer"
)

func main() {
	if len(os.Args) == 4 {	
		file_name := os.Args[3]
		if os.Args[1] == "compile" {
			if os.Args[2] == "--log" { // will print generated html
				bin_source, err := ioutil.ReadFile(file_name)
				if err != nil {
					fmt.Print("Couldn't open file: ", err)
					os.Exit(1)
				}
				source := string(bin_source) // source: yaml file content

				tokens := lex.GenerateTokens(source)
				fmt.Print(tokens)
			}
			if os.Args[2] == "--make" {
				// will create a file and put the parse result inside
			}
		}
	}
}