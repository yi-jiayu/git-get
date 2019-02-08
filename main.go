package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"syscall"
)

var urlRegexp = regexp.MustCompile(`^(?:(?:ssh|git|https?|ftps?)://(?:[\w-]+@)?([a-zA-Z0-9.-]+)(?::\d+)?/|file:///|(?:[\w-]+@)?([a-zA-Z0-9.-]+):|/)([\w./-]+).git/?$`)

// destinationPath returns the path that a repository should be cloned to relative to GITPATH.
func destinationPath(u string) string {
	m := urlRegexp.FindStringSubmatch(u)
	if m == nil {
		return ""
	}
	target := path.Join(m[1:]...)
	return target
}

func die(msg string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", a...)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		die("usage: %s <repository>", filepath.Base(os.Args[0]))
	}

	repo := os.Args[1]
	dest := destinationPath(repo)
	if dest == "" {
		die("invalid url: %v", repo)
	}

	gitpath := os.Getenv("GITPATH")
	if gitpath == "" {
		die("GITPATH is not set")
	}

	target := filepath.FromSlash(path.Join(gitpath, dest))
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
