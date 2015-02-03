#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <pthread.h>

typedef void* ThreadParam;
typedef void* (*Thread)(void*);

void *print_number(int* num) {
    while (1) {
        printf("Tallet er: %d\n", *num);
        sleep(1);
    }
	return NULL;
}


void start_thread(int tall) {
	pthread_t t;
	int *tal = malloc(sizeof(int*));
	*tal = tall;
	pthread_create(&t, NULL, (Thread)print_number, tal);
}


int main() {
	start_thread(123);
	start_thread(321);

	pthread_exit(NULL);
}

