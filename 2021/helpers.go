package advent2021

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// GetInput grabs the plaintext puzzle input using a session cookie header
func GetInputByDay(day int) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}
	session := os.Getenv("SESSION_TOKEN")

	url := "https://adventofcode.com/2021/day/" + fmt.Sprint(day) + "/input"
	cookieHeader := "session=" + session
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("cookie", cookieHeader)

	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func ParseToNumSlice(i string) []int {
	splits := strings.Split(i, "\n")

	var nums []int
	for _, v := range splits {
		num, _ := strconv.Atoi(v)
		nums = append(nums, num)
	}

	return nums
}
