package db

import (
	"testing"

	lg "github.com/GolangToolKits/go-level-logger"
	gdb "github.com/GolangToolKits/go-mysql"
)

func TestMyBlogDB_testConnection(t *testing.T) {
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
		Row: []string{"0"},
	}
	db.MockConnectSuccess = true

	//test 2
	db2 := gdb.MyDBMock{
		Host:     "localhost:3306",
		User:     "admin",
		Password: "admin",
		Database: "go_micro_blog",
	}
	db2.MockTestRow = &gdb.DbRow{
		Row: []string{"a"},
	}
	db2.MockConnectSuccess = true

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)
	type fields struct {
		DB  gdb.Database
		Log lg.Log
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				DB:  db.New(),
				Log: log,
			},
			want: true,
		},
		{
			name: "test 2",
			fields: fields{
				DB:  db2.New(),
				Log: log,
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
			if got := d.testConnection(); got != tt.want {
				t.Errorf("MyBlogDB.testConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyBlogDB_New(t *testing.T) {
	db := gdb.MyDBMock{
		Host:     "localhost:3306",
		User:     "admin",
		Password: "admin",
		Database: "go_micro_blog",
	}
	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	type fields struct {
		DB  gdb.Database
		Log lg.Log
	}
	tests := []struct {
		name   string
		fields fields
		want   BlogDB
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				DB:  db.New(),
				Log: log,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MyBlogDB{
				DB:  tt.fields.DB,
				Log: tt.fields.Log,
			}
			if got := d.New(); got == nil {
				t.Errorf("MyBlogDB.New() = %v, want %v", got, tt.want)
			}
		})
	}
}
