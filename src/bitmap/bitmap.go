package bitmap

type BitMap struct {
	runes []rune	// 1个rune是32位整型来的，也就是4个字节
	nbits int
}

func initialBitMap(nbits int) *BitMap {
	bitMap := new(BitMap)
	bitMap.nbits = nbits
	bitMap.runes = make([]rune, nbits / 32 + 1)
	return bitMap
}

func (bitMap *BitMap) Set(k int) {
	if k > bitMap.nbits {
		return
	}
	runeIndex := k / 32
	bitIndex := k % 32
	bitMap.runes[runeIndex] |= (1 << bitIndex)
}

func (bitMap *BitMap) Get(k int) bool {
	if k > bitMap.nbits {
		return false
	}
	runeIndex := k / 32
	bitIndex := k % 32
	return (bitMap.runes[runeIndex] & (1 << bitIndex)) != 0
}


