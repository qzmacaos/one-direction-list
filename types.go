package one_direction_list

//Node - структура для элемента списка
type Node struct {
	Name string
	Next *Node
}

//OneDirectionList - структура для объукта списка
type OneDirectionList struct {
	Root *Node
}
