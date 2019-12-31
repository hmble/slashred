package slashred

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func PrintHeader(resp *http.Response) {
	fmt.Println(resp.Status)

	fmt.Println("--------REQUEST HEADER--------")
	fmt.Printf("Ratelimit Used : %s\n", resp.Header.Get("X-Ratelimit-Used"))
	fmt.Printf("Ratelimit Remaining : %s\n", resp.Header.Get("X-Ratelimit-Remaining"))
	fmt.Printf("Ratelimit Reset : %s\n", resp.Header.Get("X-Ratelimit-Reset"))
	fmt.Println("--------REQUEST HEADER--------")

}

func SaveResponse(r io.Reader, filepath string) {
	var data map[string]interface{}
	body, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal("Error in reading response body")
	}
	er := json.Unmarshal([]byte(body), &data)

	if er != nil {
		log.Fatal("Error in printing json")
	}

	//j, jerr := json.MarshalIndent(&data, "", "    ")
	j, jerr := json.Marshal(&data)

	if jerr != nil {
		log.Fatal("Error in pretty printing json")
	}

	ToFile(filepath, j)

}

func ToFile(filepath string, b []byte) {
	f, err := os.Create(filepath)

	if err != nil {
		log.Fatal("Error in creating file")
	}

	defer f.Close()

	out, err := f.Write(b)

	if err != nil {
		log.Fatal("Error in writing bytes to file")
	}

	fmt.Printf("Wrote %d files\n", out)

}
