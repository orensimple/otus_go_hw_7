package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 3 {
		log.Printf("invalid number of arguments needs 2, have: %d", len(os.Args))
		os.Exit(13)
	}
	pathEnv := os.Args[1]
	program := os.Args[2]

	file, err := ioutil.ReadDir(pathEnv)
	if err != nil {
		log.Printf("Can't read dir: %v %s", err, file)
		os.Exit(13)
	}
	cmd := exec.Command(program)
	for _, files := range file {
		f, err := ioutil.ReadFile(pathEnv + "/" + files.Name())
		if err != nil {
			log.Printf("Can't read file: %v %s", err, f)
			os.Exit(13)
		}
		envVar := files.Name() + "=" + string(f)
		cmd.Env = append(cmd.Env,
			envVar,
		)
	}

	out, err := cmd.Output()
	if err != nil {
		log.Printf("Can't write file: %v %s", err, out)
		os.Exit(13)
	}
}
