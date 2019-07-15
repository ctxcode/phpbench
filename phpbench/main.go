package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"
)

var originalCode = ""
var _filepath = ""
var _originalFilepath = ""
var _benchFilepath = ""
var _tag = ""
var _version = "0.1.2"

func main() {

	// Get file from argv
	if len(os.Args) <= 1 {
		fmt.Println("Missing filename param")
		os.Exit(0)
	}
	fn := os.Args[1]

	if fn == "version" {
		fmt.Println("Version " + _version)
		os.Exit(0)
	}

	if !strings.HasSuffix(fn, ".php") {
		fmt.Println("File must be a .php file")
		os.Exit(0)
	}

	fullpath, err := filepath.Abs(fn)
	if err != nil {
		fmt.Println("Cant generate absolute file path")
		os.Exit(0)
	}

	_benchFilepath = (fullpath[:len(fullpath)-4]) + ".phpbench.php"
	_originalFilepath = (fullpath[:len(fullpath)-4]) + ".original.php"
	_filepath = fullpath
	_tag = "return include('" + (_benchFilepath) + "');"

	injectCode, err := Asset("assets/inject.php")
	if err != nil {
		fmt.Println("Cant find inject code")
		os.Exit(0)
	}

	fmt.Println("File: " + fullpath)

	if _, err := os.Stat(fullpath); os.IsNotExist(err) {
		fmt.Println("File not found")
		os.Exit(0)
	}

	// Catch CTRL+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		placeBackOriginal()
		os.Exit(0)
	}()

	// Copy original code
	code, err := ioutil.ReadFile(_filepath)

	if err != nil {
		fmt.Println("Cant read file")
		os.Exit(0)
	}

	originalCode = string(code)
	ioutil.WriteFile(_originalFilepath, []byte(code), 0644)

	// Insert tag
	newCode := string(injectCode) + "\n" + _tag
	ioutil.WriteFile(_filepath, []byte(newCode), 0644)

	//
	updateNewCode()

	// Interval
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				getOriginalCode()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	// Web server
	startWebServer()
}

func placeBackOriginal() {

	fmt.Println("Putting back original file")

	if originalCode == "" {
		fmt.Println("ERROR ...")
		return
	}

	ioutil.WriteFile(_filepath, []byte(originalCode), 0644)
	os.Remove(_benchFilepath)
	os.Remove(_originalFilepath)
}

func updateNewCode() {

	if !strings.Contains(originalCode, "<?php") {
		fmt.Println("Cant find <?php tag in code")
		os.Exit(0)
	}

	newCode := addBenchFunctions(originalCode)

	ioutil.WriteFile(_benchFilepath, []byte(newCode), 0644)
}

func getOriginalCode() {

	code, err := ioutil.ReadFile(_originalFilepath)

	if err != nil {
		fmt.Println("Cant read file")
		os.Exit(0)
	}

	newCode := string(code)

	if newCode != originalCode {
		originalCode = newCode
		updateNewCode()
	}
}
