package db

const (
	blogTest = "select count(*) from blog "

	insertUser = "insert into user (email, password, first_name, last_name, image, " +
		" role_id, active) values (?, ?, ?, ?, ?, ?, ?) "

	updateUser = " UPDATE user SET password = ?, first_name = ?, last_name = ?, " +
		" image = ? " +
		" where id = ? "

	selectUser = "SELECT id, email, first_name, last_name, image, role_id, active " +
		" from  user " +
		" where email = ? "

	selectUserList = "SELECT id, email, first_name, last_name, image, role_id, active " +
		" from  user "

	enableUser = "UPDATE user SET active = true " +
		"where id = ?"

	disableUser = "UPDATE user SET active = false " +
		"where id = ?"

	insertRole = "insert into role (name) " +
		" values (?) "

	selectRole = "SELECT id, name " +
		" from role " +
		" where name = ? "

	selectRoleList = "SELECT id, name " +
		" from role "

	deleteRole = " DELETE from role " +
		" where id = ? "

	insertBlog = "insert into blog (name, content, user_id, active, entered) " +
		" values (?, ?, ?, ?, ?) "

	updateBlog = " UPDATE blog SET name = ?, content = ?, updated = ? " +
		" where id = ? "

	selectBlog = " SELECT id, name, content, user_id, active, entered, updated " +
		" from  blog " +
		" where id = ? "
)
