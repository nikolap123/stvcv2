package commandFormatter

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type Command struct {
	Command string
	Args    []string
	Result  CommandResult
	Next    *Command
}

type CommandResult struct {
	CommandExp    string
	CommandResult string
}

func (R *Command) setArgs(args []string) { R.Args = args }
func (R *Command) setName(name string)   { R.Command = name }

func (R *Command) Exec() {

	if R.Command == "ares-inspect" {
		// TODO: This should be used with CommandContext to set cmd.Process.Kill timeout
		cmd := exec.Command(R.Command, R.Args...)

		var outb bytes.Buffer
		cmd.Stdout = &outb
		cmd.Start()

		time.Sleep(5 * time.Second)

		R.Result.CommandExp = R.Command + " " + strings.Join(R.Args, " ")
		R.Result.CommandResult = outb.String()

	} else {
		fmt.Println(R.Command, R.Args)

		//out, err := exec.Command(R.Command, R.Args...).Output()

		//fmt.Println(string(out))
		//
		R.Result.CommandExp = R.Command + " " + strings.Join(R.Args, " ")
		R.Result.CommandResult = "Testasdagtasda"
		//
		//if err != nil {
		//	R.Next = nil
		//}
	}

	if R.Next != nil {
		R.Next.Exec()
	}
}

func (R *Command) GetResult(result *[]CommandResult) {

	*result = append((*result), R.Result)

	if R.Next != nil {
		R.Next.GetResult(result)
	}

	return

}
