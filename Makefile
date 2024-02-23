init: 
	npm install
	make generate

generate:
	npx tailwindcss -i ./internal/assets/styles/tailwind.css -o ./dist/styles.css
	templ generate

build: generate
	go build -o ./dist/ ./cmd/server/

dev:
	air -c air.toml

templ-watch:
	templ generate -watch --proxy="http://localhost:3000"

css-watch:
	npx tailwindcss -i ./internal/assets/styles/tailwind.css -o ./dist/styles.css --watch
