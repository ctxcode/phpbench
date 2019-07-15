package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
)

type entry struct {
	Ms int
}

type codeLine struct {
	Code             string
	Entries          []entry
	AverageMs        int
	AverageMsDisplay float64
	TotalMs          int
	TotalMsDisplay   float64
	LastMs           int
	LastMsDisplay    float64
}

var codeLines = map[string]*codeLine{}
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

func addBenchFunctions(code string) string {

	newCode := ""
	lastWord := ""
	inString := false
	stringEndChar := ""
	timerCount := 0
	prevChar := ""
	lastToken := ""
	tokens := []string{}
	allowNewCode := true
	allowNewCodeAfterNextBracket := false
	allowNewCodeAfterNextSemiColon := false
	lastCodeLine := ""
	inCommentSingleLine := false
	inCommentMultiLine := false

	for _, char := range code {

		c := string(char)

		lastCodeLine += c

		if inString {
			if c == stringEndChar && prevChar != "\\" {
				inString = false
				stringEndChar = ""
			}
			prevChar = c
			continue
		}

		if inCommentMultiLine {
			if c == "/" && prevChar == "*" {
				inCommentMultiLine = false
			}
			prevChar = c
			continue
		}

		if inCommentSingleLine {
			if c == "\n" {
				inCommentSingleLine = false
			}
			prevChar = c
			continue
		}

		if unicode.IsLetter(char) {
			lastWord += c
		} else {

			if lastWord == "php" {
				newCode += lastCodeLine
				lastCodeLine = ""
			}

			if lastWord == "namespace" {
				allowNewCode = false
				allowNewCodeAfterNextSemiColon = true
			}

			if lastWord == "class" || lastWord == "function" || lastWord == "if" || lastWord == "else" || lastWord == "elseif" || lastWord == "for" {
				lastToken = lastWord
				allowNewCode = false
				if lastWord != "class" {
					allowNewCodeAfterNextBracket = true
				}
			}

			if c == "*" && prevChar == "/" {
				inCommentMultiLine = true
			}

			if c == "/" && prevChar == "/" {
				inCommentSingleLine = true
			}

			if c == "{" {
				tokens = append(tokens, lastToken)
				if allowNewCodeAfterNextBracket {
					allowNewCodeAfterNextBracket = false
					allowNewCode = true
				}
				newCode += lastCodeLine
				lastCodeLine = ""
			}

			if c == "}" {
				allowNewCode = true
				tokens = tokens[:len(tokens)-1]
				if len(tokens) > 0 {
					token := tokens[len(tokens)-1]
					if token != "class" {
						allowNewCode = false
					}
				}
				newCode += lastCodeLine
				lastCodeLine = ""
			}

			if c == "'" || c == "\"" {
				inString = true
				stringEndChar = c
			}

			if c == ";" {
				if allowNewCode {
					timerCount++
					newCode += "\\PhpBench::startTimer(" + (strconv.Itoa(timerCount)) + ");\n"
				}
				newCode += lastCodeLine + "\n"
				if allowNewCode {
					escCodeLine := strings.Replace(lastCodeLine, "'", "", -1)
					re := regexp.MustCompile(`\r?\n`)
					escCodeLine = re.ReplaceAllString(escCodeLine, " ")
					escCodeLine = strings.TrimSpace(escCodeLine)
					newCode += "\\PhpBench::timeCode(" + (strconv.Itoa(timerCount)) + ", trim('" + escCodeLine + "'));\n"
				}
				if allowNewCodeAfterNextSemiColon {
					allowNewCodeAfterNextSemiColon = false
					allowNewCode = true
				}
				lastCodeLine = ""
			}

			// Reset lastWord
			lastWord = ""
		}

		prevChar = c
	}

	newCode += lastCodeLine

	return newCode
}

func startServer() {

	r := gin.Default()

	t, _ := loadTemplate()
	r.SetHTMLTemplate(t)

	r.POST("/data", func(c *gin.Context) {

		code := c.PostForm("code")
		ms := c.PostForm("ms")

		fmt.Println(code)
		fmt.Println(ms)

		msInt, _ := strconv.Atoi(ms)

		if _, ok := codeLines[code]; !ok {
			codeLines[code] = &codeLine{
				Code:      code,
				LastMs:    0,
				TotalMs:   0,
				AverageMs: 0,
			}
		}

		count := len(codeLines[code].Entries) + 1
		totalMs := codeLines[code].TotalMs + msInt

		averageMs := totalMs / count

		codeLines[code].LastMs = msInt
		codeLines[code].LastMsDisplay = msDisplay(msInt)
		codeLines[code].TotalMs = totalMs
		codeLines[code].TotalMsDisplay = msDisplay(totalMs)
		codeLines[code].AverageMs = averageMs
		codeLines[code].AverageMsDisplay = msDisplay(averageMs)

		codeLines[code].Entries = append(codeLines[code].Entries, entry{Ms: msInt})

		c.JSON(200, gin.H{
			"success": "true",
		})
	})

	r.GET("/", func(c *gin.Context) {

		lines := []*codeLine{}
		total := 0
		for _, val := range codeLines {
			lines = append(lines, val)
			total += val.LastMs
		}

		sort.Slice(lines[:], func(i, j int) bool {
			return lines[i].AverageMs > lines[j].AverageMs
		})

		c.HTML(http.StatusOK, "templates/overview.html", gin.H{
			"Lines": lines,
			"Total": msDisplay(total),
		})

	})

	r.Run(":3001")

}

func msDisplay(nr int) float64 {

	res := math.Round(float64(nr)/10) / 100

	return res
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for _, name := range AssetNames() {
		if !strings.HasSuffix(name, ".html") {
			continue
		}
		html, _ := Asset(name)
		t, _ = t.New(name).Parse(string(html))
	}
	return t, nil
}
