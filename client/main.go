package client

import(
	"log"
	"net/http"
	"encoding/json"
	"texter/api"
	"fmt"
	"strconv"
)
type request struct {
	Key string `json:"key"`
	Value string `json:"value"`
	Line string `json:"line"`
}
func get(w http.ResponseWriter, r *http.Request) {
	K, ok := r.URL.Query()["key"]
	if ok != true {
		w.Write([]byte("key not found"))
	} else {
		parsed := api.Init()
		data := api.Get(parsed, K[0])
		Json, err := json.Marshal(request{Key: data.Key, Value: data.Value, Line: strconv.Itoa(data.Line)})
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(string(Json)))
	}

}
type SetPair struct {
	Value string `json:"Value"`
	Key string `json:"Key"`
}
func set(w http.ResponseWriter, r *http.Request) {
	K, ok := r.URL.Query()["data"]
	if ok != true {
		w.Write([]byte("Data not found"))
	} else {
		s := string(K[0])
    	data := SetPair{}
    	json.Unmarshal([]byte(s), &data)
		api.Set(api.Pair{Key: data.Key, Value: data.Value})
		parsed := api.Init()
		get := api.Get(parsed, data.Key)
		json, err := json.Marshal(api.Pair{Line: get.Line, Value: get.Value, Key: get.Key})
		if err != nil {
			panic(err)
		}
		w.Write([]byte(string(json)))
	}
}
func del(w http.ResponseWriter, r *http.Request) {
	K, ok := r.URL.Query()["key"]
	if ok != true {
		w.Write([]byte("Data not found"))
	} else {
		api.Del(K[0])
	}
}
func Serve(port string) {
	fmt.Println("Server running on port "+ port)
	http.HandleFunc("/get",get)
	http.HandleFunc("/set",set)
	http.HandleFunc("/del",del)
	err := http.ListenAndServe(":" +port, nil)
	if err != nil {
		log.Fatal(err)
	}
}