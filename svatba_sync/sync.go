package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	config              = map[string]string{}
	configPath          = "~/.svatba/sync_config.json"
	syncIntervalSeconds = 5
)

func main() {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Přemku, startujem, tady je konfigurace: ")
	fmt.Println(config)

	var results []ApiResult

	for {
		results, err = GetAPIResults()
		if err == nil {

			uploaded, err := UploadedFiles()
			if err == nil {
				for _, v := range results {
					if uploaded[v.UID] == false {
						err = UploadFile(v)
						if err != nil {
							fmt.Println("Přemku, chyba! 😰")
							fmt.Println(err)
						}
					}
				}
			} else {
				fmt.Println("Přemku, chyba! 😰")
				fmt.Println(err)
			}
		} else {
			fmt.Println("Přemku, chyba! 😰")
			fmt.Println(err)
		}

		time.Sleep(time.Duration(syncIntervalSeconds) * time.Second)
	}
}

func UploadFile(f ApiResult) error {
	fmt.Printf("🍺 Synchronizuji soubor %s v čase %s\n", f.UID, time.Now().Format(time.RFC3339))

	out, err := os.Create(config["targetDir"] + "/" + f.UID + ".jpg")
	defer out.Close()

	resp, err := http.Get(config["baseUrl"] + f.Original)
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func UploadedFiles() (map[string]bool, error) {
	ret := map[string]bool{}

	f, err := os.Open(config["targetDir"])
	if err != nil {
		return nil, err
	}
	defer f.Close()

	names, err := f.Readdirnames(-1)
	if err != nil {
		return nil, err
	}

	for _, v := range names {
		if strings.HasSuffix(v, ".jpg") {
			name := strings.Split(v, ".")[0]
			ret[name] = true
		}
	}
	return ret, nil
}

func GetAPIResults() ([]ApiResult, error) {
	var results []ApiResult
	res, err := http.Get(config["baseUrl"] + "/api/print")
	if err != nil {
		return results, err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return results, err
	}

	err = json.Unmarshal(data, &results)
	if err != nil {
		return results, err
	}
	return results, nil
}

type ApiResult struct {
	UID      string `json:"uid"`
	Original string `json:"original"`
}
