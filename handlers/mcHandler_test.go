package handlers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	lg "github.com/GolangToolKits/go-level-logger"
	db "github.com/Ulbora/go-micro-blog/db"
	m "github.com/Ulbora/go-micro-blog/managers"
)

func TestMCHandler_New(t *testing.T) {

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
	tests := []struct {
		name   string
		fields fields
		want   Handler
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Log: log,
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
			if got := h.New(); got == nil {
				t.Errorf("MCHandler.New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCHandler_processAPIKey(t *testing.T) {

	r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("apiKey", "testkey")

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
		r *http.Request
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
				Log:    log,
				APIKey: "testkey",
			},
			args: args{
				r: r,
			},
			want: true,
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
			if got := h.processAPIKey(tt.args.r); got != tt.want {
				t.Errorf("MCHandler.processAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCHandler_processAPIAdminKey(t *testing.T) {

	r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("APIAdminKey", "testkey")

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
		r *http.Request
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
				Log:         log,
				APIAdminKey: "testkey",
			},
			args: args{
				r: r,
			},
			want: true,
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
			if got := h.processAPIAdminKey(tt.args.r); got != tt.want {
				t.Errorf("MCHandler.processAPIAdminKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCHandler_CheckContent(t *testing.T) {

	r, _ := http.NewRequest("POST", "/ffllist", nil)
	r.Header.Set("Content-Type", "application/json")

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
		r *http.Request
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
				Log: log,
			},
			args: args{
				r: r,
			},
			want: true,
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
			if got := h.checkContent(tt.args.r); got != tt.want {
				t.Errorf("MCHandler.CheckContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMCHandler_SetContentType(t *testing.T) {

	w := httptest.NewRecorder()

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
			},
			args: args{
				w: w,
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
			h.setContentType(tt.args.w)
			if w.Header().Get("Content-Type") != "application/json" {
				t.Fail()
			}
		})
	}
}

func TestMCHandler_ProcessBody(t *testing.T) {
	type testObj struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		Location int64  `json:"location"`
	}

	aJSON := ioutil.NopCloser(bytes.NewBufferString(`{"id": 1, "name":"test", "location": 5}`))
	r, _ := http.NewRequest("POST", "/ffllist", aJSON)

	aJSON2 := ioutil.NopCloser(bytes.NewBufferString(`{"id": 1, "name":"test", "location": "5"}`))
	r2, _ := http.NewRequest("POST", "/ffllist", aJSON2)

	r3, _ := http.NewRequest("POST", "/ffllist", nil)

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
		r   *http.Request
		obj any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
		objID int64
		objName string
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Log: log,
			},
			args: args{
				r:   r,
				obj: &testObj{},
			},
			want:    true,
			wantErr: false,
			objID: 1,
			objName: "test",
		},
		{
			name: "test 2",
			fields: fields{
				Log: log,
			},
			args: args{
				r:   r2,
				obj: testObj{},
			},
			want:    false,
			wantErr: true,
			objID: 0,
			objName: "",
		},
		{
			name: "test 3",
			fields: fields{
				Log: log,
			},
			args: args{
				r:   r3,
				obj: testObj{},
			},
			want:    false,
			wantErr: true,
			objID: 0,
			objName: "",
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
			// got, err := h.processBody(tt.args.r, tt.args.obj)
			got, err := h.processBody(tt.args.r, tt.args.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("MCHandler.ProcessBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MCHandler.ProcessBody() = %v, want %v", got, tt.want)
			}
			if got && (tt.args.obj.(*testObj).ID != tt.objID || tt.args.obj.(*testObj).Name != tt.objName){
				t.Fail()
			}
		})
	}
}
