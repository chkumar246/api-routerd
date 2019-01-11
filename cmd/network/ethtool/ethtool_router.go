// SPDX-License-Identifier: Apache-2.0

package ethtool

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func RouterConfigureEthtool(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	link := vars["link"]
	command := vars["command"]
	ethtool := Ethtool{
		Link: link,
		Action: command,
	}

	switch r.Method {
	case "GET":

		err := ethtool.GetEthTool(rw)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

	case "POST":

		e := new(Ethtool)
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		ethtool.Property = e.Property
		ethtool.Value = e.Value

		err = ethtool.SetEthTool(rw)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func RegisterRouterEthtool(n *mux.Router) {
	e := n.PathPrefix("/ethtool").Subrouter().StrictSlash(false)
	e.HandleFunc("/{link}/{command}", RouterConfigureEthtool)
}