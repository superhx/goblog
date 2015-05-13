package blog

import (
	"testing"
)

func TestBlog(t *testing.T) {
	blog := &Blog{parallels: make(chan bool, articleParallelCount)}
	blog.Transform()
}
