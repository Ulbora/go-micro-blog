package managers

import (
	db "github.com/Ulbora/go-micro-blog/db"
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

// Manager Manager
type Manager interface {
	// reads config from db or creates default if none
	AddUser(u *db.User) *ResponseID  //logic
	UpdateUser(u *db.User) *Response //logic
	// //h--GetUser(email string) *db.User
	// //h--GetUserList() *[]db.User
	// //h--EnableUser(uid int64) bool
	// //h--DisableUser(uid int64) bool

	AddRole(name string) *ResponseID //logic
	// //h--GetRole(name string) *db.Role
	// //h--GetRoleList() *[]db.Role
	// //h--DeleteRole(id int64) bool

	AddBlog(b *db.Blog) *ResponseID  //logic
	UpdateBlog(b *db.Blog) *Response //logic
	// //h--GetBlog(id int64) *db.Blog
	// //h--GetBlogsByName(name string, start int64, end int64) *[]db.Blog
	// //h--GetBlogList(start int64, end int64) *[]db.Blog
	// //h--ActivateBlog(id int64) bool
	// //h--DeactivateBlog(id int64) bool

	AddLike(l *db.Like) *Response        //logic
	RemoveLike(uid, bid int64) *Response //logic
	// //h--ViewLikes(bid int64) *[]db.Like

	AddComment(c *db.Comment) *ResponseID //logic
	UpdateComment(c *db.Comment) *Response    //logic
	// //h--GetCommentList(bid, start, end int64) *[]db.Comment
	// //h--ActivateComment(id int64) bool
	// //h--DeactivateComment(id int64) bool

	AddUserAuth(a *db.UserAuth) *ResponseID //logic
	// //h--GetUserAuthList(bid, start, end int64) *[]db.UserAuth

	// //no--AddConfig(b *db.Config) *ResponseID //logoc
	// //h--UpdateConfig(u *db.Config) bool //logic
	GetConfig() *db.Config
}
