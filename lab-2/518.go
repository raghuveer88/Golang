package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var counter int = 0
var proxycounter int = 0

//PROXY BEGIN skipLine
func skipLine() {
	if counter > 0 {
		counter = counter - 1
		proxycounter = proxycounter - 1
	} else {
		counter = 0
		proxycounter = proxycounter - 1
	}
}

//PROXY END skipLine

//PROXY BEGIN splitLine
func splitLine(line string) {
	//token := strings.TrimSpace(line)
	token := line
	token1 := ""
	line1 := ""
	for i := 0; i < len(token); i++ {
		if token[i] == ' ' || token[i] == '	' {
			continue
		} else {
			token1 = token1 + string(token[i])
		}
	}
	//fmt.Println(token1)
	if len(token1) > 11 {
		if token1[:12] == "//PROXYBEGIN" {
			proxycounter = 0
			line1 = strings.TrimSpace(token[13:])
			fmt.Printf("%-28s\t", line1)
			fmt.Print("1\t\t")
		}
	}
	if len(token1) > 9 {
		if token1[:10] == "//PROXYEND" {
			fmt.Println(proxycounter - 1)
		}
	}

}

//PROXY END splitLine

//PROXY BEGIN      main
func main() {
	fmt.Println("enter your file name")
	var filename string
	fmt.Scan(&filename)
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	fmt.Println("-----------------------")
	fmt.Println("Proxy name\t\tNumber of methods\tLOC")

	for scanner.Scan() {
		line := scanner.Text()
		counter++
		proxycounter++
		if line == "" {
			skipLine()

		}

		splitLine(line)
	}

	fmt.Println("total LOC -    ", counter)
	defer file.Close()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

//PROXY END      main
