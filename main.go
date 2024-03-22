package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func HandlerBD(w http.ResponseWriter, r *http.Request) {
	conn, err := ConnectDB("MYURL")
	if err != nil {
		fmt.Print("HAHA")
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	numbers, err := GetManyRowsByLogin(conn, "Ivan")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range numbers {
		message := []byte(v)
		_, err = w.Write(message)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func HandlerTime(w http.ResponseWriter, r *http.Request) {
	time2 := time.Now()
	// time3 := time2.Format(time.RFC822)
	w.Write([]byte(time2.Format(time.RFC822)))
}

type Mes struct {
	message string
}

func (m Mes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, m.message)
	w.Write([]byte(m.message))

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/db", HandlerBD)
	mux.HandleFunc("/time", HandlerTime)
	mux.Handle("/hello", Mes{message: "om a hum"})
	http.ListenAndServe("localhost:8080", mux)

}
