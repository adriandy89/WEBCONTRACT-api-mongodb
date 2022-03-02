package filescontroller

import (
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/errorservice"
	"WEBCONTRACT-api-mongodb/services/messageservice"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var id string = vars["id"]

	Openfile, err := os.Open("./docs/" + id) //Open the file to be downloaded later

	if err != nil {
		errorservice.ErrorMessage(w, "El Archivo no existe en el servidor", 400)
		return
	}
	defer Openfile.Close() //Close after function return

	tempBuffer := make([]byte, 512)                       //Create a byte array to read the file later
	Openfile.Read(tempBuffer)                             //Read the file into  byte
	FileContentType := http.DetectContentType(tempBuffer) //Get file header

	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	Filename := id

	//Set the headers
	w.Header().Set("Content-Type", FileContentType+";"+Filename)
	w.Header().Set("Content-Length", FileSize)

	Openfile.Seek(0, 0)  //We read 512 bytes from the file already so we reset the offset back to 0
	io.Copy(w, Openfile) //'Copy' the file to the client
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		errorservice.ErrorMessage(w, "Error leyendo el archivo enviado!", 400)
		return
	}

	defer file.Close()

	// Create file
	var dst *os.File
	var err2 error
	name := handler.Filename
	if !fileExist("./docs/" + name) {
		dst, err2 = os.Create("./docs/" + handler.Filename)
		if err2 != nil {
			errorservice.ErrorMessage(w, "Error en el servidor", 500)
			return
		}
	} else {
		name = "(1)" + handler.Filename
		for fileExist("./docs/" + name) {
			name = getNewName(name)
		}
		dst, err2 = os.Create("./docs/" + name)
		if err2 != nil {
			errorservice.ErrorMessage(w, "Error en el servidor", 500)
			return
		}
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		errorservice.ErrorMessage(w, "Error en el servidor", 500)
		return
	}
	defer dst.Close()

	messageservice.SuccesMessage(w, name, 200)
	//fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func fileExist(ruta string) bool {
	if _, err := os.Stat(ruta); os.IsNotExist(err) {
		return false
	}
	return true
}

func getNewName(name string) string {
	if len(name) > 3 {
		if name[2] == ')' && name[0] == '(' {
			numb, err := strconv.Atoi(string(name[1]))
			if err != nil {
				name = "(1)" + name
			} else {
				num := fmt.Sprint(numb + 1)
				temp := RemoveChar(name, 3)
				name = "(" + string(num) + ")" + temp

			}
		} else if name[3] == ')' && name[0] == '(' {
			numb, err := strconv.Atoi(string(name[1]))
			numb2, err2 := strconv.Atoi(string(name[2]))
			if err != nil || err2 != nil {
				name = "(1)" + name
			} else {
				num2 := ""
				num := ""
				if numb2 == 9 {
					num2 = fmt.Sprint(0)
					num = fmt.Sprint(numb + 1)
				} else {
					num2 = fmt.Sprint(numb2 + 1)
					num = fmt.Sprint(numb)
				}
				temp := RemoveChar(name, 4)
				name = "(" + string(num) + string(num2) + ")" + temp

			}
		} else {
			name = "(1)" + name
		}
	} else {
		name = "(1)" + name
	}
	return name
}

func RemoveChar(input string, cant int) string {
	if len(input) <= 1 {
		return ""
	}
	return input[cant:]
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" || rol == "Gestionador" {

		var fileName models.Word
		err := json.NewDecoder(r.Body).Decode(&fileName)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}
		errDelete := os.Remove("./docs/" + fileName.Word)

		if errDelete != nil {
			errorservice.ErrorMessage(w, "Error elimando archivo"+errDelete.Error(), 500)
			return
		} else {
			messageservice.SuccesMessage(w, "Archivo elimando correctamente", 200)
			return
		}
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}
