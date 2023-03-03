package db

import (
	"strconv"
	"time"
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

// AddUserAuth AddUserAuth
func (d *MyBlogDB) AddUserAuth(c *UserAuth) (bool, int64) {
	var suc bool
	var id int64
	if !d.testConnection() {
		d.DB.Connect()
	}
	if c != nil {
		var a []any
		a = append(a, c.AuthType, c.UserID, time.Now())
		suc, id = d.DB.Insert(insertUserAugh, a...)
		d.Log.Debug("suc in add userAuth", suc)
		d.Log.Debug("id in add userAuth", id)
	}
	return suc, id
}

// GetUserAuthList GetUserAuthList
func (d *MyBlogDB) GetUserAuthList(uid, start, end int64) *[]UserAuth {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []UserAuth{}
	var a []any
	a = append(a, uid, start, end)
	rows := d.DB.GetList(selectUserAuthList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseUserAuthRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

func (d *MyBlogDB) parseUserAuthRow(foundRow *[]string) *UserAuth {
	var rtn UserAuth
	d.Log.Debug("foundRow in UserAuth", *foundRow)
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get userAuth", err)
		if err == nil {
			uid, err := strconv.ParseInt((*foundRow)[2], 10, 64)
			if err == nil {
				eTime, _ := time.Parse(timeFormat, (*foundRow)[3])
				rtn.ID = id
				rtn.UserID = uid
				rtn.AuthType = (*foundRow)[1]
				rtn.Entered = eTime
			}
		}
	}
	return &rtn
}
