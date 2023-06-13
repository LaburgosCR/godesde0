package users

import (
	"fmt"
	"time"

	"github.com/godesde0/models"
)

func AlraUsuario() {
	/*models de tipo User y dentro de models User puedo usar todas sus propiedades*/
	usuario := new(models.User)
	usuario.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(usuario)

}
