package db

import (
	"testing"

	lg "github.com/GolangToolKits/go-level-logger"
	gdb "github.com/GolangToolKits/go-mysql"
)

func TestMyBlogDB_AddConfig(t *testing.T) {

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
		c *Config
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
				c: &Config{
					AllowAutoPost:    true,
					AllowAutoComment: true,
				},
			},
			want:  true,
			want1: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MyBlogDB{
				DB:  tt.fields.DB,
				Log: tt.fields.Log,
			}
			d.DB.Connect()
			got, got1 := d.AddConfig(tt.args.c)
			if got != tt.want {
				t.Errorf("MyBlogDB.AddConfig() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MyBlogDB.AddConfig() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMyBlogDB_UpdateConfig(t *testing.T) {

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
		c *Config
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
				c: &Config{
					AllowAutoPost:    false,
					AllowAutoComment: false,
					ID:               1,
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
			if got := d.UpdateConfig(tt.args.c); got != tt.want {
				t.Errorf("MyBlogDB.UpdateConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyBlogDB_GetConfig(t *testing.T) {

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

	db.MockRows1 = &gdb.DbRows{
		Rows: [][]string{{"1", "true", "true"},
			{"2", "true", "true"}},
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
		want   *[]Config
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
			d.DB.Connect()
			if got := d.GetConfig(); len(*got) != 2 {
				t.Errorf("MyBlogDB.GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
