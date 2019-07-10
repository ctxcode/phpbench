package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type entry struct {
	ms int
}

type codeLine struct {
	code      string
	entries   []entry
	averageMs int
	totalMs   int
}

var codeLines = map[string]*codeLine{}
var originalCode = ""
var _filepath = ""
var _benchFilepath = ""
var _tag = ""

func main() {

	// Get file from argv
	if len(os.Args) <= 1 {
		fmt.Println("Missing filename param")
		os.Exit(0)
	}
	fn := os.Args[1]

	if !strings.HasSuffix(fn, ".php") {
		fmt.Println("File must be a .php file")
		os.Exit(0)
	}

	fullpath, err := filepath.Abs(fn)
	if err != nil {
		fmt.Println("Cant generate absolute file path")
		os.Exit(0)
	}

	benchFullpath := (fullpath[:len(fullpath)-4]) + ".bench.php"
	_filepath = fullpath
	_benchFilepath = benchFullpath
	_tag = "<?php include('" + (_benchFilepath) + "'); ?>"

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

	getOriginalCode()
	updateOriginalCode()
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
	startServer()
}

func placeBackOriginal() {

	fmt.Println("Putting back original file")

	if originalCode == "" {
		fmt.Println("ERROR ...")
		return
	}

	ioutil.WriteFile(_filepath, []byte(originalCode), 0644)
	os.Remove(_benchFilepath)
}

func updateNewCode() {

	if !strings.Contains(originalCode, "<?php") {
		fmt.Println("Cant find <?php tag in code")
		os.Exit(0)
	}

	injectCode, err := Asset("assets/inject.php")
	if err != nil {
		fmt.Println("Cant find inject code")
		os.Exit(0)
	}

	newCode := string(injectCode) + originalCode

	ioutil.WriteFile(_benchFilepath, []byte(newCode), 0644)
}

func updateOriginalCode() {
	newCode := _tag + originalCode
	ioutil.WriteFile(_filepath, []byte(newCode), 0644)
}

func getOriginalCode() {

	code, err := ioutil.ReadFile(_filepath)

	if err != nil {
		fmt.Println("Cant read file")
		os.Exit(0)
	}

	originalCode = string(code)
	originalCode = strings.Replace(originalCode, _tag, "", 1)
}

func startServer() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.POST("/data", func(c *gin.Context) {

		code := c.PostForm("code")
		ms := c.PostForm("ms")

		msInt, _ := strconv.Atoi(ms)

		if _, ok := codeLines[code]; !ok {
			codeLines[code] = &codeLine{
				code:      code,
				totalMs:   0,
				averageMs: 0,
			}
		}

		count := len(codeLines[code].entries) + 1
		totalMs := codeLines[code].totalMs + msInt

		averageMs := totalMs / count

		codeLines[code].totalMs = totalMs
		codeLines[code].averageMs = averageMs

		c.JSON(200, gin.H{
			"success": "true",
		})
	})

	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "overview.html", gin.H{
			"lines": codeLines,
		})

	})

	r.Run(":3001") // listen and serve on 0.0.0.0:8080

}
