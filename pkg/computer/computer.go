package computer

type Instruction struct {
	Operation interface{}
	Value     interface{}
}

type Computer struct {
	Pc, Acc           int
	Instructions      []Instruction
	ChangedOperations map[int]interface{}
	ChangedValues     map[int]interface{}
	Executed          map[interface{}]bool
	Operations        map[interface{}]func(*Computer, *Instruction)
	Status            map[interface{}]interface{}
}

func NewComputer(instructions []Instruction, operations map[interface{}]func(*Computer, *Instruction)) Computer {
	return Computer{
		Pc:                0,
		Acc:               0,
		Instructions:      instructions,
		ChangedOperations: make(map[int]interface{}),
		ChangedValues:     make(map[int]interface{}),
		Executed:          make(map[interface{}]bool),
		Operations:        operations,
		Status:            make(map[interface{}]interface{}),
	}
}

func (c *Computer) HasExecuted() bool {
	return c.Executed[c.Pc]
}

func (c *Computer) ChangeInstructionOperation(i int, op interface{}) {
	c.ChangedOperations[i] = op
	c.ChangedOperations[i], c.Instructions[i].Operation = c.Instructions[i].Operation, c.ChangedOperations[i]
}

func (c *Computer) ChangeInstructionValue(i int, val interface{}) {
	c.ChangedValues[i] = val
	c.ChangedValues[i], c.Instructions[i].Value = c.Instructions[i].Value, c.ChangedValues[i]
}

func (c *Computer) Reset(resetMemory, resetStatus bool) {
	c.Pc = 0
	c.Acc = 0
	c.Executed = make(map[interface{}]bool)
	if resetMemory {
		for i, v := range c.ChangedOperations {
			c.ChangeInstructionOperation(i, v)
		}
		c.ChangedOperations = make(map[int]interface{})
		for i, v := range c.ChangedValues {
			c.ChangeInstructionValue(i, v)
		}
		c.ChangedValues = make(map[int]interface{})
	}
	if resetStatus {
		c.Status = make(map[interface{}]interface{})
	}
}

func (c *Computer) Solve(requirement func(*Computer) bool) int {
	for !requirement(c) {
		c.Executed[c.Pc] = true
		instruction := c.Instructions[c.Pc]
		c.Operations[instruction.Operation](c, &instruction)
	}

	return c.Acc
}
