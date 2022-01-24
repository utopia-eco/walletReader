package services

import (
	"github.com/utopia-eco/walletReader/utils"
	"io/ioutil"
	"net/http"
	"time"
)

var HttpClient = &http.Client{
	Timeout: time.Second * 10,
}

func SendHttpGetRequest(url string) (*[]byte, error) {
	resp, err := HttpClient.Get(url)
	if err != nil {
		utils.Logger.Error("SendHttpGetRequest err: %v", err)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.Logger.Error("SendHttpGetRequest read body err: %v", err)
		return nil, err
	}
	return &body, err
}
