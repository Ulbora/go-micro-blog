package db

import "time"

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

// User User
type User struct {
	ID               int64  `json:"id"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Image            []byte `json:"image"`
	RoleID           int64  `json:"roleId"`
	Active           bool   `json:"active"`
	DisabledForCause bool   `json:"disabledForCause"`
}

// Role Role
type Role struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Blog Blog
type Blog struct {
	ID      int64     `json:"id"`
	Name    string    `json:"name"`
	Content string    `json:"content"`
	UserID  int64     `json:"userId"`
	Active  bool      `json:"active"`
	Entered time.Time `json:"entered"`
	Updated time.Time `json:"updated"`
}

// Like Like
type Like struct {
	BlogID int64 `json:"blogId"`
	UserID int64 `json:"userId"`
}

// Comment Comment
type Comment struct {
	ID     int64  `json:"id"`
	BlogID int64  `json:"blogId"`
	UserID int64  `json:"userId"`
	Text   string `json:"text"`
	Active bool   `json:"active"`
}

// UserAuth UserAuth
type UserAuth struct {
	ID       int64     `json:"id"`
	AuthType string    `json:"type"`
	UserID   int64     `json:"userId"`
	Entered  time.Time `json:"entered"`
}

// Config Config
type Config struct {
	ID               int64 `json:"id"`
	AllowAutoPost    bool  `json:"allowAutoPost"`
	AllowAutoComment bool  `json:"allowAutoComment"`
}

// Rule Rule
type Rule struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}

// Tos Tos
type Tos struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}

// About About
type About struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}
