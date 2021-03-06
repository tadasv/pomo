package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"syscall"
	"time"
)

var (
	configDirectory            = os.Getenv("HOME") + "/.config/pomo"
	currentTimerConfigFilePath = configDirectory + "/current"
)

func usage() {
	fmt.Printf(`usage: %s <command> [options]

commands:

 status
   print current timer if one is set

 start [duration]
   starts a new timer. Overwriter existing one if one is already running.
   If duration is specified it must follow Go duration format, e.g. 20m, 1h etc.

 stop
   stops current timer

`, os.Args[0])
	os.Exit(-1)
}

func readCurrentTimer() (time.Time, error) {
	data, err := ioutil.ReadFile(currentTimerConfigFilePath)
	if err != nil {
		return time.Time{}, err
	}

	return time.Parse(time.RFC3339, strings.TrimSpace(string(data)))
}

func setTimer(t time.Time) error {
	timestamp := t.Format(time.RFC3339)
	return ioutil.WriteFile(currentTimerConfigFilePath, []byte(timestamp), 0644)
}

func stopCurrentTimer() error {
	return os.Remove(currentTimerConfigFilePath)
}

func main() {
	command := "status"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	os.MkdirAll(configDirectory, 0755)

	switch command {
	case "start":
		duration := time.Minute * 25
		if len(os.Args) == 3 {
			d, err := time.ParseDuration(os.Args[2])
			if err != nil {
				panic(err)
			}
			duration = d
		}
		if err := setTimer(time.Now().Add(duration)); err != nil {
			panic(err)
		}
		fallthrough
	case "status":
		timer, err := readCurrentTimer()
		if err != nil {
			if pe, ok := err.(*os.PathError); ok && pe.Err == syscall.ENOENT {
				fmt.Printf("no timer set\n")
			} else {
				panic(err)
			}
		} else {
			duration := timer.Sub(time.Now())
			fmt.Printf("%s\n", duration.Truncate(time.Second).String())
		}
	case "stop":
		if err := stopCurrentTimer(); err != nil {
			if pe, ok := err.(*os.PathError); !ok || pe.Err != syscall.ENOENT {
				panic(err)
			}
		}
	default:
		usage()
	}
}
