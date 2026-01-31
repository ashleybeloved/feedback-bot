package utils

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type BanList struct {
	Cache    map[int64]time.Time
	Filename string
}

var BanListCache = &BanList{}

func (bl *BanList) Load(filename string) error {
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

func (bl *BanList) save() {
	file, err := os.Create(bl.Filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.Encode(bl.Cache)
}

func (bl *BanList) BanUser(id int64) {
	bl.Cache[id] = time.Now()
	bl.save()
}

func (bl *BanList) UnbanUser(id int64) {
	delete(bl.Cache, id)
	bl.save()
}

func (bl *BanList) IsBanned(id int64) bool {
	_, ok := bl.Cache[id]
	if ok {
		return true
	}

	return false
}
