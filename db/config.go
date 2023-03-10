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

// AddConfig AddConfig
func (d *MyBlogDB) AddConfig(c *Config) (bool, int64) {
	var suc bool
	var id int64
	if !d.testConnection() {
		d.DB.Connect()
	}
	if c != nil {
		var a []any
		a = append(a, c.AllowAutoPost, c.AllowAutoComment)
		suc, id = d.DB.Insert(insertConfig, a...)
		d.Log.Debug("suc in add config", suc)
		d.Log.Debug("id in add config", id)
	}
	return suc, id
}

// UpdateConfig UpdateConfig
func (d *MyBlogDB) UpdateConfig(c *Config) bool {
	var suc bool
	if !d.testConnection() {
		d.DB.Connect()
	}
	if c != nil {
		var a []any
		a = append(a, c.AllowAutoPost, c.AllowAutoComment, c.ID)
		suc = d.DB.Update(updateConfig, a...)
		d.Log.Debug("suc in update config", suc)
	}
	return suc
}

// GetConfig GetConfig
func (d *MyBlogDB) GetConfig() *[]Config {
	if !d.testConnection() {
		d.DB.Connect()
	}
	var rtn = []Config{}
	var a []any
	a = append(a)
	rows := d.DB.GetList(selectConfigList, a...)
	if rows != nil && len(rows.Rows) != 0 {
		foundRows := rows.Rows
		for r := range foundRows {
			foundRow := foundRows[r]
			rowContent := d.parseConfigRow(&foundRow)
			rtn = append(rtn, *rowContent)
		}
	}
	return &rtn
}

func (d *MyBlogDB) parseConfigRow(foundRow *[]string) *Config {
	var rtn Config
	d.Log.Debug("foundRow in config", *foundRow)
	if len(*foundRow) > 0 {
		id, err := strconv.ParseInt((*foundRow)[0], 10, 64)
		d.Log.Debug("id err in get config", err)
		if err == nil {
			aap, err := strconv.ParseBool((*foundRow)[1])
			if err == nil {
				aac, err := strconv.ParseBool((*foundRow)[2])
				if err == nil {
					rtn.ID = id
					rtn.AllowAutoPost = aap
					rtn.AllowAutoComment = aac
				}
			}
		}
	}
	return &rtn
}
