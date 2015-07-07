//--------------------------------------------
// web/webserv.go
//
// Handles serving and authenticating for all
// the webpages.
//
// All the handler functions are declared in
// this file.
//--------------------------------------------

package webserv

import (
	"NetworkObserver/auth"
	"html/template"
	"net/http"
)

// All URLs default to this function
func Root(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid == true {
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	} else {
		servePageStatic(w, r, "html/login.html")
	}
}

// Handles URLs referencing dashboard
func Dashboard(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid == true {
		servePageStatic(w, r, "html/dashboard.html")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// Validates a login attempt or redirects the user to
// an error page which redirects them to "Root"
func CheckLogin(w http.ResponseWriter, r *http.Request) {
	uname := r.FormValue("username")
	pword := r.FormValue("password")

	// Authenticate user credentials
	authenticated := auth.CheckCredentials(uname, pword)

	if authenticated == true {
		auth.SetSessionID(w)
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	} else {
		servePageStatic(w, r, "html/error.html")
	}
}

//--------------------------------------------
// Dashboard page handler functions
// The following four functions serve dynamic
// pages needed for the dashboard.
//
// Note: Currently static pages, will be dynamic
// later.
//--------------------------------------------
func Settings(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid == true {
		servePageStatic(w, r, "htm/dashboard/settings.html")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func Configure(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid == true {
		servePageStatic(w, r, "htm/dashboard/config.html")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func CurrentTest(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid == true {
		servePageStatic(w, r, "htm/dashboard/test.html")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func Results(w http.ResponseWriter, r *http.Request) {
	valid := auth.CheckSessionID(r)

	if valid == true {
		servePageStatic(w, r, "htm/dashboard/results.html")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

// Serves a static page
func servePageStatic(w http.ResponseWriter, r *http.Request, pageName string) {
	t, _ := template.ParseFiles(pageName)
	t.Execute(w, nil)
}