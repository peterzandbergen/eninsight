package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	if u == nil {
		url, _ := user.LoginURL(ctx, "/")
		fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
		return
	}
	url, _ := user.LogoutURL(ctx, "/")
	fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// if statement redirects all invalid URLs to the root homepage.
	// Ex: if URL is http://[YOUR_PROJECT_ID].appspot.com/FOO, it will be
	// redirected to http://[YOUR_PROJECT_ID].appspot.com.
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	if u == nil {
		fmt.Fprintln(w, "Not logged in.")
		return
	}

	fmt.Fprintf(w, "Logged in: %s\n", u.String())
	if u.Admin {
		fmt.Fprint(w, "is admin")
	} else {
		fmt.Fprint(w, "NO admin")
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/welcome", welcomeHandler)
	appengine.Main() // Starts the server to receive requests
}
