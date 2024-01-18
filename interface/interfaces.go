package main

type Shaper interface {
	Area() float32
}

type ILength interface {
	Length() int
}
type Square struct {
	side float32
}

func main() {
	var shap Shaper
	s := new(Square)
	shap = s
	if v, ok := shap.(*Square); ok {
		println(v.Length())
	}
}

func (this *Square) Area() float32 {
	return 1.1
}

func (this *Square) Length() int {
	return 1
}
