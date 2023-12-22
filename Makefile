run/todo:
	@go run ./cmd/todo

#build:
# 	@go build ./cmd/todo

#если захочешь тестить по одно прогграме в коде(будет создавать json файл) пиши эти команды в консоле

add:
	@go run ./cmd/todo -add

complete:
	@go run ./cmd/todo -complete=1

delete:
	@go run ./cmd/todo -del=1

list:
	@go run ./cmd/todo -list
