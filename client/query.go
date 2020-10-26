package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/vitecoin/zi/api"
)

func getQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	K, ok := r.URL.Query()["key"]
	if ok != true {
		w.Write([]byte("key not found"))
	} else {
		if K[0] != "++pd" {
			parsed := api.Init()
			data := api.Get(parsed, K[0], true)
			Json, err := json.Marshal(request{Key: data.Key, Value: data.Value, Line: strconv.Itoa(data.Line)})
			if err != nil {
				log.Fatal(err)
			}
			w.Write([]byte(string(Json)))
		}

	}
}

func setQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	K, ok := r.URL.Query()["data"]
	if ok != true {
		w.Write([]byte("Data not found"))
	} else {
		s := string(K[0])
		data := SetPair{}
		json.Unmarshal([]byte(s), &data)
		if data.Key != "++pd" {
			api.Set(api.Pair{Key: data.Key, Value: data.Value}, true)
			parsed := api.Init()
			get := api.Get(parsed, data.Key, false)
			json, err := json.Marshal(api.Pair{Line: get.Line, Value: get.Value, Key: get.Key})
			if err != nil {
				panic(err)
			}
			w.Write([]byte(string(json)))
		}

	}
}

func dumpQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	K, ok := r.URL.Query()["data"]
	P, okP := r.URL.Query()["path"]

	if ok != true || okP != true {
		w.Write([]byte("Data or path not found"))
	} else {
		s := K[0]
		data := SetPair{}
		json.Unmarshal([]byte(s), &data)
		if data.Key != "++pd" {
			api.Dump(data.Key, data.Value, P[0], true)
			w.Write([]byte("OK"))
		}
	}
}

func delQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	K, ok := r.URL.Query()["key"]
	if ok != true {
		w.Write([]byte("Key not found"))
	} else {
		if K[0] != "++pd" {

			api.Del(K[0], true)
			w.Write([]byte("OK"))
		}
	}
}

func getAllQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	var unFiltered []api.Pair
	json.Unmarshal([]byte(api.GetAll()), &unFiltered)
	var filtered []api.Pair
	for _, p := range unFiltered {
		if p.Key != "++pd" {
			filtered = append(filtered, p)
		}
	}
	result, err := json.Marshal(filtered)
	if err != nil {
		fmt.Println(err)
	}
	if string(result) == "null" {
		w.Write([]byte("[]"))
	} else {
		w.Write(result)

	}
}

func bindQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	url, okUrl := r.URL.Query()["url"]
	key, okKey := r.URL.Query()["key"]
	if okUrl == true && okKey == true {
		if key[0] != "++pd" {

			api.Bind(key[0], url[0], true)
			w.Write([]byte("OK"))
		}

	} else {
		w.Write([]byte("Key or url not found"))
	}

}

func renameQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	new, okUrl := r.URL.Query()["new"]
	origin, okKey := r.URL.Query()["origin"]
	if okUrl == true && okKey == true {
		if origin[0] != "++pd" || new[0] != "++pd" {

			api.Rename(origin[0], new[0], true)
			w.Write([]byte("OK"))
		}

	} else {
		w.Write([]byte("New or origin not found"))
	}

}

func getrowQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	key, ok := r.URL.Query()["key"]
	if ok == false {
		w.Write([]byte("Key not found"))
	} else {
		if key[0] != "++pd" {

			r := api.GetRow(api.Init(), key[0])
			data, _ := json.Marshal(r)
			w.Write([]byte(data))
		}
	}

}
