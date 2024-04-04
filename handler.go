package main

import (
	"html/template"
	"net/http"
)

//handlers

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login", nil)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "logout", nil)
}

func accountSettingsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "account_settings", nil)
}

func updateAccountSettingsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "update_account_settings", nil)
}

func emailSettingsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "email_settings", nil)
}

func updateEmailSettingsHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "update_email_settings", nil)
}

func quotaHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "quota", nil)
}

func increaseQuotaHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "increase_quota", nil)
}

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "services", nil)
}

func wifiHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "wifi", nil)
}

func vpnHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "vpn", nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	// Parse navbar template
	navbar, err := template.ParseFiles("templates/nav.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute navbar template
	err = navbar.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse main template
	t, err := template.ParseFiles("templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute main template
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
