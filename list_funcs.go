package one_direction_list

//New -
func New(name string) OneDirectionList{
	return  OneDirectionList{
		Root: &Node{Name: name},
	}
}

//Add -
func (odl OneDirectionList) Add(name string, afterNode *Node)  *Node{

	newNode := &Node{Name: name}

	if afterNode == nil {
		afterNode = odl.Root
	}

	newNode.Next = afterNode.Next
	afterNode.Next = newNode

	return newNode
}

//Print -
func (odl OneDirectionList) Print(){
	PrintRecoursive(odl.Root)
}
