package a

import (
	_ "unsafe" // for go:linkname
)

//go:linkname say github.com/go-project-action/go-linkname/src/b.Hi
func say(name string) string {
	return "hello, " + name
}
