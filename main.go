package main

import (
	"github.com/astaxie/beego"
	"gopkg.in/go-ini/ini.v1"
	"github.com/lib/pq"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/viper"
)

func main() {
	println("Beego version:", beego.VERSION)
	println("ini version:", ini.Version)
	var a,e = pq.ParseURL("postgres://bob:secret@1.2.3.4:5432/mydb?sslmode=verify-full")
	println("parse a: ", a);
	println("parse e: ", e);
	gin.Default();
	viper.SetDefault("ContentDir", "content")
	proto.String("hello")
}