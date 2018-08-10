package mgw

// User model
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

// EdgeProxy model
type EdgeProxy struct {
	ID        string `json:"id"`
	ProxyName string `json:"proxyName"`
	TargetURL string `json:"targetURL"`
}
