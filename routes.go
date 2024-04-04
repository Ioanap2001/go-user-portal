package main

import "net/http"

func setupRoutes() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/account/settings", accountSettingsHandler)
	http.HandleFunc("/account/settings/update", updateAccountSettingsHandler)
	http.HandleFunc("/email/settings", emailSettingsHandler)
	http.HandleFunc("/email/settings/update", updateEmailSettingsHandler)
	http.HandleFunc("/quota", quotaHandler)
	http.HandleFunc("/quota/increase", increaseQuotaHandler)
	http.HandleFunc("/services", servicesHandler)
	http.HandleFunc("/services/wifi", wifiHandler)
	http.HandleFunc("/services/vpn", vpnHandler)
}
