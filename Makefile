CC=gcc
CFLAGS=-I. -fPIC
OBJ=obj
SRC=sha3
TARGET=libbitcore_hash.so

SOURCES := $(wildcard $(SRC)/*.c)
SOURCES := $(filter-out sha3/md_helper.c, $(SOURCES))
OBJS :=  $(patsubst $(SRC)/%.c, $(OBJ)/%.o, $(SOURCES)) $(OBJ)/bitcore.o


$(TARGET): $(OBJS) 
	$(CC) -o $@ $^ $(CFLAGS) -shared 
	cp $@ /usr/lib

$(OBJ)/%.o: $(SRC)/%.c | $(OBJ)
	$(CC) -c -o $@ $< $(CFLAGS)

$(OBJ):
	mkdir obj

#$(OBJ)/timetravel.o: timetravel.c
#	$(CC) -c -o $@ $< $(CFLAGS)

$(OBJ)/bitcore.o: bitcore.c
	$(CC) -c -o $@ $< $(CFLAGS)

.PHONY: clean
clean:
	rm -Rf $(OBJ) *.so 

