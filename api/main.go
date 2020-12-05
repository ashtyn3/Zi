package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	cto "github.com/ashtyn3/zi/crypto"
)

type Pair struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
	Line  int    `json:"Line"`
}

func Init() []Pair {
	list := []Pair{}
	b64v, err := ioutil.ReadFile("dump.zi")
	content := string(b64v)
	if err != nil {
		log.Fatal(err)
	}
	file := content
	broken := strings.Split(file, "\n")
	for i, item := range broken {
		item = cto.B64_dec(item)
		var trimmed string
		if strings.Trim(item, trimmed) != "" {
			var line string = item
			parsed := strings.Fields(line)
			if len(parsed) != 0 {

				k := parsed[0]
				v := strings.Join(parsed[1:], " ")
				fullPair := Pair{Key: k, Value: v, Line: i + 1}
				list = append(list, fullPair)
			}

		}
	}
	return list
}

func ModInit(content string) []Pair {
	list := []Pair{}
	broken := strings.Split(content, "\n")
	for i, item := range broken {
		item = cto.B64_dec(item)
		var trimmed string
		if strings.Trim(item, trimmed) != "" {
			var line string = item
			parsed := strings.Fields(line)
			if len(parsed) != 0 {

				k := parsed[0]
				v := strings.Join(parsed[1:], " ")
				fullPair := Pair{Key: k, Value: v, Line: i + 1}
				list = append(list, fullPair)
			}

		}
	}
	return list
}

func Get(data []Pair, key string, print bool) Pair {
	if print == true {
		fmt.Println("Query(GET): " + key)
	}
	if strings.Contains(key, "*") == true {
		key = strings.Replace(key, "*", "", 1)
		entry := strings.Split(key, ":")
		for _, item := range data {
			if item.Key == "*"+entry[0] {
				res, err := http.Get(item.Value + "/getall")
				if err != nil {
					log.Fatal(err)
				}
				data, _ := ioutil.ReadAll(res.Body)
				res.Body.Close()
				var built []Pair
				json.Unmarshal(data, &built)
				return Get(built, entry[1], true)
				// return item
			}
		}
	} else if strings.HasPrefix(key, "^") == true {
		var found []Pair

		for _, item := range data {
			if item.Key == strings.Replace(key, "^", "", 1) {
				f, err := ioutil.ReadFile(item.Value)
				if err != nil {
					log.Fatal(err)
				}
				parsed := ModInit(string(f))
				for _, v := range parsed {
					found = append(found, v)
				}
			}
		}
		sj, _ := json.Marshal(found)
		return Pair{Key: strings.Replace(key, "^", "", 1), Value: string(sj)}
	} else {
		matched := []Pair{}
		for _, item := range data {
			if item.Key == key {
				matched = append(matched, item)
			}
		}
		if len(matched) != 0 {
			return matched[len(matched)-1]
		}
	}
	return Pair{Key: "", Value: "", Line: 0}

}

// GetRow gets every version of key
func GetRow(data []Pair, key string) []Pair {
	matched := []Pair{}
	for _, item := range data {
		if item.Key == key {
			matched = append(matched, item)
		}
	}
	fmt.Println("Query(GET): " + key)
	return matched
}

func Set(item Pair, verbose bool) {
	f, err := os.OpenFile("dump.zi",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString("\n" + cto.B64_enc(item.Key+" "+item.Value)); err != nil {
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

func Del(key string, print bool) {
	parsed := Init()
	if strings.HasPrefix(key, "^") == true {
		for _, item := range parsed {
			if "^"+item.Key == key {
				os.Remove(item.Value)
			}
		}
	}
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
	if print == true {
		fmt.Println("Query(DELETE): " + key)

	}
}

func GetAll() string {
	all := Init()
	json, _ := json.Marshal(all)
	fmt.Println("Query(GET): *")
	return string(json)
}

func Bind(title, url string, print bool) {
	res, err := http.Get(url + "/getall")
	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	var built []Pair
	json.Unmarshal([]byte(data), &built)
	r, _ := json.Marshal(built)
	if string(r) != "[]" {
		Set(Pair{Value: url, Key: "*" + title}, false)
	}
	if print == true {
		fmt.Println("Action(BIND): " + title + ":" + url)
	}
}

func Rename(origin string, new string, print bool) {
	d := Get(Init(), origin, false)
	Del(origin, false)
	Set(Pair{Key: new, Value: d.Value}, false)
	if print == true {
		fmt.Println("Query(RENAME): " + origin + " to " + new)
	}
}

func Dump(k, v, path string, print bool) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString("\n" + cto.B64_enc(k+" "+v))
	Set(Pair{Key: k, Value: path}, false)
	if print == true {
		fmt.Println("Query(DUMP): " + "+" + k + ":" + path)
	}
}

func Validate(pass string, new bool) string {
	if new == true {
		f, err := os.OpenFile("dump.zi", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}
		f.WriteAt([]byte(cto.B64_enc("++pd "+pass)), 0)
	} else {
		pair := Get(Init(), "++pd", false)
		if pair.Value == pass {
			return "ok"
		}
	}
	return "bad"
}
