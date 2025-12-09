package routes

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func DiceHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(6) + 1
	w.Write([]byte((fmt.Sprintf("%d", n))))
}
