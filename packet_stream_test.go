package main

import "testing"

func TestPacketStreamOne(t *testing.T) {
	ps := NewPacketStream("D2FE28")
	first3 := ps.readBits(3)
	next3 := ps.readBits(3)
	next5 := ps.readBits(5)
	secondNext5 := ps.readBits(5)
	last5 := ps.readBits(5)
	empty3 := ps.readBits(3)

	end := ps.readBits(1)

	if first3 != 6 {
		t.Errorf("Failed to read correct first 3 bits, first 3 %x", first3)
	} else if next3 != 4 {
		t.Errorf("Failed to read correct next 3 bits, next %x", next3)
	} else if next5 != 0b10111 {
		t.Error("Failed to read next 5 bits")
	} else if secondNext5 != 0b11110 {
		t.Errorf("Failed to read middle 5 bits")
	} else if last5 != 0b00101 {
		t.Errorf("Failed to read last 5 bits")
	} else if empty3 != 0 {
		t.Errorf("Last 3 was not empty")
	} else if end != -1 {
		t.Errorf("Did not return end marker (-1) correctly")
	}
}

func TestPacketStreamTwo(t *testing.T) {
	ps := NewPacketStream("38006F45291200")

	version := ps.readBits(3)
	packetType := ps.readBits(3)
	lengthTypeId := ps.readBits(1)
	length := ps.readBits(15)
	firstSubPacket := ps.readBits(11)
	secondSubPacket := ps.readBits(16)

	if version != 1 {
		t.Error("Version not correct")
	} else if packetType != 6 {
		t.Error("Packet type incorrect")
	} else if lengthTypeId != 0 {
		t.Error("Length type id incorrect")
	} else if length != 27 {
		t.Error("Length incorrect")
	} else if firstSubPacket != 1674 {
		t.Error("1st sub packet incorrect")
	} else if secondSubPacket != 21028 {
		t.Error("2nd sub packet incorrect")
	}
}
