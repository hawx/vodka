package main

import (
	"strings"
	"regexp"
	"io/ioutil"
	"github.com/hoisie/mustache"
	"sort"

	"github.com/hawx/vodka/parser"
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


type FunctionData struct {
	Name         string
	Sig          string
	Description  string
	Example      interface{}
	Source       interface{}
}

func (f *FunctionDoc) Data() FunctionData {
	r := new(FunctionData)

	r.Name        = f.name
	r.Sig         = strings.Replace(f.Meta("sig"), "->", "→", -1)
	r.Description = f.String()
	if f.HasMeta("example") {
		r.Example   = f.Meta("example")
	}
	if f.src != "" {
		r.Source    = f.src
	}

	return *r
}

type GroupData struct {
	Group      string
	GroupID    string
	Functions  *[]FunctionData
}

func NewGroupData(name string) GroupData {
	r := new(GroupData)
	r.Group = name
	r.GroupID = strings.Replace(name, " ", "", -1)
	r.Functions = new([]FunctionData)
	return *r
}

func (g GroupData) AddFunctionData(fd FunctionData) {
	*g.Functions = append(*g.Functions, fd)
}

// Allow sorting functions by name
func (g GroupData) Len() int {
	return len(*g.Functions)
}

func (g GroupData) Less(i, j int) bool {
	return (*g.Functions)[i].Name < (*g.Functions)[j].Name
}

func (g GroupData) Swap(i, j int) {
	(*g.Functions)[i], (*g.Functions)[j] = (*g.Functions)[j], (*g.Functions)[i]
}

// Allow sorting GroupDatas by name
type GroupDatas []GroupData

func (g GroupDatas) Len() int      { return len(g) }
func (g GroupDatas) Swap(i, j int) { g[i], g[j] = g[j], g[i] }

func (g GroupDatas) Less(i, j int) bool {
	return g[i].Group < g[j].Group
}

func (g *GroupDatas) Add(val GroupData) {
	*g = append(*g, val)
}

func Doc(input []string, output string) {
	tmpl, _ := mustache.ParseFile("doc.mustache")

	for _, file := range input {
	  contents, _ := ioutil.ReadFile(file)
		list := parser.FullParse(string(contents))
		docs := split(*list)

		groupDatas := map[string]GroupData{}

		// Set up all groups first
		for _, d := range docs {
			groupDatas[d.grp] = NewGroupData(d.grp)
		}

		// Now add functions
		for _, d := range docs {
			groupDatas[d.grp].AddFunctionData(d.Data())
		}

		// Now gather map values for array
		finalDatas := new(GroupDatas)

		for _, v := range groupDatas {
			sort.Sort(v)
			finalDatas.Add(v)
		}

		sort.Sort(finalDatas)

		data := map[string]interface{}{
			"Title": "Docs",
			"Groups": finalDatas,
		}

		str := tmpl.Render(data)
		ioutil.WriteFile(output, []byte(str), 0644)
	}
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
func split(list parser.Tokens) []FunctionDoc {
	docs := []FunctionDoc{}
	group := ""

	for i := 0; i < len(list); i++ {
		if list[i].Key == "comment" {
			if strings.HasPrefix(list[i].Val, "group: ") {
				group = list[i].Val[7:]

			} else {
				idx, doc := collectComments(list, i)
				i = idx
				// note that i has been changed by collectComments
				if list[i].Key == "str" {
					name := list[i].Val
					i++
					if list[i].Key == "stm" {
						docs = append(docs, FunctionDoc{name, doc, list[i].Val, group})

					} else if list[i].Key == "fun" && list[i].Val == "__document__" {
						docs = append(docs, FunctionDoc{name, doc, "", group})
					}
				}
			}
		}
	}

	return docs
}

func collectComments(list parser.Tokens, idx int) (int, string) {
	doc := ""
	for i := idx; i < len(list); i++ {
		if list[i].Key == "comment" {
			doc += strings.TrimSpace(list[i].Val) + " \n"
		} else {
			return i, doc
		}
	}
	return len(list), doc
}
