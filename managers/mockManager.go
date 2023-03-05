package managers

import (
	db "github.com/Ulbora/go-micro-blog/db"
	//lg "github.com/GolangToolKits/go-level-logger"
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

// MockManager MockManager
type MockManager struct {
	MockAddUser       ResponseID
	MockUpdateUser    Response
	MockAddRole       ResponseID
	MockAddBlog       ResponseID
	MockUpdateBlog    Response
	MockAddLike       Response
	MockRemoveLike    Response
	MockAddComment    ResponseID
	MockUpdateComment Response
	MockAddUserAuth   ResponseID
	MockConfig        db.Config
}

// New New
func (m *MockManager) New() Manager {
	return m
}

// AddUser AddUser
func (m *MockManager) AddUser(u *db.User) *ResponseID {
	return &m.MockAddUser
}

// UpdateUser UpdateUser
func (m *MockManager) UpdateUser(u *db.User) *Response {
	return &m.MockUpdateUser
}

// AddRole AddRole
func (m *MockManager) AddRole(name string) *ResponseID {
	return &m.MockAddRole
}

// AddBlog AddBlog
func (m *MockManager) AddBlog(b *db.Blog) *ResponseID {
	return &m.MockAddBlog
}

// UpdateBlog UpdateBlog
func (m *MockManager) UpdateBlog(b *db.Blog) *Response {
	return &m.MockUpdateBlog
}

// AddLike AddLike
func (m *MockManager) AddLike(l *db.Like) *Response {
	return &m.MockAddLike
}

// RemoveLike RemoveLike
func (m *MockManager) RemoveLike(uid, bid int64) *Response {
	return &m.MockRemoveLike
}

// //h--ViewLikes(bid int64) *[]db.Like

// AddComment AddComment
func (m *MockManager) AddComment(c *db.Comment) *ResponseID {
	return &m.MockAddComment
}

// UpdateComment UpdateComment
func (m *MockManager) UpdateComment(c *db.Comment) *Response {
	return &m.MockUpdateComment
}

// AddUserAuth AddUserAuth
func (m *MockManager) AddUserAuth(a *db.UserAuth) *ResponseID {
	return &m.MockAddUserAuth
}

// GetConfig GetConfig
func (m *MockManager) GetConfig() *db.Config {
	return &m.MockConfig
}
