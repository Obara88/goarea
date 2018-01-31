package goarea

import "math"

//Pi é uma proporção numerica definida pela relação entre o perimetro de um circunferencia e seu diametro
const Pi = 3.1416

//Circ blabla
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect blablabla
func Rect(base, altura float64) float64 {
	return base * altura
}

// _TrianguloEq blabla
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
