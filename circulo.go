package figuras

import "math"


type Circulo struct{
	Radio float32 
	
}

func (c *Circulo) Area() float32{
	return math.Pi*(c.Radio*c.Radio)
}

func (c *Circulo) Perimetro()float32{
	return 2*math.Pi*c.Radio
}