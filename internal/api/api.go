package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gmburov/service-autocompletion/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// AutocompleteResponse — ответ API автодополнения
type AutocompleteResponse struct {
	Suggestions []string `json:"suggestions"`
}

// Handler — HTTP handler для API автодополнения
type Handler struct {
	cfg *config.Config
}

// NewHandler создаёт новый Handler
func NewHandler(cfg *config.Config) *Handler {
	return &Handler{cfg: cfg}
}

// Router возвращает chi.Router с зарегистрированными маршрутами
func (h *Handler) Router() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Get("/autocomplete", h.autocomplete)

	return r
}

func (h *Handler) autocomplete(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.URL.Query().Get("q"))
	if query == "" {
		writeError(w, http.StatusBadRequest, "parameter q is required and cannot be empty")
		return
	}

	limit := h.cfg.DefaultLimit
	if s := r.URL.Query().Get("limit"); s != "" {
		n, err := strconv.Atoi(s)
		if err != nil || n < 1 || n > 100 {
			writeError(w, http.StatusBadRequest, "parameter limit must be an integer between 1 and 100")
			return
		}
		limit = n
	}

	// TODO: вызов pipeline — пока возвращаем пустой массив
	suggestions := []string{}
	if len(suggestions) > limit {
		suggestions = suggestions[:limit]
	}

	resp := AutocompleteResponse{Suggestions: suggestions}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		writeError(w, http.StatusInternalServerError, "internal server error")
		return
	}
}

func writeError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": message})
}
