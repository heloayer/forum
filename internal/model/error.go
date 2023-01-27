package model

type ErrorData struct { // структура заполняется из функции ErrorHandler из пакета pkg/error
	StatusText string // текст кода статуса
	StatusCode int    // сам код статуса
}
