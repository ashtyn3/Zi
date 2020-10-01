package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"zi/api"
)

var clear map[string]func() //create a map for storing clear funcs

func CLEAR() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported...")
	}
}
func Do() {
	// data := []string{}
	for i := 0; i < 1; i-- {
		fmt.Printf("\b> ")
		cmd := bufio.NewScanner(os.Stdin)
		if cmd.Scan() {
			line := cmd.Text()
			if strings.Trim(line, "") == "" {
				fmt.Printf("")
			} else {
				parsed := strings.Fields(line)
				if parsed[0] == "help" {
					fmt.Println("HELP:")
					fmt.Println("GET: [KEY]")
					fmt.Println("SET: [KEY] [VALUE]")
					fmt.Println("DEL: [KEY]")
				} else if parsed[0] == "SET" {
					if len(parsed) >= 3 {
						k := parsed[1]
						v := strings.Join(parsed[2:], " ")
						api.Set(api.Pair{Key: k, Value: v}, true)
					} else {
						fmt.Println("ERROR: Expected Key and Value but only got one.")
					}

				} else if parsed[0] == "GET" {
					if parsed[1] == "*" {
						var strJson []api.Pair
						json.Unmarshal([]byte(api.GetAll()), &strJson)
						data, err := json.MarshalIndent(strJson, "", "  ")
						if err != nil {
							log.Fatal(err)
						}

						fmt.Println(string(data))
					} else {
						p := api.Init()
						data, _ := json.Marshal(api.Get(p, parsed[1], true))
						fmt.Println(string(data))
					}
				} else if parsed[0] == "DEL" {
					api.Del(parsed[1], true)
				} else if parsed[0] == "clear" {
					CLEAR()
					CallClear()
				} else if parsed[0] == "bind" {
					api.Bind(parsed[1], parsed[2], true)
				} else if parsed[0] == "GET_ROW" {
					p := api.Init()
					data, _ := json.Marshal(api.GetRow(p, parsed[1]))
					fmt.Println(string(data))
				} else if parsed[0] == "RENAME" {
					api.Rename(parsed[1], parsed[2], true)
				} else {
					fmt.Println("ERROR: Bad command " + line)
				}
			}
		}
	}
}
