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

// AddLike AddLike
func (m *SysManager) AddLike(l *db.Like) *Response {
	var rtn Response
	if l != nil {
		us := m.DB.GetUserByID(l.UserID)
		if us == nil || us.Active {
			suc := m.DB.AddLike(l)
			if suc {
				rtn.Success = true
			}
		}
	}
	return &rtn
}

// RemoveLike RemoveLike
func (m *SysManager) RemoveLike(uid, bid int64) *Response {
	var rtn Response
	if uid != 0 && bid != 0 {
		us := m.DB.GetUserByID(uid)
		if us == nil || us.Active {
			suc := m.DB.RemoveLike(uid, bid)
			if suc {
				rtn.Success = true
			}
		}
	}
	return &rtn
}
