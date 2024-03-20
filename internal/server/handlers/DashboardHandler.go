package handlers

import (
	"fmt"
	"net/http"
)

var networkHtml = `
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>Dashboard</title>
    <link href="css/style.css" rel="stylesheet" />
    <script
      src="https://unpkg.com/htmx.org@1.9.10"
      integrity="sha384-D1Kt99CQMDuVetoL1lrYwg5t+9QdHe7NLX/SoJYkXDFfX37iInKRy5xLSi8nO7UC"
      crossorigin="anonymous"
    ></script>
  <script src="https://unpkg.com/@phosphor-icons/web"></script>
  </head>
  <body>
    <p
      hx-get="/watchers/network"
      hx-trigger="every 2s"
      hx-target="#network-results"
    hx-indicator="#spinner"
    >
      Query Results for the website
    </p>
<i class="ph ph-download-simple" id="spinner"></i>
    <div id="network-results">
  </div>
  </body>
</html>
`

func DashboardHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, networkHtml)
	}
}
