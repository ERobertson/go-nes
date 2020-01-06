package hardware

type instruction struct {
	opcode      string
	cycles      uint8
	instruction func(cpu)
}
