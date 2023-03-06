package managers

import (
	"reflect"
	"testing"

	lg "github.com/GolangToolKits/go-level-logger"
	gdb "github.com/GolangToolKits/go-mysql"
	db "github.com/Ulbora/go-micro-blog/db"
)

func TestSysManager_AddBlog(t *testing.T) {

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
			name: "test 1",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
			},
			args: args{
				b: &db.Blog{
					Name:   "test",
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
			m := &SysManager{
				DB:               tt.fields.DB,
				Log:              tt.fields.Log,
				allowAutoPost:    tt.fields.allowAutoPost,
				allowAutoComment: tt.fields.allowAutoComment,
			}
			if got := m.AddBlog(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SysManager.AddBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSysManager_UpdateBlog(t *testing.T) {

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
			name: "test 1",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
			},
			args: args{
				b: &db.Blog{
					Name:   "test",
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
			m := &SysManager{
				DB:               tt.fields.DB,
				Log:              tt.fields.Log,
				allowAutoPost:    tt.fields.allowAutoPost,
				allowAutoComment: tt.fields.allowAutoComment,
			}
			if got := m.UpdateBlog(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SysManager.UpdateBlog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSysManager_GetBlogList(t *testing.T) {

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
		Rows: [][]string{{"1", "test blog entry", "some test blog stuff", "4", "true", "2023-03-01 00:01:14", ""},
			{"2", "test blog entry 333", "some test blog stuff", "4", "false", "2023-03-01 00:01:14", "2023-03-01 00:01:14"}},
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
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
			},
			args: args{
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
			if got := m.GetBlogList(tt.args.start, tt.args.end); len(*got) != 1 {
				t.Errorf("SysManager.GetBlogList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSysManager_GetBlogByName(t *testing.T) {


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
		Rows: [][]string{{"1", "test blog entry", "some test blog stuff", "4", "true", "2023-03-01 00:01:14", ""},
			{"2", "test blog entry 333", "some test blog stuff", "4", "false", "2023-03-01 00:01:14", "2023-03-01 00:01:14"}},
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
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
			},
			args: args{
				name: "test",
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
			if got := m.GetBlogByName(tt.args.name, tt.args.start, tt.args.end); len(*got) != 1 {
				t.Errorf("SysManager.GetBlogByName() = %v, want %v", got, tt.want)
			}
		})
	}
}
