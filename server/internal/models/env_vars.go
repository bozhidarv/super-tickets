package models

type EnvVars struct {
	dbUrl string
	port string
	jwtKey string
}

func (e *EnvVars) SetDbUrl(url string) {
	e.dbUrl = url
}

func (e *EnvVars) DbUrl() string {
	return e.dbUrl
}

func (e *EnvVars) SetPort(port string) {
	e.port = port
}

func (e *EnvVars) Port() string {
	return e.port
}

func (e *EnvVars) SetJwtKey(key string) {
	e.jwtKey = key
}
func (e *EnvVars) JwtKey() string {
	return e.jwtKey
}

type ReadOnlyEnvVars struct {
	EnvVars
}

func (e *ReadOnlyEnvVars) SetDbUrl(url string) {
	panic("SetDbUrl is called on ReadOnlyEnvVars")
}

func (e *ReadOnlyEnvVars) SetPort(url string) {
	panic("SetPort is called on ReadOnlyEnvVars")
}

func (e *ReadOnlyEnvVars) SetJwtKey(url string) {
	panic("SetJwtKey is called on ReadOnlyEnvVars")
}


