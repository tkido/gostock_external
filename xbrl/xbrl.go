package xbrl

import (
	"encoding/xml"
	"errors"
	"io"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// XBRL instance
type XBRL struct {
	XMLName  xml.Name
	Contexts []Context `xml:"context"`
	Facts    []Fact    `xml:",any"`
}

// Context represents <ixbrl:context> tag
type Context struct {
	XMLName xml.Name
	ID      string `xml:"id,attr"`
	Instant string `xml:"period>instant"`
	Start   string `xml:"period>startDate"`
	End     string `xml:"period>endDate"`
}

// Fact represents each fact
type Fact struct {
	XMLName    xml.Name
	Name       string
	Value      string
	ContextRef string
	UnitRef    string
	Decimals   string
	Scale      string
	Sign       string
	Nil        bool
}

var reTag = regexp.MustCompile(`<.*?>`)

// UnmarshalXML implements xml.Unmarshaler interface
func (f *Fact) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v struct {
		XMLName    xml.Name
		Name       string `xml:"name,attr"`
		Value      string `xml:",innerxml"`
		ContextRef string `xml:"contextRef,attr"`
		UnitRef    string `xml:"unitRef,attr"`
		Decimals   string `xml:"decimals,attr"`
		Scale      string `xml:"scale,attr"`
		Sign       string `xml:"sign,attr"`
		Nil        string `xml:"nil,attr"`
	}
	if err := d.DecodeElement(&v, &start); err != nil {
		return err
	}

	// parse xsi:nil
	isNil := v.Nil == "true"

	// use name attribute if exists, or use XML tag's local name
	var name string
	if v.Name != "" {
		name = trimNameSpace(v.Name)
	} else {
		name = v.XMLName.Local
	}
	*f = Fact{v.XMLName, name, v.Value, v.ContextRef, v.UnitRef, v.Decimals, v.Scale, v.Sign, isNil}
	return nil
}

// Text returns text
func (f *Fact) Text() string {
	return reTag.ReplaceAllString(f.Value, "")
}

// Float returns float
func (f *Fact) Float() (fl float64, err error) {
	var sign float64
	switch f.Sign {
	case "":
		sign = 1.0
	case "-":
		sign = -1.0
	default:
		return 0.0, errors.New("factToFloat: unknown sign")
	}
	s := strings.Replace(f.Value, ",", "", -1)
	fl, err = strconv.ParseFloat(s, 64)
	if err != nil {
		return
	}
	scale := 0
	if f.Scale != "" {
		scale, err = strconv.Atoi(f.Scale)
		if err != nil {
			return
		}
	}
	fl = sign * fl * math.Pow10(scale)
	return
}

func trimNameSpace(s string) string {
	i := strings.IndexRune(s, ':')
	if i == -1 {
		return s
	}
	return s[i+1:]
}

// UnmarshalXBRL unmarshal io.Reader into *XBRL or return error
func UnmarshalXBRL(xbrl *XBRL, reader io.Reader) error {
	decoder := xml.NewDecoder(reader)
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			// Ingore HTML tags
			if se.Name.Space == "http://www.w3.org/1999/xhtml" {
				continue
			}
			switch n := se.Name.Local; n {
			case "xbrl": // Vanila XBRL
				if err := decoder.DecodeElement(&xbrl, &se); err != nil {
					return err
				}
				return nil
			// inlineXBRL
			case "hidden", "resources", "references", "unit", "header":
				continue
			case "context":
				var ctx Context
				if err := decoder.DecodeElement(&ctx, &se); err != nil {
					return err
				}
				xbrl.Contexts = append(xbrl.Contexts, ctx)
			default:
				var fact Fact
				if err := decoder.DecodeElement(&fact, &se); err != nil {
					return err
				}
				xbrl.Facts = append(xbrl.Facts, fact)
			}
		}
	}
	return nil
}
