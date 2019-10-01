package giturls

import (
	"net/url"
	"reflect"
	"testing"
)

var tests []*Test

type Test struct {
	in      string
	wantURL *url.URL
	wantStr string // expected result of reserializing the URL; empty means same as "in".
}

func NewTest(in, transport, user, host, path, str string) *Test {
	var userinfo *url.Userinfo
	if user != "" {
		userinfo = url.User(user)
	}
	if str == "" {
		str = in
	}

	return &Test{
		in: in,
		wantURL: &url.URL{
			Scheme: transport,
			Host:   host,
			Path:   path,
			User:   userinfo,
		},
		wantStr: str,
	}
}

func init() {
	// https://www.kernel.org/pub/software/scm/git/docs/git-clone.html
	tests = []*Test{
		NewTest(
			"user@host.xz:path/to/repo.git/",
			"ssh", "user", "host.xz", "path/to/repo.git/",
			"ssh://user@host.xz/path/to/repo.git/",
		),
		NewTest(
			"host.xz:path/to/repo.git/",
			"ssh", "", "host.xz", "path/to/repo.git/",
			"ssh://host.xz/path/to/repo.git/",
		),
		NewTest(
			"host.xz:/path/to/repo.git/",
			"ssh", "", "host.xz", "/path/to/repo.git/",
			"ssh://host.xz/path/to/repo.git/",
		),
		NewTest(
			"host.xz:path/to/repo-with_specials.git/",
			"ssh", "", "host.xz", "path/to/repo-with_specials.git/",
			"ssh://host.xz/path/to/repo-with_specials.git/",
		),
		NewTest(
			"git://host.xz/path/to/repo.git/",
			"git", "", "host.xz", "/path/to/repo.git/",
			"",
		),
		NewTest(
			"git://host.xz:1234/path/to/repo.git/",
			"git", "", "host.xz:1234", "/path/to/repo.git/",
			"",
		),
		NewTest(
			"http://host.xz/path/to/repo.git/",
			"http", "", "host.xz", "/path/to/repo.git/",
			"",
		),
		NewTest(
			"http://host.xz:1234/path/to/repo.git/",
			"http", "", "host.xz:1234", "/path/to/repo.git/",
			"",
		),
		NewTest(
			"https://host.xz/path/to/repo.git/",
			"https", "", "host.xz", "/path/to/repo.git/",
			"",
		),
		NewTest(
			"https://host.xz:1234/path/to/repo.git/",
			"https", "", "host.xz:1234", "/path/to/repo.git/",
			"",
		),
		NewTest(
			"ftp://host.xz/path/to/repo.git/",
			"ftp", "", "host.xz", "/path/to/repo.git/",
			"",
		),
		NewTest(
			"ftp://host.xz:1234/path/to/repo.git/",
			"ftp", "", "host.xz:1234", "/path/to/repo.git/",
			"",
		),
		NewTest(
			"ftps://host.xz/path/to/repo.git/",
			"ftps", "", "host.xz", "/path/to/repo.git/",
			"",
		),
		NewTest(
			"ftps://host.xz:1234/path/to/repo.git/",
			"ftps", "", "host.xz:1234", "/path/to/repo.git/",
			"",
		),
		NewTest(
			"rsync://host.xz/path/to/repo.git/",
			"rsync", "", "host.xz", "/path/to/repo.git/",
			"",
		),
		NewTest(
			"ssh://user@host.xz:1234/path/to/repo.git/",
			"ssh", "user", "host.xz:1234", "/path/to/repo.git/",
			"",
		),
		NewTest(
			"ssh://host.xz:1234/path/to/repo.git/",
			"ssh", "", "host.xz:1234", "/path/to/repo.git/",
			"",
		),
		NewTest(
			"ssh://host.xz/path/to/repo.git/",
			"ssh", "", "host.xz", "/path/to/repo.git/",
			"",
		),
		NewTest(
			"git+ssh://host.xz/path/to/repo.git/",
			"git+ssh", "", "host.xz", "/path/to/repo.git/",
			"",
		),
		NewTest(
			"/path/to/repo.git/",
			"file", "", "", "/path/to/repo.git/",
			"file:///path/to/repo.git/",
		),
		NewTest(
			"file:///path/to/repo.git/",
			"file", "", "", "/path/to/repo.git/",
			"",
		),
	}
}

func TestParse(t *testing.T) {
	for _, tt := range tests {
		got, err := Parse(tt.in)
		if err != nil {
			t.Errorf("Parse(%q) = unexpected err %q, want %q", tt.in, err, tt.wantURL)
			continue
		}
		if !reflect.DeepEqual(got, tt.wantURL) {
			t.Errorf("Parse(%q) = %q, want %q", tt.in, got, tt.wantURL)
		}
		str := got.String()
		if str != tt.wantStr {
			t.Errorf("Parse(%q).String() = %q, want %q", tt.in, str, tt.wantStr)
		}
	}
}
