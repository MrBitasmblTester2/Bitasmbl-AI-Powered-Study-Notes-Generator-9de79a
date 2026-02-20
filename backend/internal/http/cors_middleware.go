package http
import "net/http"
func WithCORS(next http.Handler)http.Handler{return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){w.Header().Set("Access-Control-Allow-Origin","*");if r.Method==http.MethodOptions{return};next.ServeHTTP(w,r)})}