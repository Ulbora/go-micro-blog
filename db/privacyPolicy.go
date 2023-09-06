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

// AddPrivacyPolicy AddPrivacyPolicy
func (d *MyBlogDB) AddPrivacyPolicy(ab *PrivacyPolicy) (bool, int64) {
	var suc bool
	var id int64
	if !d.testConnection() {
		d.DB.Connect()
	}
	if ab != nil {
		var a []any
		a = append(a, ab.Content)
		suc, id = d.DB.Insert(insertPrivacyPolicy, a...)
		d.Log.Debug("suc in add PrivacyPolicy", suc)
		d.Log.Debug("id in add PrivacyPolicy", id)
	}
	return suc, id
}

// UpdatePrivacyPolicy UpdatePrivacyPolicy
func (d *MyBlogDB) UpdatePrivacyPolicy(ab *PrivacyPolicy) bool {
	var suc bool
	if !d.testConnection() {
		d.DB.Connect()
	}
	if ab != nil {
		var a []any
		a = append(a, ab.Content, ab.ID)
		suc = d.DB.Update(updatePrivacyPolicy, a...)
		d.Log.Debug("suc in update PrivacyPolicy", suc)
	}
	return suc
}

// GetPrivacyPolicy GetPrivacyPolicy
func (d *MyBlogDB) GetPrivacyPolicy() *[]PrivacyPolicy {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []PrivacyPolicy{}
	var a []any
	a = append(a)
	rows := d.DB.GetList(selectPrivacyPolicy, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parsePrivacyPolicyRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

func (d *MyBlogDB) parsePrivacyPolicyRow(foundRow *[]string) *PrivacyPolicy {
	var rtn PrivacyPolicy
	d.Log.Debug("foundRow in Tos", *foundRow)
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get PrivacyPolicy", err)
		if err == nil {
			rtn.ID = id
			rtn.Content = (*foundRow)[1]
		}
	}
	return &rtn
}
