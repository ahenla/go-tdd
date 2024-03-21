package blogposts

import (
	"io/fs"
	"testing/fstest"
)

type Post struct {
}

func NewPostsFromFS(filesSystem fstest.MapFS) []Post {
	dir, _ := fs.ReadDir(filesSystem, ".")
	var posts []Post

	for range dir {
		posts = append(posts, Post{})
	}

	return posts
}
