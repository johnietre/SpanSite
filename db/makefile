.PHONY: all

all: compilec compilers link clean

compilec:
	gcc -c src/main.c

compilers:
	cargo build --release

link:
	gcc -o main main.o ./target/release/libdb.so

clean:
	rm main.o