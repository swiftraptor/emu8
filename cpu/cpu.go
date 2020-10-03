package cpu

type Cpu struct {
	IndexRegister uint8
	ProgramCounter uint8
	Memory [4096]uint8
	Stack [16]uint16
	StackPointer uint8
	DelayTimer uint8
}
