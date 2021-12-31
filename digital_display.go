package main

import "math"

func stringToMap(str string) map[byte]struct{} {
	var empty struct{}
	result := make(map[byte]struct{})
	for i := 0; i < len(str); i++ {
		result[str[i]] = empty
	}

	return result
}

func byteSliceToMap(slice []byte) map[byte]struct{} {
	var empty struct{}
	result := make(map[byte]struct{})
	for i := 0; i < len(slice); i++ {
		result[slice[i]] = empty
	}

	return result
}

func except(set1 map[byte]struct{}, set2 map[byte]struct{}) []byte {
	var result []byte
	for k := range set1 {
		if !contains(set2, k) {
			result = append(result, k)
		}
	}

	return result
}

func containsAll(set1 map[byte]struct{}, set2 map[byte]struct{}) bool {
	for k := range set2 {
		if !contains(set1, k) {
			return false
		}
	}

	return true
}

func contains(set1 map[byte]struct{}, val byte) bool {
	if _, ok := set1[val]; ok {
		return true
	} else {
		return false
	}
}

func deduceTopRightAndSixIndex(sixLengths []map[byte]struct{}, one map[byte]struct{}) (byte, int) {
	for i, sixLength := range sixLengths {
		// Six is the only 6 length not including one.
		diff := except(one, sixLength)
		if len(diff) == 1 {
			return diff[0], i
		}
	}

	return 0, 0
}

func deduceBottomRight(topRight byte, one map[byte]struct{}) byte {
	for k := range one {
		if k != topRight {
			return k
		}
	}

	return 0
}

func deduceBottomLeft(two map[byte]struct{}, five map[byte]struct{}, topRight byte) byte {
	diff := except(two, five)
	for _, b := range diff {
		if b != topRight {
			return b
		}
	}

	return 0
}

func deduceTop(four map[byte]struct{}, eight map[byte]struct{}, bottomLeft byte, bottom byte) byte {
	findTopDiff := except(eight, four)
	for i := 0; i < len(findTopDiff); i++ {
		val := findTopDiff[i]
		if val != bottomLeft && val != bottom {
			return val
		}
	}

	return 0
}

func deduceBottom(four map[byte]struct{}, seven map[byte]struct{}, eight map[byte]struct{}, bottomLeft byte) byte {
	bottomAndBottomLeft := except(byteSliceToMap(except(eight, seven)), four)
	for i := 0; i < len(bottomAndBottomLeft); i++ {
		val := bottomAndBottomLeft[i]
		if val != bottomLeft {
			return val
		}
	}

	return 0
}

func mapDisplayOutputToInt(positionMap map[byte]int, output []string, displayNumbersEncoded [10]byte) int {
	var result int
	for index, str := range output {
		var val byte
		for i := 0; i < len(str); i++ {
			b := str[i]
			shift := positionMap[b]
			val += 1 << (6 - shift)
		}

		for i := 0; i < len(displayNumbersEncoded); i++ {
			num := displayNumbersEncoded[i]
			if val == num {
				result += i * (int(math.Pow10(3 - index)))
			}
		}
	}

	return result
}
