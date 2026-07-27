package blogposts

import (
	"bufio"
	"io"
)

type Post struct {
	Title       string
	Description string
}

func newPost(postFile io.Reader) (Post, error) {
	var (
		titleLine       string = ""
		descriptionLine string = ""
	)

	scanner := bufio.NewScanner(postFile)
	if scanner.Scan() {
		titleLine = scanner.Text()
	}

	if scanner.Scan() {
		descriptionLine = scanner.Text()
	}

	return Post{Title: titleLine[7:], Description: descriptionLine[13:]}, nil
}
