package lang

import(
	"fmt"
	"strings"
	"regexp"
)
func paramParser(params string) {
	pArr := strings.Fields(params)
	Type := pArr[0]
	// value := pArr[1]
	if (Type == "%s") {
		fmt.Println("hi")
	}
}
func Do(prog []string) {
	for i := 0; i < len(prog); i++ {
		line := prog[i]
		chars := strings.Split(line, "")
		stat := 0;
		var matched string
		for c := 0; c < len(chars); c++ {
			char := chars[c]
			if char == "(" {
				stat += 1
			}
			if char == ")" {
				stat += -1
			}
			if char == ")" && stat == 0 {
				matched = line[1:c]
			}
		}
		r := regexp.MustCompile(`\s+`)
		var data []string
		data = r.Split(matched,-1)
		if data[0] == "C" {
			params := strings.Join(data[1:], " ")
			paramParser(params)
		}
	}
}