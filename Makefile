OBJ_NAME = 
LDFLAGS = "-w -s"
clock:
	$(eval OBJ_NAME += clock)
	cd ./cmd/clock; go build -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../../bin; 

LDFLAGS = "-w -s"
netcat:
	$(eval OBJ_NAME += netcat)
	cd ./cmd/netcat; go build -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../../bin; 

LDFLAGS = "-w -s"
spinner:
	$(eval OBJ_NAME += spinner)
	cd ./cmd/spinner; go build -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../../bin; 

LDFLAGS = "-w -s"
basics:
	$(eval OBJ_NAME += basics)
	cd ./cmd/basics; go build -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../../bin; 

LDFLAGS = "-w -s"
challenges:
	$(eval OBJ_NAME += challenges)
	cd ./cmd/challenges; go build -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../../bin; 
LDFLAGS = "-w -s"

mutexes:
	$(eval OBJ_NAME += mutexes)
	cd ./cmd/mutexes; go build -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../../bin; 

producer:
	$(eval OBJ_NAME += producer)
	cd ./cmd/producer-consumer; go build -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../../bin; 

chan:
	$(eval OBJ_NAME += chan)
	cd ./cmd/chan; go build -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../../bin; 

goroutines:
	$(eval OBJ_NAME += goroutines)
	cd ./cmd/routines; go build -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../../bin; 

performance:
	$(eval OBJ_NAME += performance)
	cd ./cmd/performance; go build -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../../bin; 

dining:
	$(eval OBJ_NAME += dining)
	cd ./cmd/dining; go build -ldflags $(LDFLAGS) -o $(OBJ_NAME); mv $(OBJ_NAME) ../../bin; 
