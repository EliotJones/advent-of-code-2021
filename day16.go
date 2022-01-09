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
	typeId          int
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

func NewRootPacket() *day16Packet {
	return &day16Packet{
		packetType: day16RootType,
		bitLength:  0,
		children:   make([]*day16Packet, 0),
	}
}

func NewOperatorPacket(version int, length int, lengthType int, lengthFieldBitLength int, typeId int) *day16Packet {
	return &day16Packet{
		packetType:      day16OperatorType,
		children:        make([]*day16Packet, 0),
		bitLength:       0,
		headerBitLength: 6 + lengthFieldBitLength,
		length:          length,
		lengthType:      lengthType,
		version:         version,
		typeId:          typeId,
	}
}

func NewLiteralPacket(version int) *day16Packet {
	return &day16Packet{
		packetType: day16LiteralType,
		bitLength:  6,
		version:    version,
	}
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

/*
   Packets with type ID 0 are sum packets - their value is the sum of the values of their sub-packets. If they only have a single sub-packet, their value is the value of the sub-packet.
   Packets with type ID 1 are product packets - their value is the result of multiplying together the values of their sub-packets. If they only have a single sub-packet, their value is the value of the sub-packet.
   Packets with type ID 2 are minimum packets - their value is the minimum of the values of their sub-packets.
   Packets with type ID 3 are maximum packets - their value is the maximum of the values of their sub-packets.
   Packets with type ID 5 are greater than packets - their value is 1 if the value of the first sub-packet is greater than the value of the second sub-packet; otherwise, their value is 0. These packets always have exactly two sub-packets.
   Packets with type ID 6 are less than packets - their value is 1 if the value of the first sub-packet is less than the value of the second sub-packet; otherwise, their value is 0. These packets always have exactly two sub-packets.
   Packets with type ID 7 are equal to packets - their value is 1 if the value of the first sub-packet is equal to the value of the second sub-packet; otherwise, their value is 0. These packets always have exactly two sub-packets.
*/
func interpretOperation(packet *day16Packet) int {
	if packet.packetType == day16LiteralType {
		return packet.literalValue
	}

	if packet.packetType == day16RootType {
		if len(packet.children) > 1 {
			panic("multiple root children?")
		}

		return interpretOperation(packet.children[0])
	}

	var result int
	// Sum type
	if packet.typeId == 0 {
		for _, child := range packet.children {
			result += interpretOperation(child)
		}

		return result
	}

	if len(packet.children) == 0 {
		return 0
	}

	// Product type
	if packet.typeId == 1 {
		result = 1
		for _, child := range packet.children {
			result *= interpretOperation(child)
		}

		return result
	}

	// Min and max
	if packet.typeId == 2 || packet.typeId == 3 {
		result = interpretOperation(packet.children[0])

		for i, child := range packet.children {
			if i == 0 {
				continue
			}

			val := interpretOperation(child)
			if packet.typeId == 2 && val < result {
				result = val
			} else if packet.typeId == 3 && val > result {
				result = val
			}
		}

		return result
	}

	// Bools
	if packet.typeId >= 5 && packet.typeId <= 7 {
		val1 := interpretOperation(packet.children[0])
		val2 := interpretOperation(packet.children[1])

		if packet.typeId == 5 && val1 > val2 {
			return 1
		} else if packet.typeId == 5 {
			return 0
		}

		if packet.typeId == 6 && val1 < val2 {
			return 1
		} else if packet.typeId == 6 {
			return 0
		}

		if val1 == val2 {
			return 1
		}

		return 0
	}

	return 0
}

func day16(input string) ([]int, *day16Packet) {
	fmt.Println("Running for input", input)
	ps := NewPacketStream(input)

	states := make(genericStack, 0)

	states.Push(NewRootPacket())

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
				states.Push(NewLiteralPacket(int(packetVersion)))
			} else {
				lengthType, length, bits := readOperatorLengthTypeAndValueAndBitsRead(ps)

				if lengthType == -1 || length == -1 {
					break
				}

				states.Push(NewOperatorPacket(int(packetVersion), length, lengthType, bits, int(packetType)))
			}

			continue
		}

		if packet.packetType == day16LiteralType {
			var val int

			// Read 5 bit chunks until end, signalled by first bit being 0.
			for {
				chunk := ps.readBits(5)

				if chunk == -1 {
					// If encountering end of stream in some content then exit removing this incomplete node.
					endOfStream = true
					break
				}

				val = (val << 4) + (int(chunk) & 0b1111)

				isLast := (chunk >> 4 & 1) == 0
				packet.bitLength += 5

				if isLast {
					packet.literalValue = val
					break
				}
			}

			if endOfStream {
				// Remove the incomplete node prior to exiting.
				states.Pop()
				break
			}

			popCurrentAndUpdateParent(packet, &states)
		}
	}

	root := states.Peek().(*day16Packet)

	versions := make([]int, 0)
	recursivelyGatherVersions(root, &versions)

	return versions, root
}

func day16p1() {
	scanner, err := scannerForFile("inputs/day16.txt")
	if err != nil {
		panic(err)
	}

	scanner.Scan()
	line := scanner.Text()

	versions, root := day16(line)

	result := intSliceSum(versions)
	part2Result := interpretOperation(root)

	fmt.Println("Result p1 is", result, "p2 is", part2Result)
}

func day16p2Input(input string) int {
	_, root := day16(input)

	result := interpretOperation(root)

	return result
}
