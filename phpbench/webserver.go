package main

import (
	"encoding/json"
	"html/template"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type resultSet struct {
	CodeLines       map[string]*codeLine
	CodeLinesSorted []*codeLine
	TotalMs         int
}

type entry struct {
	Ms int
}

type codeLine struct {
	Filename string
	Code     string
	LineNr   string
	Entries  []entry
	Start    int
	End      int

	Ms               int
	MsDisplay        float64
	AverageMsDisplay float64
}

var resultSets = map[string]*resultSet{}
var lastResultSetKey = ""

// var codeLines = map[string]*codeLine{}

func startWebServer() {

	r := gin.Default()

	t, _ := loadTemplate()
	r.SetHTMLTemplate(t)

	r.POST("/data", func(c *gin.Context) {

		setNr := c.PostForm("setNr")
		filename := c.PostForm("filename")

		key := c.PostForm("key")
		lineNr := c.PostForm("lineNr")
		code := c.PostForm("code")
		ms := c.PostForm("ms")

		start := c.PostForm("start")
		end := c.PostForm("end")

		lastResultSetKey = setNr

		msInt, _ := strconv.Atoi(ms)
		startInt, _ := strconv.Atoi(start)
		endInt, _ := strconv.Atoi(end)

		if _, ok := resultSets[setNr]; !ok {
			resultSets[setNr] = &resultSet{
				CodeLines: map[string]*codeLine{},
			}
		}

		if _, ok := resultSets[setNr].CodeLines[key]; !ok {
			resultSets[setNr].CodeLines[key] = &codeLine{
				Filename:         filename,
				LineNr:           lineNr,
				Code:             code,
				Ms:               0,
				MsDisplay:        0,
				AverageMsDisplay: 0,
				Start:            startInt,
				End:              endInt,
			}
		}

		lineRef := resultSets[setNr].CodeLines[key]

		count := len(lineRef.Entries) + 1

		lineRef.Ms += msInt
		lineRef.MsDisplay = msDisplay(lineRef.Ms)
		lineRef.AverageMsDisplay = msDisplay(lineRef.Ms / count)

		lineRef.End = endInt

		lineRef.Entries = append(lineRef.Entries, entry{Ms: msInt})

		c.JSON(200, gin.H{
			"success": "true",
		})
	})

	r.GET("/", func(c *gin.Context) {

		for _, rset := range resultSets {
			lines := []*codeLine{}
			total := 0

			for _, val := range rset.CodeLines {

				lines = append(lines, val)
				total += val.Ms
			}

			sort.Slice(lines[:], func(i, j int) bool {
				return lines[i].Ms > lines[j].Ms
			})

			rset.CodeLinesSorted = lines
			rset.TotalMs = total
		}

		resultSetsJson, err := json.Marshal(resultSets)

		if err != nil {
			println(err)
		}

		c.HTML(http.StatusOK, "templates/overview.html", gin.H{
			"ResultSetsJson": template.HTML(resultSetsJson),
		})

	})

	r.Run(":3001")

}

func msDisplay(nr int) float64 {

	res := math.Round(float64(nr)/10) / 100

	return res
}

func rawContent(str []byte) template.HTML {
	return template.HTML(str)
}

var templateFunctions = template.FuncMap{
	"raw": rawContent,
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for _, name := range AssetNames() {
		if !strings.HasSuffix(name, ".html") {
			continue
		}
		html, _ := Asset(name)
		t, _ = t.New(name).Funcs(templateFunctions).Parse(string(html))
	}
	return t, nil
}
