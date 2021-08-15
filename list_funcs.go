package one_direction_list

import "fmt"

//New - Cоздает новый список
func New(name string) OneDirectionList{
	return  OneDirectionList{
		Root: &Node{Name: name},
	}
}

//Add - Добавляет новый элемент в конец списка
func (odl OneDirectionList) Add(name string)  *Node {
	var (
		current *Node
	)

	// если список без корня то просто делаем первый элемент
	if odl.Root == nil {
		odl.Root = &Node{
			Name: name,
		}
		return odl.Root
	}
	//находи последний элемент и присоединяем новый элемент к нему
	current = odl.Root
	for {
		if current.Next == nil {
			current.Next = &Node{
				Name: name,
			}
			return current.Next
		}
		current = current.Next
	}
}

//Add - Добавляет новый элемент после указанного если он nil то в конец списка
func (odl OneDirectionList) AddAfter(name string, afterNode *Node)  *Node {

	newNode := &Node{Name: name}

	if afterNode == nil {
		return odl.Add(name)
	}

	newNode.Next = afterNode.Next
	afterNode.Next = newNode

	return newNode
}

//Print - Распечатывает весь список
func (odl OneDirectionList) Print(){
	var (
		current *Node
		length int64
	)
	current = odl.Root
	for {
		if current != nil{
			length++
			current.Print()
			current = current.Next
		}else{
			break
		}
	}
	fmt.Println("Длинна списка:",length)
}
