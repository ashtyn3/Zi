package main

import (
	"texter/client"
	// "math/rand"
	"texter/util"
	// "strconv"
	"fmt"
	"os/exec"
	"os"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	args := os.Args[1:]
	// isometric2
	hero := figure.NewFigure("Texter","larry3d", true)
	hero.Print()
	fmt.Println("")
	// idk,_:= util.Contain(args, "hi")
	if i , stat := util.Find(args, "serve"); stat == true {
		index := i
		if index + 1 <= len(args) - 1 {
			if args[index + 1] != ""{
				if _ , stat := util.Find(args, "--background"); stat == true {
					port := args[index + 1]
					ok := exec.Command("god", "-r", ".", "-l", "the.log", "texter", "serve",port)
					ok.Run()
					fmt.Println("Server running on port "+ port)
				} else {
					client.Serve(args[index + 1])
				}
			} else {
				client.Serve("9090")
			}
		} else {
			client.Serve("9090")
		}
	} else if _ , stat := util.Find(args, "init"); stat == true {
		f, err := os.Create("./dump.txter")
		if err != nil {
			panic(err)
		}
		f.WriteString("")
	} else {
		fmt.Println("Help:")
		fmt.Println("\t- Serve: Serve server on port 9090 by default.")
	}

}