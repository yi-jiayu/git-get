# git-get
Clone git repositories into namespaced directories like go get

## Example
```console
$ export GITPATH=$HOME/git
$ git get https://github.com/golang/go.git
Cloning into '/Users/yijiayu/git/github.com/golang/go'...
...
$ git get https://github.com/yi-jiayu/bus-eta-bot.git
Cloning into '/Users/yijiayu/git/github.com/yi-jiayu/bus-eta-bot'...
...
$ git get https://gitlab.com/yi-jiayu/bus-eta-bot.git
Cloning into '/Users/yijiayu/git/gitlab.com/yi-jiayu/bus-eta-bot'...
...
$ tree -L 3 $GITPATH
/Users/yijiayu/git
├── github.com
│   ├── golang
│   │   └── go
│   └── yi-jiayu
│       └── bus-eta-bot
└── gitlab.com
    └── yi-jiayu
        └── bus-eta-bot

10 directories, 0 files
```

## Installation

With a working Go toolchain:
```
go build -o ~/bin/git-get
```
Replace `~/bin` with any directory on your `$PATH`.

`git` will make any executable prefixed with `git-` on your `$PATH` available as a `git` subcommand, so you will be able to run `git-get` as `git get`.

## URL support

- [x] GitHub
- [x] GitLab
- [ ] Generic
