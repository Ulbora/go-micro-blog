package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	mux "github.com/GolangToolKits/grrt"
	db "github.com/Ulbora/go-micro-blog/db"
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

// AddUserAuth AddUserAuth
func (h *MCHandler) AddUserAuth(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	bcOk := h.checkContent(r)
	if !bcOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var al db.UserAuth
		as, err := h.processBody(r, &al)
		h.Log.Debug("as: ", as)
		h.Log.Debug("err: ", err)
		if !as || err != nil || !h.processAPIKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			ar := h.Manager.AddUserAuth(&al)
			h.Log.Debug("ar: ", ar)
			if ar.Success && ar.ID != 0 {
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(ar)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

// GetUserAuthList GetUserAuthList
func (h *MCHandler) GetUserAuthList(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	vars := mux.Vars(r)
	h.Log.Debug("vars: ", len(vars))
	if vars != nil && len(vars) == 3 && h.processAPIAdminKey(r) {
		var uidStr = vars["uid"]
		var stStr = vars["start"]
		var edStr = vars["end"]
		uid, iderr := strconv.ParseInt(uidStr, 10, 64)
		st, sterr := strconv.ParseInt(stStr, 10, 64)
		ed, ederr := strconv.ParseInt(edStr, 10, 64)

		if iderr == nil && sterr == nil && ederr == nil {
			blg := h.DB.GetUserAuthList(uid, st, ed)
			w.WriteHeader(http.StatusOK)
			resJSON, _ := json.Marshal(blg)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
