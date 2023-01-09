package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getCurrencyHandler(state state) func(c *gin.Context) {

	// /currency/:symbol
	return func(c *gin.Context) {

		symbol := c.Param("symbol")

		log.Info(fmt.Sprintf("asked real time price for %s", symbol))

		stateRes := make(chan stateRes)
		if symbol == "all" {
			state.reqChan <- stateReq{
				symbol: "",
			}
		} else {
			state.reqChan <- stateReq{
				symbol: symbol,
			}
		}

		res := <-stateRes

		if res.err != nil {
			c.JSON(http.StatusOK,
				gin.H{
					"error": res.err.Error(),
				})
			return
		}

		if symbol == "all" {
			response := []currencyResponse{}
			for name, tick := range res.ticks {
				response = append(response, state.tickToCurrencyResponse(name, tick))
			}

			c.JSON(http.StatusOK,
				gin.H{
					"currencies": response,
				})

			return
		}

		c.JSON(http.StatusOK, state.tickToCurrencyResponse(symbol, res.ticks[symbol]))
	}
}
