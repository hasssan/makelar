package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

const (
	httpPort = ":8080"
)

func main() {
	path := flag.String("path", "", "path to hugo site directory")
	flag.Parse()

	fmt.Println("Server run on port: ", httpPort)
	http.HandleFunc("/bwahahaha", func(w http.ResponseWriter, r *http.Request) {
		runCmd(path)
		fmt.Fprintf(w, "all command done")
	})
	if err := http.ListenAndServe(httpPort, nil); err != nil {
		fmt.Println("error cuk")
	}
}

func runCmd(path *string) {
	if err := runGitPull(path); err != nil {
		log.Fatal(err)
	}
	if err := runHugo(path); err != nil {
		log.Fatal(err)
	}
}

func runHugo(path *string) error {
	cmd := exec.Command("hugo")
	cmd.Dir = *path
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Println("hugo build")
	fmt.Println("----------")
	fmt.Printf("%s\n", out.String())
	return nil
}

func runGitPull(path *string) error {
	cmd := exec.Command("git", "pull")
	cmd.Dir = *path
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Println("--------")
	fmt.Println("git pull")
	fmt.Println("--------")
	fmt.Printf("%s\n", out.String())
	return nil
}
