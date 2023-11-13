package models

type User struct {
	PassKey string `json:"passkey"`
	Host    string `json:"host"`
}

func (u *User) GetPassKey() string {
	if u == nil {
		return ""
	}
	return u.PassKey
}

func (u *User) GetHost() string {
	if u == nil {
		return ""
	}
	return u.Host
}
