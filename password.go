package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Estructura de la clase Password
type Password struct {
	longitud   int
	contraseña string
}

// Constructor por defecto
func NewPasswordDefecto() *Password {
	return &Password{
		longitud:   8,
		contraseña: generarPassword(8),
	}
}

// Constructor con la longitud que se pasa como parámetro
func NewPassword(longitud int) *Password {
	return &Password{
		longitud:   longitud,
		contraseña: generarPassword(longitud),
	}
}

// Método para verificar si la contraseña es fuerte
func (p *Password) esFuerte() bool {
	if p.longitud < 8 {
		return false
	}

	mayusculas := 0
	minusculas := 0
	numeros := 0

	for _, char := range p.contraseña {
		if 'A' <= char && char <= 'Z' {
			mayusculas++
		} else if 'a' <= char && char <= 'z' {
			minusculas++
		} else if '0' <= char && char <= '9' {
			numeros++
		}
	}

	return mayusculas > 2 && minusculas > 1 && numeros > 5
}

// Método para generar una contraseña aleatoria
func generarPassword(longitud int) string {
	rand.Seed(time.Now().UnixNano())

	caracteres := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	var contraseña strings.Builder

	for i := 0; i < longitud; i++ {
		randomIndex := rand.Intn(len(caracteres))
		contraseña.WriteByte(caracteres[randomIndex])
	}

	return contraseña.String()
}

// Métodos get y set para longitud y contraseña
func (p *Password) GetLongitud() int {
	return p.longitud
}

func (p *Password) SetLongitud(longitud int) {
	p.longitud = longitud
	p.contraseña = generarPassword(longitud)
}

func (p *Password) GetContraseña() string {
	return p.contraseña
}

func (p *Password) SetContraseña(contraseña string) {
	p.contraseña = contraseña
}

func main() {
	var longitud, tamañoArray int

	// Solicitar la longitud de los Passwords
	fmt.Print("Ingrese la longitud de los Passwords: ")
	fmt.Scan(&longitud)

	// Solicitar el tamaño del array
	fmt.Print("Ingrese el tamaño del array de Passwords: ")
	fmt.Scan(&tamañoArray)

	// Crear un array de Passwords
	passwords := make([]*Password, tamañoArray)

	// Crear un array de booleanos para almacenar si el password es fuerte
	esFuerteArray := make([]bool, tamañoArray)

	// Llenar el array de Passwords
	for i := 0; i < tamañoArray; i++ {
		passwords[i] = NewPassword(longitud)
		esFuerteArray[i] = passwords[i].esFuerte()
	}

	// Mostrar contraseñas y si son fuertes o no
	for i := 0; i < tamañoArray; i++ {
		fmt.Printf("Contraseña%d: %s %v\n", i+1, passwords[i].GetContraseña(), esFuerteArray[i])
	}
}
