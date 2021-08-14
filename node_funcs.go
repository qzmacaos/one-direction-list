package one_direction_list

import "fmt"

//Print -
func (node Node) Print()  {
	fmt.Println(node.Name)
}

//PrintRecoursive -
func PrintRecoursive(node *Node){
	if node != nil{
		node.Print()
		PrintRecoursive(node.Next)
	}
}