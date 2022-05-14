dbu:
	docker-compose down && docker-compose build && docker-compose up

down:
	docker-compose down

build:
	go build -o build/shortCat

clean:
	rm -rf build/

dev:
	go run main.go

prod:
	./build/shortCat

test:
	./scripts/test.sh
