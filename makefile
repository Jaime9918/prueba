central: clean
	go run Central/main.go

lab1: 
	go run Laboratorio1/main.go

lab2: 
	go run Laboratorio2/main.go

lab3: 
	go run Laboratorio3/main.go

lab4: 
	go run Laboratorio4/main.go

clean:
	if [ -f SOLICITUDES.txt ]; then rm SOLICITUDES.txt -R; fi