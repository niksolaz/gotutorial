package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Info struct {
	Price     string `json:"price"`
	Color     string `json:"color"`
	Car       string `json:"car"`
	Is_active string `json:"isactive"`
}

func main() {
	getProductsHandler := http.HandlerFunc(getProducts)
	http.Handle("/products", getProductsHandler)
	http.ListenAndServe(":8080", nil)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	color := query.Get("color")       //color="red"
	price := query.Get("price")       //price="29"
	car := query.Get("car")           //car="volvo"
	isactive := query.Get("isactive") //isactive="true"
	w.WriteHeader(200)
	objectQuery := Info{Color: color, Price: price, Car: car, Is_active: isactive}

	byteArray, err := json.MarshalIndent(objectQuery, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
	w.Write([]byte(byteArray))

}
