package db

// AddUser AddUser
func (d *MyBlogDB) AddUser(u *User) (bool, int64) {
	var suc bool
	var id int64
	if !d.testConnection() {
		d.DB.Connect()
	}
	if u != nil {
		var a []any
		a = append(a, u.Email, u.Password, u.FirstName, u.LastName, u.Image, u.RoleID, u.Active)
		suc, id = d.DB.Insert(insertUser, a...)
		d.Log.Debug("suc in add user", suc)
		d.Log.Debug("id in add user", id)
	}

	return suc, id
}

// UpdateUser UpdateUser
func (d *MyBlogDB) UpdateUser(u *User) bool {
	var suc bool
	if !d.testConnection() {
		d.DB.Connect()
	}
	if u != nil {
		var a []any
		a = append(a, u.Password, u.FirstName, u.LastName, u.Image, u.ID)
		suc = d.DB.Update(updateUser, a...)
		d.Log.Debug("suc in update user", suc)
	}
	return suc
}
