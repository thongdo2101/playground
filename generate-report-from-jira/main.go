package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

type Task struct {
	Name       string   `json:"name"`
	Status     string   `json:"status"`
	Project    string   `json:"project"`
	Members    []string `json:"members"`
	StartDate  string   `json:"startDate"`
	EndDate    string   `json:"endDate"`
	StoryPoint float64  `json:"storyPoint"`
	Extra      bool     `json:"extra"`
}

func main() {
  yearMonth := "2022-04-"
	bug := "🐞"
	records := readCsvFile("./data.csv")
	tasks := make([]Task, 0)
	for i, row := range records {
		if i == 0 {
			continue
		}
		isBug := false
		if row[0] == "Bug" {
			isBug = true
		}
		name := row[1]
		if isBug {
			name = bug + " " + name
		}
		rawStatus := row[16]
		status := rawStatus
		if rawStatus == "Done" {
			status = "Completed"
		} else if rawStatus == "To Do" {
			status = "Todo"
		}
		project := row[2]
		if project == "PODLovers" {
			project = "POD Lovers"
		}
		members := make([]string, 0)
		member := row[5]
    // mapping name on jira with report
		if member == "Thong Do" {
			member = "Đỗ Văn Thông"
		} else if member == "Nguyễn Đình Quân" {
			member = "Đình Quân"
		} else if member == "Nguyen Tai Dat" {
			member = "Nguyễn Tài Đạt"
		} else if member == "Ha Quang Huy" {
			member = "Hà Quang Huy"
		}
    //
		startDateTemp := strings.Split(row[7], "/")
		startDate := yearMonth + startDateTemp[0]
		endDateTemp := strings.Split(row[7], "/")
		endDate := yearMonth + endDateTemp[0]
		members = append(members, member)
		storyPoint := 0.0
		if row[10] != "" {
			storyPoint, _ = strconv.ParseFloat(row[10], 64)
		}
		extra := false
		if strings.Contains(row[13], "Extra") {
			extra = true
		}
		tasks = append(tasks, Task{
			Name:       name,
			Status:     status,
			Project:    project,
			Members:    members,
			StartDate:  startDate,
			EndDate:    endDate,
			StoryPoint: storyPoint,
			Extra:      extra,
		})
	}
	// fmt.Println(tasks)
	data, _ := json.Marshal(tasks)
	err := ioutil.WriteFile("output.json", data, 0644)
	if err != nil {
		panic(err)
	}
}
