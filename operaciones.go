package figuras

import "fmt"

type Figura interface{
	Area() float32
	Perimetro() float32
}

func CalcularArea(figura Figura){
	fmt.Println( figura.Area());
}
   
func CalcularPerimetro(figura Figura){
	   fmt.Println(figura.Perimetro())
}