package contracts

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthUser struct {
	Email string `json:"email"`
}
type SignUpResponse struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	User         AuthUser `json:"user"`
}
