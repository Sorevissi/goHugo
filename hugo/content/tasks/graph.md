
---
menu:
    after:
        name: graph
        weight: 1
title: Построение графа
---

{{< mermaid >}}
graph TD
  Node1([ellipse Node1])
  Node1 --> Node3
  Node1 --> Node4
  Node1 --> Node5
  Node2{rhombus Node2}
  Node2 --> Node4
  Node2 --> Node5
  Node2 --> Node7
  Node2 --> Node9
  Node3{rhombus Node3}
  Node3 --> Node1
  Node3 --> Node2
  Node3 --> Node5
  Node3 --> Node6
  Node3 --> Node8
  Node3 --> Node10
  Node4[rect Node4]
  Node4 --> Node1
  Node4 --> Node2
  Node4 --> Node3
  Node4 --> Node5
  Node4 --> Node7
  Node4 --> Node8
  Node4 --> Node9
  Node4 --> Node10
  Node5[rect Node5]
  Node5 --> Node3
  Node5 --> Node7
  Node5 --> Node9
  Node6((circle Node6))
  Node6 --> Node1
  Node6 --> Node3
  Node6 --> Node4
  Node6 --> Node5
  Node6 --> Node7
  Node6 --> Node8
  Node7[rect Node7]
  Node7 --> Node1
  Node7 --> Node4
  Node7 --> Node9
  Node8(round-rect Node8)
  Node8 --> Node1
  Node8 --> Node4
  Node8 --> Node5
  Node8 --> Node9
  Node8 --> Node10
  Node9([ellipse Node9])
  Node9 --> Node2
  Node9 --> Node4
  Node9 --> Node6
  Node9 --> Node7
  Node9 --> Node8
  Node10{rhombus Node10}
  Node10 --> Node3
  Node10 --> Node4
  Node10 --> Node5
  Node10 --> Node6
  Node10 --> Node7
  Node10 --> Node9

{{< /mermaid >}}
