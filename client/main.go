package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/vitecoin/zi/api"
)

type request struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Line  string `json:"line"`
}

func get(w http.ResponseWriter, r *http.Request) {
	p := api.Get(api.Init(), "++pd", false)
	if p.Value != "" {
		pass, Pok := r.URL.Query()["auth"]
		if Pok == true {
			valid := api.Validate(pass[0], false)
			if valid == "ok" {
				getQuery(w, r)
			} else {
				w.Write([]byte("bad"))
			}
		} else {
			w.Write([]byte("bad"))
		}
	} else {
		getQuery(w, r)
	}
}

type SetPair struct {
	Value string `json:"Value"`
	Key   string `json:"Key"`
}

func set(w http.ResponseWriter, r *http.Request) {
	p := api.Get(api.Init(), "++pd", false)
	if p.Value != "" {
		pass, Pok := r.URL.Query()["auth"]
		if Pok == true {
			valid := api.Validate(pass[0], false)
			if valid == "ok" {
				setQuery(w, r)
			} else {
				w.Write([]byte("bad"))
			}
		} else {
			w.Write([]byte("bad"))
		}
	} else {
		setQuery(w, r)
	}
}
func dump(w http.ResponseWriter, r *http.Request) {
	p := api.Get(api.Init(), "++pd", false)
	if p.Value != "" {
		pass, Pok := r.URL.Query()["auth"]
		if Pok == true {
			valid := api.Validate(pass[0], false)
			if valid == "ok" {
				dumpQuery(w, r)
			} else {
				w.Write([]byte("bad"))
			}
		} else {
			w.Write([]byte("bad"))
		}
	} else {
		dumpQuery(w, r)
	}
}
func del(w http.ResponseWriter, r *http.Request) {
	p := api.Get(api.Init(), "++pd", false)
	if p.Value != "" {
		pass, Pok := r.URL.Query()["auth"]
		if Pok == true {
			valid := api.Validate(pass[0], false)
			if valid == "ok" {
				delQuery(w, r)
			} else {
				w.Write([]byte("bad"))
			}
		} else {
			w.Write([]byte("bad"))
		}
	} else {
		delQuery(w, r)
	}
}
func getAll(w http.ResponseWriter, r *http.Request) {
	p := api.Get(api.Init(), "++pd", false)
	if p.Value != "" {
		pass, Pok := r.URL.Query()["auth"]
		if Pok == true {
			valid := api.Validate(pass[0], false)
			if valid == "ok" {
				getAllQuery(w, r)
			} else {
				w.Write([]byte("bad"))
			}
		} else {
			w.Write([]byte("bad"))
		}
	} else {
		getAllQuery(w, r)
	}
}
func bind(w http.ResponseWriter, r *http.Request) {
	p := api.Get(api.Init(), "++pd", false)
	if p.Value != "" {
		pass, Pok := r.URL.Query()["auth"]
		if Pok == true {
			valid := api.Validate(pass[0], false)
			if valid == "ok" {
				bindQuery(w, r)
			} else {
				w.Write([]byte("bad"))
			}
		} else {
			w.Write([]byte("bad"))
		}
	} else {
		bindQuery(w, r)
	}
}
func rename(w http.ResponseWriter, r *http.Request) {
	p := api.Get(api.Init(), "++pd", false)
	if p.Value != "" {
		pass, Pok := r.URL.Query()["auth"]
		if Pok == true {
			valid := api.Validate(pass[0], false)
			if valid == "ok" {
				renameQuery(w, r)
			} else {
				w.Write([]byte("bad"))
			}
		} else {
			w.Write([]byte("bad"))
		}
	} else {
		renameQuery(w, r)
	}
}
func getrow(w http.ResponseWriter, r *http.Request) {
	p := api.Get(api.Init(), "++pd", false)
	if p.Value != "" {
		pass, Pok := r.URL.Query()["auth"]
		if Pok == true {
			valid := api.Validate(pass[0], false)
			if valid == "ok" {
				getrowQuery(w, r)
			} else {
				w.Write([]byte("bad"))
			}
		} else {
			w.Write([]byte("bad"))
		}
	} else {
		getrowQuery(w, r)
	}
}
func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

// Serve starts database server.
func Serve(port string) {
	if _, err := os.Stat("dump.zi"); err != nil {
		ioutil.WriteFile("dump.zi", []byte(""), 0644)
	}
	fmt.Println("Server running on port " + port)
	fmt.Printf("Public address: http://%s:%s\n", getOutboundIP(), port)
	fmt.Printf("Localhost address: http://127.0.0.1:%s\n", port)
	http.HandleFunc("/get", get)
	http.HandleFunc("/set", set)
	http.HandleFunc("/del", del)
	http.HandleFunc("/getrow", getrow)
	http.HandleFunc("/getall", getAll)
	http.HandleFunc("/bind", bind)
	http.HandleFunc("/rename", rename)
	http.HandleFunc("/dump", dump)
	http.HandleFunc("/", root)
	http.ListenAndServe(":"+port, nil)
	// if err != nil {
	// log.Fatal(err)
	// }
}
