package idgenerator

import (
	"hash/crc32"
	"hash/fnv"
)

type GeneratorStruct struct {
	GenFunc func(title string) uint32
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

func (g *GeneratorStruct) GenerateID(title string) uint32 {
	return g.GenFunc(title)
}

func (g *GeneratorStruct) GetFunc() func(title string) uint32 {
	return g.GenFunc
}
