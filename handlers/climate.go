package handlers

import (
	"bigtable_api/entity"
	"bigtable_api/usecase"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ClimateHandler struct {
	usecase *usecase.ClimateUsecase
}

func NewClimateHandler(climateUsecase *usecase.ClimateUsecase) *ClimateHandler {
	return &ClimateHandler{usecase: climateUsecase}
}

func (h *ClimateHandler) ReadClimateData(ctx *gin.Context) {
	start := time.Now()
	dataType := ctx.Query("type")
	areaID := ctx.Query("area_id")
	date := ctx.Query("date")
	version := ctx.Query("version")
	regexp := ctx.Query("regexp")
	count := ctx.Query("count")

	log.Println(dataType, areaID)

	var prefixes []string
	if dataType == "" {
		log.Printf("error reading prefix. No datatype provided")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "failed", "error": "no datatype provided"})
		return
	}
	prefixes = append(prefixes, dataType)

	var areas, dates []string

	if areaID != "" {
		areas = strings.Split(areaID, ",")
	}

	if date != "" {
		dates = strings.Split(date, ",")
	}

	if len(areas) == 0 && len(dates) > 0 {
		log.Printf("error reading prefix. Missing area_id")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "failed", "error": "missing area_id"})
		return
	}

	if len(areas) == 1 {
		prefixes = append(prefixes, areas[0])
		if len(dates) == 1 {
			prefixes = append(prefixes, dates[0])
		}
	}

	filters := make(map[string]string)
	if version != "" {
		filters["version"] = version
	}

	if regexp != "" {
		filters["regexp"] = regexp
	}

	var output []entity.BigtableOutput
	var err error

	if len(areas) > 1 || len(dates) > 1 {
		for _, date := range dates {
			layout := "2006-01-02 15:04:05"
			if _, err := time.Parse(layout, date); err != nil {
				log.Printf("error reading data: incomplete date")
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "failed", "error": "incomplete date"})
				return
			}
		}
		output, err = h.usecase.Read(ctx, "climate_data", filters, areas, dates)
		if err != nil {
			log.Printf("error reading areas: %s and dates: %s. Error: %v", areas, dates, err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err})
			return
		}
	} else {
		output, err = h.usecase.ReadPrefix(ctx, "climate_data", filters, prefixes...)
		if err != nil {
			log.Printf("error reading prefix %s/%s. Error: %v", dataType, areaID, err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "failed", "error": err})
			return
		}
	}

	result := make(map[string]interface{})
	result["result"] = output

	if count == "true" {
		result["count"] = len(output)
	}
	result["status"] = "success"

	log.Printf("Request successful. Datatype: %s, areas: %s, dates: %s Time taken: %v.", dataType, areaID, date, time.Since(start))
	ctx.JSON(http.StatusOK, result)
}
