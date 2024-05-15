package users

import (
	"fmt"
	"github.com/hannder92/godesde0/modelos"
	"time"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(1, "Pablo", 31, time.Now(), true)
	fmt.Println(u)
}
