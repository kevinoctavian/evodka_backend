APP_NAME = evodka_backend
MAIN_FILE = main.go
BUILD_DIR = bin

MIGRATIONS_FOLDER = ./platform/migrations
DB_NAME = evodka_app
DB_USER = postgres
DATABASE_URL = postgres://$(DB_USER)@localhost/$(DB_NAME)?sslmode=disable

# Default target
all: build run

# Build binary
build:
	@echo "🔨 Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	@echo "✅ Build complete: $(BUILD_DIR)/$(APP_NAME)"

# Run app without building
run:
	@echo "🚀 Running $(APP_NAME)..."
	@go run $(MAIN_FILE)

# Build and run
dev: build
	@echo "🚀 Starting $(APP_NAME)..."
	@./$(BUILD_DIR)/$(APP_NAME)

# Clean build files
clean:
	@echo "🧹 Cleaning..."
	@rm -rf $(BUILD_DIR)
	@echo "✅ Clean complete"

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down


seed:
	psql -h localhost -p 5432 -U$(DB_USER) -d $(DB_NAME) -a -f platform/seeds/001_seed_user_table.sql
	psql -h localhost -p 5432 -U$(DB_USER) -d $(DB_NAME) -a -f platform/seeds/002_seed_book_table.sql

.PHONY: all build run dev watch clean
