package one_direction_list

type Node struct {
	Name string
	Next *Node
}

type OneDirectionList struct {
	Root *Node
}
