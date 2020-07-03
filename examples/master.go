package main

import (
	"fmt"
	"log"

	"github.com/parikshitg/gitdb"
)

var db *gitdb.GitDB

func main() {

	log.SetFlags(log.Ltime | log.Lshortfile)

	var err error
	var blogs = []Blog{
		{Title: "This is my first blog", Body: "This is first blog body"},
		{Title: "This is my second blog", Body: "This is second blog body"},
		{Title: "This is my third blog", Body: "This is third blog body"},
		{Title: "This is my forth blog", Body: "This is forth blog body"},
		{Title: "This is my fifth blog", Body: "This is fifth blog body"},
	}

	// Open db

	db, err = gitdb.Open("master.db")
	if err != nil {
		log.Fatal("Error in gitdb.Open: ", err)
	}

	// Creating Blog

	var id string
	for _, blog := range blogs {
		id, err = db.Create("BLOG", &blog)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Read blog(last blog)
	var b Blog
	err = db.Read("BLOG", id, &b)
	if err != nil {
		log.Println("Read Error : ", err)
	}

	fmt.Println("blog: ", b)

	// Update Blog
	var b2 Blog
	b2 = Blog{Title: "FIFTH", Body: "FIFTHBODY"}

	err = db.Update("BLOG", id, &b2)
	if err != nil {
		log.Fatal("Update Error : ", err)
	}

	// read last blog again to see updates
	err = db.Read("BLOG", id, &b)
	if err != nil {
		log.Fatal("Read Error : ", err)
	}

	fmt.Println("blog: ", b)

	// Delete last blog
	err = db.Delete("BLOG", id)
	if err != nil {
		log.Fatal("Delete error :", err)
	}

	// read last blog again to see deletion (gives read error)
	err = db.Read("BLOG", id, &b)
	if err != nil {
		log.Println("Read Error : ", err)
	}

	fmt.Println("blog: ", b)
}
