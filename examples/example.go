package main

import (
	"fmt"
	"log"

	command "github.com/drpaneas/commandline"
)

func main() {

	output, err := command.Run("ls -l -a -h | grep example")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)

}
