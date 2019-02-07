package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"syscall"
)

var urlRegexp = regexp.MustCompile(`^(?:git@|https://)(github\.com|gitlab\.com)[:/]([a-zA-Z0-9-]+)/([a-zA-Z0-9-]+).git$`)

func die(msg string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", a...)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		die("usage: %s <repository>", filepath.Base(os.Args[0]))
	}

	repo := os.Args[1]
	slug := urlRegexp.FindStringSubmatch(repo)
	if slug == nil {
		die("invalid url")
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
		die("error creating directory %s: %v", target, err)
	}

	git, err := exec.LookPath("git")
	if err != nil {
		die("couldn't find git executable: %v", err)
	}

	err = syscall.Exec(git, []string{"git", "clone", repo, target}, os.Environ())
	if err != nil {
		die("failed to exec git clone: %v", err)
	}
}
