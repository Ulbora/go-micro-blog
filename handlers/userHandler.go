package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	mux "github.com/GolangToolKits/grrt"
	db "github.com/Ulbora/go-micro-blog/db"
	m "github.com/Ulbora/go-micro-blog/managers"
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

// AddUser AddUser
func (h *MCHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	bcOk := h.checkContent(r)
	if !bcOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var bl db.User
		bs, err := h.processBody(r, &bl)
		h.Log.Debug("bs: ", bs)
		h.Log.Debug("err: ", err)
		if !bs || err != nil || !h.processAPIKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			br := h.Manager.AddUser(&bl)
			h.Log.Debug("br: ", br)
			if br.Success && br.ID != 0 {
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(br)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

// UpdateUser UpdateUser
func (h *MCHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	bcOk := h.checkContent(r)
	if !bcOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var ubl db.User
		ubs, err := h.processBody(r, &ubl)
		h.Log.Debug("bs: ", ubs)
		h.Log.Debug("err: ", err)
		if !ubs || err != nil || !h.processAPIKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			ur := h.Manager.UpdateUser(&ubl)
			h.Log.Debug("br: ", ur)
			if ur.Success {
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(ur)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

// GetUser GetUser
func (h *MCHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	vars := mux.Vars(r)
	h.Log.Debug("vars: ", len(vars))
	if vars != nil && len(vars) == 1 && h.processAPIKey(r) {
		var em = vars["email"]
		usr := h.DB.GetUser(em)
		w.WriteHeader(http.StatusOK)
		resJSON, _ := json.Marshal(usr)
		fmt.Fprint(w, string(resJSON))
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// GetUserList GetUserList
func (h *MCHandler) GetUserList(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	if h.processAPIAdminKey(r) {
		blg := h.DB.GetUserList()
		w.WriteHeader(http.StatusOK)
		resJSON, _ := json.Marshal(blg)
		fmt.Fprint(w, string(resJSON))
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// EnableUser EnableUser
func (h *MCHandler) EnableUser(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	bcOk := h.checkContent(r)
	if !bcOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var ubl db.User
		ubs, err := h.processBody(r, &ubl)
		h.Log.Debug("bs: ", ubs)
		h.Log.Debug("err: ", err)
		if !ubs || err != nil || !h.processAPIAdminKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			ur := h.DB.EnableUser(ubl.ID)
			var res m.Response
			res.Success = ur
			h.Log.Debug("br: ", ur)
			if res.Success {
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(res)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

// DisableUser DisableUser
func (h *MCHandler) DisableUser(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	bcOk := h.checkContent(r)
	if !bcOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var ubl db.User
		ubs, err := h.processBody(r, &ubl)
		h.Log.Debug("bs: ", ubs)
		h.Log.Debug("err: ", err)
		if !ubs || err != nil || !h.processAPIAdminKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			ur := h.DB.DisableUser(ubl.ID)
			var res m.Response
			res.Success = ur
			h.Log.Debug("br: ", ur)
			if res.Success {
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(res)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}
