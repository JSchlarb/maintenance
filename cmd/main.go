package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	println("Starting ...")

	println("Exposing: /_health [No Content]")
	http.HandleFunc("/_health", healthHandler)

	println("Exposing: /** [503]")
	http.HandleFunc("/", anyRequestHandler)

	port := os.Getenv("MAINTENANCE_PORT")

	if port == "" {
		port = ":8081"
	}

	if !strings.Contains(port, ":") {
		port = fmt.Sprintf(":%s", port)
	}

	println(fmt.Sprintf("Listen to %s", port))
	err := http.ListenAndServe(port, nil)

	if err != nil {
		panic(err)
	}
	println("Started ...")
}

func anyRequestHandler(w http.ResponseWriter, r *http.Request) {
	println(fmt.Sprintf("Requested: %s", r.URL.Path))

	// copied from https://gist.github.com/pitch-gist/2999707
	html :=`
<!doctype html>
<title>Site Maintenance</title>
<style>
    body { text-align: center; padding: 150px; }
    h1 { font-size: 50px; }
    body { font: 20px Helvetica, sans-serif; color: #333; }
    article { display: block; text-align: left; width: 650px; margin: 0 auto; }
    a { color: #dc8100; text-decoration: none; }
    a:hover { color: #333; text-decoration: none; }
</style>

<article>
    <h1>We&rsquo;ll be back soon!</h1>
    <div>
        <p>Sorry for the inconvenience but we&rsquo;re performing some maintenance at the moment. We&rsquo;ll be back online shortly!</p>
    </div>
</article>
`
	w.WriteHeader(503)
	_, _ = w.Write([]byte(html))
}
func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(204)
}
