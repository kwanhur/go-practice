package combine

import "fmt"

type People struct {

}

type Teacher struct {
	People
}

// call here, did not know combine with who, so call ShowB only by self
func (p *People) ShowA()  {
	fmt.Println("show a")
	p.ShowB()
}

func (p *People) ShowB()  {
	fmt.Println("show b")
}


func (t *Teacher)ShowB()  {
	fmt.Println("teacher show b")
}