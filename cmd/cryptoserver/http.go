package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getCurrencyHandler(state state) func(c *gin.Context) {

	// /currency/:symbol
	return func(c *gin.Context) {

		symbol := c.Param("symbol")

		stateRes := make(chan stateRes)
		symbolToAsk := ""
		if symbol != "all" {
			symbolToAsk = symbol
		}

		state.reqChan <- stateReq{
			symbol:  symbolToAsk,
			resChan: stateRes,
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
