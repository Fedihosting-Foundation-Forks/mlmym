package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

var version string
var watch = flag.Bool("w", false, "watch for file changes")
var addr = flag.String("addr", ":80", "http service address")
var md goldmark.Markdown
var templates map[string]*template.Template

type AddHeaderTransport struct {
	T          http.RoundTripper
	ForwardFor string
}

func (adt *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", "Mlmym")
	if adt.ForwardFor != "" {
		req.Header.Add("X-Forwarded-For", adt.ForwardFor)
		req.Header.Add("X-Real-IP", adt.ForwardFor)
	}
	return adt.T.RoundTrip(req)
}

func NewAddHeaderTransport(remoteAddr string) *AddHeaderTransport {
	var forwardFor string
	if host, _, err := net.SplitHostPort(remoteAddr); err == nil {
		if ip := net.ParseIP(host); ip != nil {
			if !ip.IsPrivate() {
				forwardFor = ip.String()
			}
		}
	}
	return &AddHeaderTransport{
		T:          http.DefaultTransport,
		ForwardFor: forwardFor,
	}
}

func init() {
	md = goldmark.New(goldmark.WithExtensions(
		extension.Linkify,
		extension.Table,
	))
	templates = make(map[string]*template.Template)
	if !*watch {
		for _, name := range []string{"index.html", "login.html", "frontpage.html", "root.html", "settings.html", "xhr.html", "create_comment.html", "block.html"} {
			t := template.New(name).Funcs(funcMap)
			glob, err := t.ParseGlob("templates/*")
			if err != nil {
				fmt.Println(err)
				continue
			}
			templates[name] = glob
		}
	}
	if data, err := os.ReadFile("VERSION"); err == nil {
		version = string(data)
	}
}
func RemoteAddr(r *http.Request) string {
	if r.Header.Get("CF-Connecting-IP") != "" {
		return r.Header.Get("CF-Connecting-IP")
	}
	return r.RemoteAddr
}
func middleware(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		//remoteAddr := r.RemoteAddr
		//if r.Header.Get("CF-Connecting-IP") != "" {
		//	remoteAddr = r.Header.Get("CF-Connecting-IP")
		//}
		//if ps.ByName("host") != "" && !IsLemmy(ps.ByName("host"), remoteAddr) {
		//	http.Redirect(w, r, "/", 301)
		//	return
		//}
		n(w, r, ps)
	}
}
func main() {
	flag.Parse()
	log.Println("serve", *addr)
	router := GetRouter()
	err := http.ListenAndServe(*addr, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
