run/todo:
	@go run ./cmd/todoMain


#если захочешь тестить по одно прогграме в коде(будет создавать json файл) пиши эти команды в консоле

add:
	go run ./cmd/todoMain -add

complete:
	go run ./cmd/todoMain -complete=1

delete:
	go run ./cmd/todoMain -del=1

list:
	go run ./cmd/todoMain/ -list
