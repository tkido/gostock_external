package my

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func decorateTd(s string) string {
	if strings.HasPrefix(s, "-") {
		return fmt.Sprintf(`<span class="minus">%s</span>`, s)
	}
	return s
}
func toString(x interface{}) string {
	if f, ok := x.(float64); ok {
		return Round(f)
	}
	return fmt.Sprint(x)
}

var scales = [7]string{"", "K", "M", "B", "T", "Q", "q"}

// Round is Round
func Round(f float64) string {
	buf := bytes.Buffer{}
	if f < 0 {
		buf.WriteString("-") // sign
	}
	s := strconv.FormatFloat(math.Abs(f), 'f', 0, 64)
	size := len(s)
	if size < 3 {
		s += "00" // defend out of bounds
	}
	mod := size % 3
	buf.WriteString(s[:mod]) // head
	if mod != 0 {
		buf.WriteString(".") // separator
	}
	buf.WriteString(s[mod:3])           // tail
	buf.WriteString(scales[(size-1)/3]) // scale
	return buf.String()
}

// Table is Table
type Table struct {
	buf bytes.Buffer
}

// NewTable is NewTable
func NewTable(title string, classes ...string) *Table {
	t := &Table{bytes.Buffer{}}
	t.Write("<h3>")
	t.Write(title)
	t.Write("</h3>\n")

	var cs string
	if len(classes) != 0 {
		cs = fmt.Sprintf(` class="%s"`, strings.Join(classes, ","))
	}
	t.Write(fmt.Sprintf(`<table border="2"%s><tbody>`, cs))
	t.Write("\n")
	return t
}

func (t *Table) tr(start, end string, tds []interface{}) {
	t.Write("<tr>")
	for _, td := range tds {
		t.Write(start)
		t.Write(decorateTd(toString(td)))
		t.Write(end)
	}
	t.Write("</tr>\n")
}

// Th adds <tr><th>foo</th><th>bar</th>...</tr>
func (t *Table) Th(ths ...interface{}) {
	t.tr("<th>", "</th>", ths)
}

// Td adds <tr><td>foo</td><td>bar</td>...</tr>
func (t *Table) Td(tds ...interface{}) {
	t.tr("<td>", "</td>", tds)
}

func (t *Table) String() string {
	t.Write("</tbody></table>\n")
	return t.buf.String()
}

func (t *Table) Write(s string) {
	t.buf.WriteString(s)
}
