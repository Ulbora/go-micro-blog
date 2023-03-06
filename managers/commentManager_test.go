package managers

import (
	"reflect"
	"testing"

	lg "github.com/GolangToolKits/go-level-logger"
	gdb "github.com/GolangToolKits/go-mysql"
	db "github.com/Ulbora/go-micro-blog/db"
)

func TestSysManager_AddComment(t *testing.T) {

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

	mdb.MockRow2 = &gdb.DbRow{
		//Row: []string{"0"},
		Row: []string{"1", "test blog entry 222", "some test blog stuff", "4", "true", "2023-03-01 00:01:14", ""},
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
			name: "test 1",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
			},
			args: args{
				c: &db.Comment{
					UserID: 1,
					BlogID: 4,
					Text:   "comment test",
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
			if got := m.AddComment(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SysManager.AddComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSysManager_UpdateComment(t *testing.T) {

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

	mdb.MockRow2 = &gdb.DbRow{
		//Row: []string{"0"},
		Row: []string{"1", "test blog entry 222", "some test blog stuff", "4", "true", "2023-03-01 00:01:14", ""},
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
			name: "test 1",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
			},
			args: args{
				c: &db.Comment{
					UserID: 1,
					BlogID: 4,
					Text:   "comment test",
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
			if got := m.UpdateComment(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SysManager.UpdateComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSysManager_GetCommentList(t *testing.T) {

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

	mdb.MockRows1 = &gdb.DbRows{
		Rows: [][]string{{"1", "some test blog stuff", "4", "5", "true"},
			{"2", "some test blog stuff", "4", "6", "false"}},
	}

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
		bid   int64
		start int64
		end   int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *[]db.Comment
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
				bid:2,
				start: 0,
				end:   100,
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
			if got := m.GetCommentList(tt.args.bid, tt.args.start, tt.args.end);  len(*got) != 1{
				t.Errorf("SysManager.GetCommentList() = %v, want %v", got, tt.want)
			}
		})
	}
}
