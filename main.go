package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

var urlRegexp = regexp.MustCompile(`^(?:git@|https://)(github\.com|gitlab\.com)[:/]([a-zA-Z0-9-]+)/([a-zA-Z0-9-]+).git$`)

func die(msg string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", a...)
	os.Exit(1)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		die("usage: %s <repository>", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	repo := args[1]
	slug := urlRegexp.FindStringSubmatch(repo)
	if slug == nil {
		die("invalid url")
		os.Exit(1)
	}
	gitpath := os.Getenv("GITPATH")
	if gitpath == "" {
		die("GITPATH is not set")
	}
	host := slug[1]
	user := slug[2]
	project := slug[3]
	target := filepath.Join(gitpath, host, user, project)
	err := os.MkdirAll(target, 0777)
	if err != nil {
		die("error creating directory %s", target)
	}
	cmd := exec.Command("git", "clone", repo, target)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}
