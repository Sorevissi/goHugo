package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

const templateBinary = `
---
menu:
    after:
        name: binary_tree
        weight: 2
title: Построение сбалансированного бинарного дерева
---

{{< mermaid >}}
graph TD
%s
{{< /mermaid >}}
`

func BinaryWorker() {
	counter := 5
	tree := GenerateTree(counter)

	t := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-t.C:
			tree.Insert(counter)

			mermaidText := fmt.Sprintf(templateBinary, tree.ToMermaid())

			err := os.WriteFile("/app/static/tasks/binary.md", []byte(mermaidText), 06440)
			if err != nil {
				log.Println(err)
			}
			counter++
			if counter == 100 {
				counter = 5
				tree = GenerateTree(counter)
			}
		}
	}
}
