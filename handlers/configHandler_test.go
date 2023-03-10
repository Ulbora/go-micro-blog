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

	db "github.com/Ulbora/go-micro-blog/db"
	m "github.com/Ulbora/go-micro-blog/managers"
)

func TestMCHandler_UpdateConfig(t *testing.T) {

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
	mdb.MockConnectSuccess = true
	mdb.MockUpdateSuccess1 = true

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id":1, "allowAutoPost": true, "allowAutoComment": true}`))

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
				Log: log,
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
				Log: log,
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
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				Log: log,
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
				DB: &db.MyBlogDB{
					DB:  &mdb,
					Log: log,
				},
				Log: log,
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
			h.UpdateConfig(tt.args.w, tt.args.r)

			var res m.Response
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if (tt.ww.Code != tt.code || res.Success != tt.suc) {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_GetConfig(t *testing.T) {


	var mg m.MockManager
	
	var cf = db.Config{
		ID:               2,
		AllowAutoPost:    true,
		AllowAutoComment: true,
	}

	mg.MockConfig = cf

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	r, _ := http.NewRequest("GET", "/ffllist", nil)
	r.Header.Set("apiAdminKey", "1234")
	
	w := httptest.NewRecorder()



	r2, _ := http.NewRequest("GET", "/ffllist", nil)
	r2.Header.Set("apiAdminKey", "12344")
	
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
					//DB:  &mdb,
					Log: log,
				},
				Log:         log,
				APIAdminKey: "1234",
				Manager:     mg.New(),
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
					//DB:  &mdb,
					Log: log,
				},
				Log:         log,
				APIAdminKey: "1234",
				Manager:     mg.New(),
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
			h.GetConfig(tt.args.w, tt.args.r)

			var res db.Config
			body, _ := ioutil.ReadAll(w.Result().Body)
			json.Unmarshal(body, &res)
			if (tt.ww.Code != tt.code || res.AllowAutoPost != tt.suc) {
				t.Fail()
			}
		})
	}
}
