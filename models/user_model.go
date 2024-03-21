package models

import (
	"echo-rest/db"
	"net/http"
	"strconv"

	validator "github.com/go-playground/validator/v10"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" validate:"required"`
	Age      string `json:"age" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	UserType int    `json:"usertype" validate:"required"`
}

func FetchAllUsers() (Response, error) {
	var obj User
	var arrobj []User
	var res Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "SELECT * FROM users"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.ID, &obj.Name, &obj.Age, &obj.Address, &obj.Email, &obj.Password, &obj.UserType)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func StoreUser(name string, age string, address string, email string, password string, usertype int) (Response, error) {
	var res Response

	v := validator.New()

	peg := User{
		Name:     name,
		Age:      age,
		Address:  address,
		Email:    email,
		Password: password,
		UserType: usertype,
	}

	err := v.Struct(peg)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "INSERT users (name, age, address, email, password, usertype) VALUES (?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, age, address, email, password, usertype)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]interface{}{
		"id":       lastInsertedId,
		"name":     name,
		"age":      age,
		"address":  address,
		"email":    email,
		"password": password,
		"usertype": usertype,
	}

	return res, nil
}

func UpdateUser(id int, name string, age string, address string, email string, password string, usertype int) (Response, error) {
	var res Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "UPDATE users SET name = ?, age = ?, address = ?, email = ?, password = ?, usertype = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, age, address, email, password, usertype, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Update successful (rows affected: " + strconv.FormatInt(rowsAffected, 10) + ")"
	res.Data = map[string]interface{}{
		"id":       id,
		"name":     name,
		"age":      age,
		"address":  address,
		"email":    email,
		"password": password,
		"usertype": usertype,
	}

	return res, nil
}

func DeleteUser(id int) (Response, error) {
	var res Response

	con := db.CreateCon()
	defer con.Close()

	sqlStatement := "DELETE FROM users WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
