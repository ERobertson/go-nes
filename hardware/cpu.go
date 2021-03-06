package hardware

//Reference: http://nesdev.com/6502.txt
type cpu struct {
	cpuStatus
}
type status uint8

//In sequence from bit 0-7
const (
	C cpuStatus = iota << 1 //carry flag
	Z                       //zero flag
	I                       //interrupt enable/disable
	D                       //decimal mode status flag
	B                       //software interrupt
	U                       //unused should always be high
	V                       //overflow flag
	S                       //sign flag
)

type registers struct {
	A  uint8
	X  uint8
	PC uint16
	SP uint8
}

//addressing modes
/*
 * Immediate
 * Absolute and zero-page absolute
 * Implied
 * Accumulator
 * Indexed and zero page indexed
 * Indirect
 * Pre-indexed indirect
 */
