package blogposts

import (
	"bufio"
	"io"
	"strings"
)

const (
	titlePrefix       = "Title: "
	descriptionPrefix = "Description: "
	tagsPrefix        = "Tags: "
	sep               = ", "
	end               = "\n"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
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
	tags := strings.Split(readLine(tagsPrefix), sep)
	body, err := readBody(scanner)
	if err != nil {
		return Post{}, err
	}

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        body,
	}, nil
}

func readBody(scanner *bufio.Scanner) (string, error) {
	scanner.Scan() // to skip "---" line

	var builder strings.Builder
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteString(end)
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	body := strings.TrimSuffix(builder.String(), end)

	return body, nil
}
