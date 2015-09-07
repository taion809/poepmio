package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

type Chat struct {
	handle   *os.File
	buffer   *bufio.Scanner
	Filename string
}

type ChatMessage struct {
	Received time.Time
	Receiver string
	Message  string
}

const WHISPER_PREFIX = "@"

func NewChatReader(filename string) (*Chat, error) {
	chat := &Chat{Filename: filename}
	err := chat.Init()

	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (c *Chat) Init() error {
	file, err := os.Open(c.Filename)

	if err != nil {
		return err
	}

	buffer := bufio.NewScanner(file)

	c.handle = file
	c.buffer = buffer

	return nil
}

func (c *Chat) Parse() ([]ChatMessage, error) {
	err := c.Init()
	if err != nil {
		return err
	}

	for c.buffer.Scan() {
		line := c.buffer.Text()
		if line == "" {
			continue
		}

		split := strings.SplitN(line, " ", 8)

		if len(split) < 8 {
			continue
		}

		if !strings.HasPrefix(split[7], WHISPER_PREFIX) {
			continue
		}

		log.Println(split[7][1:])
	}

	return nil
}

func (c *Chat) Close() error {
	var err error
	if c != nil && c.handle != nil {
		err = c.handle.Close()
	}

	return err
}
