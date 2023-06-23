server: 
	go run ./main.go
dockerrun:
	cd docker && docker-compose up
dockerbuild:
	cd docker && docker-compose up --build
