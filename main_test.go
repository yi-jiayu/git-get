package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_destinationPath(t *testing.T) {
	testCases := map[string][]string{
		"host.xz/path/to/repo": {
			"ssh://user@host.xz:1234/path/to/repo.git/",
			"ssh://user@host.xz/path/to/repo.git/",
			"ssh://host.xz:1234/path/to/repo.git/",
			"ssh://host.xz/path/to/repo.git/",

			"git://host.xz:1234/path/to/repo.git/",
			"git://host.xz/path/to/repo.git/",

			"https://host.xz/path/to/repo.git/",
			"https://host.xz:1234/path/to/repo.git/",
			"http://host.xz/path/to/repo.git/",
			"http://host.xz:1234/path/to/repo.git/",

			"ftps://host.xz:1234/path/to/repo.git/",
			"ftps://host.xz/path/to/repo.git/",
			"ftp://host.xz:1234/path/to/repo.git/",
			"ftp://host.xz/path/to/repo.git/",

			"host.xz:path/to/repo.git/",
			"user@host.xz:path/to/repo.git/",
		},
		"subdomain.host.xz/path/to/repo": {
			"ssh://user@subdomain.host.xz:1234/path/to/repo.git/",
			"ssh://user@subdomain.host.xz/path/to/repo.git/",
			"ssh://subdomain.host.xz:1234/path/to/repo.git/",
			"ssh://subdomain.host.xz/path/to/repo.git/",

			"git://subdomain.host.xz:1234/path/to/repo.git/",
			"git://subdomain.host.xz/path/to/repo.git/",

			"https://subdomain.host.xz/path/to/repo.git/",
			"https://subdomain.host.xz:1234/path/to/repo.git/",
			"http://subdomain.host.xz/path/to/repo.git/",
			"http://subdomain.host.xz:1234/path/to/repo.git/",

			"ftps://subdomain.host.xz:1234/path/to/repo.git/",
			"ftps://subdomain.host.xz/path/to/repo.git/",
			"ftp://subdomain.host.xz:1234/path/to/repo.git/",
			"ftp://subdomain.host.xz/path/to/repo.git/",

			"subdomain.host.xz:path/to/repo.git/",
			"user@subdomain.host.xz:path/to/repo.git/",
		},
	}
	for expected, urls := range testCases {
		for _, u := range urls {
			t.Run(u, func(t *testing.T) {
				assert.Equal(t, expected, destinationPath(u))
			})
		}
	}
}
