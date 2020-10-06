package system

import (
	log "github.com/sirupsen/logrus"
	"github.com/zenazn/goji/web"
	"net/http"
)

func (application *Application) ApplyAuth(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		log.Println("Applied authorization filter!")

		emailId := r.URL.Query().Get("emailId")

		if (emailId != "") {
			c.Env["AuthFailed"] = false
			c.Env["emailId"] = emailId
		} else {
			c.Env["AuthFailed"] = true
			c.Env["emailId"] = ""
		}

		h.ServeHTTP(w, r)

	}
	return http.HandlerFunc(fn)
}
