run: build
	@./bin/app

build:
	@templ generate && go build -o bin/app .

css:
	npx tailwindcss -i views/css/app.css -o public/styles.css --watch

templ:
	templ generate --watch --proxy=http://localhost:4000