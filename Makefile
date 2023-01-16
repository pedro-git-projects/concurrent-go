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
