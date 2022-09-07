package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Message struct {
	Message string
	OrderID string
}

type Product_ struct {
	Title      string
	Qty        string
	Price      string
	SellerName string
}

type Order struct {
	Token              string
	Products           []Product_
	TotalPrice         string
	ShipmentAddressID  string
	CardLastFourDigits string
	PaymentStatus      string
	OrderStatus        string
	OrderTime          string
}

func InsertOrder(o Order) (string, error) {
	body, _ := json.Marshal(o)

	bodyReader := bytes.NewReader(body)
	requestUrl := "http://127.0.0.1:3002/insertOrder"

	req, err := http.NewRequest(http.MethodPost, requestUrl, bodyReader)
	if err != nil {
		fmt.Println("client: could not create request", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("client: error making http request: ", err)
		return "", err
	}

	//resp
	b, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
		return "", readErr
	}

	defer res.Body.Close()

	var resp Message
	if err := json.Unmarshal([]byte(b), &resp); err != nil {
		fmt.Println("p-unmarshal err", err)
		return "", err
	}

	return resp.OrderID, nil
}
