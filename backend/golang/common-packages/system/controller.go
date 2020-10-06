package system

import (
	log "github.com/sirupsen/logrus"
	"github.com/zenazn/goji/web"
	"net"
	"net/http"
	"reflect"
	"github.com/tb/task-logger/backend/golang/common-packages/log"
	"encoding/json"
)
type Controller struct {
}

type Application struct {
}

func GetIP(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

func GetProxyIP(r *http.Request) string {
	if ipProxy := r.Header.Get("X-FORWARDED-FOR"); len(ipProxy) > 0 {
		return ipProxy
	} else {
		return EMPTY_STRING
	}
}

func GetProxyPort(r *http.Request) string {
	if ipProxyPort := r.Header.Get("X-FORWARDED-PORT"); len(ipProxyPort) > 0 {
		return ipProxyPort
	} else {
		return EMPTY_STRING
	}
}

func GetProxyProtocol(r *http.Request) string {
	if ipProxyProtocol := r.Header.Get("X-FORWARDED-PROTO"); len(ipProxyProtocol) > 0 {
		return ipProxyProtocol
	} else {
		return EMPTY_STRING
	}
}

func (application *Application) Route(controller interface{}, route string, isPublic bool, roles []string) interface{} {

	fn := func(c web.C, w http.ResponseWriter, r *http.Request) {
		if !isPublic && c.Env["AuthFailed"].(bool) {
			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusUnauthorized)

			response := make(map[string]interface{})
			response["message"] = UnauthorisedErr.Error()
			errResponse, _ := json.Marshal(response)
			w.Write(errResponse)
		} else {
			logger := tblog.GetDefaultLogger()
			var l *log.Entry

			l = logger.WithFields(log.Fields{
				"url":       r.Host,
				"uri":       r.RequestURI,
				"userAgent": r.UserAgent(),
				"method":    r.Method,
			})

			methodValue := reflect.ValueOf(controller).MethodByName(route)
			methodInterface := methodValue.Interface()

			method := methodInterface.(func(c web.C, w http.ResponseWriter, r *http.Request, l *log.Entry) ([]byte, error))
			result, err := method(c, w, r, l)

			if c.Env["Content-Type"] != nil {
				w.Header().Set("Content-Type", c.Env["Content-Type"].(string))
			} else {
				w.Header().Set("Content-Type", "application/json")
			}

			if (err != nil) {
				response := make(map[string]interface{})

				if (IsFunctionalError(err)) {
					response["message"] = err.Error()
					w.WriteHeader(http.StatusPreconditionFailed)
				} else {
					response["message"] = InternalServerError.Error()
					w.WriteHeader(http.StatusInternalServerError)
				}

				errResponse, _ := json.Marshal(response)
				w.Write(errResponse)
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write(result)
			}

		}
	}
	return fn
}
