package mathplus

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

const binaryName = "./mathplus_test_bin"

func TestMain(m *testing.M) {
	// Compilar el binario antes de ejecutar los tests CLI
	build := exec.Command("go", "build", "-o", binaryName, "../cmd/mathplus")
	if output, err := build.CombinedOutput(); err != nil {
		println("Error al compilar el binario:", string(output))
		os.Exit(1)
	}
	// Ejecutar tests
	code := m.Run()

	// Limpiar: eliminar binario al finalizar
	os.Remove(binaryName)

	os.Exit(code)
}

func TestCLIAdd(t *testing.T) {
	cmd := exec.Command(binaryName, "add", "-p", "2", "3")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Error ejecutando CLI: %v", err)
	}
	result := strings.TrimSpace(string(output))
	if result != "5" {
		t.Errorf("CLI add = %v; want 5", result)
	}
}

func TestCLIMul(t *testing.T) {
	cmd := exec.Command(binaryName, "mul", "-p", "2", "4", "5")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Error ejecutando CLI: %v", err)
	}
	result := strings.TrimSpace(string(output))
	if result != "40" {
		t.Errorf("CLI mul = %v; want 40", result)
	}
}
