package advent2023

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type GetInputParams struct {
	day      string
	fileName string
}

func GetInputByDay(input GetInputParams) (error, string) {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err), ""
	}

	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("could not get current working directory: %w", err), ""
	}

	filepath := dir + "/" + input.day + "." + input.fileName + ".txt"

	if _, err := os.Stat(filepath); err == nil {
		content, err := os.ReadFile(filepath)
		if err != nil {
			return fmt.Errorf("failed to read existing file: %w", err), ""
		}
		return nil, string(content)
	}

	client := &http.Client{}

	n, _ := strconv.Atoi(input.day)
	url := "https://adventofcode.com/2023/day/" + strconv.Itoa(n) + "/input"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err), ""
	}

	cookie := os.Getenv("AOC_USER_COOKIE")
	cookieHeader := "session=" + cookie
	req.Header.Add("cookie", cookieHeader)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to perform request: %w", err), ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status), ""
	}

	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err), ""
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err), ""
	}

	fmt.Printf("File downloaded successfully and saved to %s\n", filepath)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err), ""
	}

	return nil, string(body)
}
