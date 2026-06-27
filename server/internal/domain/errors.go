package domain


import "errors"

var (
	DBErrorUserNotFound = errors.New("Пользователь не найден в базе")
	DBErrorInvalidCreds = errors.New("Неправильные логин или пароль")
	DBErrorUserAlreadyExists = errors.New("Пользователь уже есть")
)