package validators

type oauthValidatorFunc func(token string) bool
type oauthValidators map[string]oauthValidatorFunc

func (ouv oauthValidators) register(providerName string, f oauthValidatorFunc) {
	OAuthValidators[providerName] = f
}

var OAuthValidators = make(oauthValidators)
