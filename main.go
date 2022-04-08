package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type UploadSession struct {
	Data Data `json:data`
}
type Data struct {
	SessionId string `json:sessionId`
}

func main() {
	fileName := "test2605.psd"
	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjIxMDYxNjUsInVzZXJJRCI6ImMxNWYzNmQyLTU5NGItNDQwNC1hMTJjLTUxYmQ4NzM0MGQ2ZiIsImVtYWlsIjoidGhvbmdkdkBnb2Rncm91cC5jbyIsInJvbGUiOiJTVVBFUl9BRE1JTiIsInpvbmUiOiIifQ.oJGvFj_KDmiOGyXKjHLaidli_ZQ9y5775D94fE8WkvU"
	//
	client := resty.New()
	startUrl := "http://localhost:8080/api/v1/files/upload-session/start"
	startRes, err := client.R().SetHeader("Authorization", token).Post(startUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	var uploadSession UploadSession
	err = json.Unmarshal(startRes.Body(), &uploadSession)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(uploadSession)
	fileContent, err := ioutil.ReadFile("./data/test2605.psd")
	if err != nil {
		fmt.Println(err)
		return
	}
	buffer := bytes.Buffer{}
	buffer.Write(fileContent)
	bytesArr := buffer.Bytes()
	size := 10 * 1024 * 1024
	numberOfPartitions := len(bytesArr) / size
	if len(bytesArr)%size != 0 {
		numberOfPartitions++
	}
	i := 0
	for i < numberOfPartitions {
		start := i * size
		end := (i + 1) * size
		if i == numberOfPartitions-1 {
			end = len(bytesArr) - 1
		}
		data := bytesArr[start:end]
		appendUrl := "http://localhost:8080/api/v1/files/upload-session/append"
		newClient := resty.New()
		appendRes, err := newClient.R().SetHeader("Authorization", token).
			SetBody(data).
			SetHeader("UploadSessionId", fmt.Sprintf("%v", uploadSession.Data.SessionId)).
			Post(appendUrl)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(appendRes)
		i++
	}
	finishUrl := "http://localhost:8080/api/v1/files/upload-session/finish"
	finishClient := resty.New()
	finishBody := map[string]interface{}{
		"sessionId": fmt.Sprintf("%v", uploadSession.Data.SessionId),
		"fileName":  fileName,
		"path":      fmt.Sprintf("/Test1/%s", fileName),
	}
	finishBodyPayload, err := json.Marshal(finishBody)
	if err != nil {
		fmt.Println(err)
	}
	finishRes, err := finishClient.R().SetHeader("Authorization", token).SetBody(finishBodyPayload).Post(finishUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(finishRes)
}
