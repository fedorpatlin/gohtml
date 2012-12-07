// ostest.go
package main

import (
	"container/list"
	"fmt"
	"os"
	"path"
)

const (
	HTML  = "html"
	HEAD  = "head"
	TITLE = "title"
	BODY  = "body"
	H1    = "h1"
)

type tagpm map[string]string
type tagvalue interface {
	render() string
}

func (p tagpm) render() string {
	var paramlist string
	paramlist = ""
	for k, v := range p {
		paramlist += " " + k + "=" + "\"" + v + "\""
	}
	return paramlist
}

type htmltag struct {
	name   string
	params tagpm
	value  *list.List
}

func construct(name string, params tagpm, value ...interface{}) htmltag {
	var tag htmltag
	var vallist *list.List
	vallist = list.New()
	for _, v := range value {
		vallist.PushBack(v)
	}
	tag.name = name
	tag.params = params
	tag.value = vallist
	return tag
}
func (t htmltag) render() string {
	var opentag, closetag, tagval string
	opentag = "<" + t.name
	if len(t.params) > 0 {
		opentag += " " + t.params.render()
	}
	opentag += ">"
	closetag = "</" + t.name + ">"
	tagval = ""
	for i := t.value.Front(); i != nil; i = i.Next() {
		switch t := i.Value.(type) {
		case string:
			tagval += t
		case tagvalue:
			tagval += t.render()
		default:
			panic("unknown tag type")
		}
	}
	return opentag + tagval + closetag
}
func main() {
	var progname, progpath string
	os.Chdir("/")
	if len(os.Args) > 0 {
		progpath = os.Args[0]
		progname = path.Base(progpath)
	} else {
		panic("No params!")
	}
	phead := construct(HEAD, nil,
		construct(TITLE, nil, "Path to program"))
	pbody := construct(BODY, nil,
		construct(H1, tagpm{"param": "value", "param2": "value2"}, path.Clean(path.Dir(progpath))+string(os.PathSeparator)+progname))
	page := construct(HTML, nil, phead, pbody)
	fmt.Println(page.render())
}
