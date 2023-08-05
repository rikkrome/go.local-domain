package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Specific handling for `go.rikkrome/tokyo`.
	http.HandleFunc("/tokyo/", func(w http.ResponseWriter, r *http.Request) {
		t := `<!DOCTYPE html>
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
        <meta name="go-import" content="go.rikkrome/tokyo git https://github.com/rikkrome/tokyo">
    </head>
    <body>
        Hello Go Conference!
    </body>
</html>`
		w.Write([]byte(t))
		w.WriteHeader(http.StatusOK)
	})

	// Defaults to no meta tag page.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := `<!DOCTYPE html>
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    </head>
    <body>
        This is the default page, and does not have meta tag with "go-import".
    </body>
</html>`
		w.Write([]byte(t))
		w.WriteHeader(http.StatusOK)
	})

	// If you need to test this locally, do the following:
	//   - Run `mkcert go.rikkrome` in this directory
	//   - Update `/etc/hosts` to add the following line
	//       127.0.0.1 go.rikkrome
	//     This adds a dummy address of `go.rikkrome` to loopback to your
	//     localhost.
	//   - Use the following `http.ListenAndServeTLS` to set up TLS server
	//     instead of simple `http.ListenAndServe`
	//   - Run this file with `go run ./main.go`
	log.Fatal(http.ListenAndServeTLS(":"+os.Getenv("PORT"), "go.rikkrome.pem", "go.rikkrome-key.pem", nil))
	// If you are deploying this on some server with extra TLS setup, you would
	// need to use the following.
	// log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
