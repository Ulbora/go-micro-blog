package db

import (
	"testing"

	lg "github.com/GolangToolKits/go-level-logger"
	gdb "github.com/GolangToolKits/go-mysql"
)

func TestMyBlogDB_AddLike(t *testing.T) {

	db := gdb.MyDB{
		Host:     "localhost:3306",
		User:     "admin",
		Password: "admin",
		Database: "go_micro_blog",
	}

	// db := gdb.MyDBMock{
	// 	Host:     "localhost:3306",
	// 	User:     "admin",
	// 	Password: "admin",
	// 	Database: "go_micro_blog",
	// }
	// db.MockTestRow = &gdb.DbRow{
	// 	//Row: []string{"0"},
	// 	Row: []string{},
	// }
	// db.MockConnectSuccess = true
	// db.MockInsertID1 = 1
	// db.MockInsertSuccess1 = true

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	type fields struct {
		DB  gdb.Database
		Log lg.Log
	}
	type args struct {
		l *Like
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
				l: &Like{
					UserID: 12,
					BlogID: 1,
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
			got, got1 := d.AddLike(tt.args.l)
			if got != tt.want {
				t.Errorf("MyBlogDB.AddLike() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MyBlogDB.AddLike() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
