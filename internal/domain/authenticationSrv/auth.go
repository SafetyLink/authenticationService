package authenticationSrv

type Authentication interface {
}

type AuthenticationSrv struct {
}

func NewAuthenticationService() Authentication {
	return &AuthenticationSrv{}
}

func (a *AuthenticationSrv) Login() {

}
