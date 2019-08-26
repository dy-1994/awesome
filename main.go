package main

import (
	"github.com/labstack/echo"
	"fmt"
	"errors"
	"awesome/middleware"
	"github.com/golang/glog"
	"awesome/network"
)

func main() {
	go StartEcho()
	server := network.Tcp{}
	err := server.Service(1234)
	if err != nil {
		glog.Fatal(err)
	}

	//ch := network.NewChannel(conn)
	//poller.Start(conn, netpoll.EventRead, func() {
	//	go Receive(ch)
	//})
}

func StartEcho() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.NewMiddleware())
	e.GET("/", func(context echo.Context) error {
		fmt.Println("call /")
		return nil
	})
	e.GET("login", login)

	err := e.Start(":9090")
	if err != nil {
		glog.Errorln(err)
	}
}


func login(context echo.Context) error {
	return fmt.Errorf("%d", 1+2)
}

func division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}