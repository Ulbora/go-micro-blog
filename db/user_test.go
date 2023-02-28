package db

import (
	"testing"

	lg "github.com/GolangToolKits/go-level-logger"
	gdb "github.com/GolangToolKits/go-mysql"
)

func TestMyBlogDB_AddUser(t *testing.T) {

	// db := gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "admin",
	// 	Password: "admin",
	// 	Database: "go_micro_blog",
	// }

	db := gdb.MyDBMock{
		Host:     "localhost:3306",
		User:     "admin",
		Password: "admin",
		Database: "go_micro_blog",
	}
	db.MockTestRow = &gdb.DbRow{
		//Row: []string{"0"},
		Row: []string{},
	}
	db.MockConnectSuccess = true
	db.MockInsertID1 = 1
	db.MockInsertSuccess1 = true

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	type fields struct {
		DB  gdb.Database
		Log lg.Log
	}
	type args struct {
		u *User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  int64
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				DB:  db.New(),
				Log: log,
			},
			args: args{
				u: &User{
					Email:     "test@test.com",
					Password:  "tst",
					FirstName: "test",
					LastName:  "tester",
					RoleID:    1,
					Active:    true,
				},
			},
			want:  true,
			want1: 1,
		},
		// {
		// 	name: "test 2",
		// 	fields: fields{
		// 		DB:  db.New(),
		// 		Log: log,
		// 	},
		// 	args: args{
		// 		u: &User{
		// 			Email:     "test@test.com",
		// 			Password:  "tst",
		// 			FirstName: "test",
		// 			LastName:  "tester",
		// 			RoleID:    1,
		// 			Active:    true,
		// 		},
		// 	},
		// 	want: true,
		// 	want1: 1,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MyBlogDB{
				DB:  tt.fields.DB,
				Log: tt.fields.Log,
			}
			d.DB.Connect()
			got, got1 := d.AddUser(tt.args.u)
			if got != tt.want {
				t.Errorf("MyBlogDB.AddUser() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MyBlogDB.AddUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMyBlogDB_UpdateUser(t *testing.T) {

	// db := gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "admin",
	// 	Password: "admin",
	// 	Database: "go_micro_blog",
	// }

	db := gdb.MyDBMock{
		Host:     "localhost:3306",
		User:     "admin",
		Password: "admin",
		Database: "go_micro_blog",
	}
	db.MockTestRow = &gdb.DbRow{
		//Row: []string{"0"},
		Row: []string{},
	}
	db.MockConnectSuccess = true
	
	db.MockUpdateSuccess1 = true

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	type fields struct {
		DB  gdb.Database
		Log lg.Log
	}
	type args struct {
		u *User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				DB:  db.New(),
				Log: log,
			},
			args: args{
				u: &User{
					ID:        12,
					Password:  "tsterttt",
					FirstName: "testfff",
					LastName:  "testerfff",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MyBlogDB{
				DB:  tt.fields.DB,
				Log: tt.fields.Log,
			}
			d.DB.Connect()
			if got := d.UpdateUser(tt.args.u); got != tt.want {
				t.Errorf("MyBlogDB.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
