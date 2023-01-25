package ctl

func Register(username string, password string) string {
	if username == "" || password == "" {
		return "null"
	} 
	return username + password
}

func Login(username string, password string) bool{
	u := "admin"
	p := "pass"
	if username == "" || password == "" {
		return false
	} 
	if username == u || password == p {
		return true
	}
	return false
}