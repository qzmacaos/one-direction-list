package one_direction_list

import "fmt"

//Print - Распечатывает содержимое элемента списка
func (node Node) Print()  {
	fmt.Println(node.Name)
}

//PrintRecoursive - Распечатывает все элементы списка начиная с указанного
func PrintRecoursive(node *Node){
	if node != nil{
		node.Print()
		PrintRecoursive(node.Next)
	}
}