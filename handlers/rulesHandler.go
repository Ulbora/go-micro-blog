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

// AddRule AddRule
func (h *MCHandler) AddRule(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	rucOk := h.checkContent(r)
	if !rucOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var ru db.Rule
		bs, err := h.processBody(r, &ru)
		h.Log.Debug("bs: ", bs)
		h.Log.Debug("err: ", err)
		if !bs || err != nil || !h.processAPIAdminKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			rsuc, rid := h.DB.AddRule(&ru)
			h.Log.Debug("rsuc: ", rsuc)
			h.Log.Debug("rid: ", rid)
			if rsuc && rid != 0 {
				w.WriteHeader(http.StatusOK)
				var rres m.ResponseID
				rres.ID = rid
				rres.Success = rsuc
				resJSON, _ := json.Marshal(rres)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

// UpdateRule UpdateRule
func (h *MCHandler) UpdateRule(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	ruucOk := h.checkContent(r)
	if !ruucOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var uru db.Rule
		rbs, err := h.processBody(r, &uru)
		h.Log.Debug("bs: ", rbs)
		h.Log.Debug("err: ", err)
		if !rbs || err != nil || !h.processAPIAdminKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			rsuc := h.DB.UpdateRule(&uru)
			h.Log.Debug("rsuc: ", rsuc)
			if rsuc {
				w.WriteHeader(http.StatusOK)
				var rures m.Response
				rures.Success = rsuc
				resJSON, _ := json.Marshal(rures)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

// GetRules GetRules
func (h *MCHandler) GetRules(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	var rtn *db.Rule
	if !h.processAPIAdminKey(r) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	} else {
		cfg := h.DB.GetRule()
		if cfg != nil && len(*cfg) > 0 {
			rtn = &(*cfg)[0]
		} else {
			rtn = &db.Rule{}
		}
		w.WriteHeader(http.StatusOK)
		resJSON, _ := json.Marshal(rtn)
		fmt.Fprint(w, string(resJSON))
	}
}
