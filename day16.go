package main

import "fmt"

type day16Packet struct {
	packetType      int
	children        []*day16Packet
	bitLength       int
	literalValue    int
	length          int
	lengthType      int
	headerBitLength int
	version         int
}

const day16RootType = 0
const day16LiteralType = 1
const day16OperatorType = 2

const day16OperatorRawBitCountType = 6
const day16OperatorChunkCountType = 7

func intSliceSum(s []int) int {
	var result int
	for _, v := range s {
		result += v
	}

	return result
}

func readOperatorLengthTypeAndValueAndBitsRead(ps *packetStream) (int, int, int) {
	lengthTypeId := ps.readBits(1)

	if lengthTypeId == 0 {
		operatorLength := int(ps.readBits(15))
		return day16OperatorRawBitCountType, operatorLength, 16
	} else {
		operatorLength := int(ps.readBits(11))
		return day16OperatorChunkCountType, operatorLength, 12
	}
}

func popCurrentAndUpdateParent(current *day16Packet, states *genericStack) {
	(*states).Pop()

	if current.packetType == day16OperatorType {
		current.bitLength += current.headerBitLength
	}

	parent := states.Peek().(*day16Packet)

	parent.children = append(parent.children, current)
	parent.bitLength += current.bitLength
}

func isOperator(i interface{}) bool {
	return i != nil && i.(*day16Packet).packetType == day16OperatorType
}

func recursivelyGatherVersions(packet *day16Packet, versions *[]int) {
	if packet.packetType != day16RootType {
		*versions = append(*versions, packet.version)
	}

	if packet.children != nil {
		for _, child := range packet.children {
			recursivelyGatherVersions(child, versions)
		}
	}
}

func day16(input string) []int {
	fmt.Println("Running for input", input)
	ps := NewPacketStream(input)

	states := make(genericStack, 0)

	states.Push(&day16Packet{
		packetType: day16RootType,
		bitLength:  0,
		children:   make([]*day16Packet, 0),
	})

	endOfStream := false
	for {
		packet := states.Peek().(*day16Packet)

		if packet.packetType == day16OperatorType {
			if packet.lengthType == day16OperatorChunkCountType && len(packet.children) == packet.length {
				popCurrentAndUpdateParent(packet, &states)
				continue
			} else if packet.lengthType == day16OperatorRawBitCountType && packet.bitLength >= packet.length {
				popCurrentAndUpdateParent(packet, &states)
				continue
			}
		}

		if packet.packetType == day16RootType || packet.packetType == day16OperatorType {
			packetVersion := ps.readBits(3)
			packetType := ps.readBits(3)

			if packetVersion == -1 || packetType == -1 {
				endOfStream = true

				if packet.packetType == day16OperatorType {
					// Remove incomplete operator.
					states.Pop()
				}

				break
			}

			if packetType == 4 {
				states.Push(&day16Packet{
					packetType: day16LiteralType,
					bitLength:  6,
					length:     -1,
					lengthType: -1,
					version:    int(packetVersion),
				})
			} else {
				lengthType, length, bits := readOperatorLengthTypeAndValueAndBitsRead(ps)

				if lengthType == -1 || length == -1 {
					break
				}

				states.Push(&day16Packet{
					packetType:      day16OperatorType,
					children:        make([]*day16Packet, 0),
					bitLength:       0,
					headerBitLength: 6 + bits,
					length:          length,
					lengthType:      lengthType,
					version:         int(packetVersion),
				})
			}

			continue
		}

		if packet.packetType == day16LiteralType {
			// Read 5 bit chunks until end
			for {
				chunk := ps.readBits(5)

				if chunk == -1 {
					endOfStream = true
					break
				}

				isLast := (chunk >> 4 & 1) == 0
				packet.bitLength += 5

				if isLast {
					break
				}
			}

			if endOfStream {
				states.Pop()
				break
			}

			popCurrentAndUpdateParent(packet, &states)
		}
	}

	versions := make([]int, 0)
	recursivelyGatherVersions(states.Peek().(*day16Packet), &versions)

	return versions
}

func day16p1() {
	scanner, err := scannerForFile("inputs/day16.txt")
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	line := scanner.Text()

	versions := day16(line)

	result := intSliceSum(versions)

	fmt.Println("Result is", result)
}
