//堆栈
package Lstruct

import "errors"

type Stack []interface{}

func (s *Stack) Push (x interface{}){
	*s = append(*s, x)
}

func (s *Stack) Pop() (interface{}, error){
	temp := *s
	if len(temp) == 0{
		return nil, errors.New("Can't pop an empty stack")
	}
	x := temp[len(temp) - 1]
	*s = temp[:len(temp) - 1]
	return x, nil
}

func (s Stack) Top() (interface{}, error){
	if len(s) == 0 {
		return nil, errors.New("Can't top en empty stack")
	}
	return s[len(s) - 1], nil
}

func (s *Stack) Len()int{
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return (len(*s) == 0)
}