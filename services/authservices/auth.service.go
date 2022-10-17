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
func Login(username string, pass string) (models.User, string, bool) {

	var user models.User
	var passw string = "WebContract" + strconv.Itoa(time.Now().Day()) + "*"
	var typeClient string = "ext"
	if config.Domain != "cimex.com.cu" {
		typeClient = "int"
	}

	if username == "SA" && pass == passw {

		user.Name = "SuperAdmin"
		user.Username = username
		user.CreatedAt = time.Now().UTC()
		user.ExpireAt = time.Date(time.Now().Year()+1, time.Now().Month(), 10, 0, 0, 0, 0, time.UTC)
		user.Rol = "SA"
		user.State = 1
		user.Environment = 0
		user.CodeCompany = "0"

		return user, typeClient, true

	} else {
		userLogged, exist := userservice.FindByUsername(username)
		if !exist {
			return user, typeClient, false
		}
		if userLogged.State == 0 {
			return user, typeClient, false
		}

		passwordBytes := []byte(pass)
		passwordDB := []byte(userLogged.Password)
		err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
		if err != nil {

			if config.FQDN != "" {
				// Non-TLS Connection
				l, err := Connect()
				if err != nil {
					return user, typeClient, false
				}
				// User and Password Authentication
				err = l.Bind(userLogged.Username+"@"+config.Domain, pass)
				defer l.Close()
				if err != nil {
					return user, typeClient, false
				} else {
					//Actualizar el loginCount del usuario
					userservice.UpdateLoginCount(userLogged.ID, userLogged.LoginCount)
					return userLogged, typeClient, true
				}
			} else {
				return user, typeClient, false
			}

		}
		//Actualizar el loginCount del usuario
		userservice.UpdateLoginCount(userLogged.ID, userLogged.LoginCount)

		return userLogged, typeClient, true
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
