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

// GetConfig GetConfig
func (m *SysManager) GetConfig() *db.Config {
	var rtn *db.Config
	cf := m.DB.GetConfig()
	if len(*cf) > 0 {
		rtn = &(*cf)[0]
	}
	return rtn
}

// // UpdateConfig UpdateConfig
// func (m *SysManager) UpdateConfig(u *db.Config) bool {


// 	return false
// }
