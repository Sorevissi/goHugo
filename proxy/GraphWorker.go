package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type Node struct {
	ID    int
	Name  string
	Form  string
	Links []*Node
}

func (n *Node) getForm() string {
	formNode := map[string]string{
		"circle":     fmt.Sprintf("  %s((%s %s))\n", n.Name, n.Form, n.Name),
		"rect":       fmt.Sprintf("  %s[%s %s]\n", n.Name, n.Form, n.Name),
		"ellipse":    fmt.Sprintf("  %s([%s %s])\n", n.Name, n.Form, n.Name),
		"round-rect": fmt.Sprintf("  %s(%s %s)\n", n.Name, n.Form, n.Name),
		"rhombus":    fmt.Sprintf("  %s{%s %s}\n", n.Name, n.Form, n.Name),
	}

	return formNode[n.Form]
}

const templateGraph = `
---
menu:
    after:
        name: graph
        weight: 1
title: Построение графа
---

{{< mermaid >}}
graph TD
%s
{{< /mermaid >}}
`

func GraphWorker() {
	t := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-t.C:
			graphContent := generateRandomGraph()
			err := os.WriteFile("/app/static/tasks/graph.md", []byte(graphContent), 06440)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func generateRandomGraph() string {
	nodeCount := rand.Intn(26) + 5

	var nodes []Node

	for i := 1; i <= nodeCount; i++ {
		node := Node{
			ID:    i,
			Name:  fmt.Sprintf("Node%d", i),
			Form:  getRandomForm(),
			Links: nil,
		}
		nodes = append(nodes, node)
	}

	for i := range nodes {
		for j := range nodes {
			if i != j && rand.Float32() < 0.5 {
				nodes[i].Links = append(nodes[i].Links, &nodes[j])
			}
		}
	}

	var mermaidNodes string

	for _, node := range nodes {
		mermaidNodes += node.getForm()
		for _, link := range node.Links {
			mermaidNodes += fmt.Sprintf("  %s --> %s\n", node.Name, link.Name)
		}
	}

	return fmt.Sprintf(templateGraph, mermaidNodes)
}

func getRandomForm() string {
	form := []string{"circle", "rect", "ellipse", "round-rect", "rhombus"}
	return form[rand.Intn(len(form))]
}
