package oauth

type OAuthDomain struct {}

func (f *OAuthDomain) Name() string {
	return "an OAuth provider"
}

func (f *OAuthDomain) DomainId() string {
	return "oauth/null"
}