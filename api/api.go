package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	// doc.json
	_ "github.com/kautsarady/forex/docs"
	"github.com/kautsarady/forex/httputil"
	"github.com/kautsarady/forex/model"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Foreign Currency Exchange API
// @version 1.0
// @description This foreign currency exchange api documentation
// @contact.name kautsarady
// @contact.email kautsarady@gmail.com

// Controller api
type Controller struct {
	DAO    *model.DAO
	Router *gin.Engine
}

// Make api controller
func Make(dao *model.DAO) *Controller {
	ctr := &Controller{DAO: dao, Router: gin.Default()}
	api := ctr.Router.Group("/api")
	{
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		api.POST("/input/:from/:to", ctr.HandleInput())
		api.GET("/track", ctr.HandleTrack())
		api.GET("/trend/:from/:to", ctr.HandleTrend())
		api.POST("/add/:from/:to", ctr.HandleAdd())
		api.DELETE("/remove/:from/:to", ctr.HandleRemove())
	}
	return ctr
}

// HandleInput godoc
// @Summary Input foreign currency exchange
// @ID input-forex
// @Accept json
// @Produce json
// @Param from path string true "FROM currency ex: USD"
// @Param to path string true "TO currency ex: JPY"
// @Param rate query string true "forex rate ex: 1.11" Format(number)
// @Param date query string true "forex date ex: 2018-11-13" Format(date)
// @Success 200 {string} string	"input exchange data $from - $to success"
// @Failure 400 {object} httputil.HTTPError
// @Router /api/input/{from}/{to} [post]
func (ctr *Controller) HandleInput() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		from := strings.ToUpper(ctx.Param("from"))
		to := strings.ToUpper(ctx.Param("to"))

		strRate, ok := ctx.GetQuery("rate")
		if !ok {
			httputil.NewError(ctx, http.StatusBadRequest, errors.New("Should provide form key-value \"rate\""))
			return
		}

		strDate, ok := ctx.GetQuery("date")
		if !ok {
			httputil.NewError(ctx, http.StatusBadRequest, errors.New("Should provide form key-value \"date\""))
			return
		}

		// Parsing "rate" (string) to float
		rate, err := strconv.ParseFloat(strRate, 64)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, errors.New("Found \"rate\" is not a valid number: "+err.Error()))
			return
		}

		// Parsing "date" (string) to time
		date, err := time.Parse("2006-01-02", strDate)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, errors.New("Found date is not a valid time/date format: "+err.Error()))
			return
		}

		// Input for currency A to B with rate = R
		if err := ctr.DAO.Input(
			model.Exchange{
				From: from,
				To:   to,
				Rate: rate,
				Date: date,
			}); err != nil {
			httputil.NewError(ctx, http.StatusInternalServerError, errors.New("internal server error"))
			ctx.Error(err)
			return
		}

		// Input for currency B to A with rate = 1/R
		if err := ctr.DAO.Input(
			model.Exchange{
				From: to,
				To:   from,
				Rate: 1 / rate,
				Date: date,
			}); err != nil {
			httputil.NewError(ctx, http.StatusInternalServerError, errors.New("internal server error"))
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("input exchange data %s - %s success", from, to)})
	}
}

// HandleTrack godoc
// @Summary Track foreign currency exchange
// @ID track-forex
// @Accept json
// @Produce json
// @Param date query string true "forex date ex: 2018-11-13" Format(date)
// @Success 200 {array} model.Exchange
// @Failure 400 {object} httputil.HTTPError
// @Router /api/track [get]
func (ctr *Controller) HandleTrack() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		strDate, ok := ctx.GetQuery("date")
		if !ok {
			httputil.NewError(ctx, http.StatusBadRequest, errors.New("Should provide query \"date\" (format: yyyy-mm-dd) to be tracked"))
			return
		}

		// Parsing "date" (string) to time
		date, err := time.Parse("2006-01-02", strDate)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, errors.New("Found date is not a valid time/date format: "+err.Error()))
			return
		}

		exs, err := ctr.DAO.Track(date)
		if err != nil {
			httputil.NewError(ctx, http.StatusInternalServerError, errors.New("internal server error"))
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, exs)
	}
}

// HandleTrend godoc
// @Summary Show trend of foreign currency exchange
// @ID trend-forex
// @Accept json
// @Produce json
// @Param from path string true "FROM currency ex: USD"
// @Param to path string true "TO currency ex: JPY"
// @Param avg query string true "average rate ex: 1.133" Format(number)
// @Param vrn query string true "variance rate ex: 0.333" Format(number)
// @Success 200 {array} model.Exchange
// @Failure 400 {object} httputil.HTTPError
// @Router /api/trend/{from}/{to} [get]
func (ctr *Controller) HandleTrend() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		from := strings.ToUpper(ctx.Param("from"))
		to := strings.ToUpper(ctx.Param("to"))

		strAvg, ok := ctx.GetQuery("avg")
		if !ok {
			httputil.NewError(ctx, http.StatusBadRequest, errors.New("Should provide query \"avg\""))
			return
		}

		strVrn, ok := ctx.GetQuery("vrn")
		if !ok {
			httputil.NewError(ctx, http.StatusBadRequest, errors.New("Should provide query \"vrn\""))
			return
		}

		// Parsing "avg" (string) to float
		avg, err := strconv.ParseFloat(strAvg, 64)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, errors.New("Found \"avg\" is not a valid number: "+err.Error()))
			return
		}

		// Parsing "vrn" (string) to float
		vrn, err := strconv.ParseFloat(strVrn, 64)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, errors.New("Found \"vrn\" is not a valid number: "+err.Error()))
			return
		}

		exs, err := ctr.DAO.Trend(from, to, avg-vrn)
		if err != nil {
			httputil.NewError(ctx, http.StatusInternalServerError, errors.New("internal server error"))
			ctx.Error(err)
			return
		}

		if exs == nil {
			httputil.NewError(ctx, http.StatusSeeOther, errors.New("there is no correspond data found"))
			return
		}

		ctx.JSON(http.StatusOK, exs)
	}
}

// HandleAdd godoc
// @Summary Add foreign currency exchange
// @ID add-forex
// @Accept json
// @Produce json
// @Param from path string true "FROM currency ex: USD"
// @Param to path string true "TO currency ex: JPY"
// @Success 200 {string} string	"adding exchange data $from - $to success"
// @Failure 400 {object} httputil.HTTPError
// @Router /api/add/{from}/{to} [post]
func (ctr *Controller) HandleAdd() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		from := strings.ToUpper(ctx.Param("from"))
		to := strings.ToUpper(ctx.Param("to"))

		// Add forex from A to B
		if err := ctr.DAO.Add(
			model.Exchange{
				From: from,
				To:   to,
			}); err != nil {
			httputil.NewError(ctx, http.StatusInternalServerError, errors.New("internal server error"))
			ctx.Error(err)
			return
		}

		// Add forex from B to A
		if err := ctr.DAO.Add(
			model.Exchange{
				From: to,
				To:   from,
			}); err != nil {
			httputil.NewError(ctx, http.StatusInternalServerError, errors.New("internal server error"))
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("adding exchange data %s - %s success", from, to)})
	}
}

// HandleRemove godoc
// @Summary Remove foreign currency exchange
// @ID remove-forex
// @Accept json
// @Produce json
// @Param from path string true "FROM currency ex: USD"
// @Param to path string true "TO currency ex: JPY"
// @Success 200 {string} string	"removing exchange data $from - $to success"
// @Failure 400 {object} httputil.HTTPError
// @Router /api/remove/{from}/{to} [delete]
func (ctr *Controller) HandleRemove() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		from := strings.ToUpper(ctx.Param("from"))
		to := strings.ToUpper(ctx.Param("to"))

		// Remove forex from A to B
		if err := ctr.DAO.Remove(
			model.Exchange{
				From: from,
				To:   to,
			}); err != nil {
			httputil.NewError(ctx, http.StatusInternalServerError, errors.New("internal server error"))
			ctx.Error(err)
			return
		}

		// Remove forex from B to A
		if err := ctr.DAO.Remove(
			model.Exchange{
				From: to,
				To:   from,
			}); err != nil {
			httputil.NewError(ctx, http.StatusInternalServerError, errors.New("internal server error"))
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("removing exchange data %s - %s success", from, to)})
	}
}
