package main

import "fmt"

const day16NextPacketState int = 1
const day16InLiteralState int = 2
const day16InOperatorState int = 3
const day16InOperatorContentState int = 4
const day16NextPacketOperatorState int = 5

const day16OperatorRawBitCountType = 6
const day16OperatorChunkCountType = 7

func day16(input string) []int {
	fmt.Println("Running for input", input)
	ps := NewPacketStream(input)

	versions := make([]int, 0)

	states := make(intStack, 0)

	states.Push(day16NextPacketState)

	var operatorLength, operatorRead, operatorCountType, operatorChunkReadCount int
	for {
		state := states.Peek()
		if state == day16NextPacketState || state == day16NextPacketOperatorState {
			if state == day16NextPacketOperatorState {
				if operatorCountType == day16OperatorRawBitCountType &&
					operatorRead == operatorLength {
					states.Pop()
					break
				} else if operatorCountType == day16OperatorChunkCountType && operatorChunkReadCount == operatorLength {
					states.Pop()
					break
				}
			}

			packetVersion := ps.readBits(3)
			packetType := ps.readBits(3)

			operatorRead += 6

			if packetVersion == -1 || packetType == -1 {
				break
			}

			versions = append(versions, int(packetVersion))

			if packetType == 4 {
				states.Push(day16InLiteralState)
			} else {
				states.Push(day16InOperatorState)
			}
		} else if state == day16InLiteralState {
			// Read 5 bit chunks until end
			for {
				chunk := ps.readBits(5)
				operatorRead += 5
				isLast := (chunk >> 4 & 1) == 0

				if isLast {
					break
				}
			}

			operatorChunkReadCount++
			states.Pop()
		} else if state == day16InOperatorState {
			lengthTypeId := ps.readBits(1)

			if lengthTypeId == 0 {
				operatorLength = int(ps.readBits(15))
				operatorCountType = day16OperatorRawBitCountType
			} else {
				operatorLength = int(ps.readBits(11))
				operatorCountType = day16OperatorChunkCountType
			}

			operatorRead = 0

			states.Pop()
			states.Push(day16InOperatorContentState)
		} else {
			if operatorCountType == day16OperatorRawBitCountType && operatorRead == operatorLength {
				states.Pop()
			} else if operatorCountType == day16OperatorChunkCountType && operatorChunkReadCount == operatorLength {
				states.Pop()
			}
			states.Push(day16NextPacketOperatorState)
		}
	}

	return versions
}
