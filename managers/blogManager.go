package managers

import (
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

// AddBlog AddBlog
func (m *SysManager) AddBlog(b *db.Blog) *ResponseID {
	var rtn ResponseID
	if b != nil {
		us := m.DB.GetUserByID(b.UserID)
		if us == nil || us.Active {
			b.Active = m.allowAutoPost
			suc, id := m.DB.AddBlog(b)
			if suc {
				rtn.Success = true
				rtn.ID = id
			}
		}
	}
	return &rtn
}

// UpdateBlog UpdateBlog
func (m *SysManager) UpdateBlog(b *db.Blog) *Response {
	var rtn Response
	if b != nil {
		us := m.DB.GetUserByID(b.UserID)
		if us == nil || us.Active {
			suc := m.DB.UpdateBlog(b)
			if suc {
				rtn.Success = true
			}
		}
	}
	return &rtn
}

// // GetBlogList GetBlogList
// func (m *SysManager) GetBlogList(start int64, end int64) *[]db.Blog {
// 	var rtn = []db.Blog{}
// 	bl := m.DB.GetBlogList(start, end)
// 	if bl != nil {
// 		for i := range *bl {
// 			if (*bl)[i].Active {
// 				rtn = append(rtn, (*bl)[i])
// 			}
// 		}
// 	}
// 	return &rtn
// }

// GetBlogByName GetBlogByName
func (m *SysManager) GetBlogByName(name string, start int64, end int64) *[]db.Blog {
	var rtn = []db.Blog{}
	bl := m.DB.GetBlogsByName(name, start, end)
	if bl != nil {
		for i := range *bl {
			if (*bl)[i].Active {
				rtn = append(rtn, (*bl)[i])
			}
		}
	}
	return &rtn
}
