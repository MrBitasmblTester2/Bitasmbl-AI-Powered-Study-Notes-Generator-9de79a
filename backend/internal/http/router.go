package http
import("net/http""backend/internal/services")
func NewRouter(s services.AIService)*http.ServeMux{mux:=http.NewServeMux();mux.HandleFunc("/api/notes",GenerateNotes);return mux}