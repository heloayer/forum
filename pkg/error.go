package pkg

import (
	"html/template"
	"log"
	"net/http"

	"forum/internal/model"
)

func ErrorHandler(w http.ResponseWriter, status int) {
	data := model.ErrorData{ 
		StatusText: http.StatusText(status), 
		StatusCode: status,                 
	}
	tmpl, err := template.ParseFiles("./templates/html/error.html") 
	if err != nil {                                                
		log.Println(err)                                                                              
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) 
		return
	}
	w.WriteHeader(status)       
	err = tmpl.Execute(w, data) 
	if err != nil {
		log.Println(err) 
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return

	}
}
