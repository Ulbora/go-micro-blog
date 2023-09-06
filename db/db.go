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

const (
	timeFormat     = "2006-01-02 15:04:05"
	dateOnlyFormat = "2006-01-02"
)

// BlogDB BlogDB
type BlogDB interface {
	//users
	AddUser(u *User) (bool, int64)
	UpdateUser(u *User) bool
	GetUser(email string) *User
	GetUserByID(id int64) *User
	GetUserList() *[]User
	GetUnactivatedUserList() *[]User
	GetBannedUserList() *[]User
	EnableUser(uid int64) bool
	DisableUser(uid int64) bool
	DisableUserForCause(uid int64) bool
	ReinstateBannedUser(uid int64) bool

	//Role
	AddRole(name string) (bool, int64)
	GetRole(name string) *Role
	GetRoleList() *[]Role
	DeleteRole(id int64) bool

	//Blog
	AddBlog(b *Blog) (bool, int64)
	UpdateBlog(u *Blog) bool
	GetBlog(id int64) *Blog
	GetBlogsByName(name string, start int64, end int64) *[]Blog
	GetBlogList(start int64, end int64) *[]Blog
	GetActiveBlogList(start int64, end int64) *[]Blog
	ActivateBlog(id int64) bool
	DeactivateBlog(id int64) bool
	DeleteBlog(id int64) bool

	//likes
	AddLike(l *Like) bool
	RemoveLike(uid, bid int64) bool
	ViewLikes(bid int64) *[]Like

	//comments
	AddComment(b *Comment) (bool, int64)
	UpdateComment(u *Comment) bool
	GetCommentList(bid, start, end int64) *[]Comment
	ActivateComment(id int64) bool
	DeactivateComment(id int64) bool

	//user auth
	AddUserAuth(b *UserAuth) (bool, int64)
	GetUserAuthList(uid, start, end int64) *[]UserAuth

	//config
	AddConfig(b *Config) (bool, int64)
	UpdateConfig(u *Config) bool
	GetConfig() *[]Config

	//rules
	AddRule(r *Rule) (bool, int64)
	UpdateRule(r *Rule) bool
	GetRule() *[]Rule

	//Tos
	AddTos(t *Tos) (bool, int64)
	UpdateTos(t *Tos) bool
	GetTos() *[]Tos

	//About
	AddAbout(ab *About) (bool, int64)
	UpdateAbout(ab *About) bool
	GetAbout() *[]About

	//Home
	AddHome(ab *Home) (bool, int64)
	UpdateHome(ab *Home) bool
	GetHome() *[]Home

	//PrivacyPolicy
	AddPrivacyPolicy(ab *PrivacyPolicy) (bool, int64)
	UpdatePrivacyPolicy(ab *PrivacyPolicy) bool
	GetPrivacyPolicy() *[]PrivacyPolicy
}
