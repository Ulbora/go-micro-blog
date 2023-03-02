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

// AddComment AddComment
func (d *MyBlogDB) AddComment(c *Comment) (bool, int64) {
	var suc bool
	var id int64
	if !d.testConnection() {
		d.DB.Connect()
	}
	if c != nil {
		var a []any
		a = append(a, c.Text, c.UserID, c.BlogID, c.Active)
		suc, id = d.DB.Insert(insertComment, a...)
		d.Log.Debug("suc in add comment", suc)
		d.Log.Debug("id in add comment", id)
	}
	return suc, id
}

// UpdateComment UpdateComment
func (d *MyBlogDB) UpdateComment(c *Comment) bool {
	var suc bool
	if !d.testConnection() {
		d.DB.Connect()
	}
	if c != nil {
		var a []any
		a = append(a, c.Text, c.ID)
		suc = d.DB.Update(updateComment, a...)
		d.Log.Debug("suc in update comment", suc)
	}
	return suc
}

// GetCommentList GetCommentList
func (d *MyBlogDB) GetCommentList(bid, start, end int64) *[]Comment {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []Comment{}
	var a []any
	a = append(a, bid, start, end)
	rows := d.DB.GetList(selectCommentList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseCommentRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

// ActivateComment ActivateComment
func (d *MyBlogDB) ActivateComment(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)

	rtn := d.DB.Update(activateComment, a...)
	return rtn
}

// DeactivateComment DeactivateComment
func (d *MyBlogDB) DeactivateComment(id int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, id)

	rtn := d.DB.Update(deactivateComment, a...)
	return rtn
}

func (d *MyBlogDB) parseCommentRow(foundRow *[]string) *Comment {
	var rtn Comment
	d.Log.Debug("foundRow in comment", *foundRow)
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get comment", err)
		if err == nil {
			uid, err := strconv.ParseInt((*foundRow)[2], 10, 64)
			if err == nil {
				bid, err := strconv.ParseInt((*foundRow)[3], 10, 64)
				if err == nil {
					active, err := strconv.ParseBool((*foundRow)[4])
					if err == nil {
						rtn.ID = id
						rtn.UserID = uid
						rtn.BlogID = bid
						rtn.Active = active
						rtn.Text = (*foundRow)[1]
					}
				}
			}
		}
	}
	return &rtn
}
