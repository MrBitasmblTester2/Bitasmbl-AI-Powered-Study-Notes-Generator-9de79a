package main
import "net/http"
func main(){mux:=http.NewServeMux();mux.HandleFunc("/health",func(w http.ResponseWriter,r *http.Request){w.Write([]byte("ok"))});http.ListenAndServe(":8080",mux)}