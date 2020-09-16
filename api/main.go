package api

import(
	"io/ioutil"
	"log"
	"strings"
	// "fmt"
	"os"
	// "zi/util"
	// "strconv"
)

type Pair struct {
	Key string `json:"Key"`
	Value string `json:"Value"`
	Line int `json:"Line"`
}

func Init() []Pair {
	list := []Pair{}
	content,err := ioutil.ReadFile("dump.zi")
	if err != nil {
		log.Fatal(err)
	}
	file := string(content)
	broken := strings.Split(file, "\n")
	for i, item := range broken {
		var trimmed string;
		if strings.Trim(item, trimmed) != "" {
			var line string = item
			parsed := strings.Fields(line)
			k := parsed[0]
			v := strings.Join(parsed[1:], " ")
			fullPair := Pair{Key: k, Value: v, Line: i + 1}
			list = append(list, fullPair)
		}
	}
	return list
}

func Get(data []Pair,key string) Pair {
	for _, item := range data {
		if(item.Key == key) {
			return item
		}
	}
	return Pair{Key: "", Value: "", Line:0}
}

func Set(item Pair) {
	f, err := os.OpenFile("dump.zi",
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString("\n"+item.Key + " " + item.Value); err != nil {
		log.Println(err)
	}
}

func Del(key string) {
	parsed := Init()
	for i := 0; i < len(parsed); i++ {
		if  parsed[i].Key == key {
			parsed = append(parsed[:i], parsed[i+1:]...)
			i--
	}
	}
	f,err := os.Create("dump.zi")
	if err != nil {
		panic(err)
	}
	f.WriteString("")
	for _, item := range parsed {
		Set(item)
	}
}