package facade

import (
	"fmt"
)

const (
	BOOT_ADDRESS = 0
	BOOT_SECTOR  = 0
	SECTOR_SIZE  = 1023
)

type CPU struct{}

func (c *CPU) Freeze() {
	fmt.Println("CPU frozen")
}

func (c *CPU) Jump(position int) {
	fmt.Println("CPU jumps to", position)
}

func (c *CPU) Execute() {
	fmt.Println("CPU executes")

}

type Memory struct{}

func (this *Memory) Load(position int64, data []byte) {
	fmt.Println("Memory loads")
}

type HardDrive struct{}

func (hd *HardDrive) Read(lba int, size int) []byte {
	fmt.Println("HardDrive reads")
	return make([]byte, 0)
}

type Computer struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputer() *Computer {
	cpu := new(CPU)
	memory := new(Memory)
	hardDrive := new(HardDrive)
	return &Computer{cpu: cpu, memory: memory, hardDrive: hardDrive}

}

func (this *Computer) Start() {
	this.cpu.Freeze()
	this.memory.Load(BOOT_ADDRESS, this.hardDrive.Read(BOOT_SECTOR, SECTOR_SIZE))
	this.cpu.Jump(BOOT_ADDRESS)
	this.cpu.Execute()
}

func main() {
	computer := NewComputer()
	computer.Start()
}
