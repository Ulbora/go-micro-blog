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

// AddTerms AddTerms
func (h *MCHandler) AddTerms(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	tucOk := h.checkContent(r)
	if !tucOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var ti db.Tos
		ts, err := h.processBody(r, &ti)
		h.Log.Debug("bs: ", ts)
		h.Log.Debug("err: ", err)
		if !ts || err != nil || !h.processAPIAdminKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			tsuc, tid := h.DB.AddTos(&ti)
			h.Log.Debug("rsuc: ", tsuc)
			h.Log.Debug("rid: ", tid)
			if tsuc && tid != 0 {
				w.WriteHeader(http.StatusOK)
				var rres m.ResponseID
				rres.ID = tid
				rres.Success = tsuc
				resJSON, _ := json.Marshal(rres)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

// UpdateTerms UpdateTerms
func (h *MCHandler) UpdateTerms(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	tuucOk := h.checkContent(r)
	if !tuucOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var utu db.Tos
		rbs, err := h.processBody(r, &utu)
		h.Log.Debug("bs: ", rbs)
		h.Log.Debug("err: ", err)
		if !rbs || err != nil || !h.processAPIAdminKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			tsuc := h.DB.UpdateTos(&utu)
			h.Log.Debug("rsuc: ", tsuc)
			if tsuc {
				w.WriteHeader(http.StatusOK)
				var tures m.Response
				tures.Success = tsuc
				resJSON, _ := json.Marshal(tures)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

// GetTerms GetTerms
func (h *MCHandler) GetTerms(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	var rtn *db.Tos
	if !h.processAPIAdminKey(r) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	} else {
		tgs := h.DB.GetTos()
		if tgs != nil && len(*tgs) > 0 {
			rtn = &(*tgs)[0]
		} else {
			rtn = &db.Tos{}
		}
		w.WriteHeader(http.StatusOK)
		resJSON, _ := json.Marshal(rtn)
		fmt.Fprint(w, string(resJSON))
	}
}
