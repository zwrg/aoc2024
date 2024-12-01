package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type Config struct {
	Session string `json:"session"`
	DataDir string `json:"data_dir"`
	Year    int    `json:"year"`
}

func loadConfig(filePath string) (Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	return config, err
}

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config.json:", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage program.exe <day>")
		return
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil || day < 1 || day > 25 {
		fmt.Println("Error: Invalid day. Must be between 1 and 25.")
		return
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", config.Year, day)
	filePath := filepath.Join(config.DataDir, strconv.Itoa(config.Year), fmt.Sprintf("day%d.txt", day))

	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Add("Cookie", "session="+config.Session)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Received status code %d\n", resp.StatusCode)
		return
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	err = os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	err = os.WriteFile(filePath, content, 0644)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	fmt.Printf("Input for day %d saved to %s\n", day, filePath)
}
