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

// AddComment AddComment
func (m *SysManager) AddComment(c *db.Comment) *ResponseID {
	var rtn ResponseID
	if c != nil {
		us := m.DB.GetUserByID(c.UserID)
		if us == nil || us.Active {
			b := m.DB.GetBlog(c.BlogID)
			if b.Active {
				c.Active = m.allowAutoComment
				suc, id := m.DB.AddComment(c)
				if suc {
					rtn.Success = true
					rtn.ID = id
				}
			}
		}
	}
	return &rtn
}

// UpdateComment UpdateComment
func (m *SysManager) UpdateComment(c *db.Comment) *Response {
	var rtn Response
	if c != nil {
		us := m.DB.GetUserByID(c.UserID)
		if us == nil || us.Active {
			b := m.DB.GetBlog(c.BlogID)
			if b.Active {
				suc := m.DB.UpdateComment(c)
				if suc {
					rtn.Success = true
				}
			}
		}
	}
	return &rtn
}
