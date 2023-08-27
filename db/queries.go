package db

const (
	blogTest = "select count(*) from blog "

	insertUser = "insert into user (email, password, first_name, last_name, image, " +
		" role_id, active) values (?, ?, ?, ?, ?, ?, ?) "

	updateUser = " UPDATE user SET password = ?, first_name = ?, last_name = ?, " +
		" image = ? " +
		" where id = ? "

	selectUser = "SELECT id, email, first_name, last_name, image, role_id, active, " +
		" disabled_for_cause " +
		" from  user " +
		" where email = ? "

	selectUserByID = "SELECT id, email, first_name, last_name, image, role_id, active, " +
		" disabled_for_cause " +
		" from  user " +
		" where id = ? "

	selectUserList = "SELECT id, email, first_name, last_name, image, role_id, active, " +
		" disabled_for_cause " +
		" from  user "

	selectUnactivatedUserList = "SELECT id, email, first_name, last_name, image, role_id, active, " +
		" disabled_for_cause " +
		" from  user " +
		" where active = false and disabled_for_cause = false "

	selectBannedUserList = "SELECT id, email, first_name, last_name, image, role_id, active, " +
		" disabled_for_cause " +
		" from  user " +
		" where active = false and disabled_for_cause = true "

	enableUser = "UPDATE user SET active = true " +
		"where id = ?"

	disableUser = "UPDATE user SET active = false " +
		"where id = ?"

	disableUserForCause = "UPDATE user SET active = false, disabled_for_cause = true " +
		"where id = ?"

	reinstateBannedUser = "UPDATE user SET active = true, disabled_for_cause = false " +
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
		// " order by name " +
		" where name like ? " +
		" order by entered desc " +
		" LIMIT ?, ?  "

	selectActiveBlogList = " SELECT b.id, b.name, b.content, b.user_id, b.active, b.entered, b.updated " +
		" from  blog b	" +
		" inner join user u " +
		" on b.user_id = u.id " +
		" where u.active and b.active " +
		" order by b.entered desc " +
		" LIMIT ?, ? "

	selectBlogList = " SELECT id, name, content, user_id, active, entered, updated " +
		" from  blog " +
		// " order by name " +
		" order by entered desc" +
		" LIMIT ?, ? "

	activateBlog = "UPDATE blog SET active = true " +
		"where id = ?"

	deactivateBlog = "UPDATE blog SET active = false " +
		"where id = ?"

	deleteBlog = "DELETE from blog " +
		" where id = ? "

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

	insertRule = "insert into rules (content) " +
		" values (?) "

	updateRule = " UPDATE rules SET content = ? " +
		" where id = ? "

	selectRules = " SELECT id, content " +
		" from  rules "

	insertTos = "insert into tos (content) " +
		" values (?) "

	updateTos = " UPDATE tos SET content = ? " +
		" where id = ? "

	selectTos = " SELECT id, content " +
		" from  tos "

	insertAbout = "insert into about (content) " +
		" values (?) "

	updateAbout = " UPDATE about SET content = ? " +
		" where id = ? "

	selectAbout = " SELECT id, content " +
		" from  about "
)
