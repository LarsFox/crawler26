package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/LarsFox/crawler26/go/entities"
	"github.com/LarsFox/crawler26/go/vk"
)

var (
	vkAPIToken string
)

func flags() {
	flag.StringVar(&vkAPIToken, "vk_api_token", "", "Access token to VK with necessary scope")
	flag.Parse()
}

type extractor interface {
	GetComments(count int64) ([]*entities.Comment, error)
}

func main() {
	flags()

	vkExtractor := vk.NewExtractor(vkAPIToken)
	extractAll(map[string]extractor{
		"vk": vkExtractor,
	})
}

func extractAll(xs map[string]extractor) {
	wg := &sync.WaitGroup{}
	for name, x := range xs {
		wg.Add(1)
		go extract(name, x, wg)
	}
	wg.Wait()
	log.Println("done!")
}

func extract(name string, x extractor, wg *sync.WaitGroup) {
	defer wg.Done()
	var count int64 = 1000
	comments, err := x.GetComments(count)
	if err != nil {
		log.Println(err)
		return
	}

	for _, c := range comments {
		log.Println(c.Text)
	}

	f, err := os.Create(fmt.Sprintf("output/%s.json", name))
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	json.NewEncoder(f).Encode(comments)
}
