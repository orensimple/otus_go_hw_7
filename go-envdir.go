package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	pathEnv := os.Args[1]
	program := os.Args[2]
	fmt.Println(pathEnv)
	fmt.Println(program)

	file, err := ioutil.ReadDir(pathEnv)
	if err != nil {
		fmt.Println(err)
	}
	cmd := exec.Command(program)
	for _, files := range file {
		file, err := os.Open(pathEnv + "/" + files.Name())
		if err != nil {
			log.Printf("Can't create/open file: %v %s", err, file)
		}
		f, err := ioutil.ReadAll(file)
		if err != nil {
			log.Printf("Can't read file: %v %s", err, f)
		}
		envVar := files.Name() + "=" + string(f)
		cmd.Env = append(cmd.Env,
			envVar,
		)
		fmt.Println(envVar)
	}

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ENV: %s\n", out)
}
