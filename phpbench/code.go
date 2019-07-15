package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

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
	lineNr := 0

	for _, char := range code {

		c := string(char)

		if c == "\n" {
			lineNr++
		}

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
					newCode += "\\PhpBench::timeCode(" + (strconv.Itoa(timerCount)) + ", __FILE__, " + strconv.Itoa(lineNr) + ", trim('" + escCodeLine + "'));\n"
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
