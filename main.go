package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	//getForcastData()
	go scheduler()
	e := echo.New()
	e.Get("/", index)
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.Secure())

	e.Static("/", "front-end")
	e.Run(standard.New(":8000"))
}

func index(c echo.Context) error {
	return c.File("front-end/index.html")
}

func scheduler() {
	fmt.Println("do it")
	for range time.NewTicker(time.Minute * 5).C {
		// here
		fmt.Println("do it again")
	}
}

func getForcastData() {
	var data = "var data = "
	response, err := http.Get("https://api.forecast.io/forecast//40.7127,-74.0059")
	if err != nil {
		log.Error("%s", err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Error("%s", err)
		}
		data += string(contents) + ";"
		err = ioutil.WriteFile("front-end/data.js", []byte(data), 0755)
		if err != nil {
			log.Error("%s", err)
		}
	}
}
