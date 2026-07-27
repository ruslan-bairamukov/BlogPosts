package blogposts

import (
	"bufio"
	"io"
	"strings"
)

const (
	titlePrefix       = "Title: "
	descriptionPrefix = "Description: "
)

type Post struct {
	Title       string
	Description string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func(fieldPrefix string) string {
		if scanner.Scan() {
			return strings.TrimPrefix(scanner.Text(), fieldPrefix)
		}

		return ""
	}

	title := readLine(titlePrefix)
	description := readLine(descriptionPrefix)

	return Post{Title: title, Description: description}, nil
}
