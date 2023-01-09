# Crypto Server
It makes use of Free API from [https://api.hitbtc.com/](https://api.hitbtc.com/)
* A micro-service with the following endpoint
* `GET /currency/{symbol}`
    * Returns the real-time crypto prices of the given currency symbol
    * Sample Response:
    ```json
    {
        "id": "ETH",
        "fullName": "Ethereum",
        "ask": "0.054464",
        "bid": "0.054463",
        "last": "0.054463",
        "open": "0.057133",
        "low": "0.053615",
        "high": "0.057559",
        "feeCurrency": "BTC"
    }
    ```
* `GET /currency/all`
    * Returns the real-time crypto prices of all the supported currencies
    * Sample Response
    ```json
    {
        "currencies": [
            {
                "id": "ETH",
                "fullName": "Ethereum",
                "ask": "0.054464",
                "bid": "0.054463",
                "last": "0.054463",
                "open": "0.057133",
                "low": "0.053615",
                "high": "0.057559",
                "feeCurrency": "BTC"
            },
            {
                "id": "BTC",
                "fullName": "Bitcoin",
                "ask":"7906.72",
                "bid":"7906.28",
                "last":"7906.48",
                "open":"7952.3",
                "low":"7561.51",
                "high":"8107.96",
                "feeCurrency": "USD"
            }
        ]
    }
    ```

---
## Steps to run
```bash
$ go install
# it will install the cryptoserver named binary
$ cryptoserver -h
Usage of cryptoserver:
  -c string
    	config file path
  -debug
    	log debug messages
  -port int
    	server port to run this server (default 8000)
  -s value
    	symbols to subscribe for ticker, default: BTCUSDT ETHBTC
  -v	log on console
  -vv int
    	verbose logs level
```