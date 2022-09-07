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

type GetUserAddressById struct {
	Token     string
	AddressId string
}

type AddressResp struct {
	Message string
	Address Address
}

type Address struct {
	Id              string `bson:"_id"`
	Title           string
	Name            string
	Surname         string
	PhoneNumber     string
	Province        string
	County          string
	DetailedAddress string
}

func GetDetailedAddressByID(token, addressID string) (*Address, error) {
	addressReq := GetUserAddressById{
		Token:     token,
		AddressId: addressID,
	}
	body, _ := json.Marshal(addressReq)

	bodyReader := bytes.NewReader(body)
	requestUrl := dotEnv.GoDotEnvVariable("GETUSERADDRESSBYID")

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

	var resp AddressResp
	if err := json.Unmarshal([]byte(b), &resp); err != nil {
		fmt.Println("unmarshal err", err)
		return nil, err
	}

	return &resp.Address, nil
}
