package request

// AuthLoginReq login request struct
type AuthLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
