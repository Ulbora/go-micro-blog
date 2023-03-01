package db

import "strconv"

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

// AddRole AddRole
func (d *MyBlogDB) AddRole(name string) (bool, int64) {
	var suc bool
	var id int64
	if !d.testConnection() {
		d.DB.Connect()
	}
	if name != "" {
		var a []any
		a = append(a, name)
		suc, id = d.DB.Insert(insertRole, a...)
		d.Log.Debug("suc in add user", suc)
		d.Log.Debug("id in add user", id)
	}
	return suc, id
}

// GetRole GetRole
func (d *MyBlogDB) GetRole(name string) *Role {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []any
	a = append(a, name)
	row := d.DB.Get(selectRole, a...)
	rtn := d.parseRoleRow(&row.Row)

	return rtn
}

// GetRoleList GetRoleList
func (d *MyBlogDB) GetRoleList() *[]Role {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []Role{}
	var a []any
	rows := d.DB.GetList(selectRoleList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseRoleRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

// DeleteRole DeleteRole
func (d *MyBlogDB) DeleteRole(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)

	rtn := d.DB.Delete(deleteRole, a...)
	return rtn
}

func (d *MyBlogDB) parseRoleRow(foundRow *[]string) *Role {
	var rtn Role
	d.Log.Debug("foundRow in Role", *foundRow)
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get role", err)
		if err == nil {
			rtn.ID = id
			rtn.Name = (*foundRow)[1]
		}
	}
	return &rtn
}
