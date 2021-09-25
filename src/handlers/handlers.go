package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VicFlores/src/middlewares"
)

func HandleHome(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello from Home")
}

func HandleShop(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello from Shop") // recibe un response - mensaje
}

func PostUser(res http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)

	var mockData middlewares.User

	err := decoder.Decode(&mockData)

	if err != nil {
		fmt.Fprintf(res, "error: %v", err)
		return
	}

	response, err := mockData.ToJSON()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(response)
}
