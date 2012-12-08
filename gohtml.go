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

type htmltag struct {
	name       string
	attributes tagattr
	content    *list.List
}

type tagattr map[string]string

func (p tagattr) String() string {
	var paramlist string
	paramlist = ""
	for k, v := range p {
		paramlist += " " + String(k) + "=" + "\"" + String(v) + "\""
	}
	return paramlist
}

func construct(name string, params tagattr, value ...interface{}) htmltag {
	var tag htmltag
	var vallist *list.List
	vallist = list.New()
	for _, v := range value {
		vallist.PushBack(v)
	}
	tag.name = name
	tag.attributes = params
	tag.content = vallist
	return tag
}

func (t htmltag) String() string {
	var opentag, closetag, tagcontent string
	opentag = "<" + t.name
	if len(t.attributes) > 0 {
		opentag += " " + t.attributes.String()
	}
	opentag += ">"
	closetag = "</" + t.name + ">"
	tagcontent = ""
	for i := t.content.Front(); i != nil; i = i.Next() {
		tagcontent += String(i.Value)
	}
	return opentag + tagcontent + closetag
}

func String(value interface{}) string {
	var result string
	result = ""
	switch t := value.(type) {
	case string:
		result += escape(t)
	case htmltag:
		result += t.String()
	default:
		panic("unknown tag type")
	}
	return result
}

func escape(s string) string {
	return s
}

func main() {
	var progname, progpath string
	if len(os.Args) > 0 {
		progpath = os.Args[0]
		progname = path.Base(progpath)
	} else {
		panic("No params!")
	}
	phead := construct(HEAD, nil,
		construct(TITLE, nil, "Path to program"))
	pbody := construct(BODY, nil,
		construct(H1, tagattr{"id": "myfile", "class": "executable"}, path.Clean(path.Dir(progpath))+string(os.PathSeparator)+progname))
	page := construct(HTML, nil, phead, pbody)
	fmt.Println(page.String())
}
