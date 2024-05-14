package types

type LoginRequest struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}


type ErrorResponse struct {
    Error string `json:"error"`
}