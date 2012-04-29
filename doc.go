package main

import (
	"strings"
	"regexp"
	"io/ioutil"
)

type FunctionDoc struct {
	name  string
	doc   string
	src   string
	grp   string
}

func (f *FunctionDoc) String() string {
	sig, _ := regexp.Compile("^[a-z]+:")
	s := ""
	for _, line := range f.lines() {
		if !sig.MatchString(line) {
			s += line
		}
	}
	return s
}

func (f *FunctionDoc) lines() []string {
	return strings.Split(f.doc, "\n")
}

func (f *FunctionDoc) Meta(name string) string {
	for _, line := range f.lines() {
		if strings.HasPrefix(line, name + ":") {
			return strings.TrimSpace(line[len(name)+1:])
		}
	}
	return ""
}

func (f *FunctionDoc) HasMeta(name string) bool {
	return f.Meta(name) != ""
}

func (f *FunctionDoc) Html() string {
	s := "<li>\n"
	s += "  <h1 id='" + f.name + "'>" + f.name + "</h1>\n"
	if f.HasMeta("sig") {
		s += "  <h2>" + strings.Replace(f.Meta("sig"), "->", "→", 1) + "</h2>\n"
	}
	s += "  <p>" + f.String() + "<p>\n"
	if f.HasMeta("example") {
		s += "  <div class='example'>\n"
		s += "    <h3>Example</h3>\n"
		s += "    <pre><code>" + f.Meta("example") + "</code></pre>\n"
		s += "  </div>\n"
	}
	if f.src != "" {
		s += "  <div class='source'>\n"
		s += "    <h3>Source</h3>\n"
		s += "    <pre><code>" + f.src + "</code></pre>\n"
		s += "  </div>\n"
	}
	s += "</li>\n"

	return s
}

// Generates documentation for the given list of input files, writing the output
// to the output directory given.
func Doc(input []string, output string) {
	s := "<!DOCTYPE html>\n" +
		"<html>\n" +
		"  <head>\n" +
		"    <meta charset='utf-8' />" +
		"    <title>Docs</title>\n" +
		"    <link rel='stylesheet' href='style.css' type='text/css' />\n" +
		"  </head>\n" +
		"  <body>\n" +
		"    <header>\n" +
		"      <h1>Docs</h1>\n" +
		"    </header>\n"

	for _, file := range input {
		contents, _ := ioutil.ReadFile(file)
		list := FullParse(string(contents))

		docs := split(*list)

		html := ""
		group := docs[0].grp
		s += "<div class='summary'>\n"
		if group != "" {
			s += "<h3>" + group + "</h3>\n"
			html += "<h1 class='group'>" + group + "</h1>"
		}
		s += "<ul>\n"
		for _, d := range docs {
			if d.grp != group {
				group = d.grp
				s += "</ul>\n<h3>" + group + "</h3>\n<ul>\n"
				html += "<h1 class='group'>" + group + "</h1>"
			}
			s += "<li><a href='#" + d.name + "'>" + d.name + "</a></li>\n"
			html += d.Html()
		}
		s += "  <div style='clear:both;'></div>\n"
		s += "</ul>\n</div>\n"
		s += "<ul class='docs'>\n" + html + "</ul>\n"
	}

	s += "</body>\n</html>"

	ioutil.WriteFile(output, []byte(s), 0644)
	// println(s)
}

// Takes the parsed list of tokens and splits them into function definitions
// with the associated doc-comments, each "chunk" will have the pattern:
//
//   comment(s) string block "define"
//
// or
//
//   comment(s) string "__document__"
//
func split(list Tokens) []FunctionDoc {
	docs := []FunctionDoc{}
	group := ""

	for i := 0; i < len(list); i++ {
		if list[i].key == "comment" {
			if strings.HasPrefix(list[i].val, "group: ") {
				group = list[i].val[7:]

			} else {
				idx, doc := collectComments(list, i)
				i = idx
				// note that i has been changed by collectComments
				if list[i].key == "str" {
					name := list[i].val
					i++
					if list[i].key == "stm" {
						docs = append(docs, FunctionDoc{name, doc, list[i].val, group})

					} else if list[i].key == "fun" && list[i].val == "__document__" {
						docs = append(docs, FunctionDoc{name, doc, "", group})
					}
				}
			}
		}
	}

	return docs
}

func collectComments(list Tokens, idx int) (int, string) {
	doc := ""
	for i := idx; i < len(list); i++ {
		if list[i].key == "comment" {
			doc += strings.TrimSpace(list[i].val) + " \n"
		} else {
			return i, doc
		}
	}
	return len(list), doc
}