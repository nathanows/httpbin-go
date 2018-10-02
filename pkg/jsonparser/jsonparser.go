package jsonparser

import (
	"encoding/json"
	"strings"
)

// Container - an internal structure that holds a reference to the core interface map of the parsed
// json. Use this container to move context.
type Container struct {
	object interface{}
}

// Data - Return the contained data as an interface{}.
func (g *Container) Data() interface{} {
	if g == nil {
		return nil
	}
	return g.object
}

// Val - Return the contained data as a string value
func (g *Container) String() string {
	if g == nil {
		return ""
	}

	str, ok := g.object.(string)
	if !ok {
		return ""
	}
	return str
}

// Path - Search for a value using dot notation.
func (g *Container) Path(path string) *Container {
	return g.Search(strings.Split(path, ".")...)
}

// Search - Attempt to find and return an object within the JSON structure by specifying the
// hierarchy of field names to locate the target. If the search encounters an array and has not
// reached the end target then it will iterate each object of the array for the target and return
// all of the results in a JSON array.
func (g *Container) Search(hierarchy ...string) *Container {
	var object interface{}

	object = g.Data()
	for target := 0; target < len(hierarchy); target++ {
		if mmap, ok := object.(map[string]interface{}); ok {
			object, ok = mmap[hierarchy[target]]
			if !ok {
				return nil
			}
		} else if marray, ok := object.([]interface{}); ok {
			tmpArray := []interface{}{}
			for _, val := range marray {
				tmpCont := &Container{val}
				res := tmpCont.Search(hierarchy[target:]...)
				if res != nil {
					tmpArray = append(tmpArray, res.Data())
				}
			}
			if len(tmpArray) == 0 {
				return nil
			}
			return &Container{tmpArray}
		} else {
			return nil
		}
	}
	return &Container{object}
}

// ParseJSON - Convert a string into a representation of the parsed JSON.
func ParseJSON(sample []byte) (*Container, error) {
	var gabs Container

	if err := json.Unmarshal(sample, &gabs.object); err != nil {
		return nil, err
	}

	return &gabs, nil
}
