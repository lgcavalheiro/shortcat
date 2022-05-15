package server

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
)

var (
	templates                            *template.Template
	admUser, admPwd, admSecret, reqError string
)

func RoutesSetup(router *httprouter.Router, useDefaultFront bool) {
	router.GET("/go/:shorturl", handleShortcat)
	if useDefaultFront {
		admUser = viper.GetString("adm_user")
		admPwd = viper.GetString("adm_pwd")
		admSecret = viper.GetString("adm_secret")
		if len(admUser) == 0 || len(admPwd) == 0 || len(admSecret) == 0 {
			log.Fatal("[ROUTES SETUP] - Administrator credentials were not properly provided.")
		}

		templates = template.Must(template.New("").ParseGlob("public/html/*.html"))
		router.GET("/", handleRootGet)
		router.GET("/login", handleLoginGet)
		router.POST("/login", handleLoginPost)
		router.GET("/admin", handleAdminGet)
		router.POST("/admin", handleAdminGet)
		router.GET("/logout", handleLogoutGet)
	}
}

func handleShortcat(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	shortUrl := params.ByName("shorturl")

	redirect := resolveShortCat(shortUrl)
	if len(redirect.Error) > 0 {
		res := []byte(redirect.Error)
		w.WriteHeader(redirect.Status)
		w.Write(res)
		return
	}

	http.Redirect(w, r, redirect.Redirect, redirect.Status)
}

func handleRootGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}

func handleLoginGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := map[string]string{
		"Error": reqError,
	}
	templates.ExecuteTemplate(w, "login.html", data)
}

func handleLoginPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	user := r.Form.Get("user")
	pwd := r.Form.Get("pwd")

	redirect := resolveLogin(user, pwd)
	if len(redirect.Error) > 0 {
		reqError = redirect.Error
		handleLoginGet(w, r, nil)
		return
	}

	setSession(w)

	reqError = ""
	http.Redirect(w, r, redirect.Redirect, redirect.Status)
}

func handleAdminGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	checkLogin(w, r)
	data := resolveAdminGet()
	templates.ExecuteTemplate(w, "admin.html", data)
}

func handleLogoutGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	clearSession(w)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}

func setSession(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:  "shortcat-session",
		Value: admSecret,
		Path:  "/",
	})
}

func clearSession(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   "shortcat-session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}

func checkLogin(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("shortcat-session")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	if cookie.Value != admSecret {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
}
