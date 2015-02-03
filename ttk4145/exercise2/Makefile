CC=gcc
CFLAGS= -I -Wall

all: server client

server: server.c list.c list.h
	gcc -o server server.c list.c $(CFLAGS) -lpthread
	
client: client.c
	gcc -o client client.c $(CFLAGS)

