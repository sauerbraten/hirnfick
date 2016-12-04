package machine

type Machine struct {
	cells map[int]byte
	addr  int
}

func New() *Machine {
	return &Machine{
		cells: map[int]byte{},
	}
}

func (m *Machine) IncrAddr() { m.addr++ }

func (m *Machine) DecrAddr() { m.addr-- }

func (m *Machine) IncrByte() { m.cells[m.addr]++ }

func (m *Machine) DecrByte() { m.cells[m.addr]-- }

func (m *Machine) GetByte() byte { return m.cells[m.addr] }

func (m *Machine) PutByte(b byte) { m.cells[m.addr] = b }
