package db

const (
	blogTest = "select count(*) from blog "

	insertUser = "insert into user (email, password, first_name, last_name, image, " +
		" role_id, active) values (?, ?, ?, ?, ?, ?, ?) "

	updateUser = " UPDATE user SET password = ?, first_name = ?, last_name = ?, " +
		" image = ? " +
		" where id = ? "
)
