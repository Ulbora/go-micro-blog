package db

import (
	"io/ioutil"
	"os"
	"testing"

	lg "github.com/GolangToolKits/go-level-logger"
	gdb "github.com/GolangToolKits/go-mysql"
)

func TestMyBlogDB_AddUser(t *testing.T) {

	// db := gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "admin",
	// 	Password: "admin",
	// 	Database: "go_micro_blog",
	// }

	db := gdb.MyDBMock{
		Host:     "localhost:3306",
		User:     "admin",
		Password: "admin",
		Database: "go_micro_blog",
	}
	db.MockTestRow = &gdb.DbRow{
		//Row: []string{"0"},
		Row: []string{},
	}
	db.MockConnectSuccess = true
	db.MockInsertID1 = 1
	db.MockInsertSuccess1 = true

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	type fields struct {
		DB  gdb.Database
		Log lg.Log
	}
	type args struct {
		u *User
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
				u: &User{
					Email:     "test@test.com",
					Password:  "tst",
					FirstName: "test",
					LastName:  "tester",
					RoleID:    1,
					Active:    true,
				},
			},
			want:  true,
			want1: 1,
		},
		// {
		// 	name: "test 2",
		// 	fields: fields{
		// 		DB:  db.New(),
		// 		Log: log,
		// 	},
		// 	args: args{
		// 		u: &User{
		// 			Email:     "test@test.com",
		// 			Password:  "tst",
		// 			FirstName: "test",
		// 			LastName:  "tester",
		// 			RoleID:    1,
		// 			Active:    true,
		// 		},
		// 	},
		// 	want: true,
		// 	want1: 1,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MyBlogDB{
				DB:  tt.fields.DB,
				Log: tt.fields.Log,
			}
			d.DB.Connect()
			got, got1 := d.AddUser(tt.args.u)
			if got != tt.want {
				t.Errorf("MyBlogDB.AddUser() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MyBlogDB.AddUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMyBlogDB_UpdateUser(t *testing.T) {

	// db := gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "admin",
	// 	Password: "admin",
	// 	Database: "go_micro_blog",
	// }

	db := gdb.MyDBMock{
		Host:     "localhost:3306",
		User:     "admin",
		Password: "admin",
		Database: "go_micro_blog",
	}
	db.MockTestRow = &gdb.DbRow{
		//Row: []string{"0"},
		Row: []string{},
	}
	db.MockConnectSuccess = true

	db.MockUpdateSuccess1 = true

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	type fields struct {
		DB  gdb.Database
		Log lg.Log
	}
	type args struct {
		u *User
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
				DB:  db.New(),
				Log: log,
			},
			args: args{
				u: &User{
					ID:        12,
					Password:  "tsterttt",
					FirstName: "testfff",
					LastName:  "testerfff",
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
			if got := d.UpdateUser(tt.args.u); got != tt.want {
				t.Errorf("MyBlogDB.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyBlogDB_GetUser(t *testing.T) {
	var fileName = "../../golang.png"
	//var fileName = "../mailGmail.json"
	//var mm mailFile
	file, err := ioutil.ReadFile(fileName)
	//file, err:= os.Open(fileName)
	//file, err := os.ReadFile(fileName)
	//defer file.Close()
	var fileStr string
	if err == nil {
		fileStr = string(file)
		// err := json.Unmarshal(file, &mm)
		//fmt.Println("image err: ", err)
	}

	// db := gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "admin",
	// 	Password: "admin",
	// 	Database: "go_micro_blog",
	// }

	db := gdb.MyDBMock{
		Host:     "localhost:3306",
		User:     "admin",
		Password: "admin",
		Database: "go_micro_blog",
	}
	db.MockTestRow = &gdb.DbRow{
		//Row: []string{"0"},
		Row: []string{},
	}
	db.MockRow1 = &gdb.DbRow{
		Row: []string{"1", "test@test.com", "testfff", "tester", fileStr, "2", "0"},
	}

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	type fields struct {
		DB  gdb.Database
		Log lg.Log
	}
	type args struct {
		email string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *User
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				DB:  db.New(),
				Log: log,
			},
			args: args{
				email: "test@test.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MyBlogDB{
				DB:  tt.fields.DB,
				Log: tt.fields.Log,
			}
			d.DB.Connect()
			got := d.GetUser(tt.args.email)
			if got.Email != "test@test.com" ||
				got.FirstName != "testfff" {
				t.Errorf("MyBlogDB.GetUser() = %v, want %v", got, tt.want)
			}
			f, err := os.Create("../../golang2.png")

			if err == nil {
				defer f.Close()
				img := []byte(got.Image)
				f.Write(img)
			}
		})
	}
}

func TestMyBlogDB_GetUserList(t *testing.T) {

	// db := gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "admin",
	// 	Password: "admin",
	// 	Database: "go_micro_blog",
	// }

	db := gdb.MyDBMock{
		Host:     "localhost:3306",
		User:     "admin",
		Password: "admin",
		Database: "go_micro_blog",
	}
	db.MockTestRow = &gdb.DbRow{
		//Row: []string{"0"},
		Row: []string{},
	}
	//r1:= []string{"1", "test@test.com", "testfff", "tester", "", "2", "0"}

	db.MockRows1 = &gdb.DbRows{
		Rows: [][]string{{"1", "test@test.com", "testfff", "tester", "", "2", "0"},
			{"2", "test2@test.com", "test", "tester", "", "2", "1"}},
	}
	// db.MockRow1 = &gdb.DbRow{
	// 	Row: []string{"1", "test@test.com", "testfff", "tester", "", "2", "0"},
	// }

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	type fields struct {
		DB  gdb.Database
		Log lg.Log
	}
	tests := []struct {
		name   string
		fields fields
		want   *[]User
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				DB:  db.New(),
				Log: log,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MyBlogDB{
				DB:  tt.fields.DB,
				Log: tt.fields.Log,
			}
			d.DB.Connect()
			got := d.GetUserList()
			if len(*got) == 0 || (*got)[0].Email != "test@test.com" {
				t.Errorf("MyBlogDB.GetUserList() = %v, want %v", got, tt.want)
			}
			if len(*got) > 1 && (*got)[1].Email != "test2@test.com" {
				t.Errorf("MyBlogDB.GetUserList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyBlogDB_EnableUser(t *testing.T) {

	// db := gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "admin",
	// 	Password: "admin",
	// 	Database: "go_micro_blog",
	// }

	db := gdb.MyDBMock{
		Host:     "localhost:3306",
		User:     "admin",
		Password: "admin",
		Database: "go_micro_blog",
	}
	db.MockTestRow = &gdb.DbRow{
		//Row: []string{"0"},
		Row: []string{},
	}
	db.MockUpdateSuccess1 = true

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	type fields struct {
		DB  gdb.Database
		Log lg.Log
	}
	type args struct {
		uid int64
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
				DB:  db.New(),
				Log: log,
			},
			args: args{
				uid: 12,
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
			if got := d.EnableUser(tt.args.uid); got != tt.want {
				t.Errorf("MyBlogDB.EnableUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyBlogDB_DisableUser(t *testing.T) {

	// db := gdb.MyDB{
	// 	Host:     "localhost:3306",
	// 	User:     "admin",
	// 	Password: "admin",
	// 	Database: "go_micro_blog",
	// }


	db := gdb.MyDBMock{
		Host:     "localhost:3306",
		User:     "admin",
		Password: "admin",
		Database: "go_micro_blog",
	}
	db.MockTestRow = &gdb.DbRow{
		//Row: []string{"0"},
		Row: []string{},
	}
	db.MockUpdateSuccess1 = true

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)

	
	type fields struct {
		DB  gdb.Database
		Log lg.Log
	}
	type args struct {
		uid int64
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
				DB:  db.New(),
				Log: log,
			},
			args: args{
				uid: 12,
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
			if got := d.DisableUser(tt.args.uid); got != tt.want {
				t.Errorf("MyBlogDB.DisableUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
