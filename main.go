package main

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"io/ioutil"
	"net/http"
)

const autologUrl = "https://intra.epitech.eu/auth-b4076976be4815f632794fd00a5a6c69d1655939"

type Activity struct {
	Title     string
	TypeTitle string
}

func doRequest() []iris.Map {

	res, _ := http.Get(autologUrl + "/module/2019/B-MOO-500/NCE-5-1/?format=json")

	body := make(iris.Map)
	content, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(content, &body)

	activities, ok := body["activites"]

	if !ok {
		/*
			for k, v := range body {
				if k == "activites" {
					activities = v
				}
			}
		*/
		panic("Undefined 'activites' key")
	}

	switch x := activities.(type) {
	case []interface{}:
		fmt.Printf("got %T\n", x)
		for _, e := range x {
			events := e.(map[string]interface{})["events"]
			eventsArray := events.([]interface{})

			fmt.Print("Check ", e.(map[string]interface{})["title"], " : ")

			for _, value := range eventsArray {
				if value.(map[string]interface{})["user_status"] == "present" {
					fmt.Print("->Present")
				}
			}
			fmt.Print("\n")
		}
	default:
		fmt.Printf("I don't know how to handle %T\n", activities)
	}

	return activities.([]iris.Map)
}

func processActivities(data iris.Map) (r iris.Map) {

	return
}

func main() {
	router := iris.New()
	router.Use(logger.New())
	router.Use(recover.New())
	router.Get("/", func(c iris.Context) {
		c.JSON(iris.Map{"hello": "sailor"})
	})
	/*
		router.Get("/exp", func(c iris.Context) {
			c.JSON(doRequest())
		})
	*/
	doRequest()
	//	router.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
