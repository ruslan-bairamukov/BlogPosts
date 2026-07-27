package blogposts

import (
	"io/fs"
)

type Post struct{}

func NewPostsFromFS(fileSystem fs.FS) []Post {
	dir, _ := fs.ReadDir(fileSystem, ".")
	posts := make([]Post, 0, 2)
	for range dir {
		posts = append(posts, Post{})
	}

	return posts
}
