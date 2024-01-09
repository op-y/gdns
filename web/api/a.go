package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/op-y/gdns/cache"
	"github.com/op-y/gdns/storage"
)

func QueryA(c *gin.Context) {
	domain := c.Param("domain")
	if domain == "" {
		c.JSON(http.StatusBadRequest, "domain is required")
		return
	}

	ds := storage.Default()
	records, err := ds.QueryA(domain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "query failed")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"domain":  domain,
		"type":    "A",
		"records": records,
	})
	return
}

type RequestCreateA struct {
	Records []string `json:"records"`
}

func CreateA(c *gin.Context) {
	domain := c.Param("domain")
	if domain == "" {
		c.JSON(http.StatusBadRequest, "domain is required")
		return
	}

	req := &RequestCreateA{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if domain == "" || req.Records == nil || len(req.Records) == 0 {
		c.JSON(http.StatusBadRequest, "bad parameter")
		return
	}

	dc := cache.Default()
	ds := storage.Default()

	if err := ds.CreateA(domain, req.Records); err != nil {
		c.JSON(http.StatusInternalServerError, "create failed")
		return
	}
	if err := dc.DeleteA(domain); err != nil {
		log.Printf("delete cache failed: %s\n", err.Error())
	}
	c.JSON(200, gin.H{"message": "ok"})
}

type RequestUpdateA struct {
	Records []string `json:"records"`
}

func UpdateA(c *gin.Context) {
	domain := c.Param("domain")
	if domain == "" {
		c.JSON(http.StatusBadRequest, "domain is required")
		return
	}

	req := &RequestUpdateA{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if domain == "" || req.Records == nil || len(req.Records) == 0 {
		c.JSON(http.StatusBadRequest, "bad parameter")
		return
	}

	dc := cache.Default()
	ds := storage.Default()

	if err := ds.UpdateA(domain, req.Records); err != nil {
		c.JSON(http.StatusInternalServerError, "update failed")
	}
	if err := dc.DeleteA(domain); err != nil {
		log.Printf("delete cache failed: %s\n", err.Error())
	}
	c.JSON(200, gin.H{"message": "ok"})
	return
}

func DeleteA(c *gin.Context) {
	domain := c.Param("domain")
	if domain == "" {
		c.JSON(http.StatusBadRequest, "domain is required")
		return
	}

	dc := cache.Default()
	ds := storage.Default()

	if err := ds.DeleteA(domain); err != nil {
		c.JSON(http.StatusInternalServerError, "delete failed")
		return
	}
	if err := dc.DeleteA(domain); err != nil {
		log.Printf("delete cache failed: %s\n", err.Error())
	}
	c.JSON(200, gin.H{"message": "ok"})
	return
}
