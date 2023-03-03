package managers

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

import (
	db "github.com/Ulbora/go-micro-blog/db"

	lg "github.com/GolangToolKits/go-level-logger"
)

// SysManager SysManager
type SysManager struct {
	DB  db.BlogDB
	Log lg.Log
	// configRead       bool
	allowAutoPost    bool
	allowAutoComment bool
}

// New New
func (m *SysManager) New() Manager {
	cf := m.GetConfig()
	if cf != nil {
		m.allowAutoPost = cf.AllowAutoPost
		m.allowAutoComment = cf.AllowAutoComment
	} else {
		var c db.Config
		m.DB.AddConfig(&c)
	}
	return m
}
