package model

type User struct {
	UserID	string	`bson:"userId" json:"userId" validate:"required"`
	Area	[]map[string]string  `bson:"area" json:"area"`
}