# 二进制名称（与root.go中CliName保持一致）
BINARY_NAME := goBusyBox
# 构建输出目录
BUILD_DIR := bin
# 源代码根路径
SRC_ROOT := .
# 定义所有支持的命令
COMMANDS := echo lolcat ls
# 安装路径（系统可执行目录）
# INSTALL_PATH := /usr/local/bin

# 默认目标：构建二进制
all: build

# 构建二进制文件
build: tidy
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_ROOT)/main.go
	@echo "Build completed: $(BUILD_DIR)/$(BINARY_NAME)"

# 整理依赖（自动添加/移除模块依赖）
tidy:
	@echo "Tidying dependencies..."
	go mod tidy

# 清理构建产物
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
	@echo "Clean completed"

links: build
	@echo "Creating command symlinks..."
	@cd $(BUILD_DIR) && for cmd in $(COMMANDS); do \
		if [ ! -L $$cmd ]; then \
			ln -s $(BINARY_NAME) $$cmd; \
			echo "Created symlink: $$cmd -> $(BINARY_NAME)"; \
		fi \
	done

# 安装二进制到系统目录（需sudo权限）
# install: build
# 	@echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)..."
# 	sudo cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_PATH)/
# 	@echo "Install completed. Use '$(BINARY_NAME) --help' to get started"

# 卸载系统中的二进制
# uninstall:
# 	@echo "Uninstalling $(BINARY_NAME)..."
# 	sudo rm -f $(INSTALL_PATH)/$(BINARY_NAME)
# 	@echo "Uninstall completed"

# 格式化所有Go代码
fmt:
	@echo "Formatting code..."
	go fmt ./...
	@echo "Format completed"

# 运行测试（若后续添加测试文件，可在此扩展）
test:
	@echo "Running tests..."
	go test -v ./...
	@echo "Test completed"

# 显示帮助信息
help:
	@echo "Available commands:"
	@echo "  make           - 构建二进制（默认目标）"
	@echo "  make build     - 构建二进制文件到bin目录"
	@echo "  make tidy      - 整理Go模块依赖"
	@echo "  make clean     - 清理构建产物"
	@echo "  make install   - 安装二进制到系统目录（需sudo）"
	@echo "  make uninstall - 从系统目录卸载二进制"
	@echo "  make fmt       - 格式化所有Go代码"
	@echo "  make test      - 运行测试用例"
	@echo "  make help      - 显示此帮助信息"

.PHONY: all build tidy clean install uninstall fmt test help