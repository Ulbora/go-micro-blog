package handlers

import "net/http"

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
	parseBodyErr = "Failed to process body"
)

// Handler Handler
type Handler interface {
	AddUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	GetUserList(w http.ResponseWriter, r *http.Request)
	GetUnactivatedUserList(w http.ResponseWriter, r *http.Request)
	GetBannedUserList(w http.ResponseWriter, r *http.Request)
	EnableUser(w http.ResponseWriter, r *http.Request)
	DisableUser(w http.ResponseWriter, r *http.Request)
	DisableUserForCause(w http.ResponseWriter, r *http.Request)
	ReinstateBannedUser(w http.ResponseWriter, r *http.Request)

	AddRole(w http.ResponseWriter, r *http.Request)
	GetRole(w http.ResponseWriter, r *http.Request)
	GetRoleList(w http.ResponseWriter, r *http.Request)
	DeleteRole(w http.ResponseWriter, r *http.Request)

	AddBlog(w http.ResponseWriter, r *http.Request)
	UpdateBlog(w http.ResponseWriter, r *http.Request)
	GetBlog(w http.ResponseWriter, r *http.Request)
	GetBlogByName(w http.ResponseWriter, r *http.Request)
	GetBlogList(w http.ResponseWriter, r *http.Request)
	GetAdminBlogList(w http.ResponseWriter, r *http.Request)
	ActivateBlog(w http.ResponseWriter, r *http.Request)
	DectivateBlog(w http.ResponseWriter, r *http.Request)
	DeleteBlog(w http.ResponseWriter, r *http.Request)

	AddLike(w http.ResponseWriter, r *http.Request)
	RemoveLike(w http.ResponseWriter, r *http.Request)
	ViewLikes(w http.ResponseWriter, r *http.Request)

	AddComment(w http.ResponseWriter, r *http.Request)
	UpdateComment(w http.ResponseWriter, r *http.Request)
	GetCommentList(w http.ResponseWriter, r *http.Request)
	GetCommentAdminList(w http.ResponseWriter, r *http.Request)
	ActivateComment(w http.ResponseWriter, r *http.Request)
	DectivateComment(w http.ResponseWriter, r *http.Request)

	AddUserAuth(w http.ResponseWriter, r *http.Request)
	GetUserAuthList(w http.ResponseWriter, r *http.Request)

	UpdateConfig(w http.ResponseWriter, r *http.Request)
	GetConfig(w http.ResponseWriter, r *http.Request)

	//GetRules(w http.ResponseWriter, r *http.Request)
	//SetRules(w http.ResponseWriter, r *http.Request)

	//GetTerms(w http.ResponseWriter, r *http.Request)
	//SetTerms(w http.ResponseWriter, r *http.Request)

	//GetAbout(w http.ResponseWriter, r *http.Request)
	//SetAbout((w http.ResponseWriter, r *http.Request))

	SetLogLevel(w http.ResponseWriter, r *http.Request)
}
