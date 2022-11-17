package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

const (
	markName  = "Cli Reminder Tool"
	markValue = "1"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage <hh:mm> <test/message> ")
		os.Exit(0)
	}
	now := time.Now()

	w := when.New(nil)
	w.Add(common.All...)
	w.Add(en.All...)

	t, err := w.Parse(os.Args[1], now)

	if err != nil {
		fmt.Println("Unable to Parse")
	}

	if now.After(t.Time) {
		fmt.Println("Set the future time")
		os.Exit(1)
	}

	diff := t.Time.Sub(now)
	if os.Getenv(markName) == markValue {
		time.Sleep(diff)
		beeep.Alert("Reminder ", strings.Join(os.Args[2:], " "), " ")
		if err != nil {
			fmt.Println(err)
		}
	} else {

		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue))
		cmd.Start()

		fmt.Println("Reminder will be show in :", diff.Round(time.Second))

	}

}
