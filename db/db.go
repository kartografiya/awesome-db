package db

import (
	"log"
	
	"github.com/kr/pretty"
)

type Video struct {
	Title string
	Director string
	Year int
}

type DB []Video

func Init() DB {
	return DB{}
}

func (d *DB) Add(v Video) {
	*d = append(*d, v)
}

func (d DB) Size() int {
	return len(d)
}

func (d *DB) String() string {
	log.Print("ouch")
	return pretty.Sprint(*d)
}
