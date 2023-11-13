package models

// [
// {
// 	"host": "github",
// 	"profiles":[
// 		{
// 			"active":true,
// 			"id":"test name",
// 			"passkey":"asddsadds"
// 		}
// 	]
// },
// {
// 	"host": "gitlab",
// 	"profiles":[
// 		{
// 			"active":true,
// 			"id":"test name",
// 			"passkey":"asddsadds"
// 		}
// 	]
// }
// ]

type Config struct {
	Host     string     `json:"host"`
	Profiles []*Profile `json:"profiles"`
}

type Profile struct {
	Active   bool   `json:"active"`
	Name     string `json:"name"`
	Id       string `json:"id"`
	PassKey  string `json:"passkey"`
	Email    string `json:"email"`
	UserName string `json:"username"`
}
