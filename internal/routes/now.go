package routes

import (
	"net/http"
	"time"
)

func NowHandler(w http.ResponseWriter, r *http.Request) {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	now := time.Now().In(loc)

	w.Write([]byte(now.Format(time.RFC3339)))
}
