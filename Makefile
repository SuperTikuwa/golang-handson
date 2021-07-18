init:
	go build -o ./util/installer/installer ./util/installer/main.go
	util/installer/installer

remove:
	go build -o ./util/remover/remover ./util/remover/main.go
	util/remover/remover