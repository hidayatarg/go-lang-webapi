package main

import (
	"log"
	db "./db"
)

func main()  {
	log.Printf("Hello GoLang PostgreSql Implementation working...")
	db.Connect()
}
