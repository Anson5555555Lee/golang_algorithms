package bits

func GetBit(num int, position uint) bool {
	return (num>>position)&1 == 1
}

func SetBit(num int, position uint) int {
	return num | 1<<position
}

func ClearBit(num int, position uint) int {
	mask := ^(1 << position)

	return num & mask
}

func UpdateBit(num int, position uint, bitValue int) int {
	// clear bit first
	temp := ClearBit(num, position)
	return temp | (bitValue << position)
}
