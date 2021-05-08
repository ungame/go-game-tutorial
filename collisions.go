package main

import (
	"fmt"
	"math"
)

// Referências
// - https://developer.mozilla.org/pt-BR/docs/Games/Techniques/2D_collision_detection
// - https://medium.com/davidpaniz/matematica-para-jogos-parte-1-colisao-de-circulos-e1119e030976
//

type Circle struct {
	center Vector
	radius float64
}

// Collides calcula a colisão entre dois circulos.
// Calculando a distância entre dois pontos:
// Considere o elemento A na posição (x: 10, y: 20)
// e o elemento B na posição (x: 20, y: 30).
// Aplicar a formula: Bx - Ax e By - Ay
// Resultado de Bx - Ax: (20 - 10) = 10
// Resultado de By - Ay: (30 - 20) = 10
// Formula para calcular a distância:
// √ (Bx - Ax)² + (By - Ay)²
// Resultado:
// √ (10² + 10²) = 14,14213562373095
// Para verificar se ocorreu uma colisão
// basta calcular se a distância entre os dois
// elementos é menor ou igual a soma do raio deles.
func Collides(c1, c2 Circle) bool {
	s := math.Pow(c2.center.x-c1.center.x, 2) + math.Pow(c2.center.y-c1.center.y, 2)
	dist := math.Sqrt(s)
	return dist <= c1.radius+c2.radius
}

func CheckCollisions(elements []*Element) error {
	for i := 0; i < len(elements) - 1; i++ {
		for j := i + 1; j < len(elements); j++ {
			before := elements[i]
			next := elements[j]

			for _, c1 := range before.collisions {
				for _, c2 := range next.collisions {

					if Collides(c1, c2) && before.active && next.active {
						fmt.Println("Collides")
						err := before.Collision(next)
						if err != nil {
							return err
						}
						err = next.Collision(before)
						if err != nil {
							return err
						}
					}

				}
			}

		}
	}

	return nil
}