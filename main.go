package main

import (
	"bufio"
	"eleDB/db"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var DB db.IServer

func main() {
	println("SEEMIS is starting...")

	DB = *db.New("./db.eledb", false, 60)
	println(DB.Get("test") + "\n")

	reader := bufio.NewReader(os.Stdin)

	go func() {
		for {
			print("SEEMIS $ ")

			text, err := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			throw(err)

			for _, v := range commands {
				if v.Caller == strings.Split(text, " ")[0] {
					res := v.Function(text)
					println(res)
					break
				}
			}
		}
	}()

	sc := make(chan os.Signal, 1)
	waitForExit(sc)

	println("\n\n===\n\nTo prevent from data loss, I'm not exiting without confirmation.  " +
		"Interrupt again to confirm the exit.\nThe current state of the DB will be saved in, at most, 20 seconds.")

	sc = make(chan os.Signal, 1)
	waitForExit(sc)
}

func waitForExit(sc chan os.Signal) {
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}