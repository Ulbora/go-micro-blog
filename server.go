package main

/*
 Copyright (C) 2023 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2023 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.

*/

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	lg "github.com/GolangToolKits/go-level-logger"
	gdb "github.com/GolangToolKits/go-mysql"
	mux "github.com/GolangToolKits/grrt"
	db "github.com/Ulbora/go-micro-blog/db"
	han "github.com/Ulbora/go-micro-blog/handlers"
	m "github.com/Ulbora/go-micro-blog/managers"
)

func main() {
	var dbHost string
	var dbUser string
	var dbPassword string
	var dbName string

	var apiAdminKey string
	var apiKey string

	if os.Getenv("DB_HOST") != "" {
		dbHost = os.Getenv("DB_HOST")
	} else {
		dbHost = "localhost:3306"
	}

	if os.Getenv("DB_USER") != "" {
		dbUser = os.Getenv("DB_USER")
	} else {
		dbUser = "admin"
	}

	if os.Getenv("DB_PASSWORD") != "" {
		dbPassword = os.Getenv("DB_PASSWORD")
	} else {
		dbPassword = "admin"
	}

	if os.Getenv("DB_DATABASE") != "" {
		dbName = os.Getenv("DB_DATABASE")
	} else {
		dbName = "go_micro_blog"
	}

	if os.Getenv("API_ADMIN_KEY") != "" {
		apiAdminKey = os.Getenv("API_ADMIN_KEY")
	} else {
		apiAdminKey = "54211789991515"
	}

	if os.Getenv("API_KEY") != "" {
		apiKey = os.Getenv("API_KEY")
	} else {
		apiKey = "557444414141"
	}

	var l lg.Logger
	log := l.New()
	log.SetLogLevel(lg.AllLevel)
	var mdb gdb.MyDB
	mdb.Host = dbHost
	mdb.User = dbUser
	mdb.Password = dbPassword
	mdb.Database = dbName
	mdb.Connect()

	var mbdb db.MyBlogDB
	mbdb.DB = mdb.New()
	mbdb.Log = log

	bdbi := mbdb.New()

	var man m.SysManager
	man.DB = bdbi
	man.Log = log

	bm := man.New()

	var bh han.MCHandler
	bh.DB = bdbi
	bh.Log = log
	bh.Manager = bm
	bh.APIKey = apiKey
	bh.APIAdminKey = apiAdminKey

	router := mux.NewRouter()
	port := "3000"
	envPort := os.Getenv("PORT")
	if envPort != "" {
		portInt, _ := strconv.Atoi(envPort)
		if portInt != 0 {
			port = envPort
		}
	}

	h := bh.New()

	router.HandleFunc("/rs/blog/add", h.AddBlog).Methods("POST")
	router.HandleFunc("/rs/blog/update", h.UpdateBlog).Methods("PUT")
	router.HandleFunc("/rs/blog/get/id/{id}", h.GetBlog).Methods("GET")
	router.HandleFunc("/rs/blog/get/name/{name}/{start}/{end}", h.GetBlogByName).Methods("GET")
	router.HandleFunc("/rs/blog/list/{start}/{end}", h.GetBlogList).Methods("GET")
	router.HandleFunc("/rs/blog/admin/list/{start}/{end}", h.GetAdminBlogList).Methods("GET")
	router.HandleFunc("/rs/blog/activate", h.ActivateBlog).Methods("PUT")
	router.HandleFunc("/rs/blog/deactivate", h.DectivateBlog).Methods("PUT")

	//router.HandleFunc("/rs/loglevel", h.SetLogLevel).Methods("POST")

	fmt.Println("Micor-Blog server is running on port " + port + "!")

	// log.SetLogLevel(lg.OffLevel)

	http.ListenAndServe(":"+port, router)
}

// go mod init github.com/Ulbora/go-micro-blog
