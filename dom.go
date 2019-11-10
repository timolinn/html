package html

// NodeType DOM element type
type NodeType int

const (
	// TextNode represents 	DOM text node
	TextNode NodeType = iota
	// HTMLElement represents 	DOM text node
	HTMLElement
)

// Node The DOM made up of a tree of
// Nodes
type Node struct {
	Children []Node
	NodeType NodeType
	Data     string
	TagName  string
}

// Text returns a DOM TextNode
func Text(data string) Node {
	return Node{
		Children: []Node{},
		NodeType: TextNode,
		Data:     data,
	}
}

// Element returns a DOM TextNode
func Element(tagName string, attrs AttrMap, children []Node) Node {
	return Node{
		TagName:  tagName,
		Children: children,
		NodeType: HTMLElement,
	}
}
