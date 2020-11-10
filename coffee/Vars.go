package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	coffeeShops []CoffeeShop
	ID          int = 0
)

type CoffeeShop struct {
	ID       string `json: "id"`
	ShopName string `json: "shopname"`
	Location string `json: "location"`
	Menu     *Menu  `json: "menu"`
}

type Menu struct {
	Item1 string `json: "item1"`
	Item2 string `json: "item2"`
	Item3 string `json: "item3"`
}

func getList(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(coffeeShops)
}

func addShop(res http.ResponseWriter, req *http.Request) {
	ID = ID + 1
	res.Header().Set("Content-Type", "application/json")
	var shop CoffeeShop
	json.NewDecoder(req.Body).Decode(&shop)
	shop.ID = strconv.Itoa(ID)
	coffeeShops = append(coffeeShops, shop)
}

func updateShop(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var shop CoffeeShop
	json.NewDecoder(req.Body).Decode(&shop)
	params := mux.Vars(req)

	for i, val := range coffeeShops {
		if val.ID == params["id"] {
			shop.ID = params["id"]
			coffeeShops[i] = shop
			return
		}
	}
}

func getShop(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for _, value := range coffeeShops {
		if value.ID == params["id"] {
			json.NewEncoder(res).Encode(value)
			return
		}
	}
}

func deleteShop(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for index, val := range coffeeShops {
		if val.ID == params["id"] {
			coffeeShops = append(coffeeShops[:index], coffeeShops[index+1:]...)
		}
	}
}
