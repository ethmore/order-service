package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"order/dotEnv"
	"time"
)

type Token struct {
	Token string
}

type Item struct {
	Id         string
	Title      string
	TotalPrice string
}

type CartInfo struct {
	Id             string
	Items          []Item
	TotalCartPrice string
}

type CartResp struct {
	Message  string
	CartInfo CartInfo
}

type Product struct {
	Id  string
	Qty string
}

type CartProductsResp struct {
	Message  string
	Products []Product
}

func GetCartInfo(token string) (*CartInfo, error) {
	cartReq := Token{
		Token: token,
	}
	body, _ := json.Marshal(cartReq)

	bodyReader := bytes.NewReader(body)
	requestUrl := dotEnv.GoDotEnvVariable("GETCARTINFO")

	req, err := http.NewRequest(http.MethodPost, requestUrl, bodyReader)
	if err != nil {
		fmt.Println("client: could not create request", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("client: error making http request: ", err)
		return nil, err
	}

	//resp
	b, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
		return nil, readErr
	}

	defer res.Body.Close()

	var resp CartResp
	if err := json.Unmarshal([]byte(b), &resp); err != nil {
		fmt.Println("unmarshal err", err)
		return nil, err
	}

	return &resp.CartInfo, nil
}

func GetCartProducts(token string) ([]Product, error) {
	cartReq := Token{
		Token: token,
	}
	body, _ := json.Marshal(cartReq)

	bodyReader := bytes.NewReader(body)
	requestUrl := dotEnv.GoDotEnvVariable("GETCARTPRODUCTS")

	req, err := http.NewRequest(http.MethodPost, requestUrl, bodyReader)
	if err != nil {
		fmt.Println("client: could not create request", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("client: error making http request: ", err)
		return nil, err
	}

	//resp
	b, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
		return nil, readErr
	}

	defer res.Body.Close()

	var resp CartProductsResp
	if err := json.Unmarshal([]byte(b), &resp); err != nil {
		fmt.Println("unmarshal err", err)
		return nil, err
	}

	return resp.Products, nil
}
