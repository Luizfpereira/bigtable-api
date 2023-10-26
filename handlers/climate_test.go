package handlers_test

import (
	"bigtable_api/database"
	"bigtable_api/entity"
	"bigtable_api/handlers"
	"bigtable_api/repository"
	"bigtable_api/router"
	"bigtable_api/usecase"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type ClimateHandlersSuite struct {
	suite.Suite
	router *gin.Engine
}

func TestClimateHandlersSuite(t *testing.T) {
	suite.Run(t, new(ClimateHandlersSuite))
}

func (c *ClimateHandlersSuite) SetupTest() {
	ctx := context.Background()
	clientInstance, err := database.GetClientSingleton(ctx)
	if err != nil {
		c.Suite.T().Fatal(err)
	}
	repo := repository.NewClimateRepository(clientInstance)
	usecase := usecase.NewClimateUsecase(repo)
	climateHandler := handlers.NewClimateHandler(usecase)
	router := router.InitializeRouter(climateHandler)
	c.router = router
}

type output struct {
	Result []entity.BigtableOutput `json:"result"`
	Count  int                     `json:"count"`
	Status string                  `json:"status"`
}

func (c *ClimateHandlersSuite) TestReadPrefix() {

	c.T().Run("testing prefix without type", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/read/climate-data", nil)
		c.Nil(err)
		w := httptest.NewRecorder()
		c.router.ServeHTTP(w, req)
		c.Equal(http.StatusBadRequest, w.Code)
	})

	c.T().Run("testing prefix weather and area A327734", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/read/climate-data?type=w&area_id=A327734", nil)
		c.Nil(err)
		w := httptest.NewRecorder()
		c.router.ServeHTTP(w, req)
		resultBytes, err := io.ReadAll(w.Body)
		c.Nil(err)
		var out output
		err = json.Unmarshal(resultBytes, &out)
		c.Nil(err)

		c.NotEmpty(out.Result)

		for _, value := range out.Result[:len(out.Result)/4] {
			subKeys := strings.Split(value.Key, "/")
			c.Equal("A327734", subKeys[1])
		}

		c.Equal(http.StatusOK, w.Code)
	})

	c.T().Run("testing prefix weather and date", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/read/climate-data?type=w&date=2023-10-10", nil)
		c.Nil(err)
		w := httptest.NewRecorder()
		c.router.ServeHTTP(w, req)
		resultBytes, err := io.ReadAll(w.Body)
		c.Nil(err)
		var out output
		err = json.Unmarshal(resultBytes, &out)
		c.Nil(err)
		c.Empty(out.Result)
		c.Equal(http.StatusBadRequest, w.Code)
	})

	c.T().Run("testing prefix weather and area A327734 with incomplete date", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/read/climate-data?type=w&area_id=A327734&date=2023-10-10", nil)
		c.Nil(err)
		w := httptest.NewRecorder()
		c.router.ServeHTTP(w, req)
		resultBytes, err := io.ReadAll(w.Body)
		c.Nil(err)
		var out output
		err = json.Unmarshal(resultBytes, &out)
		c.Nil(err)

		c.NotEmpty(out.Result)

		mapDate := make(map[string]struct{})

		for _, value := range out.Result {
			subKeys := strings.Split(value.Key, "/")
			mapDate[subKeys[2]] = struct{}{}
		}

		c.Greater(len(mapDate), 1)
		c.Equal(http.StatusOK, w.Code)
	})

	c.T().Run("testing prefix weather and area A327734 with date and regexp", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/read/climate-data?type=w&area_id=A327734&date=2023-10-10&regexp=.*00:00", nil)
		c.Nil(err)
		w := httptest.NewRecorder()
		c.router.ServeHTTP(w, req)
		resultBytes, err := io.ReadAll(w.Body)
		c.Nil(err)
		var out output
		err = json.Unmarshal(resultBytes, &out)
		c.Nil(err)

		c.NotEmpty(out.Result)

		hour := make(map[string]struct{})

		for _, v := range out.Result {
			dateString := strings.Split(v.Key, "/")[2]
			time := strings.Split(dateString, " ")[1]
			hms := strings.Split(time, ":")

			c.Equal("00", hms[1])
			c.Equal("00", hms[2])

			hour[hms[0]] = struct{}{}
		}
		c.Greater(len(hour), 1)

		c.Equal(http.StatusOK, w.Code)
	})

	c.T().Run("testing prefix weather and area A327734 with date and 2 versions", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/read/climate-data?type=w&area_id=A327734&date=2023-10-10 00:00:00&version=2", nil)
		c.Nil(err)
		w := httptest.NewRecorder()
		c.router.ServeHTTP(w, req)
		resultBytes, err := io.ReadAll(w.Body)
		c.Nil(err)
		var out output
		err = json.Unmarshal(resultBytes, &out)
		c.Nil(err)

		c.NotEmpty(out.Result)
		count := 0

		for _, v := range out.Result {
			if v.Key == "w/A327734/2023-10-10 00:00:00" {
				count++
			}
		}
		c.Equal(count, 2)
		c.Equal(http.StatusOK, w.Code)
	})

	c.T().Run("testing prefix weather and area A327734 with complete date and count", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/read/climate-data?type=w&area_id=A327734&date=2023-10-10&count=true", nil)
		c.Nil(err)
		w := httptest.NewRecorder()
		c.router.ServeHTTP(w, req)
		resultBytes, err := io.ReadAll(w.Body)
		c.Nil(err)
		var out output
		err = json.Unmarshal(resultBytes, &out)
		c.Nil(err)

		c.NotEmpty(out.Result)
		c.NotEmpty(out.Count)
		c.Equal(len(out.Result), out.Count)
		c.Equal(http.StatusOK, w.Code)
	})

}

// chamar duas areas e um range de datas
//chamar chave completa com count true, version 2

func (c *ClimateHandlersSuite) TestReadTwoAreasCompleteDate() {
	req, err := http.NewRequest(http.MethodGet, "/read/climate-data?type=w&area_id=A327734,A327735&date=2023-10-10 00:00:00", nil)
	c.Nil(err)
	w := httptest.NewRecorder()
	c.router.ServeHTTP(w, req)
	resultBytes, err := io.ReadAll(w.Body)
	c.Nil(err)
	var out output
	err = json.Unmarshal(resultBytes, &out)
	c.Nil(err)

	c.NotEmpty(out.Result)
	c.Equal(http.StatusOK, w.Code)

	mapKey := make(map[string]struct{})

	for _, v := range out.Result {
		mapKey[v.Key] = struct{}{}
	}
	c.Equal(2, len(mapKey))
}

func (c *ClimateHandlersSuite) TestReadAreaRangeDates() {
	req, err := http.NewRequest(http.MethodGet, "/read/climate-data?type=w&area_id=A327734&date=2023-10-10 00:00:00,2023-10-11 00:00:00", nil)
	c.Nil(err)
	w := httptest.NewRecorder()
	c.router.ServeHTTP(w, req)
	resultBytes, err := io.ReadAll(w.Body)
	c.Nil(err)
	var out output
	err = json.Unmarshal(resultBytes, &out)
	c.Nil(err)

	c.NotEmpty(out.Result)
	c.Equal(http.StatusOK, w.Code)

	mapKey := make(map[string]struct{})

	for _, v := range out.Result {
		mapKey[v.Key] = struct{}{}
	}
	c.GreaterOrEqual(len(mapKey), 24)
}

func (c *ClimateHandlersSuite) TestReadAreaRangeDatesWithRegexp() {
	req, err := http.NewRequest(http.MethodGet, "/read/climate-data?type=w&area_id=A327734&date=2023-10-10 00:00:00,2023-10-11 00:00:00&regexp=.*00:00", nil)
	c.Nil(err)
	w := httptest.NewRecorder()
	c.router.ServeHTTP(w, req)
	resultBytes, err := io.ReadAll(w.Body)
	c.Nil(err)
	var out output
	err = json.Unmarshal(resultBytes, &out)
	c.Nil(err)

	c.NotEmpty(out.Result)
	c.Equal(http.StatusOK, w.Code)

	mapKey := make(map[string]struct{})

	for _, v := range out.Result {
		mapKey[v.Key] = struct{}{}
		dateString := strings.Split(v.Key, "/")[2]
		time := strings.Split(dateString, " ")[1]
		hms := strings.Split(time, ":")

		c.Equal("00", hms[1])
		c.Equal("00", hms[2])
	}
	c.GreaterOrEqual(len(mapKey), 23)
}

func (c *ClimateHandlersSuite) TestReadTwoAreasRangeDatesRegexp() {
	req, err := http.NewRequest(http.MethodGet, "/read/climate-data?type=w&area_id=A327734,A327735&date=2023-10-10 00:00:00,2023-10-11 00:00:00&regexp=.*00:00", nil)
	c.Nil(err)
	w := httptest.NewRecorder()
	c.router.ServeHTTP(w, req)
	resultBytes, err := io.ReadAll(w.Body)
	c.Nil(err)
	var out output
	err = json.Unmarshal(resultBytes, &out)
	c.Nil(err)

	c.NotEmpty(out.Result)
	c.Equal(http.StatusOK, w.Code)

	mapAreasRange := make(map[string][]struct{})

	for _, v := range out.Result {
		area := strings.Split(v.Key, "/")[1]
		mapAreasRange[area] = append(mapAreasRange[area], struct{}{})
		dateString := strings.Split(v.Key, "/")[2]
		time := strings.Split(dateString, " ")[1]
		hms := strings.Split(time, ":")

		c.Equal("00", hms[1])
		c.Equal("00", hms[2])
	}
	c.Equal(2, len(mapAreasRange))
	for _, v := range mapAreasRange {
		c.GreaterOrEqual(len(v), 23)
	}
}

func (c *ClimateHandlersSuite) TestReadTwoAreasIncompleteDates() {
	req, err := http.NewRequest(http.MethodGet, "/read/climate-data?type=w&area_id=A327734,A327735&date=2023-10-10", nil)
	c.Nil(err)
	w := httptest.NewRecorder()
	c.router.ServeHTTP(w, req)
	resultBytes, err := io.ReadAll(w.Body)
	c.Nil(err)
	var out output
	err = json.Unmarshal(resultBytes, &out)
	c.Nil(err)

	c.Empty(out.Result)
	c.Equal(http.StatusBadRequest, w.Code)
	c.Equal("failed", out.Status)
}

func (c *ClimateHandlersSuite) TestReadTwoAreasRangeIncompleteDates() {
	req, err := http.NewRequest(http.MethodGet, "/read/climate-data?type=w&area_id=A327734,A327735&date=2023-10-10,2023-10-20", nil)
	c.Nil(err)
	w := httptest.NewRecorder()
	c.router.ServeHTTP(w, req)
	resultBytes, err := io.ReadAll(w.Body)
	c.Nil(err)
	var out output
	err = json.Unmarshal(resultBytes, &out)
	c.Nil(err)

	c.Empty(out.Result)
	c.Equal(http.StatusBadRequest, w.Code)
	c.Equal("failed", out.Status)
}

//test com duas areas e uma data incompleta
//test com duas areas e duas datas incompleta
