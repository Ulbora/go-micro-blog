package managers

import (
	"reflect"
	"testing"

	lg "github.com/GolangToolKits/go-level-logger"
	gdb "github.com/GolangToolKits/go-mysql"
	db "github.com/Ulbora/go-micro-blog/db"
)

func TestSysManager_AddUser(t *testing.T) {

	mdb := gdb.MyDBMock{
		Host:     "localhost:3306",
		User:     "admin",
		Password: "admin",
		Database: "go_micro_blog",
	}
	mdb.MockTestRow = &gdb.DbRow{
		//Row: []string{"0"},
		Row: []string{},
	}

	mdb.MockRow1 = &gdb.DbRow{
		//Row: []string{"0"},
		Row: []string{},
	}

	mdb.MockInsertID1 = 1
	mdb.MockInsertSuccess1 = true

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	type fields struct {
		DB               db.BlogDB
		Log              lg.Log
		allowAutoPost    bool
		allowAutoComment bool
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
			name: "test 1",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
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
			m := &SysManager{
				DB:               tt.fields.DB,
				Log:              tt.fields.Log,
				allowAutoPost:    tt.fields.allowAutoPost,
				allowAutoComment: tt.fields.allowAutoComment,
			}
			if got := m.AddUser(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SysManager.AddUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSysManager_UpdateUser(t *testing.T) {

	mdb := gdb.MyDBMock{
		Host:     "localhost:3306",
		User:     "admin",
		Password: "admin",
		Database: "go_micro_blog",
	}
	mdb.MockTestRow = &gdb.DbRow{
		//Row: []string{"0"},
		Row: []string{},
	}

	mdb.MockRow1 = &gdb.DbRow{
		//Row: []string{"0"},
		Row: []string{"1", "test@test.com", "testfff", "tester", "", "2", "1"},
	}

	mdb.MockUpdateSuccess1 = true

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	type fields struct {
		DB               db.BlogDB
		Log              lg.Log
		allowAutoPost    bool
		allowAutoComment bool
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
			name: "test 1",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
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
			m := &SysManager{
				DB:               tt.fields.DB,
				Log:              tt.fields.Log,
				allowAutoPost:    tt.fields.allowAutoPost,
				allowAutoComment: tt.fields.allowAutoComment,
			}
			if got := m.UpdateUser(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SysManager.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
