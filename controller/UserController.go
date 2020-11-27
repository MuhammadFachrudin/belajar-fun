package controller

import (
	"golang-mysql/table"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (db *InDB) AddUser(c *gin.Context) {
	var (
		user   table.User
		result gin.H
	)

	nama_depan := c.PostForm("nama_depan")
	nama_belakang := c.PostForm("nama_belakang")
	// alamat_id := c.DefaultPostForm("alamat_id", 1)

	user.NamaDepan = nama_depan
	user.NamaBelakang = nama_belakang
	user.AlamatID = 1

	db.DB.Create(&user)

	result = gin.H{
		"result": user,
	}

	c.JSON(http.StatusOK, result)
}

func (db *InDB) GetAllUser(c *gin.Context) {
	var (
		user   []table.User
		result gin.H
	)

	db.DB.Find(&user)
	if len(user) <= 0 {
		result = gin.H{
			"result": user,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": user,
			"count":  len(user),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (db *InDB) GetUser(c *gin.Context) {
	var (
		user   table.User
		result gin.H
	)

	id := c.Param("id")
	db.DB.First(&user, id)
	if &user != nil {
		result = gin.H{
			"result": user,
		}
	} else {
		result = gin.H{
			"code":   404,
			"result": "Data not found",
		}
	}
	c.JSON(http.StatusOK, result)
}

func (db *InDB) UpdateUser(c *gin.Context) {
	var (
		user      table.User
		newUpdate table.User
		result    gin.H
	)

	id := c.Query("id")
	nama_depan := c.PostForm("nama_depan")
	nama_belakang := c.PostForm("nama_belakang")
	alamat_id := c.PostForm("alamat_id")

	err := db.DB.First(&user, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	id_alamat, _ := strconv.Atoi(alamat_id) //Convert to Integer

	if (nama_belakang != "") && (nama_depan != "") {
		newUpdate.NamaDepan = nama_depan
		newUpdate.NamaBelakang = nama_belakang
		newUpdate.AlamatID = id_alamat

		err = db.DB.Model(&user).Updates(newUpdate).Error
		if err == nil {
			result = gin.H{
				"result": "update successfully",
			}
		} else {
			result = gin.H{
				"result": "update failed",
			}
		}
	} else {
		result = gin.H{
			"result": "cannot empty field",
		}
	}
	c.JSON(http.StatusOK, result)

}

func (db *InDB) DeleteUser(c *gin.Context) {
	var (
		user   table.User
		result gin.H
	)

	id := c.Param("id")

	err := db.DB.First(&user, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	err = db.DB.Delete(&user).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "delete successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
