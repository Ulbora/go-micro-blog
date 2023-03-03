package managers

import (
	"testing"

	lg "github.com/GolangToolKits/go-level-logger"
	gdb "github.com/GolangToolKits/go-mysql"
	db "github.com/Ulbora/go-micro-blog/db"
)

func TestSysManager_GetConfig(t *testing.T) {

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
		Rows: [][]string{{"1", "true", "true"},
			{"2", "true", "true"}},
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
	tests := []struct {
		name   string
		fields fields
		want   *db.Config
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SysManager{
				DB:               tt.fields.DB,
				allowAutoPost:    tt.fields.allowAutoPost,
				allowAutoComment: tt.fields.allowAutoComment,
			}
			if got := m.GetConfig(); got.ID != 1 {
				t.Errorf("SysManager.GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
