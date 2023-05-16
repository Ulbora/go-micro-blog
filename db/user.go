package db

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

import "strconv"

// AddUser AddUser
func (d *MyBlogDB) AddUser(u *User) (bool, int64) {
	var suc bool
	var id int64
	if !d.testConnection() {
		d.DB.Connect()
	}
	if u != nil {
		var a []any
		a = append(a, u.Email, u.Password, u.FirstName, u.LastName, u.Image, u.RoleID, u.Active)
		suc, id = d.DB.Insert(insertUser, a...)
		d.Log.Debug("suc in add user", suc)
		d.Log.Debug("id in add user", id)
	}

	return suc, id
}

// UpdateUser UpdateUser
func (d *MyBlogDB) UpdateUser(u *User) bool {
	var suc bool
	if !d.testConnection() {
		d.DB.Connect()
	}
	if u != nil {
		var a []any
		a = append(a, u.Password, u.FirstName, u.LastName, u.Image, u.ID)
		suc = d.DB.Update(updateUser, a...)
		d.Log.Debug("suc in update user", suc)
	}
	return suc
}

// GetUser GetUser
func (d *MyBlogDB) GetUser(email string) *User {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []any
	a = append(a, email)
	row := d.DB.Get(selectUser, a...)
	rtn := d.parseUserRow(&row.Row)

	return rtn
}

// GetUserByID GetUserById
func (d *MyBlogDB) GetUserByID(id int64) *User {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []any
	a = append(a, id)
	row := d.DB.Get(selectUserByID, a...)
	rtn := d.parseUserRow(&row.Row)

	return rtn
}

// GetUserList GetUserList
func (d *MyBlogDB) GetUserList() *[]User {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []User{}
	var a []any
	rows := d.DB.GetList(selectUserList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseUserRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

// GetUnactivatedUserList GetUnactivatedUserList
func (d *MyBlogDB) GetUnactivatedUserList() *[]User {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []User{}
	var a []any
	rows := d.DB.GetList(selectUnactivatedUserList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseUserRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

// GetBannedUserList GetBannedUserList
func (d *MyBlogDB) GetBannedUserList() *[]User {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []User{}
	var a []any
	rows := d.DB.GetList(selectBannedUserList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseUserRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

// EnableUser EnableUser
func (d *MyBlogDB) EnableUser(uid int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, uid)

	rtn := d.DB.Update(enableUser, a...)
	return rtn
}

// DisableUser DisableUser
func (d *MyBlogDB) DisableUser(uid int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, uid)

	rtn := d.DB.Update(disableUser, a...)
	return rtn
}

// DisableUserForCause DisableUserForCause
func (d *MyBlogDB) DisableUserForCause(uid int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, uid)

	rtn := d.DB.Update(disableUserForCause, a...)
	return rtn
}

// ReinstateBannedUser ReinstateBannedUser
func (d *MyBlogDB) ReinstateBannedUser(uid int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, uid)

	rtn := d.DB.Update(reinstateBannedUser, a...)
	return rtn
}

func (d *MyBlogDB) parseUserRow(foundRow *[]string) *User {
	var rtn User
	d.Log.Debug("foundRow in User", *foundRow)
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get user", err)
		if err == nil {
			rid, err := strconv.ParseInt((*foundRow)[5], 10, 64)
			if err == nil {
				active, err := strconv.ParseBool((*foundRow)[6])
				if err == nil {
					dc, err := strconv.ParseBool((*foundRow)[7])
					if err == nil {
						img := (*foundRow)[4]
						if img != "" {
							rtn.Image = []byte(img)
						}
						rtn.ID = id
						rtn.RoleID = rid
						rtn.Active = active
						rtn.DisabledForCause = dc
						rtn.Email = (*foundRow)[1]
						rtn.FirstName = (*foundRow)[2]
						rtn.LastName = (*foundRow)[3]
					}
				}
			}
		}
	}
	return &rtn
}
