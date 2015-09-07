package main

import (
	"fmt"
	"github.com/go-fsnotify/fsnotify"
	"log"
)

var (
	chat *Chat
)

func init() {
	var err error
	chat, err = NewChatReader("chat.log")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Hello! watching file!")

	defer chat.Close()

	chat.Parse()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					chat.Parse()
					log.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("chat.log")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
