package service_handler

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"

    "db_handler"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateNewKodePeserta(role string) string {
	var prefix, kodepeserta string
	switch role{
	case "umum":
		prefix = "PU"
	case "khusus":
		prefix = "PK"
	case "grup":
		prefix = "PG"
	}

	if !db_handler.CheckRoleExist(role) {
		kodepeserta = prefix + "0001"
	}else{
		counter_role := db_handler.GetNumberRegistered(role)
		counter_role++
        counter_role_string := strconv.Itoa(counter_role)
        num_of_zero := 4 - len(counter_role_string)
        kodepeserta = prefix
        for i:=0; i<num_of_zero; i++ {
        	kodepeserta += "0"
        }
        kodepeserta += counter_role_string
	}

	return kodepeserta
}