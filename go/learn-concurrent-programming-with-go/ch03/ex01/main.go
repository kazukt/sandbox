package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countWords(url string, frequency map[string]int) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic("Server returning error status code: " + resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)
	wordRegex := regexp.MustCompile(`[a-zA-Z]+`)
	for _, w := range wordRegex.FindAllString(string(body), -1) {
		w = strings.ToLower(w)
		frequency[w]++
	}
	fmt.Printf("Completed: %s\n", url)
}

func main() {
	frequency := make(map[string]int)
	for i := 1000; i < 1005; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countWords(url, frequency)
	}
	time.Sleep(5 * time.Second)
	for w, c := range frequency {
		fmt.Printf("%s -> %d\n", w, c)
	}
}
