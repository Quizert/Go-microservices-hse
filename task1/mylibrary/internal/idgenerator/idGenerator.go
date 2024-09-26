package idgenerator

import (
	"hash/crc32"
	"hash/fnv"
)

type Generator func(title string) uint32

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
