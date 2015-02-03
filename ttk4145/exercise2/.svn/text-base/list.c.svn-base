#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "list.h"

//Linked-list
struct node {
	int id_int;
	int socket;
	struct node *next;
};

struct node *node_init(int id, int socket) {
	struct node *new_node = malloc(sizeof(struct node));
	new_node->id_int = id;
	new_node->socket = socket;
	new_node->next = NULL;
	return new_node;
};

struct node *list_create(void) {
	struct node *root = node_init(0,0);
	return root;
};


int list_add(struct node *root, int id, int socket) {
	struct node *p;
	p = root;
	if (p==0) {
		return -1;
	}
	else {
		while (p->next !=NULL)  {
			p = p->next;
		}
	}
	p->next = node_init(id, socket);
	list_print(root);
	return 0;			
};

int list_get_socket(struct node *root, int id) { //ID's unique, return first result
	struct node *p;
	p = root;
	if (p==0) {
		return -1;
	}
	else {
		while(p != NULL) {
			if (p->id_int == id) {
				return p->socket;
			}
			p = p->next;
		}
	}
	return -1; //No results
};

int list_del(struct node *root, int id) { //ID's unique, delete first result
	struct node *curr_p, *prev_p;
	curr_p = root;
	prev_p = curr_p;
	if (curr_p==0) {
		return -1;
	}
	else {
		while((curr_p = curr_p->next) != NULL) {
			if (curr_p->id_int == id) {
				prev_p->next = curr_p->next;
				free(curr_p);
				return 0; //Success
			}
			prev_p = curr_p;
		}
	}
	return -1; //No results
};

int list_print(struct node *root) {
	struct node *p;
	p = root;
	if (p==0) {
		return -1;
	}
	printf("List of active clients:\n");
	while ((p = p->next) != NULL) {
		printf("---> Client id_int: %d\n", p->id_int);
	}
	return 0;
};

char *list_listofclients(struct node *root){
	int i;
	char *clients = malloc(sizeof(char)*MAXDATASIZE); //Uten malloc får jeg feil: returnerer adresse til lokal variabel
	for(i = 0; i<MAXDATASIZE;i++){
			clients[i] = '\0';
		}

	struct node *p;
	p = root;
	if (p==0) {
		return NULL;;
	}
	while ((p = p->next) != NULL) {
		char id_c[5] = {0};
		sprintf(id_c, "%d", p->id_int);
		strcat(clients, id_c);
		strcat(clients, "\n");
	}
	strcat(clients, "\0");
	return clients;
};	
	
		
int example(){
	printf("Demonstration of list functions\n");
	struct node *root = list_create();
	if(list_add(root, 5,3) != 0) {
		printf("List error: add\n");
	}
	char *c = list_listofclients(root);
	printf("%s", c);
	return 0;
};
	



	
