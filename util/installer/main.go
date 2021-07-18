package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
)

const steps = 9
const basePath = "./src/step"

func main() {
	name := ""

	fmt.Println("GitHubのユーザーネームを入力してください")
	fmt.Scanln(&name)

	if name == "" {
		log.Fatal("名前を入れてください")
	}

	for i := 0; i < steps; i++ {

		err := exec.Command("sh", "-c", "cd "+basePath+strconv.Itoa(i+1)+" && "+"go mod init github.com/"+name+"/go-study/"+"step"+strconv.Itoa(i+1)+" && "+"go get github.com/SuperTikuwa/testutil").Run()
		if err != nil {
			log.Fatal(err)
		}

	}

}
