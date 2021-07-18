init:
	go build -o ./util/installer/installer ./util/installer/main.go
	util/installer/installer

remove:
	go build -o ./util/remove/remove ./util/remove/main.go
	util/remover/remover