##
## EPITECH PROJECT, 2022
## Kuzzle GM
## File description:
## Makefile
##

CC = go build

TARGET = gm

SRC_COMPONENTS = main.go \

SRC =	$(SRC_COMPONENTS) \

all: $(TARGET)

build_all:
	$(CC) $(CFLAGS) -o $(TARGET) $(OBJ)

$(TARGET): build_all

clean:
	rm -f $(TARGET)

fclean: clean

re: fclean all

.PHONY: all re clean fclean
