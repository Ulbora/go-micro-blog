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

// AddBlog AddBlog
func (d *MyBlogDB) AddBlog(b *Blog) (bool, int64) {
	var suc bool
	var id int64
	if !d.testConnection() {
		d.DB.Connect()
	}
	if b != nil {
		var a []any
		a = append(a, b.Name, b.Content, b.UserID, b.Active, time.Now())
		suc, id = d.DB.Insert(insertBlog, a...)
		d.Log.Debug("suc in add blog", suc)
		d.Log.Debug("id in add blog", id)
	}

	return suc, id
}

// UpdateBlog UpdateBlog
func (d *MyBlogDB) UpdateBlog(b *Blog) bool {
	var suc bool
	if !d.testConnection() {
		d.DB.Connect()
	}
	if b != nil {
		var a []any
		a = append(a, b.Name, b.Content, time.Now(), b.ID)
		suc = d.DB.Update(updateBlog, a...)
		d.Log.Debug("suc in update blog", suc)
	}
	return suc
}

// GetBlog GetBlog
func (d *MyBlogDB) GetBlog(id int64) *Blog {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var a []any
	a = append(a, id)
	row := d.DB.Get(selectBlog, a...)
	rtn := d.parseBlogRow(&row.Row)

	return rtn
}

// GetBlogsByName GetBlogsByName
func (d *MyBlogDB) GetBlogsByName(name string, start int64, end int64) *[]Blog {
	return nil
}

// GetBlogList GetBlogList
func (d *MyBlogDB) GetBlogList(start int64, end int64) *[]Blog {
	return nil
}

// ActivateBlog ActivateBlog
func (d *MyBlogDB) ActivateBlog(id int64) bool {
	return false
}

// DeactivateBlog DeactivateBlog
func (d *MyBlogDB) DeactivateBlog(id int64) bool {
	return false
}

func (d *MyBlogDB) parseBlogRow(foundRow *[]string) *Blog {
	var rtn Blog
	d.Log.Debug("foundRow in blog", *foundRow)
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get blog", err)
		if err == nil {
			uid, err := strconv.ParseInt((*foundRow)[3], 10, 64)
			if err == nil {
				active, err := strconv.ParseBool((*foundRow)[4])
				if err == nil {
					eTime, _ := time.Parse(timeFormat, (*foundRow)[5])
					uTime, err := time.Parse(timeFormat, (*foundRow)[6])
					if err == nil {
						rtn.Updated = uTime
					}
					rtn.ID = id
					rtn.UserID = uid
					rtn.Active = active
					rtn.Name = (*foundRow)[1]
					rtn.Content = (*foundRow)[2]
					rtn.Entered = eTime

				}
			}
		}
	}
	return &rtn
}
