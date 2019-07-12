package main

import (
	"CardTaskGo/conf"
	"CardTaskGo/routers"
	"fmt"
	"net/http"
)

func main() {

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.HTTP_PORT),
		Handler:        routers.InitRouter(),
		ReadTimeout:    conf.READ_TIMEOUT,
		WriteTimeout:   conf.WRITE_TIMEOUT,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
