package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
	"text/template/parse"

	"github.com/k0kubun/pp"
)

var Version string

func main() {
	filename := os.Args[1]
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("visualforce").Parse(string(body))
	if err != nil {
		panic(err)
	}
	p := &Printer{}
	content, err := p.traverse(tmpl.Tree.Root)
	if err != nil {
		panic(err)
	}
	fmt.Printf("<apex:page>\n%s\n</apex:page>", content)
}

type Printer struct{}

func (p *Printer) traverse(node parse.Node) (string, error) {
	switch node := node.(type) {
	case *parse.ActionNode:
	case *parse.IfNode:
		cmd := node.BranchNode.Pipe.Cmds[0]
		var condition string
		switch cmd.Args[0].Type() {
		case parse.NodeField:
			condition = cmd.String()[1:]
		case parse.NodeIdentifier:
			ident := cmd.Args[0].(*parse.IdentifierNode).Ident
			//debug(cmd.Args)
			switch ident {
			case "eq":
				ident := cmd.Args[1].String()[1:]
				value := cmd.Args[2].String()
				condition = fmt.Sprintf("%s == %s", ident, value)
			}
		}
		ifContent, err := p.traverse(node.BranchNode.List)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("<apex:outputPanel rendered=\"{!%s}\">%s</apex:outputPanel>", condition, ifContent), nil
	case *parse.ListNode:
		lines := make([]string, len(node.Nodes))
		var err error
		for i, children := range node.Nodes {
			lines[i], err = p.traverse(children)
			if err != nil {
				return "", err
			}
		}
		return strings.Join(lines, "\n"), nil
	case *parse.RangeNode:
	case *parse.TemplateNode:
	case *parse.TextNode:
		return node.String(), nil
	case *parse.WithNode:
	default:
	}
	return "", nil
}

func debug(args ...interface{}) {
	pp.Println(args...)
}
