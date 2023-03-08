package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	lg "github.com/GolangToolKits/go-level-logger"
	gdb "github.com/GolangToolKits/go-mysql"
	mux "github.com/GolangToolKits/grrt"
	db "github.com/Ulbora/go-micro-blog/db"
	m "github.com/Ulbora/go-micro-blog/managers"
)

func TestMCHandler_AddBlog(t *testing.T) {

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)
	var mg m.MockManager
	var rid m.ResponseID
	rid.ID = 1
	rid.Success = true
	mg.MockAddBlog = rid
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"name":"test", "userId": 5}`))

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

	aJSON4 := ioutil.NopCloser(bytes.NewBufferString(`{"name":"test", "userId": 5}`))
	r4, _ := http.NewRequest("POST", "/ffllist", aJSON4)
	r4.Header.Set("apiKey", "1234")
	r4.Header.Set("Content-Type", "application/json")

	w4 := httptest.NewRecorder()

	var mg4 m.MockManager
	var rid4 m.ResponseID

	mg4.MockAddBlog = rid4

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
			h.AddBlog(tt.args.w, tt.args.r)

			var res m.ResponseID
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if tt.name == "test 1" && (w.Code != 200 || !res.Success) {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_UpdateBlog(t *testing.T) {

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)
	var mg m.MockManager
	var rid m.Response
	rid.Success = true
	mg.MockUpdateBlog = rid
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"name":"test", "userId": 5}`))

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

	aJSON4 := ioutil.NopCloser(bytes.NewBufferString(`{"name":"test", "userId": 5}`))
	r4, _ := http.NewRequest("POST", "/ffllist", aJSON4)
	r4.Header.Set("apiKey", "1234")
	r4.Header.Set("Content-Type", "application/json")

	w4 := httptest.NewRecorder()

	var mg4 m.MockManager
	var rid4 m.ResponseID

	mg4.MockAddBlog = rid4

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
			h.UpdateBlog(tt.args.w, tt.args.r)

			var res m.Response
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if tt.name == "test 1" && (w.Code != 200 || !res.Success) {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_GetBlog(t *testing.T) {

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
		Row: []string{"1", "test blog entry 222", "some test blog stuff", "4", "true", "2023-03-01 00:01:14", ""},
	}

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"id": "1",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	mdb2 := gdb.MyDBMock{
		Host:     "localhost:3306",
		User:     "admin",
		Password: "admin",
		Database: "go_micro_blog",
	}
	mdb2.MockTestRow = &gdb.DbRow{
		//Row: []string{"0"},
		Row: []string{},
	}
	mdb2.MockRow1 = &gdb.DbRow{
		Row: []string{"1", "test blog entry 222", "some test blog stuff", "4", "false", "2023-03-01 00:01:14", ""},
	}
	w2 := httptest.NewRecorder()

	r3, _ := http.NewRequest("GET", "/ffllist", nil)
	vars3 := map[string]string{
		"id": "n1",
	}
	r3 = mux.SetURLVars(r3, vars3)
	w3 := httptest.NewRecorder()

	r4, _ := http.NewRequest("GET", "/ffllist", nil)
	vars4 := map[string]string{
		"id":  "1",
		"idd": "2",
	}
	r4 = mux.SetURLVars(r4, vars4)
	w4 := httptest.NewRecorder()

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
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				Log: log,
			},
			args: args{
				w: w,
				r: r,
			},
		},
		{
			name: "test 2",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb2,
					Log: log,
				},
				Log: log,
			},
			args: args{
				w: w2,
				r: r,
			},
		},
		{
			name: "test 3",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb2,
					Log: log,
				},
				Log: log,
			},
			args: args{
				w: w3,
				r: r3,
			},
		},
		{
			name: "test 4",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb2,
					Log: log,
				},
				Log: log,
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
			h.GetBlog(tt.args.w, tt.args.r)

			var res db.Blog
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if tt.name == "test 1" && (w.Code != 200 || res.ID != 1) {
				t.Fail()
			}

		})
	}
}

func TestMCHandler_GetBlogByName(t *testing.T) {

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)
	var mg m.MockManager
	var bl = []db.Blog{{ID: 1, Name: "test blog entry", Content: "some test blog stuff", UserID: 4, Active: true, Entered: time.Now(), Updated: time.Now()},
		{ID: 2, Name: "test blog entry 333", Content: "some test blog stuff", UserID: 4, Active: false, Entered: time.Now(), Updated: time.Now()}}

	mg.MockBlogList = bl

	// aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"name":"test", "userId": 5}`))

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"name":  "test",
		"start": "1",
		"end":   "5",
	}
	r = mux.SetURLVars(r, vars)
	//r.Header.Set("apiKey", "1234")
	//r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r2, _ := http.NewRequest("GET", "/ffllist", nil)
	vars2 := map[string]string{
		"name":  "test",
		"start": "1c",
		"end":   "5",
	}
	r2 = mux.SetURLVars(r2, vars2)
	//r.Header.Set("apiKey", "1234")
	//r.Header.Set("Content-Type", "application/json")

	w2 := httptest.NewRecorder()

	r3, _ := http.NewRequest("GET", "/ffllist", nil)
	vars3 := map[string]string{
		//"name":  "test",
		"start": "1c",
		"end":   "5",
	}
	r3 = mux.SetURLVars(r3, vars3)
	//r.Header.Set("apiKey", "1234")
	//r.Header.Set("Content-Type", "application/json")

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
				DB: &db.MyBlogDB{
					//DB:  &mdb2,
					Log: log,
				},
				Log:     log,
				Manager: mg.New(),
			},
			args: args{
				w: w,
				r: r,
			},
		},
		{
			name: "test 2",
			fields: fields{
				DB: &db.MyBlogDB{
					//DB:  &mdb2,
					Log: log,
				},
				Log:     log,
				Manager: mg.New(),
			},
			args: args{
				w: w2,
				r: r2,
			},
		},
		{
			name: "test 3",
			fields: fields{
				DB: &db.MyBlogDB{
					//DB:  &mdb2,
					Log: log,
				},
				Log:     log,
				Manager: mg.New(),
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
			h.GetBlogByName(tt.args.w, tt.args.r)

			var res []db.Blog
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if tt.name == "test 1" && (w.Code != 200 || len(res) != 2) {
				t.Fail()
			}

		})
	}
}

func TestMCHandler_GetBlogList(t *testing.T) {

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)
	var mg m.MockManager
	var bl = []db.Blog{{ID: 1, Name: "test blog entry", Content: "some test blog stuff", UserID: 4, Active: true, Entered: time.Now(), Updated: time.Now()},
		{ID: 2, Name: "test blog entry 333", Content: "some test blog stuff", UserID: 4, Active: false, Entered: time.Now(), Updated: time.Now()}}

	mg.MockBlogList = bl

	// aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"name":"test", "userId": 5}`))

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	vars := map[string]string{
		"start": "1",
		"end":   "5",
	}
	r = mux.SetURLVars(r, vars)

	w := httptest.NewRecorder()

	r2, _ := http.NewRequest("GET", "/ffllist", nil)
	vars2 := map[string]string{
		"name":  "test",
		"start": "1c",
		"end":   "5",
	}
	r2 = mux.SetURLVars(r2, vars2)
	//r.Header.Set("apiKey", "1234")
	//r.Header.Set("Content-Type", "application/json")

	w2 := httptest.NewRecorder()

	r3, _ := http.NewRequest("GET", "/ffllist", nil)
	vars3 := map[string]string{
		//"name":  "test",
		"start": "1c",
		"end":   "5",
	}
	r3 = mux.SetURLVars(r3, vars3)
	//r.Header.Set("apiKey", "1234")
	//r.Header.Set("Content-Type", "application/json")

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
				DB: &db.MyBlogDB{
					//DB:  &mdb2,
					Log: log,
				},
				Log:     log,
				Manager: mg.New(),
			},
			args: args{
				w: w,
				r: r,
			},
		},
		{
			name: "test 2",
			fields: fields{
				DB: &db.MyBlogDB{
					//DB:  &mdb2,
					Log: log,
				},
				Log:     log,
				Manager: mg.New(),
			},
			args: args{
				w: w2,
				r: r2,
			},
		},
		{
			name: "test 3",
			fields: fields{
				DB: &db.MyBlogDB{
					//DB:  &mdb2,
					Log: log,
				},
				Log:     log,
				Manager: mg.New(),
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
			h.GetBlogList(tt.args.w, tt.args.r)

			var res []db.Blog
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if tt.name == "test 1" && (w.Code != 200 || len(res) != 2) {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_GetAdminBlogList(t *testing.T) {

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

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	r.Header.Set("apiAdminKey", "1234")
	vars := map[string]string{
		"start": "1",
		"end":   "5",
	}
	r = mux.SetURLVars(r, vars)

	w := httptest.NewRecorder()

	r3, _ := http.NewRequest("GET", "/ffllist", nil)
	r3.Header.Set("apiAdminKey", "1234")
	vars3 := map[string]string{
		"start": "1",
		"end":   "5v",
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
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				Log:         log,
				APIAdminKey: "1234",
				//Manager: mg.New(),
			},
			args: args{
				w: w,
				r: r,
			},
		},
		{
			name: "test 2",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				Log:         log,
				APIAdminKey: "12343",
				//Manager: mg.New(),
			},
			args: args{
				w: w,
				r: r,
			},
		},
		{
			name: "test 3",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				Log:         log,
				APIAdminKey: "1234",
				//Manager: mg.New(),
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
			h.GetAdminBlogList(tt.args.w, tt.args.r)

			var res []db.Blog
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if tt.name == "test 1" && (w.Code != 200 || len(res) != 2) {
				t.Fail()
			}

		})
	}
}

func TestMCHandler_ActivateBlog(t *testing.T) {

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

	mdb.MockUpdateSuccess1 = true

	// mdb.MockRows1 = &gdb.DbRows{
	// 	Rows: [][]string{{"1", "test blog entry", "some test blog stuff", "4", "true", "2023-03-01 00:01:14", ""},
	// 		{"2", "test blog entry 333", "some test blog stuff", "4", "false", "2023-03-01 00:01:14", "2023-03-01 00:01:14"}},
	// }

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	r.Header.Set("apiAdminKey", "1234")
	vars := map[string]string{
		"id": "1",
	}
	r = mux.SetURLVars(r, vars)

	w := httptest.NewRecorder()

	r2, _ := http.NewRequest("GET", "/ffllist", nil)
	r2.Header.Set("apiAdminKey", "1234")
	vars2 := map[string]string{
		"id": "1w",
	}
	r2 = mux.SetURLVars(r2, vars2)

	w2 := httptest.NewRecorder()

	r3, _ := http.NewRequest("GET", "/ffllist", nil)
	r3.Header.Set("apiAdminKey", "12343")
	vars3 := map[string]string{
		"id": "1",
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
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				Log:         log,
				APIAdminKey: "1234",
				//Manager: mg.New(),
			},
			args: args{
				w: w,
				r: r,
			},
		},
		{
			name: "test 2",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				Log:         log,
				APIAdminKey: "1234",
				//Manager: mg.New(),
			},
			args: args{
				w: w2,
				r: r2,
			},
		},
		{
			name: "test 3",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				Log:         log,
				APIAdminKey: "1234",
				//Manager: mg.New(),
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
			h.ActivateBlog(tt.args.w, tt.args.r)
		})

		var res m.Response
		body, _ := ioutil.ReadAll(w.Result().Body)
		json.Unmarshal(body, &res)
		if tt.name == "test 1" && (w.Code != 200 || !res.Success) {
			t.Fail()
		}

	}
}

func TestMCHandler_DectivateBlog(t *testing.T) {

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

	mdb.MockUpdateSuccess1 = true

	// mdb.MockRows1 = &gdb.DbRows{
	// 	Rows: [][]string{{"1", "test blog entry", "some test blog stuff", "4", "true", "2023-03-01 00:01:14", ""},
	// 		{"2", "test blog entry 333", "some test blog stuff", "4", "false", "2023-03-01 00:01:14", "2023-03-01 00:01:14"}},
	// }

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	r.Header.Set("apiAdminKey", "1234")
	vars := map[string]string{
		"id": "1",
	}
	r = mux.SetURLVars(r, vars)

	w := httptest.NewRecorder()

	r2, _ := http.NewRequest("GET", "/ffllist", nil)
	r2.Header.Set("apiAdminKey", "1234")
	vars2 := map[string]string{
		"id": "1w",
	}
	r2 = mux.SetURLVars(r2, vars2)

	w2 := httptest.NewRecorder()

	r3, _ := http.NewRequest("GET", "/ffllist", nil)
	r3.Header.Set("apiAdminKey", "12343")
	vars3 := map[string]string{
		"id": "1",
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
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				Log:         log,
				APIAdminKey: "1234",
				//Manager: mg.New(),
			},
			args: args{
				w: w,
				r: r,
			},
		},
		{
			name: "test 2",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				Log:         log,
				APIAdminKey: "1234",
				//Manager: mg.New(),
			},
			args: args{
				w: w2,
				r: r2,
			},
		},
		{
			name: "test 3",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				Log:         log,
				APIAdminKey: "1234",
				//Manager: mg.New(),
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
			h.DectivateBlog(tt.args.w, tt.args.r)

			var res m.Response
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if tt.name == "test 1" && (w.Code != 200 || !res.Success) {
				t.Fail()
			}
		})
	}
}
