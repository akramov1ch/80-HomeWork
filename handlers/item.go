package handlers

import (
	"net/http"
	"80HW/db"
	"80HW/model"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, title, body FROM items")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var items []model.Item
	for rows.Next() {
		var item model.Item
		if err := rows.Scan(&item.ID, &item.Title, &item.Body); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		items = append(items, item)
	}

	c.JSON(http.StatusOK, items)
}

func GetItem(c *gin.Context) {
	id := c.Param("id")
	row := db.DB.QueryRow("SELECT id, title, body FROM items WHERE id = $1", id)

	var item model.Item
	if err := row.Scan(&item.ID, &item.Title, &item.Body); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func CreateItem(c *gin.Context) {
	var item model.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.DB.QueryRow(
		"INSERT INTO items (title, body) VALUES ($1, $2) RETURNING id",
		item.Title, item.Body,
	).Scan(&item.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var item model.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.DB.Exec(
		"UPDATE items SET title = $1, body = $2 WHERE id = $3",
		item.Title, item.Body, id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")

	_, err := db.DB.Exec("DELETE FROM items WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
