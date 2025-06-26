![First Project](https://img.shields.io/badge/built%20with-coffee%20and%20curiosity-%23c49cff?style=flat&logo=go)

> 💡 Started as a side exercise… now it’s my first open-source package!

# mathplus

A mathematical Go library with extended basic functions, statistic functions, logic functions and trigonometric functions. Ideal for personal use or as a foundation for projects that require calcularions beyond standard `math`.

## ✨ Main functionalities

- ➕ Arithmetic: suma, resta, multiplicación, división, potencias, raíces
- 🧮 Numeric logic: factorial, primos, MCD, MCM
- 📊 Statistics: media, mediana, moda, desviación estándar
- 📐 Trigonometry: seno, coseno, tangente
- 🔢 Extremes: valor mínimo y máximo con soporte para NaN e infinitos
- 🧠 Logarithms: naturales, base 2, base 10, Log1p y Logb
- ⚙️ Auxiliar functions: valor absoluto, control de errores

And I'm open for suggestions if any important functions are missing!

## 🛠️ Installation

`go get github.com/HugoCodesig/mathplus`

## 🧪 Use examples

```go
package main

import (
    "fmt"
    "github.com/HugoCodesig/mathplus"
)

func main() {
    fmt.Println("Sum:", mathplus.Add(1, 2, 3))                     // 6
    fmt.Println("Average:", mathplus.Avg(10, 20, 30))              // 20
    fmt.Println("Median:", mathplus.Med(3, 1, 4, 2))               // 2.5
    fmt.Println("Moda:", mathplus.Mode(1, 2, 2, 3, 3, 3, 4))       // [3]
    fmt.Println("Is 17 prime:", mathplus.IsPrime(17))              // true
    fmt.Println("Factorial 5:", mathplus.Factorial(5))             // 120
    fmt.Println("Máximo:", mathplus.Max(2, 4, 6, math.Inf(1)))     // +Inf
}
```

## 📁 File structure

`arithmetic.go` – Basic arithmetic operations

`statistics.go` – Statistic functions

`logic.go` – Logic functions with floating point numbers

`extremes.go` – Minimum and maximum calculations

`logs.go` – Basic and extended logarithms

`trigonometry.go` – Trigonometric functions

`utils.go` – Auxiliar functions
