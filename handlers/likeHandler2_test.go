package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	lg "github.com/GolangToolKits/go-level-logger"
	gdb "github.com/GolangToolKits/go-mysql"
	mux "github.com/GolangToolKits/grrt"
	"github.com/Ulbora/go-micro-blog/db"
	m "github.com/Ulbora/go-micro-blog/managers"
)

func TestMCHandler_AddLike(t *testing.T) {

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)
	var mg m.MockManager
	var rid m.Response
	rid.Success = true
	mg.MockAddLike = rid
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"blogId":1, "userId": 5}`))

	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("apiKey", "1234")
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r2, _ := http.NewRequest("POST", "/ffllist", aJSON)

	w2 := httptest.NewRecorder()

	r3, _ := http.NewRequest("POST", "/ffllist", nil)
	r3.Header.Set("apiKey", "1234")
	r3.Header.Set("Content-Type", "application/json")
	w3 := httptest.NewRecorder()

	aJSON4 := ioutil.NopCloser(bytes.NewBufferString(`{"blogId":5, "userId": 5}`))
	r4, _ := http.NewRequest("POST", "/ffllist", aJSON4)
	r4.Header.Set("apiKey", "1234")
	r4.Header.Set("Content-Type", "application/json")

	w4 := httptest.NewRecorder()

	var mg4 m.MockManager
	var rid4 m.Response

	mg4.MockAddLike = rid4

	type fields struct {
		DB          db.BlogDB
		Log         lg.Log
		Manager     m.Manager
		APIKey      string
		APIAdminKey string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Log:     log,
				Manager: mg.New(),
				APIKey:  "1234",
			},
			args: args{
				w: w,
				r: r,
			},
		},
		{
			name: "test 2",
			fields: fields{
				Log:     log,
				Manager: mg.New(),
				APIKey:  "1234",
			},
			args: args{
				w: w2,
				r: r2,
			},
		},
		{
			name: "test 3",
			fields: fields{
				Log:     log,
				Manager: mg.New(),
				APIKey:  "1234",
			},
			args: args{
				w: w3,
				r: r3,
			},
		},
		{
			name: "test 4",
			fields: fields{
				Log:     log,
				Manager: mg4.New(),
				APIKey:  "1234",
			},
			args: args{
				w: w4,
				r: r4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MCHandler{
				DB:          tt.fields.DB,
				Log:         tt.fields.Log,
				Manager:     tt.fields.Manager,
				APIKey:      tt.fields.APIKey,
				APIAdminKey: tt.fields.APIAdminKey,
			}
			h.AddLike(tt.args.w, tt.args.r)

			var res m.Response
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if tt.name == "test 1" && (w.Code != 200 || !res.Success) {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_RemoveLike(t *testing.T) {

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)
	var mg m.MockManager
	var rid m.Response
	rid.Success = true
	mg.MockRemoveLike = rid
	//aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"blogId":1, "userId": 5}`))

	r, _ := http.NewRequest("DELETE", "/ffllist", nil)
	r.Header.Set("apiKey", "1234")
	r.Header.Set("Content-Type", "application/json")
	vars := map[string]string{
		"bid": "1",
		"uid": "3",
	}
	r = mux.SetURLVars(r, vars)

	w := httptest.NewRecorder()

	r2, _ := http.NewRequest("POST", "/ffllist", nil)

	w2 := httptest.NewRecorder()

	r3, _ := http.NewRequest("DELETE", "/ffllist", nil)
	r3.Header.Set("apiKey", "1234")
	r3.Header.Set("Content-Type", "application/json")
	vars3 := map[string]string{
		"bid": "1",
		"uid": "n3",
	}
	r3 = mux.SetURLVars(r3, vars3)

	w3 := httptest.NewRecorder()

	type fields struct {
		DB          db.BlogDB
		Log         lg.Log
		Manager     m.Manager
		APIKey      string
		APIAdminKey string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Log:     log,
				Manager: mg.New(),
				APIKey:  "1234",
			},
			args: args{
				w: w,
				r: r,
			},
		},
		{
			name: "test 2",
			fields: fields{
				Log:     log,
				Manager: mg.New(),
				APIKey:  "1234",
			},
			args: args{
				w: w2,
				r: r2,
			},
		},
		{
			name: "test 3",
			fields: fields{
				Log:     log,
				Manager: mg.New(),
				APIKey:  "1234",
			},
			args: args{
				w: w3,
				r: r3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MCHandler{
				DB:          tt.fields.DB,
				Log:         tt.fields.Log,
				Manager:     tt.fields.Manager,
				APIKey:      tt.fields.APIKey,
				APIAdminKey: tt.fields.APIAdminKey,
			}
			h.RemoveLike(tt.args.w, tt.args.r)

			var res m.Response
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if tt.name == "test 1" && (w.Code != 200 || !res.Success) {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_ViewLikes(t *testing.T) {

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
	// mdb.MockRow1 = &gdb.DbRow{
	// 	Row: []string{"1", "4"},
	// }
	mdb.MockRows1 = &gdb.DbRows{
		Rows: [][]string{{"1", "4"},
			{"2", "4"}},
	}

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"bid": "1",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()




	r2, _ := http.NewRequest("GET", "/ffllist", nil)
	vars2 := map[string]string{
		"bid": "1m",
	}
	r2 = mux.SetURLVars(r2, vars2)
	w2 := httptest.NewRecorder()



	r3, _ := http.NewRequest("GET", "/ffllist", nil)
	vars3 := map[string]string{
		//"bid": "1",
	}
	r3 = mux.SetURLVars(r3, vars3)
	w3 := httptest.NewRecorder()


	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	type fields struct {
		DB          db.BlogDB
		Log         lg.Log
		Manager     m.Manager
		APIKey      string
		APIAdminKey string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Log: log,
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				//Manager: mg.New(),
				APIKey: "1234",
			},
			args: args{
				w: w,
				r: r,
			},
		},
		{
			name: "test 2",
			fields: fields{
				Log: log,
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				//Manager: mg.New(),
				APIKey: "1234",
			},
			args: args{
				w: w2,
				r: r2,
			},
		},
		{
			name: "test 2",
			fields: fields{
				Log: log,
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				//Manager: mg.New(),
				APIKey: "1234",
			},
			args: args{
				w: w3,
				r: r3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MCHandler{
				DB:          tt.fields.DB,
				Log:         tt.fields.Log,
				Manager:     tt.fields.Manager,
				APIKey:      tt.fields.APIKey,
				APIAdminKey: tt.fields.APIAdminKey,
			}
			h.ViewLikes(tt.args.w, tt.args.r)

			var res []db.Like
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if tt.name == "test 1" && (w.Code != 200 || len(res) != 2) {
				t.Fail()
			}
		})
	}
}
