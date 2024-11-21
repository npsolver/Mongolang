package global

type Symbol struct {
	name string
}

type TerminalSymbol struct {
	Symbol
}

type NonTerminalSymbol struct {
	Symbol
}

func NewTerminalSymbol(name string) TerminalSymbol       { return TerminalSymbol{Symbol{name}} }
func NewNonTerminalSymbol(name string) NonTerminalSymbol { return NonTerminalSymbol{Symbol{name}} }

func (s *Symbol) GetName() string { return s.name }

func (t *TerminalSymbol) IsTerminal() bool { return true }

func (t *NonTerminalSymbol) IsTerminal() bool { return false }
