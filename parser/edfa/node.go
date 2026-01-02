package edfa

type Node struct {
	id            int
	isTerminating bool
	items         []*Item
	bridges       map[string]*Node
}

func CreateNode(id int) *Node {
	return &Node{id, false, make([]*Item, 0), make(map[string]*Node)}
}

func (n *Node) SetTerminating() {
	n.isTerminating = true
}

func (n *Node) AppendItem(item *Item) {
	n.items = append(n.items, item)
}
