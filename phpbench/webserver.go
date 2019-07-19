package main

import (
	"encoding/json"
	"html/template"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

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

type setJson struct {
	SetNr   string      `json:"setNr"`
	Entries []entryJson `json:"entries"`
}

type entryJson struct {
	Filename string `json:"filename"`
	Key      string `json:"key"`
	LineNr   string `json:"lineNr"`
	Code     string `json:"code"`
	Ms       string `json:"ms"`
	Start    string `json:"start"`
	End      string `json:"end"`
}

var resultSets = map[string]*resultSet{}
var lastResultSetKey = ""

// var codeLines = map[string]*codeLine{}

func startWebServer() {

	r := gin.Default()
	// r := gin.New()

	t, _ := loadTemplate()
	r.SetHTMLTemplate(t)

	r.POST("/set/create", func(c *gin.Context) {

		jsonData := c.PostForm("data")

		set := setJson{}

		if err := json.Unmarshal([]byte(jsonData), &set); err != nil {
			panic(err)
		}

		parseNewData(&set)

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

	// r.Run(":3001")
	s := &http.Server{
		Addr:           ":3001",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}

func parseNewData(s *setJson) {

	lastResultSetKey = s.SetNr

	resultSets[s.SetNr] = &resultSet{
		CodeLines: map[string]*codeLine{},
	}

	for _, m := range s.Entries {

		msInt, _ := strconv.Atoi(m.Ms)
		startInt, _ := strconv.Atoi(m.Start)
		endInt, _ := strconv.Atoi(m.End)

		if _, ok := resultSets[s.SetNr].CodeLines[m.Key]; !ok {
			resultSets[s.SetNr].CodeLines[m.Key] = &codeLine{
				Filename:         m.Filename,
				LineNr:           m.LineNr,
				Code:             m.Code,
				Ms:               0,
				MsDisplay:        0,
				AverageMsDisplay: 0,
				Start:            startInt,
				End:              endInt,
			}
		}

		lineRef := resultSets[s.SetNr].CodeLines[m.Key]

		count := len(lineRef.Entries) + 1

		lineRef.Ms += msInt
		lineRef.MsDisplay = msDisplay(lineRef.Ms)
		lineRef.AverageMsDisplay = msDisplay(lineRef.Ms / count)

		lineRef.End = endInt

		lineRef.Entries = append(lineRef.Entries, entry{Ms: msInt})

	}
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
