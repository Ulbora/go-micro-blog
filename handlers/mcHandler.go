package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	db "github.com/Ulbora/go-micro-blog/db"
	m "github.com/Ulbora/go-micro-blog/managers"

	lg "github.com/GolangToolKits/go-level-logger"
)

/*
 Copyright (C) 2023 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2023 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.

*/

// MCHandler MCHandler
type MCHandler struct {
	DB          db.BlogDB
	Log         lg.Log
	Manager     m.Manager
	APIKey      string
	APIAdminKey string
}

// New New
func (h *MCHandler) New() Handler {
	return h
}

func (h *MCHandler) processAPIKey(r *http.Request) bool {
	var rtn bool
	apiKey := r.Header.Get("apiKey")
	h.Log.Debug("apiKey: ", apiKey)
	h.Log.Debug("h.APIKey: ", h.APIKey)
	if apiKey == h.APIKey {
		rtn = true
	}
	return rtn
}

func (h *MCHandler) processAPIAdminKey(r *http.Request) bool {
	var rtn bool
	apiAdminKey := r.Header.Get("apiAdminKey")
	h.Log.Debug("apiAdminKey: ", apiAdminKey)
	h.Log.Debug("h.APIAdminKey: ", h.APIAdminKey)
	if apiAdminKey == h.APIAdminKey {
		rtn = true
	}
	return rtn
}

// CheckContent CheckContent
func (h *MCHandler) checkContent(r *http.Request) bool {
	var rtn bool
	cType := r.Header.Get("Content-Type")
	if cType == "application/json" {
		rtn = true
	}
	return rtn
}

// SetContentType SetContentType
func (h *MCHandler) setContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// ProcessBody ProcessBody
func (h *MCHandler) processBody(r *http.Request, obj any) (bool, error) {
	var suc bool
	var err error
	//fmt.Println("r.Body: ", r.Body)
	h.Log.Debug("r.Body: ", r.Body)
	if r.Body != nil {
		decoder := json.NewDecoder(r.Body)
		//fmt.Println("decoder: ", decoder)
		err = decoder.Decode(obj)
		//fmt.Println("decoder: ", decoder)
		if err != nil {
			//log.Println("Decode Error: ", err.Error())
			h.Log.Error("Decode Error: ", err.Error())
		} else {
			suc = true
		}
	} else {
		err = errors.New("Bad Body")
	}
	return suc, err
}
