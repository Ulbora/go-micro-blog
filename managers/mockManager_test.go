package managers

import (
	"reflect"
	"testing"
	"time"

	db "github.com/Ulbora/go-micro-blog/db"
)

func TestMockManager_AddUser(t *testing.T) {
	type fields struct {
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
	type args struct {
		u *db.User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ResponseID
	}{
		// TODO: Add test cases.
		{
			name: "test add user",
			fields: fields{
				MockAddUser: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateUser: Response{
					Success: true,
				},
				MockAddRole: ResponseID{
					ID:      1,
					Success: true,
				},
				MockAddBlog: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateBlog: Response{
					Success: true,
				},
				MockAddLike: Response{
					Success: true,
				},
				MockRemoveLike: Response{
					Success: true,
				},
				MockAddComment: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateComment: Response{
					Success: true,
				},
				MockAddUserAuth: ResponseID{
					ID:      1,
					Success: true,
				},
				MockConfig: db.Config{
					ID: 1,
				},
			},
			args: args{
				u: &db.User{
					Email: "test@test.com",
				},
			},
			want: &ResponseID{
				ID:      1,
				Success: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockManager{
				MockAddUser:       tt.fields.MockAddUser,
				MockUpdateUser:    tt.fields.MockUpdateUser,
				MockAddRole:       tt.fields.MockAddRole,
				MockAddBlog:       tt.fields.MockAddBlog,
				MockUpdateBlog:    tt.fields.MockUpdateBlog,
				MockAddLike:       tt.fields.MockAddLike,
				MockRemoveLike:    tt.fields.MockRemoveLike,
				MockAddComment:    tt.fields.MockAddComment,
				MockUpdateComment: tt.fields.MockUpdateComment,
				MockAddUserAuth:   tt.fields.MockAddUserAuth,
				MockConfig:        tt.fields.MockConfig,
			}
			if got := m.AddUser(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockManager.AddUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockManager_UpdateUser(t *testing.T) {
	type fields struct {
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
	type args struct {
		u *db.User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		// TODO: Add test cases.
		{
			name: "test add user",
			fields: fields{
				MockAddUser: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateUser: Response{
					Success: true,
				},
				MockAddRole: ResponseID{
					ID:      1,
					Success: true,
				},
				MockAddBlog: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateBlog: Response{
					Success: true,
				},
				MockAddLike: Response{
					Success: true,
				},
				MockRemoveLike: Response{
					Success: true,
				},
				MockAddComment: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateComment: Response{
					Success: true,
				},
				MockAddUserAuth: ResponseID{
					ID:      1,
					Success: true,
				},
				MockConfig: db.Config{
					ID: 1,
				},
			},
			args: args{
				u: &db.User{
					Email: "test@test.com",
				},
			},
			want: &Response{
				Success: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockManager{
				MockAddUser:       tt.fields.MockAddUser,
				MockUpdateUser:    tt.fields.MockUpdateUser,
				MockAddRole:       tt.fields.MockAddRole,
				MockAddBlog:       tt.fields.MockAddBlog,
				MockUpdateBlog:    tt.fields.MockUpdateBlog,
				MockAddLike:       tt.fields.MockAddLike,
				MockRemoveLike:    tt.fields.MockRemoveLike,
				MockAddComment:    tt.fields.MockAddComment,
				MockUpdateComment: tt.fields.MockUpdateComment,
				MockAddUserAuth:   tt.fields.MockAddUserAuth,
				MockConfig:        tt.fields.MockConfig,
			}
			if got := m.UpdateUser(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockManager.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockManager_AddRole(t *testing.T) {
	type fields struct {
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
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ResponseID
	}{
		// TODO: Add test cases.
		{
			name: "test add user",
			fields: fields{
				MockAddUser: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateUser: Response{
					Success: true,
				},
				MockAddRole: ResponseID{
					ID:      1,
					Success: true,
				},
				MockAddBlog: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateBlog: Response{
					Success: true,
				},
				MockAddLike: Response{
					Success: true,
				},
				MockRemoveLike: Response{
					Success: true,
				},
				MockAddComment: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateComment: Response{
					Success: true,
				},
				MockAddUserAuth: ResponseID{
					ID:      1,
					Success: true,
				},
				MockConfig: db.Config{
					ID: 1,
				},
			},
			args: args{
				name: "test",
			},
			want: &ResponseID{
				ID:      1,
				Success: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockManager{
				MockAddUser:       tt.fields.MockAddUser,
				MockUpdateUser:    tt.fields.MockUpdateUser,
				MockAddRole:       tt.fields.MockAddRole,
				MockAddBlog:       tt.fields.MockAddBlog,
				MockUpdateBlog:    tt.fields.MockUpdateBlog,
				MockAddLike:       tt.fields.MockAddLike,
				MockRemoveLike:    tt.fields.MockRemoveLike,
				MockAddComment:    tt.fields.MockAddComment,
				MockUpdateComment: tt.fields.MockUpdateComment,
				MockAddUserAuth:   tt.fields.MockAddUserAuth,
				MockConfig:        tt.fields.MockConfig,
			}
			if got := m.AddRole(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockManager.AddRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockManager_AddBlog(t *testing.T) {
	type fields struct {
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
	type args struct {
		b *db.Blog
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ResponseID
	}{
		// TODO: Add test cases.
		{
			name: "test add user",
			fields: fields{
				MockAddUser: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateUser: Response{
					Success: true,
				},
				MockAddRole: ResponseID{
					ID:      1,
					Success: true,
				},
				MockAddBlog: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateBlog: Response{
					Success: true,
				},
				MockAddLike: Response{
					Success: true,
				},
				MockRemoveLike: Response{
					Success: true,
				},
				MockAddComment: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateComment: Response{
					Success: true,
				},
				MockAddUserAuth: ResponseID{
					ID:      1,
					Success: true,
				},
				MockConfig: db.Config{
					ID: 1,
				},
			},
			args: args{
				b: &db.Blog{
					Name: "test",
				},
			},
			want: &ResponseID{
				ID:      1,
				Success: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockManager{
				MockAddUser:       tt.fields.MockAddUser,
				MockUpdateUser:    tt.fields.MockUpdateUser,
				MockAddRole:       tt.fields.MockAddRole,
				MockAddBlog:       tt.fields.MockAddBlog,
				MockUpdateBlog:    tt.fields.MockUpdateBlog,
				MockAddLike:       tt.fields.MockAddLike,
				MockRemoveLike:    tt.fields.MockRemoveLike,
				MockAddComment:    tt.fields.MockAddComment,
				MockUpdateComment: tt.fields.MockUpdateComment,
				MockAddUserAuth:   tt.fields.MockAddUserAuth,
				MockConfig:        tt.fields.MockConfig,
			}
			if got := m.AddBlog(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockManager.AddBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockManager_UpdateBlog(t *testing.T) {
	type fields struct {
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
	type args struct {
		b *db.Blog
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		// TODO: Add test cases.
		{
			name: "test add user",
			fields: fields{
				MockAddUser: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateUser: Response{
					Success: true,
				},
				MockAddRole: ResponseID{
					ID:      1,
					Success: true,
				},
				MockAddBlog: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateBlog: Response{
					Success: true,
				},
				MockAddLike: Response{
					Success: true,
				},
				MockRemoveLike: Response{
					Success: true,
				},
				MockAddComment: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateComment: Response{
					Success: true,
				},
				MockAddUserAuth: ResponseID{
					ID:      1,
					Success: true,
				},
				MockConfig: db.Config{
					ID: 1,
				},
			},
			args: args{
				b: &db.Blog{
					Name: "test",
				},
			},
			want: &Response{

				Success: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockManager{
				MockAddUser:       tt.fields.MockAddUser,
				MockUpdateUser:    tt.fields.MockUpdateUser,
				MockAddRole:       tt.fields.MockAddRole,
				MockAddBlog:       tt.fields.MockAddBlog,
				MockUpdateBlog:    tt.fields.MockUpdateBlog,
				MockAddLike:       tt.fields.MockAddLike,
				MockRemoveLike:    tt.fields.MockRemoveLike,
				MockAddComment:    tt.fields.MockAddComment,
				MockUpdateComment: tt.fields.MockUpdateComment,
				MockAddUserAuth:   tt.fields.MockAddUserAuth,
				MockConfig:        tt.fields.MockConfig,
			}
			if got := m.UpdateBlog(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockManager.UpdateBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockManager_AddLike(t *testing.T) {
	type fields struct {
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
	type args struct {
		l *db.Like
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		// TODO: Add test cases.
		{
			name: "test add user",
			fields: fields{
				MockAddUser: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateUser: Response{
					Success: true,
				},
				MockAddRole: ResponseID{
					ID:      1,
					Success: true,
				},
				MockAddBlog: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateBlog: Response{
					Success: true,
				},
				MockAddLike: Response{
					Success: true,
				},
				MockRemoveLike: Response{
					Success: true,
				},
				MockAddComment: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateComment: Response{
					Success: true,
				},
				MockAddUserAuth: ResponseID{
					ID:      1,
					Success: true,
				},
				MockConfig: db.Config{
					ID: 1,
				},
			},
			args: args{
				l: &db.Like{
					UserID: 1,
				},
			},
			want: &Response{
				Success: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockManager{
				MockAddUser:       tt.fields.MockAddUser,
				MockUpdateUser:    tt.fields.MockUpdateUser,
				MockAddRole:       tt.fields.MockAddRole,
				MockAddBlog:       tt.fields.MockAddBlog,
				MockUpdateBlog:    tt.fields.MockUpdateBlog,
				MockAddLike:       tt.fields.MockAddLike,
				MockRemoveLike:    tt.fields.MockRemoveLike,
				MockAddComment:    tt.fields.MockAddComment,
				MockUpdateComment: tt.fields.MockUpdateComment,
				MockAddUserAuth:   tt.fields.MockAddUserAuth,
				MockConfig:        tt.fields.MockConfig,
			}
			if got := m.AddLike(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockManager.AddLike() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockManager_RemoveLike(t *testing.T) {
	type fields struct {
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
	type args struct {
		uid int64
		bid int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		// TODO: Add test cases.
		{
			name: "test add user",
			fields: fields{
				MockAddUser: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateUser: Response{
					Success: true,
				},
				MockAddRole: ResponseID{
					ID:      1,
					Success: true,
				},
				MockAddBlog: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateBlog: Response{
					Success: true,
				},
				MockAddLike: Response{
					Success: true,
				},
				MockRemoveLike: Response{
					Success: true,
				},
				MockAddComment: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateComment: Response{
					Success: true,
				},
				MockAddUserAuth: ResponseID{
					ID:      1,
					Success: true,
				},
				MockConfig: db.Config{
					ID: 1,
				},
			},
			args: args{
				uid: 1,
				bid: 2,
			},
			want: &Response{
				Success: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockManager{
				MockAddUser:       tt.fields.MockAddUser,
				MockUpdateUser:    tt.fields.MockUpdateUser,
				MockAddRole:       tt.fields.MockAddRole,
				MockAddBlog:       tt.fields.MockAddBlog,
				MockUpdateBlog:    tt.fields.MockUpdateBlog,
				MockAddLike:       tt.fields.MockAddLike,
				MockRemoveLike:    tt.fields.MockRemoveLike,
				MockAddComment:    tt.fields.MockAddComment,
				MockUpdateComment: tt.fields.MockUpdateComment,
				MockAddUserAuth:   tt.fields.MockAddUserAuth,
				MockConfig:        tt.fields.MockConfig,
			}
			if got := m.RemoveLike(tt.args.uid, tt.args.bid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockManager.RemoveLike() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockManager_AddComment(t *testing.T) {
	type fields struct {
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
	type args struct {
		c *db.Comment
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ResponseID
	}{
		// TODO: Add test cases.
		{
			name: "test add user",
			fields: fields{
				MockAddUser: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateUser: Response{
					Success: true,
				},
				MockAddRole: ResponseID{
					ID:      1,
					Success: true,
				},
				MockAddBlog: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateBlog: Response{
					Success: true,
				},
				MockAddLike: Response{
					Success: true,
				},
				MockRemoveLike: Response{
					Success: true,
				},
				MockAddComment: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateComment: Response{
					Success: true,
				},
				MockAddUserAuth: ResponseID{
					ID:      1,
					Success: true,
				},
				MockConfig: db.Config{
					ID: 1,
				},
			},
			args: args{
				c: &db.Comment{
					UserID: 1,
				},
			},
			want: &ResponseID{
				ID:      1,
				Success: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockManager{
				MockAddUser:       tt.fields.MockAddUser,
				MockUpdateUser:    tt.fields.MockUpdateUser,
				MockAddRole:       tt.fields.MockAddRole,
				MockAddBlog:       tt.fields.MockAddBlog,
				MockUpdateBlog:    tt.fields.MockUpdateBlog,
				MockAddLike:       tt.fields.MockAddLike,
				MockRemoveLike:    tt.fields.MockRemoveLike,
				MockAddComment:    tt.fields.MockAddComment,
				MockUpdateComment: tt.fields.MockUpdateComment,
				MockAddUserAuth:   tt.fields.MockAddUserAuth,
				MockConfig:        tt.fields.MockConfig,
			}
			if got := m.AddComment(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockManager.AddComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockManager_UpdateComment(t *testing.T) {
	type fields struct {
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
	type args struct {
		c *db.Comment
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		// TODO: Add test cases.
		{
			name: "test add user",
			fields: fields{
				MockAddUser: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateUser: Response{
					Success: true,
				},
				MockAddRole: ResponseID{
					ID:      1,
					Success: true,
				},
				MockAddBlog: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateBlog: Response{
					Success: true,
				},
				MockAddLike: Response{
					Success: true,
				},
				MockRemoveLike: Response{
					Success: true,
				},
				MockAddComment: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateComment: Response{
					Success: true,
				},
				MockAddUserAuth: ResponseID{
					ID:      1,
					Success: true,
				},
				MockConfig: db.Config{
					ID: 1,
				},
			},
			args: args{
				c: &db.Comment{
					UserID: 1,
				},
			},
			want: &Response{
				Success: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockManager{
				MockAddUser:       tt.fields.MockAddUser,
				MockUpdateUser:    tt.fields.MockUpdateUser,
				MockAddRole:       tt.fields.MockAddRole,
				MockAddBlog:       tt.fields.MockAddBlog,
				MockUpdateBlog:    tt.fields.MockUpdateBlog,
				MockAddLike:       tt.fields.MockAddLike,
				MockRemoveLike:    tt.fields.MockRemoveLike,
				MockAddComment:    tt.fields.MockAddComment,
				MockUpdateComment: tt.fields.MockUpdateComment,
				MockAddUserAuth:   tt.fields.MockAddUserAuth,
				MockConfig:        tt.fields.MockConfig,
			}
			if got := m.UpdateComment(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockManager.UpdateComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockManager_AddUserAuth(t *testing.T) {
	type fields struct {
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
	type args struct {
		a *db.UserAuth
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ResponseID
	}{
		// TODO: Add test cases.
		{
			name: "test add user",
			fields: fields{
				MockAddUser: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateUser: Response{
					Success: true,
				},
				MockAddRole: ResponseID{
					ID:      1,
					Success: true,
				},
				MockAddBlog: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateBlog: Response{
					Success: true,
				},
				MockAddLike: Response{
					Success: true,
				},
				MockRemoveLike: Response{
					Success: true,
				},
				MockAddComment: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateComment: Response{
					Success: true,
				},
				MockAddUserAuth: ResponseID{
					ID:      1,
					Success: true,
				},
				MockConfig: db.Config{
					ID: 1,
				},
			},
			args: args{
				a: &db.UserAuth{
					UserID: 1,
				},
			},
			want: &ResponseID{
				ID:      1,
				Success: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockManager{
				MockAddUser:       tt.fields.MockAddUser,
				MockUpdateUser:    tt.fields.MockUpdateUser,
				MockAddRole:       tt.fields.MockAddRole,
				MockAddBlog:       tt.fields.MockAddBlog,
				MockUpdateBlog:    tt.fields.MockUpdateBlog,
				MockAddLike:       tt.fields.MockAddLike,
				MockRemoveLike:    tt.fields.MockRemoveLike,
				MockAddComment:    tt.fields.MockAddComment,
				MockUpdateComment: tt.fields.MockUpdateComment,
				MockAddUserAuth:   tt.fields.MockAddUserAuth,
				MockConfig:        tt.fields.MockConfig,
			}
			if got := m.AddUserAuth(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockManager.AddUserAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockManager_GetConfig(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name   string
		fields fields
		want   *db.Config
	}{
		// TODO: Add test cases.
		{
			name: "test add user",
			fields: fields{
				MockAddUser: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateUser: Response{
					Success: true,
				},
				MockAddRole: ResponseID{
					ID:      1,
					Success: true,
				},
				MockAddBlog: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateBlog: Response{
					Success: true,
				},
				MockAddLike: Response{
					Success: true,
				},
				MockRemoveLike: Response{
					Success: true,
				},
				MockAddComment: ResponseID{
					ID:      1,
					Success: true,
				},
				MockUpdateComment: Response{
					Success: true,
				},
				MockAddUserAuth: ResponseID{
					ID:      1,
					Success: true,
				},
				MockConfig: db.Config{
					ID: 1,
				},
			},
			want: &db.Config{
				ID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockManager{
				MockAddUser:       tt.fields.MockAddUser,
				MockUpdateUser:    tt.fields.MockUpdateUser,
				MockAddRole:       tt.fields.MockAddRole,
				MockAddBlog:       tt.fields.MockAddBlog,
				MockUpdateBlog:    tt.fields.MockUpdateBlog,
				MockAddLike:       tt.fields.MockAddLike,
				MockRemoveLike:    tt.fields.MockRemoveLike,
				MockAddComment:    tt.fields.MockAddComment,
				MockUpdateComment: tt.fields.MockUpdateComment,
				MockAddUserAuth:   tt.fields.MockAddUserAuth,
				MockConfig:        tt.fields.MockConfig,
			}
			if got := m.GetConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockManager.GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockManager_GetBlogList(t *testing.T) {

	var bl = []db.Blog{{ID: 1, Name: "test blog entry", Content: "some test blog stuff", UserID: 4, Active: true, Entered: time.Now(), Updated: time.Now()},
		{ID: 2, Name: "test blog entry 333", Content: "some test blog stuff", UserID: 4, Active: false, Entered: time.Now(), Updated: time.Now()}}

	type fields struct {
		MockAddUser       ResponseID
		MockUpdateUser    Response
		MockAddRole       ResponseID
		MockAddBlog       ResponseID
		MockUpdateBlog    Response
		MockBlogList      []db.Blog
		MockAddLike       Response
		MockRemoveLike    Response
		MockAddComment    ResponseID
		MockUpdateComment Response
		MockAddUserAuth   ResponseID
		MockConfig        db.Config
	}
	type args struct {
		start int64
		end   int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]db.Blog
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockBlogList: bl,
			},
			args: args{
				start: 0,
				end:   5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockManager{
				MockAddUser:       tt.fields.MockAddUser,
				MockUpdateUser:    tt.fields.MockUpdateUser,
				MockAddRole:       tt.fields.MockAddRole,
				MockAddBlog:       tt.fields.MockAddBlog,
				MockUpdateBlog:    tt.fields.MockUpdateBlog,
				MockBlogList:      tt.fields.MockBlogList,
				MockAddLike:       tt.fields.MockAddLike,
				MockRemoveLike:    tt.fields.MockRemoveLike,
				MockAddComment:    tt.fields.MockAddComment,
				MockUpdateComment: tt.fields.MockUpdateComment,
				MockAddUserAuth:   tt.fields.MockAddUserAuth,
				MockConfig:        tt.fields.MockConfig,
			}
			if got := m.GetBlogList(tt.args.start, tt.args.end); len(*got) != 2 {
				t.Errorf("MockManager.GetBlogList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockManager_GetBlogByName(t *testing.T) {

	var bl = []db.Blog{{ID: 1, Name: "test blog entry", Content: "some test blog stuff", UserID: 4, Active: true, Entered: time.Now(), Updated: time.Now()},
		{ID: 2, Name: "test blog entry 333", Content: "some test blog stuff", UserID: 4, Active: false, Entered: time.Now(), Updated: time.Now()}}

	type fields struct {
		MockAddUser       ResponseID
		MockUpdateUser    Response
		MockAddRole       ResponseID
		MockAddBlog       ResponseID
		MockUpdateBlog    Response
		MockBlogList      []db.Blog
		MockAddLike       Response
		MockRemoveLike    Response
		MockAddComment    ResponseID
		MockUpdateComment Response
		MockAddUserAuth   ResponseID
		MockConfig        db.Config
	}
	type args struct {
		name  string
		start int64
		end   int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]db.Blog
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				MockBlogList: bl,
			},
			args: args{
				name:  "test",
				start: 0,
				end:   5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockManager{
				MockAddUser:       tt.fields.MockAddUser,
				MockUpdateUser:    tt.fields.MockUpdateUser,
				MockAddRole:       tt.fields.MockAddRole,
				MockAddBlog:       tt.fields.MockAddBlog,
				MockUpdateBlog:    tt.fields.MockUpdateBlog,
				MockBlogList:      tt.fields.MockBlogList,
				MockAddLike:       tt.fields.MockAddLike,
				MockRemoveLike:    tt.fields.MockRemoveLike,
				MockAddComment:    tt.fields.MockAddComment,
				MockUpdateComment: tt.fields.MockUpdateComment,
				MockAddUserAuth:   tt.fields.MockAddUserAuth,
				MockConfig:        tt.fields.MockConfig,
			}
			if got := m.GetBlogByName(tt.args.name, tt.args.start, tt.args.end); len(*got) != 2 {
				t.Errorf("MockManager.GetBlogByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockManager_New(t *testing.T) {
	type fields struct {
		MockAddUser       ResponseID
		MockUpdateUser    Response
		MockAddRole       ResponseID
		MockAddBlog       ResponseID
		MockUpdateBlog    Response
		MockBlogList      []db.Blog
		MockAddLike       Response
		MockRemoveLike    Response
		MockAddComment    ResponseID
		MockUpdateComment Response
		MockAddUserAuth   ResponseID
		MockConfig        db.Config
	}
	tests := []struct {
		name   string
		fields fields
		want   Manager
	}{
		// TODO: Add test cases.
		{
			name: "test 1",

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockManager{
				MockAddUser:       tt.fields.MockAddUser,
				MockUpdateUser:    tt.fields.MockUpdateUser,
				MockAddRole:       tt.fields.MockAddRole,
				MockAddBlog:       tt.fields.MockAddBlog,
				MockUpdateBlog:    tt.fields.MockUpdateBlog,
				MockBlogList:      tt.fields.MockBlogList,
				MockAddLike:       tt.fields.MockAddLike,
				MockRemoveLike:    tt.fields.MockRemoveLike,
				MockAddComment:    tt.fields.MockAddComment,
				MockUpdateComment: tt.fields.MockUpdateComment,
				MockAddUserAuth:   tt.fields.MockAddUserAuth,
				MockConfig:        tt.fields.MockConfig,
			}
			if got := m.New(); got == nil {
				t.Errorf("MockManager.New() = %v, want %v", got, tt.want)
			}
		})
	}
}
