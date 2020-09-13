package main

import (
	"texter/api"
	"fmt"
)
func main() {
	parsed := api.Init()
	item := api.Pair{Key: "i", Value: "world"}
	api.Set(item)
	fmt.Println(api.Get(parsed, "i"))
}