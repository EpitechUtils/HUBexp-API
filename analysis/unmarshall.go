package analysis

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// parent structure
type ModuleResp struct {
	Year       string     `json:"scolaryear"`
	Activities []Activity `json:"activites"`
}

// sub activity structure
type Activity struct {
	Code   string  `json:"codeacti"`
	Type   string  `json:"type_title"`
	Title  string  `json:"title"`
	Events []Event `json:"events"`
}

// event of activity structure
type Event struct {
	Code       string      `json:"code"`
	Register   string      `json:"already_register"`
	Assistants []Assistant `json:"assistants"`
}

type Assistant struct {
	Login  string `json:"login"`
	Status string `json:"manager_status"`
}

func DoUnmarshall(response *http.Response) ModuleResp {

	// read http get response
	content, _ := ioutil.ReadAll(response.Body)

	// create module (parent json object) and unmarshall the content in it
	var module ModuleResp
	_ = json.Unmarshal(content, &module)

	return module
}
