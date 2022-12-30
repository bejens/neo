package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

// `neo build`
// `neo install`
// `neo update`
var (
	app        = kingpin.New("neo", "neo command")
	buildCmd   = app.Command("build", "neo build")
	installCmd = app.Command("install", "neo install")
	updateCmd  = app.Command("update", "neo update")
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case buildCmd.FullCommand():
		fmt.Println("neo build")
	case installCmd.FullCommand():
		fmt.Println("neo install")
	case updateCmd.FullCommand():
		fmt.Println("neo update")
	default:
		fmt.Println("unsupported command")
	}
}
