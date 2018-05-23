package main

import (
	"io/ioutil"
	"os"

	"./vm"
)

var version string = "0.0.2 alpha"

func main() {
	// fmt.Println("GoVM 0.0.1")

	if len(os.Args) < 2 {
		println("\nMissing arguments, type \"govm -h\" to get help\n")
	} else {
		if os.Args[1] == "-h" {
			println("\ngovm command help : \n")
			println("[1] govm [filename.asm]")
			println("[2] govm -h")
			println("[3] govm -v\n")
		} else if os.Args[1] == "-v" {
			println("\nGoVM version", version, "\n")
		} else {
			data, err := ioutil.ReadFile(os.Args[1])
			vm.Check(err)
			content := string(data)
			splited := vm.SpeSplit(content)

			// for i := 0; i < len(splited); i++ {
			// 	println("[", i, "] element :", splited[i])
			// }

			vm.Run(splited)

		}
	}
}
