#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <pthread.h>

typedef void* ThreadParam;
typedef void* (*Thread)(void*);

void *print_number(int *ptr) {
    int num = (int)ptr;
    while (1) {
        printf("Tallet er: %d\n", num);
        sleep(1);
    }
	return NULL;
}


void start_thread(int tall) {
	//pthread_t t;
	pthread_t *t = malloc(sizeof(pthread_t));
	pthread_create(t, NULL, (Thread)print_number((void *)tall), NULL);
}


int main() {
	printf("Starting thread 1\n");
	start_thread(123);
	printf("Starting thread 2\n");
	start_thread(321);

	pthread_exit(NULL);
}

