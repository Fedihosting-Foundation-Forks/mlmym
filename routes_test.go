package main

import (
	"testing"
)

func TestLinkRewrites(t *testing.T) {
	links := [][]string{
		{"https://lemmy.local/u/dude", "/lemmy.local/u/dude", "/u/dude"},
		{"https://lemmy.local/u/dude@lemmy.local", "/lemmy.local/u/dude", "/u/dude"},
		{"/u/dude", "/lemmy.local/u/dude", "/u/dude"},
		{"/u/dude@lemmy.world", "/lemmy.local/u/dude@lemmy.world", "/u/dude@lemmy.world"},
		{"/u/dude@lemmy.local", "/lemmy.local/u/dude", "/u/dude"},
		{"https://lemmy.world/c/dude", "/lemmy.local/c/dude@lemmy.world", "/c/dude@lemmy.world"},
		{"https://lemmy.world/u/dude", "/lemmy.local/u/dude@lemmy.world", "/u/dude@lemmy.world"},
		{"https://lemmy.world/u/dude@lemmy.world", "/lemmy.local/u/dude@lemmy.world", "/u/dude@lemmy.world"},
		{"https://lemmy.world/post/123", "/lemmy.local/post/123@lemmy.world", "/post/123@lemmy.world"},
		{"https://lemmy.world/post/123#123", "https://lemmy.world/post/123#123", "https://lemmy.world/post/123#123"},
		{"/post/123", "/lemmy.local/post/123", "/post/123"},
		{"/comment/123", "/lemmy.local/comment/123", "/comment/123"},
		{"https://lemmy.local/comment/123", "/lemmy.local/comment/123", "/comment/123"},
	}
	for _, url := range links {
		output := LemmyLinkRewrite(`href="`+url[0]+`"`, "lemmy.local", "")
		success := (output == (`href="` + url[1] + `"`))
		if !success {
			t.Errorf(`Failed rewriting multi instance link: url = %q, output = %q`, url, output)
		}
		output = LemmyLinkRewrite(`href="`+url[0]+`"`, ".", "lemmy.local")
		success = (output == (`href="` + url[2] + `"`))
		if !success {
			t.Errorf(`Failed rewriting single instance link: url = %q, output = %q`, url, output)
		}
	}
}
