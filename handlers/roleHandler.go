package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

// AddRole AddRole
func (h *MCHandler) AddRole(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	bcOk := h.checkContent(r)
	if !bcOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var rl db.Role
		rs, err := h.processBody(r, &rl)
		h.Log.Debug("rs: ", rs)
		h.Log.Debug("err: ", err)
		if !rs || err != nil || !h.processAPIAdminKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			rr := h.Manager.AddRole(rl.Name)
			h.Log.Debug("rr: ", rr)
			if rr.Success && rr.ID != 0 {
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(rr)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

// GetRole GetRole
func (h *MCHandler) GetRole(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	vars := mux.Vars(r)
	h.Log.Debug("vars: ", len(vars))
	if vars != nil && len(vars) == 1 && h.processAPIKey(r) {
		var nStr = vars["name"]
		rl := h.DB.GetRole(nStr)
		w.WriteHeader(http.StatusOK)
		resJSON, _ := json.Marshal(rl)
		fmt.Fprint(w, string(resJSON))
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// GetRoleList GetRoleList
func (h *MCHandler) GetRoleList(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	if h.processAPIKey(r) {
		rls := h.DB.GetRoleList()
		w.WriteHeader(http.StatusOK)
		resJSON, _ := json.Marshal(rls)
		fmt.Fprint(w, string(resJSON))
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// DeleteRole DeleteRole
func (h *MCHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	vars := mux.Vars(r)
	h.Log.Debug("vars: ", len(vars))
	if vars != nil && len(vars) == 1 && h.processAPIAdminKey(r) {
		var didStr = vars["id"]
		id, sterr := strconv.ParseInt(didStr, 10, 64)
		if sterr == nil {
			suc := h.DB.DeleteRole(id)
			var res m.Response
			res.Success = suc
			w.WriteHeader(http.StatusOK)
			resJSON, _ := json.Marshal(res)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
