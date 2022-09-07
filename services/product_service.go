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

type GetSellers struct {
	Token      string
	ProductIDs []string
}

type ProductAndSelller struct {
	ProductID string
	Seller    string
}

type ProductsAndSelllers struct {
	Message            string
	ProductsAndSellers []ProductAndSelller
}

func GetProductsSellers(token string, productIDs []string) ([]ProductAndSelller, error) {
	cartReq := GetSellers{
		Token:      token,
		ProductIDs: productIDs,
	}
	body, _ := json.Marshal(cartReq)

	bodyReader := bytes.NewReader(body)
	requestUrl := dotEnv.GoDotEnvVariable("GETPRODUCTSSELLERS")

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

	var resp ProductsAndSelllers
	if err := json.Unmarshal([]byte(b), &resp); err != nil {
		fmt.Println("p-unmarshal err", err)
		return nil, err
	}

	return resp.ProductsAndSellers, nil
}
