
GO_CMD=go
GO_APP=./cmd/main.go


.PHONY: run test ent_gen_schema ent_gen install clean

install:
	@echo "Checking if build-essential is installed..."
	@if ! dpkg -l | grep -q "build-essential"; then \
		echo "Installing build-essential..."; \
		sudo apt update && sudo apt install -y build-essential; \
	else \
		echo "build-essential is already installed."; \
	fi

run: install
	@echo "Running Go application with CGO_ENABLED=1..."
	# Set CGO_ENABLED=1 only for this command
	@CGO_ENABLED=1 $(GO_CMD) run $(GO_APP)

test:
	@echo "Running tests..."
	$(GO_CMD) test -v ./...

ent_gen_schema:
	@echo "Generating Ent schema..."
	$(GO_CMD) run -mod=mod entgo.io/ent/cmd/ent new $(SCHEMA) --target ./ent/schema

ent_gen:
	@echo "Generating Ent code..."
	$(GO_CMD) run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema

clean:
	@echo "Cleaning up generated files..."
	rm -f myapp
