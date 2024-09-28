package idgenerator

import (
	"hash/crc32"
	"hash/fnv"
)

type Generator struct {
	genFunc func(title string) uint32
}

func FnvID(title string) uint32 {
	h := fnv.New32a()
	_, err := h.Write([]byte(title))
	if err != nil {
		return 0
	}
	return h.Sum32()
}

func CrcID(title string) uint32 {
	return crc32.ChecksumIEEE([]byte(title))
}

func (g *Generator) GenerateID(title string) uint32 {
	return g.genFunc(title)
}

func (g *Generator) GetFunc() func(title string) uint32 {
	return g.genFunc
}

func CreateFnvGen() Generator {
	return Generator{genFunc: FnvID}
}
func CreateCrcGen() Generator {
	return Generator{genFunc: CrcID}
}
