package services

import (
	"bufio"
	"log"
	"os"
	"pieFireDire/models"
	"strings"
)

type BeefInterface interface {
	Read(filepath string) error
	Count(str string)
	Get() map[string]int
}

func NewBeefService() BeefInterface {
	b := &Beef{}
	b.CounterList = make(map[string]int)
	return b
}

type Beef struct {
	CounterList map[string]int
}

func (b *Beef) Read(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
		return models.ErrOpenFile
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			b.Count(scanner.Text())
		}
	}
	return nil
}

func (b *Beef) Count(str string) {
	str = strings.ReplaceAll(str, ".", "")
	str = strings.ReplaceAll(str, ",", "")
	strSplit := strings.Split(str, " ")
	for _, v := range strSplit {
		if v != "" {
			b.CounterList[strings.ToLower(v)] += 1
		}
	}
}

func (b *Beef) Get() map[string]int {
	return b.CounterList
}
