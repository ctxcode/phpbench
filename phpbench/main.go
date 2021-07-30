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
var _benchDir = ""
var _originalFilepath = ""
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

	///////////////////////////

	fullpath, err := filepath.Abs(fn)
	if err != nil {
		fmt.Println("Cant generate absolute file path")
		os.Exit(0)
	}

	_filepath = fullpath
	_benchDir = filepath.Dir(_filepath) + "/.phpbench"
	_originalFilepath = (fullpath[:len(fullpath)-4]) + ".original.php"

	// Check if there already is a original file
	if _, err := os.Stat(_originalFilepath); err == nil {
		fmt.Println("Original filepath already exists... something must have gone wrong. You have to fix it manually.")
		os.Exit(0)
	}

	os.Mkdir(_benchDir, 0777)

	classCode, _ := Asset("assets/PhpBench.php")
	ioutil.WriteFile(_benchDir+"/class.php", []byte(classCode), 0777)
	fmt.Println(_benchDir + "/class.php")

	// fmt.Println("File: " + fullpath)

	if _, err := os.Stat(fullpath); os.IsNotExist(err) {
		fmt.Println("File not found")
		os.Exit(0)
	}

	///////////////////////////

	// Catch CTRL+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		placeBackOriginal()
	}()

	// Copy original code
	code, err := ioutil.ReadFile(_filepath)

	if err != nil {
		fmt.Println("Cant read file")
		os.Exit(0)
	}

	originalCode = string(code)
	ioutil.WriteFile(_originalFilepath, []byte(code), 0644)

	// Set new code
	newCode := "<?php\n\n" +
		"include('" + (_benchDir) + "/class.php');\n" +
		// "include('/mnt/c/www/phpbench/assets/PhpBench.php');\n" +
		"declare (ticks = 1);\n" +
		"register_shutdown_function('\\PhpBench::send');\n" +
		"pcntl_signal(SIGINT, '\\PhpBench::send');\n" +
		"$_phpbench_result = include(phpbench_include('" + (_originalFilepath) + "'));\n" +
		// "\\PhpBench::send();\n" +
		"return $_phpbench_result;\n"
	ioutil.WriteFile(_filepath, []byte(newCode), 0644)

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

	//  Remove all .phpbench.php files
	dir, _ := os.Getwd()
	fmt.Println("Clean dir: " + dir)
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			println("Read file error...")
			println(err)
		}
		if strings.HasSuffix(path, ".phpbench.php") {
			os.Remove(path)
		}
		return nil
	})

	if err != nil {
		println("Recursive file loop error")
		println(err)
	}

	os.Remove(_originalFilepath)
	os.RemoveAll(_benchDir)

	os.Exit(0)
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
		// updateNewCode()
	}
}
