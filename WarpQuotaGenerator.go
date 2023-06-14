package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// Default Variables
var ClientID string
var SuccessCnt, FailedCnt int = 0, 0

type WarpRequestParameter struct {
	Key         string `json:"key"`
	InstallID   string `json:"install_id"`
	FcmToken    string `json:"fcm_token"`
	Referrer    string `json:"referrer"`
	WarpEnabled bool   `json:"warp_enabled"`
	Tos         string `json:"tos"`
	Type        string `json:"type"`
	Locale      string `json:"locale"`
}

func main() {
	fmt.Println("Warp Plus Quota Generator")
	fmt.Println("By: @ItsGoYoung")
	fmt.Print("Type Your Warp ID: ")
	fmt.Scan(&ClientID)

	url := fmt.Sprintf("https://api.cloudflareclient.com/v0a%s/reg", digitString(3))

	for {
		if SuccessCnt != 500 {
			installID := genString(22)
			body := WarpRequestParameter{
				Key:         fmt.Sprintf("%s=", genString(43)),
				InstallID:   installID,
				FcmToken:    fmt.Sprintf("%s:APA91b%s", installID, genString(134)),
				Referrer:    ClientID,
				WarpEnabled: false,
				Tos:         fmt.Sprintf("%s+02:00", time.Now().UTC().Format("2006-01-02T15:04:05.000")),
				Type:        "Android",
				Locale:      "es_ES",
			}

			jsonBody, err := json.Marshal(body)
			if err != nil {
				panic(err)
			}

			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
			if err != nil {
				panic(err)
			}
			req.Header.Set("Content-Type", "application/json; charset=UTF-8")
			req.Header.Set("Host", "api.cloudflareclient.com")
			req.Header.Set("Connection", "Keep-Alive")
			req.Header.Set("Accept-Encoding", "gzip")
			req.Header.Set("User-Agent", "okhttp/3.12.1")

			// do request and return response
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				panic(err)
			}

			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return
			}

			fmt.Println(string(respBody))

			if resp.StatusCode == 200 {
				SuccessCnt++
				fmt.Printf("Passed: +1GB (total: %dGB, failed: %d)\n", SuccessCnt, FailedCnt)
			} else {
				FailedCnt++
				fmt.Printf("Failed: %d\n", resp.StatusCode)
			}

			// random cooldown time between 15 to 25 seconds
			startCoolDown := rand.NewSource(15)
			endCoolDown := rand.New(startCoolDown)
			cooldown := endCoolDown.Intn(25)
			fmt.Printf("Cooldown: %d seconds\n", cooldown)
			time.Sleep(time.Duration(cooldown) * time.Second)
		} else {
			fmt.Printf("Ended : (total: %dGB, failed: %d)\n", SuccessCnt, FailedCnt)
			break
		}
	}
}

func getASCIILetters() string {
	return "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
}

func getDigits() string {
	return "0123456789"
}

func genString(stringLength int) string {
	var letters = getASCIILetters() + getDigits()
	var result string
	for i := 0; i < stringLength; i++ {
		result += string(letters[rand.Intn(len(letters))])
	}
	return result
}

func digitString(stringLength int) string {
	var digit = getDigits()
	var result string
	for i := 0; i < stringLength; i++ {
		result += string(digit[rand.Intn(len(digit))])
	}
	return result
}
