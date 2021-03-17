package http

import (
	"log"
	"net/http"
	"time"
)

func HttpServerV1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("httpserver v1"))
	})
	http.HandleFunc("/bye", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("bye bye ,this is v1 httpServer"))
	})
	log.Println("Starting v1 server ...")
	log.Fatal(http.ListenAndServe(":1210", nil))
}





type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is version 2"))
}
func sayBye(w http.ResponseWriter, r *http.Request) {
	time.Sleep(4 * time.Second)
	w.Write([]byte("bye bye ,this is v2 httpServer"))
}

func HttpServerV2(){
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)

	log.Println("Starting v2 httpserver")
	log.Fatal(http.ListenAndServe(":1210", mux))
}

func HttpServerV3(){
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)

	server := &http.Server{
		Addr:         ":1210",
		WriteTimeout: time.Second * 3,            //设置3秒的写超时
		Handler:      mux,
	}
	log.Println("Starting v3 httpserver")
	log.Fatal(server.ListenAndServe())
}

