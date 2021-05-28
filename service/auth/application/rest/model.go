package rest

type Auth struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type AccessToken struct {
	AccessToken string `json:"access_token" binding:"required"`
}
