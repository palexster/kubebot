package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-chat-bot/bot"
)

type Kubebot struct {
	token    string
	admins   map[string]bool
	channels map[string]bool
	commands map[string]bool
}

const (
	forbiddenUserMessage     string = "%s - ⚠ kubectl forbidden for user @%s\n"
	forbiddenChannelMessage  string = "%s - ⚠ Channel %s forbidden for user @%s\n"
	forbiddenCommandMessage  string = "%s - ⚠ Command %s forbidden for user @%s\n"
	forbiddenFlagMessage     string = "%s - ⚠ Flag(s) %s forbidden for user @%s\n"
	forbiddenUserResponse    string = "Sorry @%s, but you don't have permission to run this command :confused:"
	forbiddenChannelResponse string = "Sorry @%s, but I'm not allowed to run this command here :zipper_mouth_face:"
	forbiddenCommandResponse string = "Sorry @%s, but I cannot run this command."
	forbiddenFlagResponse    string = "Sorry @%s, but I'm not allowed to run one of your flags."
	okResponse               string = "Roger that!\n@%s, this is the response to your request:\n"
)

var (
	ignored = map[string]map[string]bool{
		"get": map[string]bool{
			"-f":           true,
			"--filename":   true,
			"-w":           true,
			"--watch":      true,
			"--watch-only": true,
		},
		"describe": map[string]bool{
			"-f":         true,
			"--filename": true,
		},
		"create": map[string]bool{
			"-f":         true,
			"--filename": true,
		},
		"replace": map[string]bool{
			"-f":         true,
			"--filename": true,
		},
		"patch": map[string]bool{
			"-f":         true,
			"--filename": true,
		},
		"delete": map[string]bool{
			"-f":         true,
			"--filename": true,
		},
		"edit": map[string]bool{
			"-f":         true,
			"--filename": true,
		},
		"apply": map[string]bool{
			"-f":         true,
			"--filename": true,
		},
		"logs": map[string]bool{
			"-f":       true,
			"--follow": true,
		},
		"rolling-update": map[string]bool{
			"-f":         true,
			"--filename": true,
		},
		"scale": map[string]bool{
			"-f":         true,
			"--filename": true,
		},
		"attach": map[string]bool{
			"-i":      true,
			"--stdin": true,
			"-t":      true,
			"--tty":   true,
		},
		"exec": map[string]bool{
			"-i":      true,
			"--stdin": true,
			"-t":      true,
			"--tty":   true,
		},
		"run": map[string]bool{
			"--leave-stdin-open": true,
			"-i":                 true,
			"--stdin":            true,
			"--tty":              true,
		},
		"expose": map[string]bool{
			"-f":         true,
			"--filename": true,
		},
		"autoscale": map[string]bool{
			"-f":         true,
			"--filename": true,
		},
		"label": map[string]bool{
			"-f":         true,
			"--filename": true,
		},
		"annotate": map[string]bool{
			"-f":         true,
			"--filename": true,
		},
		"convert": map[string]bool{
			"-f":         true,
			"--filename": true,
		},
	}
)

func validateFlags(arguments ...string) error {
	if len(arguments) <= 1 {
		return nil
	}

	for i := 1; i < len(arguments); i++ {
		if ignored[arguments[0]][arguments[i]] {
			return errors.New(fmt.Sprintf("Error: %s is an invalid flag", arguments[i]))
		}

	}

	return nil
}

func kubectl(command *bot.Cmd) (result bot.CmdResultV3, err error) {
	nickname := command.User.Nick

	mchan := make(chan string)
	done := make(chan bool)

	if err := checkPerms(command, nickname); err != nil {
		go sendMessage("Something went wrong", err.Error(), mchan, done)

		return bot.CmdResultV3{command.Channel, mchan, done}, nil
	}

	output := execute("kubectl", command.Args...)

	go sendMessage(fmt.Sprintf(okResponse, nickname), output, mchan, done)

	return bot.CmdResultV3{command.Channel, mchan, done}, nil
}

func checkPerms(command *bot.Cmd, nickname string) error {
	t := time.Now()
	time := t.Format(time.RFC3339)

	//TODO: use log instead of fmt
	if !kb.admins[nickname] {
		fmt.Printf(forbiddenUserMessage, time, nickname)
		return errors.New(fmt.Sprintf(forbiddenUserResponse, nickname))
	}

	if !kb.channels[command.Channel] {
		fmt.Printf(forbiddenChannelMessage, time, command.Channel, nickname)
		return errors.New(fmt.Sprintf(forbiddenChannelResponse, nickname))
	}

	if len(command.Args) > 0 && !kb.commands[command.Args[0]] {
		fmt.Printf(forbiddenCommandMessage, time, command.Args, nickname)
		return errors.New(fmt.Sprintf(forbiddenCommandResponse, nickname))
	}

	if err := validateFlags(command.Args...); err != nil {
		fmt.Printf(forbiddenFlagMessage, time, command.Args, nickname)
		return errors.New(fmt.Sprintf(forbiddenFlagResponse, nickname))
	}

	return nil
}

func sendMessage(prefix, message string, mchan chan string, done chan bool) {
	const max = 1500

	mchan <- prefix

	for len(message) > 0 {
		msg := "```\n"

		if len(message) >= max {
			msg += message[:max]
			message = message[max:]
		} else {
			msg += message
			message = ""
		}

		msg += "\n```\n"

		mchan <- msg
	}

	done <- true
}

func init() {
	bot.RegisterCommandV3(
		"kubectl",
		"Kubectl Slack integration",
		"cluster-info",
		kubectl)
}
