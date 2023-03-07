package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	mux "github.com/GolangToolKits/grrt"
	"github.com/Ulbora/go-micro-blog/db"
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

// AddLike AddLike
func (h *MCHandler) AddLike(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	bcOk := h.checkContent(r)
	if !bcOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var lk db.Like
		ls, err := h.processBody(r, &lk)
		h.Log.Debug("ls: ", ls)
		h.Log.Debug("err: ", err)
		if !ls || err != nil || !h.processAPIKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			lr := h.Manager.AddLike(&lk)
			h.Log.Debug("br: ", lr)
			if lr.Success {
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(lr)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

// RemoveLike RemoveLike
func (h *MCHandler) RemoveLike(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	vars := mux.Vars(r)
	h.Log.Debug("vars: ", len(vars))
	if vars != nil && len(vars) == 2 && h.processAPIKey(r) {
		var uidStr = vars["uid"]
		var bidStr = vars["bid"]
		uid, uerr := strconv.ParseInt(uidStr, 10, 64)
		bid, berr := strconv.ParseInt(bidStr, 10, 64)
		if uerr == nil && berr == nil {
			res := h.Manager.RemoveLike(uid, bid)
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

// ViewLikes ViewLikes
func (h *MCHandler) ViewLikes(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	vars := mux.Vars(r)
	h.Log.Debug("vars: ", len(vars))
	if vars != nil && len(vars) == 1 {
		var bidStr = vars["bid"]
		bid, biderr := strconv.ParseInt(bidStr, 10, 64)
		if biderr == nil {
			blg := h.DB.ViewLikes(bid)
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
