package xkcd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type XkcdStruct struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	Safe_title string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func Xkcd() {
	max_num, err := get_max_num()
	if err != nil {
		log.Fatal(err)
	}

	indexFile := "xkcd_index.json"
	index, err := loadIndex(indexFile)
	if err != nil {
		index = make(map[int]XkcdStruct)
		for i := 1; i <= max_num; i++ {
			comic, err := getComic(i)
			if err != nil {
				log.Printf("Failed to get comic %d: %v", i, err)
				continue
			}
			index[i] = comic
		}
		saveIndex(indexFile, index)
	}

	terms := os.Args[1:]
	if len(terms) == 0 {
		fmt.Println("Usage: xkcd <search terms>")
		return
	}
	searchIndex(index, terms)
}

func get_max_num() (int, error) {
	resp, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		return 0, errors.New("can't get the xkcd comic amount")
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return 0, errors.New("can't surf xkcd")
	}
	var res XkcdStruct
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return 0, errors.New("can't decode json")
	}
	resp.Body.Close()
	return res.Num, nil
}

func getComic(num int) (XkcdStruct, error) {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", num)
	resp, err := http.Get(url)
	if err != nil {
		return XkcdStruct{}, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return XkcdStruct{}, errors.New("can't get comic")
	}
	var comic XkcdStruct
	err = json.NewDecoder(resp.Body).Decode(&comic)
	if err != nil {
		return XkcdStruct{}, err
	}
	resp.Body.Close()
	return comic, nil
}

func saveIndex(filename string, index map[int]XkcdStruct) error {
	data, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func loadIndex(filename string) (map[int]XkcdStruct, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var index map[int]XkcdStruct
	err = json.Unmarshal(data, &index)
	if err != nil {
		return nil, err
	}
	return index, nil
}

func searchIndex(index map[int]XkcdStruct, terms []string) {
	for num, comic := range index {
		for _, term := range terms {
			if strings.Contains(comic.Title, term) || strings.Contains(comic.Safe_title, term) || strings.Contains(comic.Alt, term) {
				fmt.Printf("Comic #%d: %s\nURL: %s\n", num, comic.Title, comic.Img)
				break
			}
		}
	}
}
