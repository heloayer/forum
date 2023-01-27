package pkg

import (
	"html/template"
	"log"
	"net/http"

	"forum/internal/model"
)

func ErrorHandler(w http.ResponseWriter, status int) {
	data := model.ErrorData{ // data (struct) будет исп. для вставки в html шаблон error.html, для вывода визуализированного вывода ошибки
		StatusText: http.StatusText(status), // возвращает текст для статуса кода
		StatusCode: status,                  // сам код статуса
	}
	tmpl, err := template.ParseFiles("./templates/html/error.html") // исп. содержимое файла error.html для создания знач. Template
	if err != nil {                                                 // if upon creating Teamplte is an error (for visualisation error for not having correct URL path)
		log.Println(err)                                                                               // принтим в консоль
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) // responses to request with error string (calls for StatusText) and code status (int)
		return
	}
	w.WriteHeader(status)       // sends http response header with provided status code
	err = tmpl.Execute(w, data) // в шаблон вставляются данные data (error string, status code)
	if err != nil {
		log.Println(err) // принтим в консоль
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return

	}
}
