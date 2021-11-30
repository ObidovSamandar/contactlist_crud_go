package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/obidovsamandar/task-list-crud/dbconnector"
)

type Contact struct {
	ID        int64
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

func CreateContact(c *gin.Context) {
	var createContactBody Contact
	var id int
	if err := c.ShouldBindJSON(&createContactBody); err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Something went wrong while take request body!",
		})
		return
	}

	err := dbconnector.DBClient.Get(&id, "SELECT nextval($1)", "contactlistSquence")

	if err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "ID not found!",
		})
		return
	}

	_, err = dbconnector.DBClient.Exec("INSERT INTO contactlistdb (id, firstname, lastname, email, phone) VALUES($1,$2,$3,$4,$5);", id, createContactBody.FirstName, createContactBody.LastName, createContactBody.Email, createContactBody.Phone)

	fmt.Println(err)
	if err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Something went wrong while creating contact !",
		})
		return
	}

	c.JSON(201, gin.H{
		"error":   false,
		"message": "Contact created!",
	})
}

func DeleteContact(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)

	contact := Contact{}

	fmt.Println(id)

	err := dbconnector.DBClient.Get(&contact, "SELECT * FROM contactlistdb WHERE id=$1", id)

	if err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Contact which searching is not found!",
		})

		return

	}
	res, err := dbconnector.DBClient.Exec("DELETE FROM contactlistdb WHERE id=$1", id)

	if err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Something went wrong while deleting contact!",
		})
		return
	}

	count, _ := res.RowsAffected()
	fmt.Println(count)
	if count == 0 {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Contact did not deleted!",
		})
		return
	}

	c.JSON(200, gin.H{
		"error":   false,
		"message": "Contact deleted!",
		"id":      id,
	})

}

func GetAllContacts(c *gin.Context) {

	allContacts := []Contact{}
	err := dbconnector.DBClient.Select(&allContacts, "SELECT * FROM contactlistdb;")

	if err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Something went wrong while getting all contacts!",
		})
		return
	}

	c.JSON(200, gin.H{
		"error": false,
		"data":  allContacts,
	})
}

func GetSpecificContact(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)

	contact := Contact{}

	err := dbconnector.DBClient.Get(&contact, "SELECT * FROM contactlistdb WHERE id=$1", id)

	if err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Contact which you are finding is not found !",
		})
		return
	}

	c.JSON(200, gin.H{
		"error": false,
		"data":  contact,
	})
}

func UpdateContact(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)

	var requestBody Contact

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Please check your request body",
		})
	}

	contact := Contact{}

	err := dbconnector.DBClient.Get(&contact, "SELECT * FROM contactlistdb WHERE id=$1", id)

	if err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Contact which updating is not found!",
		})
		return
	}

	fmt.Println(contact)

	res, err := dbconnector.DBClient.Exec("UPDATE contactlistdb SET firstname=$1 lastname=$2 email=$3 phone=$4 WHERE id=$5", requestBody.FirstName, requestBody.LastName, requestBody.Email, requestBody.Phone, id)

	if err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Something whent wrong while updating contact!",
		})
		return
	}

	count, err := res.RowsAffected()

	if err != nil || count == 0 {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Contact did not updated!",
		})
	}

	c.JSON(200, gin.H{
		"error":   false,
		"message": "Contact updated",
		"data":    contact,
	})

}
