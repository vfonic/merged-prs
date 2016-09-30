package git

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

type client struct{}

func New() *client {
	return &client{}
}

func (c *client) Available() bool {
	return nil == exec.Command("git", "--version").Run()
}

func (c *client) IsRef(ref string) bool {
	return nil == exec.Command("git", "show", ref).Run()
}

func (c *client) IsRepo() bool {
	return nil == exec.Command("git", "rev-parse", "--is-inside-work-tree").Run()
}

func (c *client) GetMergeCommits(a string, b string) []*MergeCommit {
	// TODO: figure out how to fatal out
	if !c.IsRef(a) {
		log.Fatalf("%s, is not a valid ref, exiting", a)
	}
	if !c.IsRef(b) {
		log.Fatalf("%s, is not a valid ref, exiting", b)
	}

	arg := fmt.Sprintf("%s...%s", a, b)
	cmd := exec.Command("git", "log", "--merges", "--abbrev-commit", "--pretty=oneline", arg)

	out, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Something went wrong with the previous command: %v", cmd.Args)
	}

	return c.processMergeCommits(out)

	err = cmd.Start()
	if err != nil {
		log.Fatalf("Could not start command")
	}

	return c.processMergeCommits(out)
}

func (c *client) processMergeCommits(in io.Reader) []*MergeCommit {
	list := []*MergeCommit{}

	s := bufio.NewScanner(in)
	for s.Scan() {
		raw := strings.SplitN(s.Text(), " ", 2)

		// bad commit git
		if raw[0] == "" || raw[1] == "" {
			continue
		}

		cm := &MergeCommit{
			Ref:     raw[0],
			Message: raw[1],
		}
		list = append(list, cm)
	}

	return list
}
