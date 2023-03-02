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

// AddLike AddLike
func (d *MyBlogDB) AddLike(l *Like) bool {
	var suc bool
	var id int64
	if !d.testConnection() {
		d.DB.Connect()
	}
	if l != nil {
		var a []any
		a = append(a, l.UserID, l.BlogID)
		suc, id = d.DB.Insert(insertLike, a...)
		d.Log.Debug("suc in add like", suc)
		d.Log.Debug("id in add like", id)
	}
	return suc
}

// RemoveLike RemoveLike
func (d *MyBlogDB) RemoveLike(uid, bid int64) bool {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []interface{}
	a = append(a, uid, bid)

	rtn := d.DB.Delete(deleteLike, a...)
	return rtn
}

// ViewLikes ViewLikes
func (d *MyBlogDB) ViewLikes(bid int64) *[]Like {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []Like{}
	var a []any
	a = append(a, bid)
	rows := d.DB.GetList(selectLikeList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseLikeRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

func (d *MyBlogDB) parseLikeRow(foundRow *[]string) *Like {
	var rtn Like
	d.Log.Debug("foundRow in likes", *foundRow)
	if len(*foundRow) > 0 {
		uid, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get likes", err)
		if err == nil {
			bid, err := strconv.ParseInt((*foundRow)[1], 10, 64)
			if err == nil {
				rtn.UserID = uid
				rtn.BlogID = bid
			}
		}
	}
	return &rtn
}
