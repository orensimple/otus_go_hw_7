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

	files, err := ioutil.ReadDir(pathEnv)
	if err != nil {
		log.Printf("Can't read dir: %v %s", err, files)
		os.Exit(13)
	}
	cmd := exec.Command(program)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		f, err := ioutil.ReadFile(pathEnv + "/" + file.Name())
		if err != nil {
			log.Printf("Can't read file: %v %s", err, f)
			continue
		}
		envVar := file.Name() + "=" + string(f)
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
