package page

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/tkido/gostock/my"
	"github.com/tkido/gostock/my/template"
)

var (
	reDate   = regexp.MustCompile(`\(\d{4}\.\d{1,2}\)`)
	reSector = regexp.MustCompile(`(【.*?】)([^【]*)`)
	reRow    = regexp.MustCompile(`(.*?)(\d{1,3})(?:\((-?\d{1,3})\))?`)
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Row is a row of sector e.g. 講師派遣型研修66
type Row struct {
	Label       string
	Percentage  int
	ProfitRatio int
}

func newRow(s string) Row {
	sm := reRow.FindStringSubmatch(s)
	if len(sm) == 0 {
		log.Printf("invalid data as row %s", s)
		return Row{"", 0, 0}
	}
	label := sm[1]
	per, err := strconv.Atoi(sm[2])
	if err != nil {
		// log.Println(err)
	}
	var ratio int
	if len(sm) == 4 {
		ratio, err = strconv.Atoi(sm[3])
		if err != nil {
			// log.Println(err)
		}
	}
	return Row{label, per, ratio}
}

// Sector is e.g.【連結事業】講師派遣型研修66、公開講座25、他9
type Sector struct {
	ID    string
	Title string
	Rows  []Row
}

func newSector(sm []string) Sector {
	title := sm[1]
	rows := []Row{}
	ss := strings.Split(sm[2], "、")
	rest := 100

	for _, s := range ss {
		row := newRow(s)
		rest -= row.Percentage
		rows = append(rows, row)
	}
	if 0 < rest {
		label := "その他"
		rows = append(rows, Row{label, rest, 0})
	}
	id := fmt.Sprint(rand.Int())
	return Sector{id, title, rows}
}

// e.g. ['講師派遣型研修', 66],['公開講座', 25],['他', 9]
func (s Sector) graphData() string {
	ss := []string{}
	for _, r := range s.Rows {
		label := r.Label
		if p := r.ProfitRatio; p != 0 {
			label += fmt.Sprintf("(%+d)", p)
		}
		ss = append(ss, fmt.Sprintf(`['%s', %d]`, label, r.Percentage))
	}
	return strings.Join(ss, ",")
}

// Sectors is []Sector
type Sectors []Sector

func newSectors(s string) Sectors {
	sectors := Sectors{}
	sms := reSector.FindAllStringSubmatch(s, -1)
	for _, sm := range sms {
		sec := newSector(sm)
		sectors = append(sectors, sec)
	}
	return sectors
}

func (ss Sectors) graph() string {
	buf := bytes.Buffer{}
	for _, s := range ss {
		g := fmt.Sprintf(graphTmpl, s.Title, s.ID)
		buf.WriteString(g)
		buf.WriteString("\n")
	}
	return buf.String()
}

func (ss Sectors) javascript() string {
	buf := bytes.Buffer{}
	for _, s := range ss {
		js := fmt.Sprintf(jsTmpl, s.graphData(), s.ID)
		buf.WriteString(js)
		buf.WriteString("\n")
	}
	return buf.String()
}

// TODO 【連結事業】国内通信35(23)、スプリント39(5)、ヤフー9(22)、流通14(-1)、アーム1(11)、他1(-13)【海外】51(2017.3)
func addData(data map[string]string) map[string]string {
	s := data["事業"]
	data["date"] = reDate.FindString(s)
	s = reDate.ReplaceAllString(s, "")
	sectors := newSectors(s)
	data["divs"] = sectors.graph()
	data["jss"] = sectors.javascript()

	// data["rireki"] = rireki.Table(data["ID"])
	return data
}

// Write html from data
func Write(path string, data map[string]string) (err error) {
	addedData := addData(data)
	s := template.Execute(pageTmpl, addedData)
	err = my.WriteFile(path, s)
	if err != nil {
		return
	}
	return
}
