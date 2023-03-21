package main

import (
	"log"
	"os/exec"
	"time"

	"gopkg.in/telebot.v3/middleware"

	tele "gopkg.in/telebot.v3"
)

type WrapFunc func(cmd string, c tele.Context) error

func runBot() {
	pref := tele.Settings{
		Token:   *fBotKey,
		Verbose: false,
		Poller:  &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	cmds := buildCmds()
	err = b.SetCommands(cmds)
	if err != nil {
		log.Fatal(err)
	}

	// b.Use(middleware.Logger())
	// Group-scoped middleware:
	adminOnly := b.Group()
	adminOnly.Use(middleware.Whitelist(getUsersList(*fUsers)...))

	for _, cmd := range cmds {
		adminOnly.Handle("/"+cmd.Text, wrap(cmd.Text, runScript))
	}

	log.Println("telegram bot started")
	b.Start()
}
func wrap(cmd string, f WrapFunc) tele.HandlerFunc {
	return func(c tele.Context) error {
		return f(cmd, c)
	}
}
func runScript(script string, c tele.Context) error {
	path := cScriptsDir + "/" + script
	out, err := exec.Command(*fShell, path, c.Message().Payload).Output()
	if err != nil {
		return c.Reply("run script failed: " + err.Error())
	}
	ret := "run script done, result:\n" + string(out)
	return c.Reply(ret)
}
