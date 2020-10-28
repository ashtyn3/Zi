package main

import (
	"github.com/ashtyn3/zi/api"
	"github.com/ashtyn3/zi/client"

	// "math/rand"
	"github.com/ashtyn3/zi/util"
	// "strconv"
	"fmt"
	"os"
	"os/exec"

	cmd "github.com/ashtyn3/zi/command"

	"github.com/common-nighthawk/go-figure"
)

func main() {
	args := os.Args[1:]
	// isometric2
	hero := figure.NewFigure("Zi", "larry3d", true)
	hero.Print()
	fmt.Println("")
	// idk,_:= util.Contain(args, "hi")
	if i, stat := util.Find(args, "serve"); stat == true {
		index := i
		if index+1 <= len(args)-1 {
			if args[index+1] != "" {
				if _, stat := util.Find(args, "--background"); stat == true {
					port := args[index+1]
					ok := exec.Command("god", "-r", ".", "-l", "zi.log", "zi", "serve", port)
					ok.Run()
					fmt.Println("Server running on port " + port)
				} else {
					client.Serve(args[index+1])
				}
			} else {
				client.Serve("9090")
			}
		} else {
			client.Serve("9090")
		}
	} else if _, stat := util.Find(args, "init"); stat == true {
		f, err := os.Create("./dump.zi")
		if err != nil {
			panic(err)
		}
		f.WriteString("")
	} else if _, stat := util.Find(args, "--docker"); stat == true {
		fmt.Println("Server running on port 5000")
		if _, stat := util.Find(args, "--detached"); stat == true {
			starter := exec.Command("sudo", "docker", "run", "-p", "5000:9090", "-d", "--mount", "source=zi-presist,target=/app", "vitecoin/zi")
			starter.Run()

		} else {
			starter := exec.Command("sudo", "docker", "run", "-p", "5000:9090", "--mount", "source=zi-presist,target=/app", "vitecoin/zi")
			starter.Run()
		}
	} else if _, stat := util.Find(args, "run"); stat == true {
		cmd.Do()
	} else if _, stat := util.Find(args, "auth"); stat == true {
		i, _ := util.Find(args, "auth")
		if len(args)-1 < i+1 {
			fmt.Println("Please provide a password.")
		} else {
			api.Validate(args[i+1], true)
		}
	} else {
		fmt.Println("Help:")
		fmt.Println("\t- serve: Starts server server on port 9090 by default.")
		fmt.Println("\t  --background: Placed at the end of serve will run server as a daemon proccess.")
		fmt.Println("\t  --docker: Uses perfered flags to start docker container. Container MUST be installed")
		fmt.Println("\t- init: Creates dump.zi file.")
		fmt.Println("\t- run: Opens a CLI for running queries directly.")
	}
	// TODO: Add background server killing flag.
	// Use the following exec:
	// ps -ef | grep zi
	// kill [pid]
	// a very very very very secret key

	// key := []byte("a very very very very secret key") // 32 bytes
	// plaintext := []byte("some really really really long plaintext\nwow boom")
	// fmt.Printf("%s\n", plaintext)
	// ciphertext, err := cto.Encrypt(key, plaintext)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(hex.EncodeToString(ciphertext))
	// result, err := cto.Decrypt(key, ciphertext)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
