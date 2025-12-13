package main

type Box struct {
	id      string
	x, y, z int64
}

type DisjointSetUnion struct {
	parent map[string]string
}

func NewDSU(ids []string) *DisjointSetUnion {
	parent := make(map[string]string)
	for _, id := range ids {
		parent[id] = id
	}
	return &DisjointSetUnion{parent: parent}
}

func (d *DisjointSetUnion) Find(id string) string {
	if d.parent[id] != id {
		d.parent[id] = d.Find(d.parent[id])
	}

	return d.parent[id]
}

func (d *DisjointSetUnion) Union(idA, idB string) {
	rootA := d.Find(idA)
	rootB := d.Find(idB)

	if rootA != rootB {
		d.parent[rootA] = rootB
	}
}

func (d *DisjointSetUnion) IsAllConnected(id string) bool {
	currNode := d.Find(id)
	for _, id := range d.parent {
		if d.Find(id) != currNode {
			return false
		}
	}
	return true
}