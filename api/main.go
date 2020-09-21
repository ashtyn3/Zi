package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	// "zi/util"
	"net/http"
	// "strconv"
)

type Pair struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
	Line  int    `json:"Line"`
}

func Init() []Pair {
	list := []Pair{}
	content, err := ioutil.ReadFile("dump.zi")
	if err != nil {
		log.Fatal(err)
	}
	file := string(content)
	broken := strings.Split(file, "\n")
	for i, item := range broken {
		var trimmed string
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

func Get(data []Pair, key string) Pair {
	fmt.Println("Query(GET): " + key)
	for _, item := range data {
		if item.Key == key {
			return item
		}
	}
	return Pair{Key: "", Value: "", Line: 0}
}

func Set(item Pair, verbose bool) {
	f, err := os.OpenFile("dump.zi",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString("\n" + item.Key + " " + item.Value); err != nil {
		log.Println(err)
	}
	json, err := json.Marshal(item)
	if err != nil {
		log.Println(err)
	}
	if verbose == true {
		fmt.Println("Query(SET): " + string(json))
	}
}

func Del(key string) {
	parsed := Init()
	for i := 0; i < len(parsed); i++ {
		if parsed[i].Key == key {
			parsed = append(parsed[:i], parsed[i+1:]...)
			i--
		}
	}
	f, err := os.Create("dump.zi")
	if err != nil {
		panic(err)
	}
	f.WriteString("")
	for _, item := range parsed {
		Set(item, false)
	}
	fmt.Println("Query(DELETE): " + key)
}

func GetAll() string {
	all := Init()
	json, _ := json.Marshal(all)
	fmt.Println("Query(GET): *")
	return string(json)
}

func Bind(url string) {
	res, err := http.Get(url + "/getall")

	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	var built []Pair
	json.Unmarshal([]byte(data), &built)
	fmt.Println(built)
}
