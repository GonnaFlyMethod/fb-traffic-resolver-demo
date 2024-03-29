package server

import "net/http"

func ResponseAsJSONMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func CorsMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		allowedHeaders := "Accept, Accept-Language, Content-Type, Content-Length, Authorization"
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
