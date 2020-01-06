package hardware

type memory interface {
	load(address uint16) (value uint8)
	store(address uint16, value uint8)
}

type memoryContents [65536]uint8

func reset(memory mem) {

}
