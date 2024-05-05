package health

import (
	"ChelsikBot/internal/app"
	"net/http"
)

func Start(a *app.App) {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if a.Health() {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("OK"))
		}
	})

	_ = http.ListenAndServe(":8080", nil)
}
