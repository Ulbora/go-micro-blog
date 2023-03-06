package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ulbora/go-micro-blog/db"
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

// UpdateConfig UpdateConfig
func (h *MCHandler) UpdateConfig(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	bcOk := h.checkContent(r)
	if !bcOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var ucf db.Config
		ucs, err := h.processBody(r, &ucf)
		h.Log.Debug("ucs: ", ucs)
		h.Log.Debug("err: ", err)
		if !ucs || err != nil || !h.processAPIAdminKey(r) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			br := h.DB.UpdateConfig(&ucf)
			var res m.Response
			res.Success = br
			h.Log.Debug("br: ", br)
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

// GetConfig GetConfig
func (h *MCHandler) GetConfig(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)	
	if !h.processAPIAdminKey(r) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	} else {
		cfg := h.Manager.GetConfig()
		w.WriteHeader(http.StatusOK)
		resJSON, _ := json.Marshal(cfg)
		fmt.Fprint(w, string(resJSON))
	}
}
