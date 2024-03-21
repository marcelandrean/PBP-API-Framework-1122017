package controllers

import (
	"net/http"
	"strconv"

	"echo-rest/db"
	m "echo-rest/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	var user m.User
	var users []m.User
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Email, &user.Password, &user.UserType)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		users = append(users, user)
	}

	var response m.Response
	response.Status = http.StatusOK
	response.Message = "Get all users successful"
	response.Data = users

	return c.JSON(http.StatusOK, response)
}

func InsertUser(c echo.Context) error {
	con := db.CreateCon()

	name := c.FormValue("name")
	age := c.FormValue("age")
	address := c.FormValue("address")
	email := c.FormValue("email")
	password := c.FormValue("password")
	usertype := c.FormValue("usertype")

	convUsertype, _ := strconv.Atoi(usertype)

	v := validator.New()

	peg := m.User{
		Name:     name,
		Age:      age,
		Address:  address,
		Email:    email,
		Password: password,
		UserType: convUsertype,
	}

	err := v.Struct(peg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sqlStatement := "INSERT users (name, age, address, email, password, usertype) VALUES (?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := stmt.Exec(name, age, address, email, password, usertype)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	var response m.Response
	response.Status = http.StatusOK
	response.Message = "Insert successful"
	response.Data = m.User{
		ID:       int(lastInsertedId),
		Name:     name,
		Age:      age,
		Address:  address,
		Email:    email,
		Password: password,
		UserType: convUsertype,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateUser(c echo.Context) error {
	con := db.CreateCon()

	id := c.FormValue("id")
	name := c.FormValue("name")
	age := c.FormValue("age")
	address := c.FormValue("address")
	email := c.FormValue("email")
	password := c.FormValue("password")
	usertype := c.FormValue("usertype")

	convId, _ := strconv.Atoi(id)
	convUsertype, _ := strconv.Atoi(usertype)

	v := validator.New()

	peg := m.User{
		Name:     name,
		Age:      age,
		Address:  address,
		Email:    email,
		Password: password,
		UserType: convUsertype,
	}

	err := v.Struct(peg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sqlStatement := "UPDATE users SET name = ?, age = ?, address = ?, email = ?, password = ?, usertype = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := stmt.Exec(name, age, address, email, password, usertype, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var response m.Response
	response.Status = http.StatusOK
	response.Message = "Update successful (rows affected: " + strconv.FormatInt(rowsAffected, 10) + ")"
	response.Data = m.User{
		ID:       convId,
		Name:     name,
		Age:      age,
		Address:  address,
		Email:    email,
		Password: password,
		UserType: convUsertype,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteUser(c echo.Context) error {
	con := db.CreateCon()

	id := c.FormValue("id")

	sqlStatement := "DELETE FROM users WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var response m.Response
	response.Status = http.StatusOK
	response.Message = "Delete successful"
	response.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return c.JSON(http.StatusOK, response)
}
