package oauth

type OauthClient interface {
	GetToken(code string) string
	GetUserInfo(accessToken string) *OauthUserInfo
}

type OauthUserInfo struct {
	Email           string `json:"email"`
	Nickname        string `json:"nickname"`
	ProfileImageUrl string `json:"profileImageUrl"`
}