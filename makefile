central: clean
	go run Central/main.go

laboratorio1: 
	go run Laboratorio1/main.go

laboratorio2: 
	go run Laboratorio2/main.go

laboratorio3: 
	go run Laboratorio3/main.go

laboratorio4: 
	go run Laboratorio4/main.go

clean:
	if [ -f SOLICITUDES.txt ]; then rm SOLICITUDES.txt -R; fi