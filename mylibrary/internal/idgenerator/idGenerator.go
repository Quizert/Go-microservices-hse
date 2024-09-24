package idgenerator

import "hash/fnv"

func FnvID(title string) uint32 {
	h := fnv.New32a()
	_, err := h.Write([]byte(title))
	if err != nil {
		return 0
	}
	return h.Sum32()
}
