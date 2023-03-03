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

	selectUserByID = "SELECT id, email, first_name, last_name, image, role_id, active " +
		" from  user " +
		" where id = ? "

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

	selectBlogByName = " SELECT id, name, content, user_id, active, entered, updated " +
		" from  blog " +
		" order by name " +
		" where name like ? LIMIT ?, ?  "

	selectBlogList = " SELECT id, name, content, user_id, active, entered, updated " +
		" from  blog " +
		" order by name " +
		" LIMIT ?, ? "

	activateBlog = "UPDATE blog SET active = true " +
		"where id = ?"

	deactivateBlog = "UPDATE blog SET active = false " +
		"where id = ?"

	insertLike = "insert into likes (user_id, blog_id) " +
		" values (?, ?) "

	selectLikeList = "SELECT user_id, blog_id " +
		" from likes " +
		" where blog_id = ? "

	deleteLike = " DELETE from likes " +
		" where user_id = ? and blog_id = ? "

	insertComment = "insert into comment (text, user_id, blog_id, active) " +
		" values (?, ?, ?, ?) "

	updateComment = " UPDATE comment SET text = ? " +
		" where id = ? "

	selectCommentList = " SELECT id, text, user_id, blog_id, active " +
		" from  comment " +
		" where blog_id = ? LIMIT ?, ? "

	activateComment = "UPDATE comment SET active = true " +
		"where id = ?"

	deactivateComment = "UPDATE comment SET active = false " +
		"where id = ?"

	insertUserAugh = "insert into user_auth (type, user_id, date_entered) " +
		" values (?, ?, ?) "

	selectUserAuthList = " SELECT id, type, user_id, date_entered " +
		" from  user_auth " +
		" where user_id = ? LIMIT ?, ? "

	insertConfig = "insert into config (allow_auto_post, allow_auto_comment) " +
		" values (?, ?) "

	updateConfig = " UPDATE config SET allow_auto_post = ?, allow_auto_comment = ? " +
		" where id = ? "

	selectConfigList = " SELECT id, allow_auto_post, allow_auto_comment " +
		" from  config "
)
