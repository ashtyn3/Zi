package api

import(
	"io/ioutil"
	"log"
	"strings"
	// "fmt"
	"os"
	// "strconv"
)

type Pair struct {
	Key string
	Value string
	line int
}

func Init() []Pair {
	list := []Pair{}
	content,err := ioutil.ReadFile("dump.txter")
	if err != nil {
		log.Fatal(err)
	}
	file := string(content)
	broken := strings.Split(file, "\n")
	for i, item := range broken {
		var line string = item
		parsed := strings.Fields(line)
		k := parsed[0]
		v := strings.Join(parsed[1:], " ")
		fullPair := Pair{Key: k, Value: v, line: i + 1}
		list = append(list, fullPair)
	}
	return list
}

func Get(data []Pair,key string) Pair {
	for _, item := range data {
		if(item.Key == key) {
			return item
		}
	}
	return Pair{Key: "", Value: "", line:0}
}

func Set(item Pair) {
	f, err := os.OpenFile("dump.txter",
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString("\n"+item.Key + " " + item.Value); err != nil {
		log.Println(err)
	}
}