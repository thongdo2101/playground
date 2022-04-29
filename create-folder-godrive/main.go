package main

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

const path = "Teams/2022/GDM/Google"
const nameFolder = "Trọng Khoa"

// const monthFolder = "August"
// const monthFolder = "September"
const monthFolder = "May"

// const monthFolder = "November"
// const monthFolder = "December"

// const monthNumber = "08"
// const monthNumber = "09"
const monthNumber = "05"

// const monthNumber = "11"
// const monthNumber = "12"
const token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTA5NjU0MDMsInVzZXJJRCI6ImMxNWYzNmQyLTU5NGItNDQwNC1hMTJjLTUxYmQ4NzM0MGQ2ZiIsImVtYWlsIjoidGhvbmdkdkBnb2Rncm91cC5jbyIsInJvbGUiOiJTVVBFUl9BRE1JTiIsInpvbmUiOiIifQ.-V17zIadbMmUeDqqELyBwEUSu8SfvzXoP5fN3otJ7us"

func createNameFolder() {
	client := resty.New()
	url := "https://contents.godmerch.com/api/v1/files/create-folder"
	temp := fmt.Sprintf("%s/%s", path, nameFolder)
	res, err := client.R().SetHeader("Authorization", token).SetBody(map[string]interface{}{
		"path": temp,
	}).Post(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	fmt.Println(temp)
}

func createMonthFolder() {
	client := resty.New()
	url := "https://contents.godmerch.com/api/v1/files/create-folder"
	temp := fmt.Sprintf("%s/%s/%s", path, nameFolder, monthFolder)
	res, err := client.R().SetHeader("Authorization", token).SetBody(map[string]interface{}{
		"path": temp,
	}).Post(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	fmt.Println(temp)
}

func createDaysFolder() {
	for i := 1; i < 32; i++ {
		time.Sleep(500)
		client := resty.New()
		url := "https://contents.godmerch.com/api/v1/files/create-folder"
		temp := fmt.Sprintf("%s/%s/%s/%s", path, nameFolder, monthFolder, fmt.Sprintf("%s.%v", monthNumber, i))
		if i < 10 {
			temp = fmt.Sprintf("%s/%s/%s/%s", path, nameFolder, monthFolder, fmt.Sprintf("%s.0%v", monthNumber, i))
		}
		res, err := client.R().SetHeader("Authorization", token).SetBody(map[string]interface{}{
			"path": temp,
		}).Post(url)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
		fmt.Println(temp)
	}
}

func main() {
	// createNameFolder()
	createMonthFolder()
	createDaysFolder()
	fmt.Println("done")
}
