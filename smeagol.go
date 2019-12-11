package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func menu() {
	fmt.Println("1. sessions - list active sessions")
	fmt.Println("2. interact <agent> - interact with agent using given agent id")
	fmt.Println("3. listener <start,stop,configure,list> - configue and start listeners")
	fmt.Println("4. exit")
}

func usage() {
	fmt.Println("Did not understand input")
	menu()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		menu()
		fmt.Print("$>")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		if strings.Compare("exit\n", text) == 0 {
			break
		}
	}
	fmt.Println("Exiting")

}
