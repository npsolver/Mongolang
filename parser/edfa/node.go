package edfa

type Node struct {
	ID            int
	isTerminating bool
	Items         []*Item
	Bridges       map[string]*Node
	Visited       bool
}

func CreateNode(id int) *Node {
	return &Node{id, false, make([]*Item, 0), make(map[string]*Node), false}
}

func (n *Node) SetTerminating() {
	n.isTerminating = true
}

func (n *Node) IsTerminating() bool {
	return n.isTerminating
}

func (n *Node) AppendItem(item *Item) {
	n.Items = append(n.Items, item)
}
