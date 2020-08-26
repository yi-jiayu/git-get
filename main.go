package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
)

const website = "https://github.com/yi-jiayu/git-get"

var (
	version = "dev"
	commit  = "none"
)

var urlRegexp = regexp.MustCompile(`^(?:(?:ssh|git|https?|ftps?)://(?:[\w-.]+@)?([a-zA-Z0-9.-]+)(?::\d+)?/|file:///|(?:[\w-]+@)?([a-zA-Z0-9.-]+):|/)~?([\w./-]+)$`)

var showVersion bool

// destinationPath returns the path that a repository should be cloned to relative to GITPATH.
func destinationPath(u string) string {
	u = strings.TrimSuffix(u, "/")    // remove trailing slash
	u = strings.TrimSuffix(u, ".git") // remove .git suffix
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

func init() {
	flag.BoolVar(&showVersion, "v", false, "print version info and exit")
}

func main() {
	flag.Parse()

	if showVersion {
		fmt.Printf("git-get %s\ncommit %s\n%s\n", version, commit, website)
		return
	}

	if flag.NArg() < 1 {
		die("usage: %s <repository>", filepath.Base(os.Args[0]))
	}

	repo := flag.Arg(0)
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
