package cal
import (
	"errors"
)
// cal.go
// Operações matemáticas simples
func Add(a, b int) (int, error) {
	return a + b, nil
}	

func Sub(a, b int) (int, error) {
	return a - b, nil
}	

func Mul(a, b int) (int, error) {
	return a * b, nil
}	

func Div(a, b int) (int, error) {
	if (b==0) {
		return-1, errors.New("Divisão por zero")
	}
	return a / b, nil
}	
