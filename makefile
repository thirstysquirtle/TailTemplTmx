all: ./view/*_templ.go ./view/*/*_templ.go ./tmp/main 

./view/*_templ.go ./view/*/*_templ.go: ./view/*.templ ./view/*/*.templ
	templ generate

./tmp/main: ./*.go ./*/*.go
	go build -o ./tmp/main .

build:
	npx tailwindcss -i ./src/tailwindInput.css -o ./static/style.css
	go build -o ./tmp/main .



