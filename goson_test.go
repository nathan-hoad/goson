package goson

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

type Repo struct {
	Name  string
	Url   string
	Stars int
	Forks int
}

type User struct {
	Name  string
	Repos []Repo
}

var user = &User{
	Name: "Emil Sjölander",
	Repos: []Repo{
		Repo{
			Name:  "goson",
			Url:   "https://github.com/emilsjolander/goson",
			Stars: 0,
			Forks: 0,
		},
		Repo{
			Name:  "StickyListHeaders",
			Url:   "https://github.com/emilsjolander/StickyListHeaders",
			Stars: 722,
			Forks: 197,
		},
		Repo{
			Name:  "android-FlipView",
			Url:   "https://github.com/emilsjolander/android-FlipView",
			Stars: 157,
			Forks: 47,
		},
	},
}

func BenchmarkTokenizer(b *testing.B) {
	template, _ := ioutil.ReadFile("sample/user.goson")
	for i := 0; i < b.N; i++ {
		Tokenize(template)
	}
}

func BenchmarkGosonSerialization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Render("sample/user", Args{"User": user})
	}
}

func BenchmarkStdlibSerialization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(user)
	}
}
