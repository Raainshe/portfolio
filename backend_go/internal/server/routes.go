package server

import (
	"backend_go/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"fmt"
	"time"

	"github.com/coder/websocket"
	"github.com/gorilla/mux"
)

func (s *Server) RegisterRoutes() http.Handler {
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/", s.HelloWorldHandler).Methods("GET")
	router.HandleFunc("/health", s.healthHandler).Methods("GET")
	router.HandleFunc("/websocket", s.websocketHandler).Methods("GET")

	// User routes with path parameters
	router.HandleFunc("/api/users/signup", s.userSignupHandler).Methods("POST")
	router.HandleFunc("/api/users/{id}", s.getUserHandler).Methods("GET") // This supports {id}
	router.HandleFunc("/api/visitors", s.trackVisitorHandler).Methods("POST")

	// Wrap the router with CORS middleware
	return s.corsMiddleware(router)
}

func (s *Server) trackVisitorHandler(w http.ResponseWriter, r *http.Request) {

	ip := getClientIP(r)
	userAgent := r.UserAgent()

	//track the visit
	visitor, err := s.db.TrackVisitor(ip, userAgent)
	if err != nil {
		log.Printf("Failed to track visitor: %v", err)
		http.Error(w, "Failed to track visitor", http.StatusInternalServerError)
	}

	res := map[string]interface{}{
		"success":  true,
		"message":  "Visit tracked successfully",
		"location": visitor.Location,
	}

	w.Header().Set("Contet-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func getClientIP(r *http.Request) string {
	headers := []string{
		"X-Forwarded-For",
		"X-Real-IP",
		"X-Client-IP",
		"CF-Connecting-IP",
	}

	for _, header := range headers {
		if ip := r.Header.Get(header); ip != "" {
			//X forward for can contain multiple ips
			if strings.Contains(ip, ",") {
				ip = strings.Split(ip, ",")[0]
			}
			ip = strings.TrimSpace(ip)
			if ip != "" {
				return ip
			}
		}
	}

	//fallback to remote address
	if r.RemoteAddr != "" {
		if strings.Contains(r.RemoteAddr, ":") {
			return strings.Split(r.RemoteAddr, ":")[0]
		}
		return r.RemoteAddr
	}

	return "unknown"
}

func (s *Server) getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	//get user from database
	user, err := s.db.GetUser(userID)
	if err != nil {
		log.Printf("Failed to get user: %w", err)
		http.Error(w, "User not found", http.StatusNotFound)
	}

	//return the user since we have them
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("failed to encode response: %w", err)
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}
}

func (s *Server) userSignupHandler(w http.ResponseWriter, r *http.Request) {
	//Only allow Post method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//parse request body
	var signupReq models.SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&signupReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//now lets create the user model

	user := models.User{
		Username:  signupReq.Username,
		Email:     signupReq.Email,
		Password:  signupReq.Password,
		FirstName: signupReq.FirstName,
		LastName:  signupReq.LastName,
	}

	userID, err := s.db.CreateUser(user)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	//return success response
	response := models.SignupResponse{
		Message: "User Created succesfully",
		UserID:  userID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("failed to encode response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins if needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are required

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message": "Hello World"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(s.db.Health())
	if err != nil {
		http.Error(w, "Failed to marshal health check response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) websocketHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := websocket.Accept(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to open websocket", http.StatusInternalServerError)
		return
	}
	defer socket.Close(websocket.StatusGoingAway, "Server closing websocket")

	ctx := r.Context()
	socketCtx := socket.CloseRead(ctx)

	for {
		payload := fmt.Sprintf("server timestamp: %d", time.Now().UnixNano())
		if err := socket.Write(socketCtx, websocket.MessageText, []byte(payload)); err != nil {
			log.Printf("Failed to write to socket: %v", err)
			break
		}
		time.Sleep(2 * time.Second)
	}
}
