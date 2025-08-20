package utils

import "errors"

func Login(email string, password string) (string, error) {
	if email != "admin@gmail.com" && password != "admin123" {
		return "", errors.New("Email dan Password Salah !!")
	} else if email != "admin@gmail.com" && password == "admin123" {
		return "", errors.New("Email Salah !!")
	} else if email == "admin@gmail.com" && password != "admin123" {
		return "", errors.New("Password Salah !!")
	}

	return "Anda Berhasil Login !!", nil
}
