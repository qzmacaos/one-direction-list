package one_direction_list

//New - Cоздает новый список
func New(name string) OneDirectionList{
	return  OneDirectionList{
		Root: &Node{Name: name},
	}
}

//Add - Добавляет новый элемент после указанного(если не указан то после первого)
func (odl OneDirectionList) Add(name string, afterNode *Node)  *Node{

	newNode := &Node{Name: name}

	if afterNode == nil {
		afterNode = odl.Root
	}

	newNode.Next = afterNode.Next
	afterNode.Next = newNode

	return newNode
}

//Print - Распечатывает весь список
func (odl OneDirectionList) Print(){
	PrintRecoursive(odl.Root)
}
