package main

import (
    "fmt"
)

type Vertex struct {
    visited    bool
    value      string
    neighbours []*Vertex
}

func (v *Vertex) connect(vertex *Vertex) {
    v.neighbours = append(v.neighbours, vertex)
}

type Graph struct{}

func (g *Graph) dfs(vertex *Vertex) {

    if vertex.visited != true {
        vertex.visited = true
        fmt.Println(vertex)
        if len(vertex.neighbours) != 0 {
            for _, v := range vertex.neighbours {
                g.dfs(v)
            }
        } else {
            return
        }
    }

}

func (g *Graph) disconnected(vertices ...*Vertex){
   for _, v := range vertices{
      g.dfs(v)
   }
}

func main() {
    v1 := Vertex{false, "A", make([]*Vertex, 0, 5)}
    v2 := Vertex{false, "B", make([]*Vertex, 0, 5)}
    v3 := Vertex{false, "C", make([]*Vertex, 0, 5)}
    v4 := Vertex{false, "D", make([]*Vertex, 0, 5)}
    v5 := Vertex{false, "E", make([]*Vertex, 0, 5)}
    g := Graph{}
    v1.connect(&v2)
    v2.connect(&v4)
    v2.connect(&v5)
    v3.connect(&v4)
    v3.connect(&v5)
    g.dfs(&v1)
}