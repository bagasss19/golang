package helper

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(usrPwd string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usrPwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CompareHashAndPassword(usrPwd string, dbPwd string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(dbPwd), []byte(usrPwd))
	if err != nil {
		return false, err
	}
	return true, nil
}
