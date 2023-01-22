package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/aWildProgrammer/fconf"
	_ "github.com/lib/pq"
	"github.com/solenovex/it/common"
	"github.com/solenovex/it/controller"
	"github.com/solenovex/it/middleware"

	_ "net/http/pprof"

	// lhh
	_ "github.com/go-sql-driver/mysql"
)

// const (
//
//	host     = ""
//	port     = 5432
//	user     = ""
//	password = ""
//	dbname   = "demo"
//
// )

type mysql_Conf struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func runOrBuildLoadConfig() (mysql_res mysql_Conf) {
	// 使用库 go get github.com/aWildProgrammer/fconf
	c, err := fconf.NewFileConf("./conf/conf_mysql.ini")
	if err != nil {
		panic(err)
	}
	// fmt.Println(c.Int("mysql.port"))
	mysql_res.host = c.String("mysql.host")
	mysql_res.port = c.String("mysql.port")
	mysql_res.user = c.String("mysql.user")
	mysql_res.password = c.String("mysql.password")
	mysql_res.dbname = c.String("mysql.database")
	// fmt.Printf("%v \n", mysql_res)
	return
}
func init() {
	var err error

	res := runOrBuildLoadConfig() // 读取配置文件
	// fmt.Printf("%v \n", res)
	// connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)
	// common.Db, err = sql.Open("postgres", connStr)
	// common.Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3309)/test")
	str_sql := res.user + ":" + res.password + "@" + "tcp(" + res.host + ":" + res.port + ")/" + res.dbname
	fmt.Println(str_sql)
	common.Db, err = sql.Open("mysql", str_sql)
	if err != nil {
		log.Fatalln(err.Error())
	}

	ctx := context.Background()
	err = common.Db.PingContext(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Connected!")
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: &middleware.BasicAuthMiddleware{},
	}

	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./wwwroot")))) // 自定义 css link 不起作用
	http.Handle("/wwwroot/static/", http.StripPrefix("/wwwroot/static/", http.FileServer(http.Dir("wwwroot/static/")))) // 有效
	controller.RegisterRoutes()

	log.Println("Server starting...")
	go http.ListenAndServe(":8000", nil)
	server.ListenAndServe()
}
