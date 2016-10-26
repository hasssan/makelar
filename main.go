/*
Suja
Config:
	- path to hugo site directory
	- path to hugo executable
	- path to hugo output directory
	- path for url

Start:

	1. get all path
	2. wait request incoming

Request Incoming:

	1. delete output directory content
	2. runHugo and runGitPull
	3. wait for next request
*/
package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/kelseyhightower/envconfig"
)

type specs struct {
	Port      string
	HugoBin   string `envconfig:"hugo_bin"`
	HugoSite  string `envconfig:"hugo_site"`
	OutputDir string `envconfig:"output_dir"`
	URL       string `default:"/webhook"`
}

var s specs

func main() {
	if err := envconfig.Process("mlr", &s); err != nil {
		log.Fatal(err.Error())
	}

	path := s.HugoSite
	urlPath := s.URL

	port := ":" + s.Port

	fmt.Println("hugo path", s.HugoBin)

	fmt.Println("Server run on port:", s.Port)
	runCmd(path)
	fmt.Println("waiting incoming request...")

	http.HandleFunc(urlPath, func(w http.ResponseWriter, r *http.Request) {
		runCmd(path)
		fmt.Fprintf(w, "all command done")
	})
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("error cuk")
	}
}

func runCmd(path string) {
	if err := runGitPull(path); err != nil {
		log.Fatal(err)
	}
	if err := runHugo(path); err != nil {
		log.Fatal(err)
	}
}

func runHugo(path string) error {
	cmd := exec.Command(s.HugoBin)
	cmd.Dir = path
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

func runGitPull(path string) error {
	cmd := exec.Command("git", "pull")
	cmd.Dir = path
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return err
	}
	fmt.Println("git pull")
	fmt.Println("--------")
	fmt.Printf("%s\n", out.String())
	return nil
}
