package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	tele "gopkg.in/telebot.v3"
)

func getUsersList(s string) (users []int64) {
	ret := strings.Split(s, ":")
	for _, v := range ret {
		iv, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, int64(iv))
	}
	return
}

func buildCmds() (cmds []tele.Command) {
	files, err := ioutil.ReadDir(cScriptsDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		comment := getScriptComment(cScriptsDir + "/" + f.Name())
		if comment == "" {
			comment = f.Name()
		}
		cmd := tele.Command{
			Text:        f.Name(),
			Description: comment,
		}
		cmds = append(cmds, cmd)
	}
	return
}

func getScriptComment(path string) (comment string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "#!") {
			comment = line
			return
		}
	}
	return ""
}
