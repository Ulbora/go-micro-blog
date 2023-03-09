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

func TestMCHandler_AddUser(t *testing.T) {

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)
	var mg m.MockManager
	var rid m.ResponseID
	rid.ID = 1
	rid.Success = true
	mg.MockAddUser = rid
	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"email":"test", "firstName": "bob","lastName": "hope", "roleId": 5}`))

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

	aJSON4 := ioutil.NopCloser(bytes.NewBufferString(`{"firstName":"test", "roleId": 5}`))
	r4, _ := http.NewRequest("POST", "/ffllist", aJSON4)
	r4.Header.Set("apiKey", "1234")
	r4.Header.Set("Content-Type", "application/json")

	w4 := httptest.NewRecorder()

	var mg4 m.MockManager
	var rid4 m.ResponseID

	mg4.MockAddUser = rid4

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
			name: "test 2",
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
			name: "test 2",
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
			h.AddUser(tt.args.w, tt.args.r)

			var res m.ResponseID
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if tt.name == "test 1" && (w.Code != 200 || !res.Success) {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_UpdateUser(t *testing.T) {

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)
	var mg m.MockManager
	var rid m.Response
	rid.Success = true
	mg.MockUpdateUser = rid
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

	mg4.MockAddUser = rid4

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
			name: "test 3",
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
			h.UpdateUser(tt.args.w, tt.args.r)

			var res m.Response
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if tt.name == "test 1" && (w.Code != 200 || !res.Success) {
				t.Fail()
			}

		})
	}
}

func TestMCHandler_GetUser(t *testing.T) {

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
		Row: []string{"1", "test@test.com", "bob", "hope", "", "4", "true"},
	}

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	r.Header.Set("apiKey", "1234")
	vars := map[string]string{
		"email": "test@test.com",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()

	r2, _ := http.NewRequest("GET", "/ffllist", nil)
	r2.Header.Set("apiKey", "12342")
	vars2 := map[string]string{
		"email": "test@test.com",
	}
	r2 = mux.SetURLVars(r2, vars2)
	w2 := httptest.NewRecorder()

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
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				APIKey: "1234",
				Log:    log,
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
				APIKey: "1234",
				Log:    log,
			},
			args: args{
				w: w2,
				r: r2,
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
			h.GetUser(tt.args.w, tt.args.r)

			var res db.User
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if tt.name == "test 1" && (w.Code != 200 || res.ID != 1) {
				t.Fail()
			}

		})
	}
}

func TestMCHandler_GetUserList(t *testing.T) {

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
		Rows: [][]string{{"1", "test@test.com", "bob", "hope", "", "4", "true"},
			{"2", "test2@test.com", "bobby", "hope", "", "4", "true"}},
	}

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	r.Header.Set("apiAdminKey", "1234")
	// vars := map[string]string{
	// 	"start": "1",
	// 	"end":   "5",
	// }
	// r = mux.SetURLVars(r, vars)

	w := httptest.NewRecorder()

	r2, _ := http.NewRequest("GET", "/ffllist", nil)
	r2.Header.Set("apiAdminKey", "1234")
	// vars2 := map[string]string{
	// 	"start": "1",
	// 	"end":   "5",
	// }
	// r = mux.SetURLVars(r, vars)

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
				w: w2,
				r: r2,
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
			h.GetUserList(tt.args.w, tt.args.r)

			var res []db.User
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if tt.name == "test 1" && (w.Code != 200 || len(res) != 2) {
				t.Fail()
			}

		})
	}
}

func TestMCHandler_EnableUser(t *testing.T) {

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

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"email":"test", "firstName": "bob","lastName": "hope", "roleId": 5}`))

	r, _ := http.NewRequest("GET", "/ffllist", aJSON)
	r.Header.Set("apiAdminKey", "1234")
	r.Header.Set("Content-Type", "application/json")

	// vars := map[string]string{
	// 	"id": "1",
	// }
	// r = mux.SetURLVars(r, vars)

	w := httptest.NewRecorder()

	r2, _ := http.NewRequest("GET", "/ffllist", nil)
	r2.Header.Set("apiAdminKey", "1234")
	// vars2 := map[string]string{
	// 	"id": "1w",
	// }
	// r2 = mux.SetURLVars(r2, vars2)

	w2 := httptest.NewRecorder()

	r3, _ := http.NewRequest("GET", "/ffllist", nil)
	r3.Header.Set("apiAdminKey", "12343")
	r3.Header.Set("Content-Type", "application/json")

	vars3 := map[string]string{
		"id": "1",
	}
	r3 = mux.SetURLVars(r3, vars3)

	w3 := httptest.NewRecorder()

	aJSON4 := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "blogId": 4, "text":"test", "userId": 5}`))
	r4, _ := http.NewRequest("GET", "/ffllist", aJSON4)
	r4.Header.Set("Content-Type", "application/json")
	r4.Header.Set("apiAdminKey", "1234")
	// vars3 := map[string]string{
	// 	"id": "1",
	// }
	// r3 = mux.SetURLVars(r3, vars3)

	w4 := httptest.NewRecorder()

	mdb4 := gdb.MyDBMock{
		Host:     "localhost:3306",
		User:     "admin",
		Password: "admin",
		Database: "go_micro_blog",
	}
	mdb4.MockTestRow = &gdb.DbRow{
		//Row: []string{"0"},
		Row: []string{},
	}

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
		{
			name: "test 4",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb4,
					Log: log,
				},
				Log:         log,
				APIAdminKey: "1234",
				//Manager: mg.New(),
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
			h.EnableUser(tt.args.w, tt.args.r)

			var res m.Response
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if tt.name == "test 1" && (w.Code != 200 || !res.Success) {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_DisableUser(t *testing.T) {

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

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"email":"test", "firstName": "bob","lastName": "hope", "roleId": 5}`))

	r, _ := http.NewRequest("GET", "/ffllist", aJSON)
	r.Header.Set("apiAdminKey", "1234")
	r.Header.Set("Content-Type", "application/json")

	// vars := map[string]string{
	// 	"id": "1",
	// }
	// r = mux.SetURLVars(r, vars)

	w := httptest.NewRecorder()

	r2, _ := http.NewRequest("GET", "/ffllist", nil)
	r2.Header.Set("apiAdminKey", "1234")
	// vars2 := map[string]string{
	// 	"id": "1w",
	// }
	// r2 = mux.SetURLVars(r2, vars2)

	w2 := httptest.NewRecorder()

	r3, _ := http.NewRequest("GET", "/ffllist", nil)
	r3.Header.Set("apiAdminKey", "12343")
	r3.Header.Set("Content-Type", "application/json")

	vars3 := map[string]string{
		"id": "1",
	}
	r3 = mux.SetURLVars(r3, vars3)

	w3 := httptest.NewRecorder()

	aJSON4 := ioutil.NopCloser(bytes.NewBufferString(`{"id": 4, "blogId": 4, "text":"test", "userId": 5}`))
	r4, _ := http.NewRequest("GET", "/ffllist", aJSON4)
	r4.Header.Set("Content-Type", "application/json")
	r4.Header.Set("apiAdminKey", "1234")
	// vars3 := map[string]string{
	// 	"id": "1",
	// }
	// r3 = mux.SetURLVars(r3, vars3)

	w4 := httptest.NewRecorder()

	mdb4 := gdb.MyDBMock{
		Host:     "localhost:3306",
		User:     "admin",
		Password: "admin",
		Database: "go_micro_blog",
	}
	mdb4.MockTestRow = &gdb.DbRow{
		//Row: []string{"0"},
		Row: []string{},
	}

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
		{
			name: "test 4",
			fields: fields{
				DB: &db.MyBlogDB{
					DB:  &mdb4,
					Log: log,
				},
				Log:         log,
				APIAdminKey: "1234",
				//Manager: mg.New(),
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
			h.DisableUser(tt.args.w, tt.args.r)

			var res m.Response
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if tt.name == "test 1" && (w.Code != 200 || !res.Success) {
				t.Fail()
			}
		})
	}
}
