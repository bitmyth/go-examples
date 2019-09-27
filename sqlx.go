package main
import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id int
}

func main(){
	db, err := sqlx.Connect("mysql", "root:iqvUA3AydwxHQkEoiDQ7hDTBfyKxFkyR@tcp(127.0.0.1:3307)/hsh_development")
	if err != nil {
		fmt.Print(err)
		return
	}
	var user []User
	err = db.Select(&user, "select caid as id from cashiers")
	fmt.Print("users", len(user))
}
