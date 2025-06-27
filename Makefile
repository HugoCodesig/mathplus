# 💡 Makefile para Math+

APP_NAME = mathplus
CMD_DIR = ./cmd/mathplus
BIN = $(APP_NAME)
BIN_CLI_TEST = mathplus_test_bin

# 🎯 Compilar el binario principal
build:
	go build -o $(BIN) $(CMD_DIR)

# 🧪 Ejecutar todos los tests
test:
	go test ./...

# 🧪 Ejecutar solo tests de la librería
test-lib:
	go test ./internal/...

# 🧪 Ejecutar solo tests del CLI
test-cli:
	go test ./tests/cli_test.go

# ▶️ Ejecutar el CLI con parámetros de ejemplo
run:
	go run $(CMD_DIR) add -p 5 3

# 🧹 Eliminar binarios de compilación
clean:
	rm -f $(BIN) $(BIN_CLI_TEST)

# 🪄 Ejecutar build + test
all: build test
