package registry

import "strings"

type Image struct {
	Name      string
	Reference string
}

func NewImage(s string) Image {
	split := strings.Split(s, ":")
	name := ""
	ref := ""
	if len(split) == 1 {
		name = s
		ref = "latest"
	} else {
		name = split[0]
		ref = split[1]
	}
	return Image{name, ref}
}
