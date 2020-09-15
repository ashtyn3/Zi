package main

import (
	"texter/client"
	// "math/rand"
	"texter/util"
	// "strconv"
	// "strconv"
	"os"
)

func main() {
	args := os.Args[1:]
	// idk,_:= util.Contain(args, "hi")
	if i , stat := util.Find(args, "serve"); stat == true {
		index := i
		if index + 1 <= len(args) - 1 {
			if args[index + 1] != ""{
				client.Serve(args[index + 1])
			}
		}
	}
}