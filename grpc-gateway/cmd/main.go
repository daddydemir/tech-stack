package main

import (
	"fmt"
	"gateway/protogen/golang/users"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {

	user := users.User{
		UserName:  "daddydemir",
		Name:      "Mehmet",
		Surname:   "C.",
		Gender:    users.Gender_MAN,
		BirthDate: &date.Date{Year: 200, Month: 06, Day: 05},
	}

	bytes, err := protojson.Marshal(&user)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
