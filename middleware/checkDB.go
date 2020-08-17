package middleware

import (
	"net/http"

	"github.com/IanDex/twitter/db"
)

/*CheckDB f*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConection() == 0 {
			http.Error(w, "Conexión Perdida", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
