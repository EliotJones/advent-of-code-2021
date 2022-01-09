package main

type packetStream struct {
	bitOffset int
	hexData   string
}

func (v *packetStream) readBits(count int) int32 {
	var result int32
	if v.bitOffset+count > len(v.hexData)*4 {
		return -1
	}

	for i := 0; i < count; i++ {
		bitAt := v.getPacketStreamBitAt(v.bitOffset + i)
		if bitAt {
			result += (1 << (count - 1 - i))
		}
	}

	v.bitOffset += count

	return result
}

func (v *packetStream) getPacketStreamBitAt(index int) bool {
	offset := index / 4
	hexNibble := v.hexData[offset]

	var rawNibble byte
	if hexNibble > 57 {
		rawNibble = hexNibble - 55
	} else {
		rawNibble = hexNibble - asciiNumToByteAdjustment
	}

	withinNibbleIndex := byte(index % 4)

	resultBit := rawNibble >> (3 - withinNibbleIndex)

	return (resultBit & 1) == 1
}

func NewPacketStream(hex string) *packetStream {
	return &packetStream{
		hexData:   hex,
		bitOffset: 0,
	}
}
