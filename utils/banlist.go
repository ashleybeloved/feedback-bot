package utils

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Cache struct {
	Cache    map[int64]time.Time
	Filename string
}

var BanListCache = &Cache{}

func (bl *Cache) Load(filename string) error {
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		bl.Cache = make(map[int64]time.Time)
		bl.Filename = filename
		return nil
	}
	if err != nil {
		return err
	}
	defer file.Close()

	bl.Cache = make(map[int64]time.Time)
	bl.Filename = filename
	return json.NewDecoder(file).Decode(&bl.Cache)
}

func (bl *Cache) save() {
	file, err := os.Create(bl.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.Encode(bl.Cache)
}

func (bl *Cache) BanUser(id int64) {
	bl.Cache[id] = time.Now()
	bl.save()
}

func (bl *Cache) UnbanUser(id int64) {
	delete(bl.Cache, id)
	bl.save()
}

func (bl *Cache) IsBanned(id int64) bool {
	_, ok := bl.Cache[id]
	if ok {
		return true
	}

	return false
}
