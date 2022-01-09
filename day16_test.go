package main

import (
	"testing"
)

func intSlicesEqual(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func TestDay16One(t *testing.T) {
	versions, _ := day16("D2FE28")
	if !intSlicesEqual(versions, []int{6}) {
		t.Errorf("Invalid versions: %x", versions)
	}
}

func TestDay16Two(t *testing.T) {
	versions, _ := day16("38006F45291200")
	if !intSlicesEqual(versions, []int{1, 6, 2}) {
		t.Errorf("Invalid versions: %x", versions)
	}
}

func TestDay16Three(t *testing.T) {
	versions, _ := day16("EE00D40C823060")
	if !intSlicesEqual(versions, []int{7, 2, 4, 1}) {
		t.Errorf("Invalid versions: %x", versions)
	}
}

func TestDay16Four(t *testing.T) {
	versions, _ := day16("8A004A801A8002F478")
	result := intSliceSum(versions)
	if result != 16 {
		t.Errorf("Incorrect version sum %d from %x", result, versions)
	}
}

func TestDay16Five(t *testing.T) {
	versions, _ := day16("620080001611562C8802118E34")
	result := intSliceSum(versions)
	if result != 12 {
		t.Errorf("Incorrect version sum %d from %x", result, versions)
	}
}

func TestDay16Six(t *testing.T) {
	versions, _ := day16("C0015000016115A2E0802F182340")
	result := intSliceSum(versions)
	if result != 23 {
		t.Errorf("Incorrect version sum %d from %x", result, versions)
	}
}

func TestDay16Seven(t *testing.T) {
	versions, _ := day16("A0016C880162017C3686B18A3D4780")
	result := intSliceSum(versions)
	if result != 31 {
		t.Errorf("Incorrect version sum %d from %x", result, versions)
	}
}

func assertIntEqual(t *testing.T, expected int, actual int) {
	if actual != expected {
		t.Errorf("Result was incorrect, expected %d but got %d", expected, actual)
	}
}

func TestDay16PartTwoOne(t *testing.T) {
	result := day16p2Input("C200B40A82")
	assertIntEqual(t, 3, result)
}

func TestDay16PartTwoTwo(t *testing.T) {
	result := day16p2Input("04005AC33890")
	assertIntEqual(t, 54, result)
}

func TestDay16PartTwoThree(t *testing.T) {
	result := day16p2Input("880086C3E88112")
	assertIntEqual(t, 7, result)
}

func TestDay16PartTwoFour(t *testing.T) {
	result := day16p2Input("CE00C43D881120")
	assertIntEqual(t, 9, result)
}

func TestDay16PartTwoFive(t *testing.T) {
	result := day16p2Input("D8005AC2A8F0")
	assertIntEqual(t, 1, result)
}

func TestDay16PartTwoSix(t *testing.T) {
	result := day16p2Input("F600BC2D8F")
	assertIntEqual(t, 0, result)
}

func TestDay16PartTwoSeven(t *testing.T) {
	result := day16p2Input("9C005AC2F8F0")
	assertIntEqual(t, 0, result)
}

func TestDay16PartTwoEight(t *testing.T) {
	result := day16p2Input("9C0141080250320F1802104A08")
	assertIntEqual(t, 1, result)
}
