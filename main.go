package main

import (
	"net/http"
	"time"
	"xian/router"
)

func main()  {
	//model.Setup()
	//gredis.Setup()
	r := router.InitRouter()

	s := &http.Server{
		Addr:           "0.0.0.0:8089",
		Handler:        r,
		ReadTimeout:    60*time.Second,
		WriteTimeout:   60*time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
