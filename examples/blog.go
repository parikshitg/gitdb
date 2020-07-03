package main

import (
	"encoding/json"
	"fmt"

	"github.com/parikshitg/gitdb"
)

type Blog struct {
	gitdb.Header
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (b *Blog) GetHeader() *gitdb.Header {
	return &b.Header
}

func (b *Blog) Serialize() ([]byte, error) {
	return json.Marshal(b)
}

func (b *Blog) Deserialize(data []byte) error {
	return json.Unmarshal(data, b)
}

func (b *Blog) Copy() interface{} {
	var blog = &Blog{
		Title: b.Title,
		Body:  b.Body,
	}
	blog.Header.ID = b.Header.ID
	return blog
}

func (b Blog) String() string {
	s := fmt.Sprintf("Title : [%s]\n", b.Title)
	s += fmt.Sprintf("Body : [%s]\n", b.Body)
	return s
}
