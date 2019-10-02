package main

const makefileTemplate = `PROJECT={{ .Project }}

CC=g++
FLAGS=-W -Wall -std=c++17

#STATIC=-static # -static-libgcc -static-libstdc++
SFML=-lsfml-graphics -lsfml-window -lsfml-system

INCLUDE_DIRS=-Iinclude {{ range .IncDirs }}-I{{ . }} {{ end }}
LIB_DIRS=-Llib {{ range .LibDirs }}-L{{ . }} {{ end }}
LIBS={{ range .Libraries }}-l{{ . }} {{ end }}# $(SFML)

OPTIONS=$(INCLUDE_DIRS) $(LIB_DIRS) $(LIBS)

release:
	$(CC) -O2 $(FLAGS) -o build/$(PROJECT) src/main.cpp $(OPTIONS)

debug:
	$(CC) -g $(FLAGS) -o build/$(PROJECT)_debug src/main.cpp $(OPTIONS)

clean:
	rm build/*`

const maincppTemplate = `#include <iostream>

int main() {
	return 0;
}`
