package one_direction_list

import "fmt"

//Print - распечатывает содержимое элемента списка
func (node Node) Print()  {
	fmt.Println(node.Name)
}
