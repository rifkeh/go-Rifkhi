package controller

import (
	"competence/config"
	"competence/model"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetAllBarang(c echo.Context) error{
	var barang []model.Items
	keyword := c.QueryParam("keyword")
	if keyword == ""{
		if err := config.DB.Find(&barang).Error; err != nil{
			return echo.NewHTTPError(500, echo.Map{
				"message": "Failed get all items",
				"error": err.Error(),
		})
		}
		return c.JSON(200, echo.Map{
			"message": "Success get all items",
			"data": barang,
		})
	}else{
		if err := config.DB.Find(&barang, "name LIKE ?", "%"+keyword+"%").Error; err != nil{
			return echo.NewHTTPError(500, echo.Map{
				"message": "Failed get all barang",
				"error": err.Error(),
		})
		}
		if len(barang) == 0{
			return c.JSON(200, echo.Map{
				"message": "There is no items with that keyword",
		})}

		return c.JSON(200, echo.Map{
			"message": "Success get all items",
			"data": barang,
		})
	}
}

func GetBarang(c echo.Context) error{
	var barang model.Items
	id := c.Param("id")
	if err := config.DB.Find(&barang, "id = ?", id).Error; err != nil{
		return echo.NewHTTPError(500, echo.Map{
			"message": "Failed get item",
			"error": err.Error(),
	})
	}
	return c.JSON(200, echo.Map{
		"message": "Success get item",
		"data": barang,
	})
}

func CreateBarang(c echo.Context) error{
	barang := model.Items{}
	c.Bind(&barang)
	uuidWithHyphen := uuid.New()
    uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	barang.Id = uuid
	if err := config.DB.Save(&barang).Error; err != nil{
		return echo.NewHTTPError(500, echo.Map{
			"message": "Failed create item",
			"error": err.Error(),
	})
	}
	return c.JSON(200, echo.Map{
		"message": "Success create item",
		"data": barang,
	})
}

func UpdateBarang(c echo.Context) error{
	barang := model.Items{}
	c.Bind(&barang)
	id := c.Param("id")
	if err := config.DB.Where("id=?", id).Updates(&barang).Error; err != nil{
		return echo.NewHTTPError(500, echo.Map{
			"message": "Failed update item",
			"error": err.Error(),
	})
	}
	if err := config.DB.Find(&barang, "id = ?", id).Error; err != nil{
		return echo.NewHTTPError(500, echo.Map{
			"message": "Failed get item",
			"error": err.Error(),
	})}	
	return c.JSON(200, echo.Map{
		"message": "Success update item",
		"data": barang,
	})
}

func DeleteBarang(c echo.Context) error{
	barang := model.Items{}
	id := c.Param("id")
	if err := config.DB.Delete(&barang, "id = ?", id).Error; err != nil{
		return echo.NewHTTPError(500, echo.Map{
			"message": "Failed delete item",
			"error": err.Error(),
	})
	}
	return c.JSON(200, echo.Map{
		"message": "Success delete item",
		"data": barang,
	})
}

func GetBarangByKategori(c echo.Context) error{
	var barang []model.Items
	kategori := c.Param("category_id")
	if err := config.DB.Find(&barang, "category = ?", kategori).Error; err != nil{
		return echo.NewHTTPError(500, echo.Map{
			"message": "Failed get item by kategori",
			"error": err.Error(),
	})
	}
	return c.JSON(200, echo.Map{
		"message": "Success get item by kategori",
		"data": barang,
	})
}
