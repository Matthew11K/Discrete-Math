package main

import (
	"fmt"
	"math/big"
)

type Rational struct {
	Numerator, Denominator *big.Int
}

func NewRational(num, den int64) Rational {
	n := big.NewInt(num)
	d := big.NewInt(den)
	if d.Sign() == 0 {
		panic("0")
	}
	r := Rational{Numerator: n, Denominator: d}
	r.normalize()
	return r
}

func (r *Rational) normalize() {
	if r.Denominator.Sign() == 0 {
		panic("0")
	}
	if r.Numerator.Sign() == 0 {
		r.Denominator.SetInt64(1)
		return
	}
	gcd := new(big.Int).GCD(nil, nil, r.Numerator, r.Denominator)
	r.Numerator.Div(r.Numerator, gcd)
	r.Denominator.Div(r.Denominator, gcd)
	if r.Denominator.Sign() < 0 {
		r.Numerator.Neg(r.Numerator)
		r.Denominator.Neg(r.Denominator)
	}
}

func (r Rational) add(other Rational) Rational {
	numerator := new(big.Int).Mul(r.Numerator, other.Denominator)
	temp := new(big.Int).Mul(other.Numerator, r.Denominator)
	numerator.Add(numerator, temp)
	denominator := new(big.Int).Mul(r.Denominator, other.Denominator)
	return NewRational(numerator.Int64(), denominator.Int64())
}

func (r Rational) sub(other Rational) Rational {
	numerator := new(big.Int).Mul(r.Numerator, other.Denominator)
	temp := new(big.Int).Mul(other.Numerator, r.Denominator)
	numerator.Sub(numerator, temp)
	denominator := new(big.Int).Mul(r.Denominator, other.Denominator)
	return NewRational(numerator.Int64(), denominator.Int64())
}

func (r Rational) mul(other Rational) Rational {
	numerator := new(big.Int).Mul(r.Numerator, other.Numerator)
	denominator := new(big.Int).Mul(r.Denominator, other.Denominator)
	return NewRational(numerator.Int64(), denominator.Int64())
}

func (r Rational) div(other Rational) Rational {
	if other.Numerator.Sign() == 0 {
		panic("Деление на 0")
	}
	numerator := new(big.Int).Mul(r.Numerator, other.Denominator)
	denominator := new(big.Int).Mul(r.Denominator, other.Numerator)
	return NewRational(numerator.Int64(), denominator.Int64())
}

func gaussElimination(matrix [][]Rational) ([]Rational, error) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		if matrix[i][i].Numerator.Sign() == 0 {
			found := false
			for k := i + 1; k < n; k++ {
				if matrix[k][i].Numerator.Sign() != 0 {
					matrix[i], matrix[k] = matrix[k], matrix[i]
					found = true
					break
				}
			}
			if !found {
				return nil, fmt.Errorf("No solution")
			}
		}

		pivot := matrix[i][i]
		for j := i; j <= n; j++ {
			matrix[i][j] = matrix[i][j].div(pivot)
		}

		for k := i + 1; k < n; k++ {
			factor := matrix[k][i]
			for j := i; j <= n; j++ {
				matrix[k][j] = matrix[k][j].sub(matrix[i][j].mul(factor))
			}
		}
	}

	solution := make([]Rational, n)
	for i := n - 1; i >= 0; i-- {
		solution[i] = matrix[i][n]
		for j := i + 1; j < n; j++ {
			solution[i] = solution[i].sub(matrix[i][j].mul(solution[j]))
		}
	}
	return solution, nil
}

func main() {
	var n int
	fmt.Scan(&n)
	matrix := make([][]Rational, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]Rational, n+1)
		for j := 0; j <= n; j++ {
			var value int64
			fmt.Scan(&value)
			matrix[i][j] = NewRational(value, 1)
		}
	}

	solution, err := gaussElimination(matrix)
	if err != nil {
		fmt.Println("No solution")
	} else {
		for _, sol := range solution {
			fmt.Printf("%s/%s\n", sol.Numerator.String(), sol.Denominator.String())
		}
	}
}
