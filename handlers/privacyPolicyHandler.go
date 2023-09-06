package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

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

// AddPrivacyPolicy AddPrivacyPolicy
func (h *MCHandler) AddPrivacyPolicy(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	aiOk := h.checkContent(r)
	if !aiOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var ai db.PrivacyPolicy
		ts, err := h.processBody(r, &ai)
		h.Log.Debug("bs: ", ts)
		h.Log.Debug("err: ", err)
		if !ts || err != nil || !h.processAPIAdminKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			asuc, aid := h.DB.AddPrivacyPolicy(&ai)
			h.Log.Debug("rsuc: ", asuc)
			h.Log.Debug("rid: ", aid)
			if asuc && aid != 0 {
				w.WriteHeader(http.StatusOK)
				var rres m.ResponseID
				rres.ID = aid
				rres.Success = asuc
				resJSON, _ := json.Marshal(rres)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

// UpdatePrivacyPolicy UpdatePrivacyPolicy
func (h *MCHandler) UpdatePrivacyPolicy(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	auucOk := h.checkContent(r)
	if !auucOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var atu db.PrivacyPolicy
		abs, err := h.processBody(r, &atu)
		h.Log.Debug("bs: ", abs)
		h.Log.Debug("err: ", err)
		if !abs || err != nil || !h.processAPIAdminKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			asuc := h.DB.UpdatePrivacyPolicy(&atu)
			h.Log.Debug("rsuc: ", asuc)
			if asuc {
				w.WriteHeader(http.StatusOK)
				var tures m.Response
				tures.Success = asuc
				resJSON, _ := json.Marshal(tures)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

// GetPrivacyPolicy GetPrivacyPolicy
func (h *MCHandler) GetPrivacyPolicy(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	var rtn *db.PrivacyPolicy
	if !h.processAPIAdminKey(r) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	} else {
		ags := h.DB.GetPrivacyPolicy()
		if ags != nil && len(*ags) > 0 {
			rtn = &(*ags)[0]
		} else {
			rtn = &db.PrivacyPolicy{}
		}
		w.WriteHeader(http.StatusOK)
		resJSON, _ := json.Marshal(rtn)
		fmt.Fprint(w, string(resJSON))
	}
}
