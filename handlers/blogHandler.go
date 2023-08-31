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

// AddBlog AddBlog
func (h *MCHandler) AddBlog(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	bcOk := h.checkContent(r)
	if !bcOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var bl db.Blog
		bs, err := h.processBody(r, &bl)
		h.Log.Debug("bs: ", bs)
		h.Log.Debug("err: ", err)
		if !bs || err != nil || !h.processAPIKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			br := h.Manager.AddBlog(&bl)
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

// UpdateBlog UpdateBlog
func (h *MCHandler) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	bcOk := h.checkContent(r)
	if !bcOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var ubl db.Blog
		ubs, err := h.processBody(r, &ubl)
		h.Log.Debug("bs: ", ubs)
		h.Log.Debug("err: ", err)
		if !ubs || err != nil || !h.processAPIKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			br := h.Manager.UpdateBlog(&ubl)
			h.Log.Debug("br: ", br)
			if br.Success {
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(br)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

// GetBlog GetBlog
func (h *MCHandler) GetBlog(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	vars := mux.Vars(r)
	h.Log.Debug("vars: ", len(vars))
	if vars != nil && len(vars) == 1 {
		var bidStr = vars["id"]
		bid, biderr := strconv.ParseInt(bidStr, 10, 64)
		if biderr == nil {
			blg := h.DB.GetBlog(bid)
			if blg != nil && blg.Active {
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(blg)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// GetBlogByName GetBlogByName
func (h *MCHandler) GetBlogByName(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	vars := mux.Vars(r)
	h.Log.Debug("vars: ", len(vars))
	if vars != nil && len(vars) == 3 {
		var name = vars["name"]
		h.Log.Debug("name: ", name)
		var stStr = vars["start"]
		var edStr = vars["end"]
		st, sterr := strconv.ParseInt(stStr, 10, 64)
		ed, ederr := strconv.ParseInt(edStr, 10, 64)

		if sterr == nil && ederr == nil {
			blg := h.Manager.GetBlogByName(name, st, ed)
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

// GetAdminBlogByName GetAdminBlogByName
func (h *MCHandler) GetAdminBlogByName(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	avars := mux.Vars(r)
	h.Log.Debug("vars: ", len(avars))
	if avars != nil && len(avars) == 3 && h.processAPIAdminKey(r) {
		var name = avars["name"]
		h.Log.Debug("name: ", name)
		var astStr = avars["start"]
		var aedStr = avars["end"]
		st, sterr := strconv.ParseInt(astStr, 10, 64)
		ed, ederr := strconv.ParseInt(aedStr, 10, 64)

		if sterr == nil && ederr == nil {
			//blg := h.Manager.GetBlogByName(name, st, ed)
			blg := h.DB.GetBlogsByName(name, st, ed)
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

// GetBlogList GetBlogList
func (h *MCHandler) GetBlogList(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	vars := mux.Vars(r)
	h.Log.Debug("vars: ", len(vars))
	if vars != nil && len(vars) == 2 {
		var stStr = vars["start"]
		var edStr = vars["end"]
		st, sterr := strconv.ParseInt(stStr, 10, 64)
		ed, ederr := strconv.ParseInt(edStr, 10, 64)

		if sterr == nil && ederr == nil {
			blg := h.DB.GetActiveBlogList(st, ed)
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

// GetAdminBlogList GetBlogList
func (h *MCHandler) GetAdminBlogList(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	vars := mux.Vars(r)
	h.Log.Debug("vars: ", len(vars))
	if vars != nil && len(vars) == 2 && h.processAPIAdminKey(r) {
		var stStr = vars["start"]
		var edStr = vars["end"]
		st, sterr := strconv.ParseInt(stStr, 10, 64)
		ed, ederr := strconv.ParseInt(edStr, 10, 64)

		if sterr == nil && ederr == nil {
			blg := h.DB.GetBlogList(st, ed)
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

// ActivateBlog ActivateBlog
func (h *MCHandler) ActivateBlog(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	bcOk := h.checkContent(r)
	if !bcOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var bl db.Blog
		bs, err := h.processBody(r, &bl)
		h.Log.Debug("bs: ", bs)
		h.Log.Debug("err: ", err)
		if !bs || err != nil || !h.processAPIAdminKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			br := h.DB.ActivateBlog(bl.ID)
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

// DectivateBlog DectivateBlog
func (h *MCHandler) DectivateBlog(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	dbcOk := h.checkContent(r)
	if !dbcOk {
		http.Error(w, "json required", http.StatusUnsupportedMediaType)
	} else {
		var dbl db.Blog
		dbs, err := h.processBody(r, &dbl)
		h.Log.Debug("dbs: ", dbs)
		h.Log.Debug("err: ", err)
		if !dbs || err != nil || !h.processAPIAdminKey(r) {
			http.Error(w, parseBodyErr, http.StatusBadRequest)
		} else {
			br := h.DB.DeactivateBlog(dbl.ID)
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

// DeleteBlog DeleteBlog
func (h *MCHandler) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	h.setContentType(w)
	vars := mux.Vars(r)
	h.Log.Debug("vars: ", len(vars))
	if vars != nil && len(vars) == 1 && h.processAPIAdminKey(r) {
		var idStr = vars["id"]
		h.Log.Debug("id: ", idStr)
		id, iderr := strconv.ParseInt(idStr, 10, 64)
		if iderr == nil {
			dblg := h.DB.DeleteBlog(id)
			var res m.Response
			res.Success = dblg
			h.Log.Debug("dblg: ", dblg)
			if res.Success {
				w.WriteHeader(http.StatusOK)
				resJSON, _ := json.Marshal(res)
				fmt.Fprint(w, string(resJSON))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
