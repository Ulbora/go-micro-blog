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
	db "github.com/Ulbora/go-micro-blog/db"
	m "github.com/Ulbora/go-micro-blog/managers"
)

func TestMCHandler_AddRole(t *testing.T) {

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)
	var mg m.MockManager
	var rid m.ResponseID
	rid.ID = 1
	rid.Success = true
	mg.MockAddRole = rid
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"name":"test"}`))

	r, _ := http.NewRequest("POST", "/ffllist", aJSON)
	r.Header.Set("apiAdminKey", "1234")
	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r2, _ := http.NewRequest("POST", "/ffllist", aJSON)

	w2 := httptest.NewRecorder()

	r3, _ := http.NewRequest("POST", "/ffllist", nil)
	r3.Header.Set("apiAdminKey", "1234")
	r3.Header.Set("Content-Type", "application/json")
	w3 := httptest.NewRecorder()

	aJSON4 := ioutil.NopCloser(bytes.NewBufferString(`{"name":"test"}`))
	r4, _ := http.NewRequest("POST", "/ffllist", aJSON4)
	r4.Header.Set("apiAdminKey", "1234")
	r4.Header.Set("Content-Type", "application/json")

	w4 := httptest.NewRecorder()

	var mg4 m.MockManager
	var rid4 m.ResponseID

	mg4.MockAddRole = rid4

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
		code   int
		suc    bool
		len int
		ww     *httptest.ResponseRecorder
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Log:         log,
				Manager:     mg.New(),
				APIAdminKey: "1234",
			},
			args: args{
				w: w,
				r: r,
			},
			code: 200,
			suc:  true,
			len: 0,
			ww:   w,
		},
		{
			name: "test 2",
			fields: fields{
				Log:         log,
				Manager:     mg.New(),
				APIAdminKey: "1234",
			},
			args: args{
				w: w2,
				r: r2,
			},
			code: 415,
			suc:  false,
			len: 0,
			ww:   w2,
		},
		{
			name: "test 3",
			fields: fields{
				Log:         log,
				Manager:     mg.New(),
				APIAdminKey: "1234",
			},
			args: args{
				w: w3,
				r: r3,
			},
			code: 400,
			suc:  false,
			len: 0,
			ww:   w3,
		},
		{
			name: "test 4",
			fields: fields{
				Log:         log,
				Manager:     mg4.New(),
				APIAdminKey: "1234",
			},
			args: args{
				w: w4,
				r: r4,
			},
			code: 500,
			suc:  false,
			len: 0,
			ww:   w4,
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
			h.AddRole(tt.args.w, tt.args.r)

			var res m.ResponseID
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if  (tt.ww.Code != tt.code || res.Success != tt.suc) {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_GetRole(t *testing.T) {

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
		Row: []string{"1", "user"},
	}

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	r.Header.Set("apiKey", "1234")

	vars := map[string]string{
		"name": "user",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	r2, _ := http.NewRequest("GET", "/ffllist", nil)
	vars2 := map[string]string{
		//"name": "user",
	}
	r2 = mux.SetURLVars(r2, vars2)
	w2 := httptest.NewRecorder()

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
		code   int
		suc    bool
		id int64
		ww     *httptest.ResponseRecorder
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				Log:    log,
				APIKey: "1234",
			},
			args: args{
				w: w,
				r: r,
			},
			code: 200,
			suc:  true,
			id: 1,
			ww:   w,
		},
		{
			name: "test 2",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				Log: log,
			},
			args: args{
				w: w2,
				r: r2,
			},
			code: 400,
			suc:  false,
			id: 0,
			ww:   w2,
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
			h.GetRole(tt.args.w, tt.args.r)

			var res db.Role
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if  (tt.ww.Code != tt.code || res.ID != tt.id) {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_GetRoleList(t *testing.T) {

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
		Rows: [][]string{{"1", "role1"},
			{"2", "role2"}},
	}

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	r.Header.Set("apiKey", "1234")
	// vars := map[string]string{
	// 	"start": "1",
	// 	"end":   "5",
	// }
	// r = mux.SetURLVars(r, vars)

	w := httptest.NewRecorder()

	r2, _ := http.NewRequest("GET", "/ffllist", nil)
	r2.Header.Set("apiKey", "12343")

	w2 := httptest.NewRecorder()

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
		code   int
		suc    bool
		len int
		ww     *httptest.ResponseRecorder
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				Log:    log,
				APIKey: "1234",
			},
			args: args{
				w: w,
				r: r,
			},
			code: 200,
			suc:  true,
			len: 2,
			ww:   w,
		},
		{
			name: "test 2",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				Log:    log,
				APIKey: "1234",
			},
			args: args{
				w: w2,
				r: r2,
			},
			code: 400,
			suc:  false,
			len: 0,
			ww:   w2,
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
			h.GetRoleList(tt.args.w, tt.args.r)

			var res []db.Role
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if  (tt.ww.Code != tt.code || len(res) != tt.len) {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_DeleteRole(t *testing.T) {

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

	mdb.MockDeleteSuccess1 = true

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
	r2.Header.Set("apiAdminKey", "12343")
	vars2 := map[string]string{
		"id": "1",
	}
	r2 = mux.SetURLVars(r2, vars2)

	w2 := httptest.NewRecorder()




	r3, _ := http.NewRequest("GET", "/ffllist", nil)
	r3.Header.Set("apiAdminKey", "1234")
	vars3 := map[string]string{
		"id": "n1",
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
		code   int
		suc    bool
		len int
		ww     *httptest.ResponseRecorder
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				Log:    log,
				APIAdminKey: "1234",
			},
			args: args{
				w: w,
				r: r,
			},
			code: 200,
			suc:  true,
			len: 0,
			ww:   w,
		},
		{
			name: "test 2",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				Log:    log,
				APIAdminKey: "1234",
			},
			args: args{
				w: w2,
				r: r2,
			},
			code: 400,
			suc:  false,
			len: 0,
			ww:   w2,
		},
		{
			name: "test 3",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				Log:    log,
				APIAdminKey: "1234",
			},
			args: args{
				w: w3,
				r: r3,
			},
			code: 400,
			suc:  false,
			len: 0,
			ww:   w3,
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
			h.DeleteRole(tt.args.w, tt.args.r)



			var res m.Response
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if (tt.ww.Code != tt.code || res.Success != tt.suc) {
				t.Fail()
			}
		})
	}
}
