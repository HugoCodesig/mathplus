![Go Version](https://img.shields.io/badge/go-1.24.3-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![First Project](https://img.shields.io/badge/my%20first-project-%23f47fff)

> 💡 Started as a side exercise… now it’s my first open-source package!

# mathplus

![mathplus banner](assets/banner.png)

A mathematical Go library with extended basic functions, statistic functions, logic functions and trigonometric functions. Ideal for personal use or as a foundation for projects that require calcularions beyond standard `math`.

## ✨ Main functionalities

- ➕ Arithmetic: addition, subtraction, multiplication, division, power, square root
- 🧮 Numeric logic: factorial, prime checking, LCM, MCD
- 📊 Statistics: average, median, mode, standard deviation
- 📐 Trigonometry: sine, cosine, tangeant
- 🔢 Extremes: minimum and maximum values with NaN and ±Inf support
- 🧠 Logarithms: naturals, base 2, base 10, custom base, Log1p and Logb
- ⚙️ Auxiliar functions: absolute value, error control

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
