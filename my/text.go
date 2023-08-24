package my

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

// TruncateWidth truncate string with width
func TruncateWidth(s string, max int) string {
	var w int
	for p, r := range s {
		w += RuneWidth(r)
		if max < w {
			return s[0:p]
		}
	}
	return s
}

// RuneWidth is width of rune 半角 == 1, 全角 == 2
func RuneWidth(r rune) int {
	switch utf8.RuneLen(r) {
	case 1, 2:
		return 1
	case 3, 4:
		return 2
	}
	return 0
}

// Width of string 半角 == 1, 全角 == 2
func Width(s string) (n int) {
	for _, r := range s {
		n += RuneWidth(r)
	}
	return
}

// Readlines read text file and return lines
func Readlines(path string) (ss []string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		ss = append(ss, s.Text())
	}
	if err = s.Err(); err != nil {
		return
	}
	return
}

// ReadlinesMatched read text file and return lines
func ReadlinesMatched(path string, re *regexp.Regexp) ([]string, error) {
	return ReadlinesFiltered(path, func(line string) bool {
		return re.MatchString(line)
	})
}

// ReadlinesFiltered read text file and return lines
func ReadlinesFiltered(path string, filter func(string) bool) (ss []string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		if filter(line) {
			ss = append(ss, line)
		}
	}
	if err = s.Err(); err != nil {
		return
	}
	return ss, nil
}

// WriteFile is WriteFile
func WriteFile(path string, v interface{}) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return
	}
	defer f.Close()
	f.WriteString(sprint(v))
	return
}

// WriteFileForCopyPaste is WriteFileForCopyPaste
func WriteFileForCopyPaste(path string, v interface{}) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return
	}
	defer f.Close()
	f.WriteString(cpHead)
	f.WriteString(sprint(v))
	f.WriteString(cpFoot)
	return
}

func sprint(v interface{}) string {
	switch v := v.(type) {
	case []string:
		return strings.Join(v, "\n")
	default:
		return fmt.Sprint(v)
	}
}

const cpHead = `<html lang="jp">
<head>
  <meta charset="utf-8">
  <meta name="robots" content="noindex,nofollow">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>CopyPaste</title>
</head>
<body>
<p>
<button>Copy</button>
</p>
<p>
<textarea readonly style="width: 100%; height: 100%">`

const cpFoot = `</textarea>
</p>
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.4.1/jquery.min.js"></script>
<script>
window.addEventListener('load', function(){
  $('button').on('click',function(){
    $('textarea').select();
    document.execCommand('copy');
  });
});
</script>
</body>
</html>`
