package main

import (
	"fmt"
	"net/http"

	"gin-blog/pkg/setting"
	"gin-blog/routers"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HEda fuck")
}

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}


	s.ListenAndServe()
}
