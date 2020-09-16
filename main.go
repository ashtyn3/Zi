package main

import (
	"zi/client"
	// "math/rand"
	"zi/util"
	// "strconv"
	"fmt"
	"os/exec"
	"os"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	args := os.Args[1:]
	// isometric2
	hero := figure.NewFigure("Zi","larry3d", true)
	hero.Print()
	fmt.Println("")
	// idk,_:= util.Contain(args, "hi")
	if i , stat := util.Find(args, "serve"); stat == true {
		index := i
		if index + 1 <= len(args) - 1 {
			if args[index + 1] != ""{
				if _ , stat := util.Find(args, "--background"); stat == true {
					port := args[index + 1]
					ok := exec.Command("god", "-r", ".", "-l", "the.log", "zi", "serve",port)
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
		f, err := os.Create("./dump.zi")
		if err != nil {
			panic(err)
		}
		f.WriteString("")
	} else {
		fmt.Println("Help:")
		fmt.Println("\t- serve: Starts server server on port 9090 by default.")
		fmt.Println("\t  --background: Placed at the end of serve will run server as a daemon proccess.")
		fmt.Println("\t- init: Creates dump.zi file.")
	}
// TODO: Add background server killing flag.
// Use the following exec:
// ps -ef | grep zi
// kill [pid]
}