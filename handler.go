package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const pattern = `<!DOCTYPE html>
<htnl lang="ru"><>head
<meta charset="utf-8"> />
<title>Тестовый сервер</title>
</head>
<body>%s</body></html>`

func mainHandle(w http.ResponseWriter, req *http.Request) {
	var answer string

	name := req.URL.Query().Get("name")
	if len(name) == 0 {
		answer = "Укажите имя заголовка"
	} else if v := req.Header.Get(name); len(v) > 0 {
		answer = fmt.Sprintf("%s %s hello Denis", name, v)
	} else {
		answer = fmt.Sprintf("Заголовок %s не определён", name)
	}
	io.WriteString(w, answer)
}

func mainHandle1(res http.ResponseWriter, req *http.Request) {
	var out string

	if req.URL.Path == `/time` || req.URL.Path == `/time/` {
		out = time.Now().Format("2024-12-09 17:56")
	} else {
		out = fmt.Sprintf("Host: %s\nPath: %s\nMethod: %s",
			req.Host, req.URL.Path, req.Method)
	}
	s := []byte(out)
	res.Write(s)

}

func Handle2(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		fmt.Fprintf(w, "Email: %s\nName: %s",
			req.PostFormValue("email"), req.PostFormValue("name"))
		return
	}
	io.WriteString(w, "Отправьте POST запрос")
}

func Handle3(w http.ResponseWriter, req *http.Request) {
	// var answer string

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, pattern, "Сервер не поддерживает Get-запросы")
		return
	}
	fmt.Fprintf(w, pattern, "Получен Get-запрос")
}

func handleTime(w http.ResponseWriter, req *http.Request) {
	s := time.Now().Format("02.01.2006 15:26:03")
	w.Write([]byte(s))

}
func handleMain(w http.ResponseWriter, req *http.Request) {
	s := fmt.Sprintf("Method: %s\nHost: %s\nPath: %s",
		req.Method, req.Host, req.URL.Path)
	w.Write([]byte(s))

}

func handlerHello(w http.ResponseWriter, req *http.Request) {
	message := []byte("Hello, Web!")
	_, err := w.Write(message)
	if err != nil {
		log.Fatal(err)
	}

}

func handlerOmahum(w http.ResponseWriter, req *http.Request) {
	writeHandler(w, "om a hum")

}

func writeHandler(w http.ResponseWriter, message string) {
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}
func handlerHindi(w http.ResponseWriter, req *http.Request) {
	writeHandler(w, "Namaste")
}
