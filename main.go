package main

import (
	"flag"
	"log"
	"os"
)

var fBotKey = flag.String("key", "", "telegram bot api key")
var fUsers = flag.String("users", "", "telegram users ids to use this bot, like: 123:456:789")
var fShell = flag.String("shell", "sh", "shell to run script")
var fVerbose = flag.Bool("verbose", false, "verbose output")

const cScriptsDir = "./scripts"
const cTmpDir = "./tmp"

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if *fBotKey == "" || *fUsers == "" {
		flag.Usage()
		return
	}
	if _, err := os.Stat(cScriptsDir); os.IsNotExist(err) {
		os.Mkdir(cScriptsDir, os.ModePerm)
	}
	runBot()
}
