package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
			<head>
				<title>Hello, World!</title>
			</head>
			<body>
				<h1>Hello, World!</h1>
			</body>
		</html>`,
	)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe("localhost:19991", nil)
}
