package figuras

type Cuadrado struct{
	Lado float32 
	
}

func (cuadrado *Cuadrado) Area() float32{
	return cuadrado.Lado*cuadrado.Lado
}

func (c *Cuadrado) Perimetro()float32{
	return 4*c.Lado
}