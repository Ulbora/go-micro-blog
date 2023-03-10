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

// AddComment AddComment
func (h *MCHandler) AddComment(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	bcOk := h.checkContent(r)
	if !bcOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var cm db.Comment
		cs, err := h.processBody(r, &cm)
		h.Log.Debug("cs: ", cs)
		h.Log.Debug("err: ", err)
		if !cs || err != nil || !h.processAPIKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			cr := h.Manager.AddComment(&cm)
			h.Log.Debug("cr: ", cr)
			if cr.Success && cr.ID != 0 {
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(cr)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

// UpdateComment UpdateComment
func (h *MCHandler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	bcOk := h.checkContent(r)
	if !bcOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var cm db.Comment
		cms, err := h.processBody(r, &cm)
		h.Log.Debug("bs: ", cms)
		h.Log.Debug("err: ", err)
		if !cms || err != nil || !h.processAPIKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			cr := h.Manager.UpdateComment(&cm)
			h.Log.Debug("cr: ", cr)
			if cr.Success {
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(cr)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

// GetCommentList GetCommentList
func (h *MCHandler) GetCommentList(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	vars := mux.Vars(r)
	h.Log.Debug("vars: ", len(vars))
	if vars != nil && len(vars) == 3 {
		var bidStr = vars["bid"]
		var stStr = vars["start"]
		var edStr = vars["end"]
		bid, berr := strconv.ParseInt(bidStr, 10, 64)
		st, sterr := strconv.ParseInt(stStr, 10, 64)
		ed, ederr := strconv.ParseInt(edStr, 10, 64)

		if berr == nil && sterr == nil && ederr == nil {
			clg := h.Manager.GetCommentList(bid, st, ed)
			w.WriteHeader(http.StatusOK)
			resJSON, _ := json.Marshal(clg)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// GetCommentAdminList GetCommentList
func (h *MCHandler) GetCommentAdminList(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	vars := mux.Vars(r)
	h.Log.Debug("vars: ", len(vars))
	if vars != nil && len(vars) == 3 && h.processAPIAdminKey(r) {
		var bidStr = vars["bid"]
		var stStr = vars["start"]
		var edStr = vars["end"]
		bid, berr := strconv.ParseInt(bidStr, 10, 64)
		st, sterr := strconv.ParseInt(stStr, 10, 64)
		ed, ederr := strconv.ParseInt(edStr, 10, 64)

		if berr == nil && sterr == nil && ederr == nil {
			clg := h.DB.GetCommentList(bid, st, ed)
			w.WriteHeader(http.StatusOK)
			resJSON, _ := json.Marshal(clg)
			fmt.Fprint(w, string(resJSON))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// ActivateComment ActivateComment
func (h *MCHandler) ActivateComment(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	bcOk := h.checkContent(r)
	if !bcOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var cm db.Comment
		cms, err := h.processBody(r, &cm)
		h.Log.Debug("cm: ", cm)
		h.Log.Debug("bs: ", cms)
		h.Log.Debug("err: ", err)
		if !cms || err != nil || !h.processAPIAdminKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			cr := h.DB.ActivateComment(cm.ID)
			h.Log.Debug("cr: ", cr)
			var res m.Response
			res.Success = cr
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

// DectivateComment DectivateComment
func (h *MCHandler) DectivateComment(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	bcOk := h.checkContent(r)
	if !bcOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var cm db.Comment
		cms, err := h.processBody(r, &cm)
		h.Log.Debug("bs: ", cms)
		h.Log.Debug("err: ", err)
		if !cms || err != nil || !h.processAPIAdminKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			cr := h.DB.DeactivateComment(cm.ID)
			h.Log.Debug("cr: ", cr)
			var res m.Response
			res.Success = cr
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
