package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

var foodPriceMap map[string]float32 = map[string]float32{
	"RedSet":    50,
	"GreenSet":  40,
	"BlueSet":   30,
	"YellowSet": 50,
	"PinkSet":   80,
	"PurpleSet": 90,
	"OrangeSet": 120,
}

type Order struct {
	RedSet    int  `json:"RedSet"`
	GreenSet  int  `json:"GreenSet"`
	BlueSet   int  `json:"BlueSet"`
	YellowSet int  `json:"YellowSet"`
	PinkSet   int  `json:"PinkSet"`
	PurpleSet int  `json:"PurpleSet"`
	OrangeSet int  `json:"OrangeSet"`
	IsMember  bool `json:"isMember"`
}

func main() {
	jsonFile, err := os.Open("order.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	var bytes []byte
	bytes, err = ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
	}

	var order Order

	json.Unmarshal(bytes, &order)

	var totalAmount float32 = 0
	totalAmount, err = calTotalAmount(order)

	var totalDiscount float32 = 0
	var discount float32 = 0

	discount, err = memberDiscount(order)

	if err != nil {
		fmt.Println(err)
	}

	totalDiscount += discount

	discount, err = doublesDiscount(order)

	if err != nil {
		fmt.Println(err)
	}

	totalDiscount += discount

	var afterDiscount float32
	afterDiscount, err = calDiscount(totalAmount, totalDiscount)

	fmt.Printf("total amount = %.2f\ntotal discount = %.2f%%\namount after discount = %.2f", totalAmount, totalDiscount, afterDiscount)
}

func calTotalAmount(o Order) (float32, error) {
	var foodTypes int = 7
	var sum float32 = 0
	for i := 0; i < foodTypes; i++ {
		var foodName string = reflect.TypeOf(&o).Elem().Field(i).Name
		var numFood int = reflect.ValueOf(o).FieldByIndex([]int{i}).Interface().(int)
		sum += foodPriceMap[foodName] * float32(numFood)
	}
	return sum, nil
}

func calDiscount(totalAmount float32, discount float32) (float32, error) {
	if totalAmount < 0 || discount < 0 {
		return 0, errors.New("total amount or discount can not be negative.")
	}
	if discount > 100 {
		return 0, errors.New("total discount can not exceed 100.")
	}
	return totalAmount - (totalAmount * discount / 100), nil
}

func memberDiscount(o Order) (float32, error) {
	if o.IsMember {
		return 10, nil
	}
	return 0, nil
}

func doublesDiscount(o Order) (float32, error) {
	if o.OrangeSet >= 2 || o.PinkSet >= 2 || o.GreenSet >= 2 {
		return 5, nil
	}
	return 0, nil
}
