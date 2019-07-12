package webapi

import (
	"bytes"
	"demos/lib/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type Status struct {
	Enable bool `json:"enable"`
}

type WebAPICtrl interface {
	Enable(enable bool)
	IsEnable() bool
}

func SetUpWebAPIforCommon(controller WebAPICtrl) {
	http.Handle("/api/config", util.NewCORSHandler(
		func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "POST":
				bufbody := new(bytes.Buffer)
				bufbody.ReadFrom(r.Body)
				config, err := UnmarshalConfigration(bufbody.Bytes())
				if err != nil {
					http.Error(w, "Invalid json body.", http.StatusNotFound)
				} else {
					controller.Enable(config.Enable)
				}

			default:
				http.Error(w, "Not implemented.", http.StatusNotFound)
			}
		}))
	http.Handle("/api/hello", util.NewCORSHandler(
		func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				fmt.Fprintf(w, "Hello")
			default:
				http.Error(w, "Not implemented.", http.StatusNotFound)
			}
		}))
	http.Handle("/api/status", util.NewCORSHandler(
		func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "GET":
				status := Status{controller.IsEnable()}
				jsoBytes, _ := json.Marshal(status)
				w.Write(jsoBytes)
			default:
				http.Error(w, "Not implemented.", http.StatusNotFound)
			}
		}))
}
