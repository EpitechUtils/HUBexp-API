package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/lucasGras/HUBexp-API.git/analysis"
	"net/http"
	"strings"
	"fmt"
)

const autologUrl = "https://intra.epitech.eu/auth-b4076976be4815f632794fd00a5a6c69d1655939"
const currentLogin = "lucas.gras@epitech.eu"

/**
Check if a student is an assistant
*/
func isAnAssistant(login string, comp []analysis.Assistant) bool {
	for _, v := range comp {
		if strings.Compare(login, v.Login) == 0 {
			return true
		}
	}
	return false
}

/**
create API response
matching is working with:
	- Talk / Meetup
	- Workshop
	- Hackaton

TODO: Experience and HUB Project (variable) + absent or canceled (minus)
*/
func createApiResponse(module analysis.ModuleResp) iris.Map {
	// total exp of student
	var exp int

	// matching array for activity titles
	var matching = map[string]int{
		"Conference_par": 1,
		"Conference_org": 4,
		"Workshop_par":   3,
		"Workshop_org":   10,
		"Rush_par":       6,
	}

	// Process data and compute exp sum
	for _, act := range module.Activities {
		for _, eve := range act.Events {

			// If student wasn't register and wasn't the organizer, skip it
			if len(eve.Register) == 0 && !isAnAssistant(currentLogin, eve.Assistants) {
				continue
			}

			fmt.Print("Type [" + act.Type + "] : " + act.Title + "\n")

			// If the student have organized the session
			if isAnAssistant(currentLogin, eve.Assistants) {
				exp += matching[act.Type+"_org"]
				//fmt.Print(" : ORG\n")
			} else { // Or the student is just a participant
				exp += matching[act.Type+"_par"]
				//fmt.Print(" : PART\n")
			}
		}
	}

	//fmt.Print("Total exp: ", exp, "\n")

	return iris.Map{"exp": exp}
}

func doRequest() iris.Map {

	// Make http request to intranet
	res, _ := http.Get(autologUrl + "/module/2019/B-INN-000/NCE-0-1/?format=json")

	// Process JSON data
	module := analysis.DoUnmarshall(res)

	// Create API response based on JSON data
	response := createApiResponse(module)

	return response
}

func main() {
	router := iris.New()
	router.Use(logger.New())
	router.Use(recover.New())
	router.Get("/", func(c iris.Context) {
		c.JSON(iris.Map{"hello": "sailor"})
	})
	router.Get("/exp", func(c iris.Context) {
		c.JSON(doRequest())
	})
	router.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
