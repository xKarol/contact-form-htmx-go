init: 
	npm install
	make generate

generate:
	npx tailwindcss -i ./styles/tailwind.css -o ./dist/styles.css
	templ generate

build: generate
	go build -o ./tmp/ .

dev:
	air

templ-watch:
	templ generate -watch --proxy="http://localhost:3000"

css-watch:
	npx tailwindcss -i ./styles/tailwind.css -o ./dist/styles.css --watch
