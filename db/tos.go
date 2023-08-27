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

// AddTos AddTos
func (d *MyBlogDB) AddTos(t *Tos) (bool, int64) {
	var suc bool
	var id int64
	if !d.testConnection() {
		d.DB.Connect()
	}
	if t != nil {
		var a []any
		a = append(a, t.Content)
		suc, id = d.DB.Insert(insertTos, a...)
		d.Log.Debug("suc in add config", suc)
		d.Log.Debug("id in add config", id)
	}
	return suc, id
}

// UpdateTos UpdateTos
func (d *MyBlogDB) UpdateTos(t *Tos) bool {
	var suc bool
	if !d.testConnection() {
		d.DB.Connect()
	}
	if t != nil {
		var a []any
		a = append(a, t.Content, t.ID)
		suc = d.DB.Update(updateTos, a...)
		d.Log.Debug("suc in update config", suc)
	}
	return suc
}

// GetTos GetTos
func (d *MyBlogDB) GetTos() *[]Tos {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []Tos{}
	var a []any
	a = append(a)
	rows := d.DB.GetList(selectTos, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseTosRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

func (d *MyBlogDB) parseTosRow(foundRow *[]string) *Tos {
	var rtn Tos
	d.Log.Debug("foundRow in Tos", *foundRow)
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get config", err)
		if err == nil {
			rtn.ID = id
			rtn.Content = (*foundRow)[1]
		}
	}
	return &rtn
}
