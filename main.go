package main

import (
	"fmt"
	"hash/crc32"
	"sort"
)

type HashRing struct {
	nodes               []HashNode
	virtualNodesPerNode int
}

type HashNode struct {
	hash   int
	server string
}

func CreateHashRing(vnodes int) *HashRing {
	return &HashRing{
		virtualNodesPerNode: vnodes,
	}
}

func (h *HashRing) AddServer(server string) {

	for i := range h.virtualNodesPerNode {

		vnode := server + fmt.Sprintf("%d", i)
		creatHash := int(crc32.ChecksumIEEE([]byte(vnode)))

		h.nodes = append(h.nodes, HashNode{
			hash:   creatHash,
			server: server,
		})
	}

	sort.Slice(h.nodes, func(i, j int) bool {
		return h.nodes[i].hash < h.nodes[j].hash
	})
}

func (h *HashRing) RemoveServer(server string) {

	var tempNode []HashNode

	for _, node := range h.nodes {

		if node.server != server {
			tempNode = append(tempNode, node)
		}

	}

	h.nodes = tempNode
}

func (h *HashRing) GetServer(key string) string {

	hash := int(crc32.ChecksumIEEE([]byte(key)))

	for _, node := range h.nodes {
		if node.hash > hash {
			return node.server
		}
	}

	return h.nodes[0].server
}

func main() {

	virtualNodes := 10
	hashRing := CreateHashRing(virtualNodes)

	hashRing.AddServer("server1")
	hashRing.AddServer("server2")
	hashRing.AddServer("server3")

	// Initialize the Hash Ring
	str := hashRing.GetServer("test1113333224")

	fmt.Println(str)

}
