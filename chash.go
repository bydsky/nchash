package nchash

import (
	"encoding/binary"
	"errors"
	"sort"
	"strings"
)

type uint32s []uint32

func (h uint32s) Len() int           { return len(h) }
func (h uint32s) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h uint32s) Less(i, j int) bool { return h[i] < h[j] }

type Chash struct {
	keyMap     map[uint32]string
	sortedKeys []uint32
	nodeMap    map[string]int
	vnodeNum   int
}

func New(nodes []string) *Chash {
	nodeMap := make(map[string]int, len(nodes))
	for _, node := range nodes {
		nodeMap[node] = 1
	}

	return NewWithWeights(nodeMap, 160)
}

func NewWithWeights(nodeMap map[string]int, vnodeNum int) *Chash {
	chash := &Chash{
		keyMap:     make(map[uint32]string),
		sortedKeys: make([]uint32, 0),
		nodeMap:    nodeMap,
		vnodeNum:   vnodeNum,
	}
	chash.init()
	return chash
}

func (h *Chash) init() {
	for node, weight := range h.nodeMap {
		host := strings.Split(node, ":")
		crc := NewNiginxCrc()
		crc.Write([]byte(host[0]))
		crc.Write([]byte{0})
		crc.Write([]byte(host[1]))
		base := crc.crc
		var prev uint32
		prev = 0
		for j := 0; j < h.vnodeNum*weight; j++ {
			prevByte := make([]byte, 4)
			binary.LittleEndian.PutUint32(prevByte, prev)
			prev = Update(base, prevByte)
			key := uint32(prev)
			h.keyMap[key] = node
			h.sortedKeys = append(h.sortedKeys, key)
		}
	}

	sort.Sort(uint32s(h.sortedKeys))
}

func (h *Chash) hash(key string) uint32 {
	crc := NewNiginxCrc()
	crc.Write([]byte(key))
	return uint32(crc.Sum32())
}

func (h *Chash) Get(key string) (node string, err error) {
	if len(h.keyMap) == 0 {
		return "", errors.New("no node")
	}

	hash := h.hash(key)
	pos := sort.Search(len(h.sortedKeys), func(i int) bool { return h.sortedKeys[i] > hash })
	if pos == len(h.sortedKeys) {
		pos = 0
	}
	return h.keyMap[h.sortedKeys[pos]], nil
}
