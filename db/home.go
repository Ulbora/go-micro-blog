package db

import (
	"strconv"
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

// AddHome AddHome
func (d *MyBlogDB) AddHome(ab *Home) (bool, int64) {
	var suc bool
	var id int64
	if !d.testConnection() {
		d.DB.Connect()
	}
	if ab != nil {
		var a []any
		a = append(a, ab.Content)
		suc, id = d.DB.Insert(insertHome, a...)
		d.Log.Debug("suc in add Home", suc)
		d.Log.Debug("id in add Home", id)
	}
	return suc, id
}

// UpdateHome UpdateHome
func (d *MyBlogDB) UpdateHome(ab *Home) bool {
	var suc bool
	if !d.testConnection() {
		d.DB.Connect()
	}
	if ab != nil {
		var a []any
		a = append(a, ab.Content, ab.ID)
		suc = d.DB.Update(updateHome, a...)
		d.Log.Debug("suc in update Home", suc)
	}
	return suc
}

// GetHome GetHome
func (d *MyBlogDB) GetHome() *[]Home {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []Home{}
	var a []any
	a = append(a)
	rows := d.DB.GetList(selectHome, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseHomeRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

func (d *MyBlogDB) parseHomeRow(foundRow *[]string) *Home {
	var rtn Home
	d.Log.Debug("foundRow in Home", *foundRow)
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get Home", err)
		if err == nil {
			rtn.ID = id
			rtn.Content = (*foundRow)[1]
		}
	}
	return &rtn
}
