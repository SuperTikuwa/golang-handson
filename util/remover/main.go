package main

import (
	"log"
	"os/exec"
)

const basePath = "./src"

func main() {

	err := exec.Command("sh", "-c", "rm `find "+basePath+" | grep go.mod`").Run()
	if err != nil {
		log.Fatal(err)
	}

}
