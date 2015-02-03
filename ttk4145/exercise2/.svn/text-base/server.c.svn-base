#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <unistd.h>
#include <errno.h>
#include <netdb.h>
#include <arpa/inet.h>
#include <sys/wait.h>
#include <signal.h>
#include <pthread.h>
#include "list.h"

#define PORT "9001"
#define BACKLOG 10
#define MAXDATASIZE 100
#define MAXID 10


struct client_info{
	int id;
	int socket;
	struct node *root;
};


void* client_thread(void* client_info_void){
	struct client_info *my_info = (struct client_info*) client_info_void;
	int my_socket = my_info->socket;
	int my_id = my_info->id;
	
	printf("server: New client connected on socket %d at mem: %p \n", my_socket, &my_socket);
	
	if ((list_add(my_info->root, my_id, my_socket)) != 0) {
		printf("List error: adding client '%d'\n", my_id);
	}
	int i, their_id, their_socket;
	int numbytes = -1;
	char buf_in[MAXDATASIZE];
	char buf_out[MAXDATASIZE];
	char r_id[MAXID];
	
	char* client_list = list_listofclients(my_info->root);
	while(numbytes != 0){
		for(i = 0; i<MAXDATASIZE;i++){
  			buf_in[i] = '\0';
  			buf_out[i] = '\0';
  		}
		if((numbytes = recv(my_socket, buf_in, MAXDATASIZE-1, 0)) > 0){
			sscanf(buf_in, "to:%s;msg:%s;;\n", &r_id, buf_out);
			their_id = atoi(r_id);
			their_socket = list_get_socket(my_info->root, their_id);
			if (send(their_socket, buf_in, strlen(buf_in)+1, MSG_NOSIGNAL) == -1) {
				printf("server: send\n");
				sleep(1);
			}else{
				printf("server: msg sent to socket %d from socket %d \n", their_socket, my_socket);
			}
		}else if(numbytes == 0){
			printf("server: connection lost on socket %d \n", my_socket);
		}else{
			printf("server: recv\n");
			sleep(1);	
		}
	}
	list_del(my_info->root, my_id);
	printf("server: closing socket %d \n", my_socket);
	close(my_socket);
 	return NULL;
 	
}//	End client_thread

int main(int argc, char *argv[]){
	int sockfd, new_fd;
	int yes = 1;
	struct sockaddr_storage their_addr;
	struct addrinfo hints, *res, *p;
	socklen_t addr_size;

	memset(&hints, 0, sizeof hints);
	hints.ai_family = AF_INET;
	hints.ai_socktype = SOCK_STREAM;
	hints.ai_flags = AI_PASSIVE;

	//Server setup & Error handling  
  
	if(getaddrinfo(NULL, PORT, &hints, &res) != 0) {
		printf("getaddrinfo error\n");
		return 1;
	}
  
	for(p=res; p!=NULL; p->ai_next) {
		if((sockfd = socket(p->ai_family, p->ai_socktype, 0)) == -1) {
			perror("server: socket\n");
			sleep(1);
			continue;
		}
		
		if(setsockopt(sockfd, SOL_SOCKET, SO_REUSEADDR, &yes, sizeof(int)) == -1){
			perror("setsockopt");
			exit(1);
		}
		
		if(bind(sockfd, res->ai_addr, res->ai_addrlen) == -1) {
			close(sockfd);
			perror("server:bind");
			sleep(1);
			continue;
		}
		break;
	}//End for-loop

	if(p==NULL) {
		printf("server: failed to bind\n");
		return 2;
	}

	freeaddrinfo(res); //Finished with res


  	if(listen(sockfd, BACKLOG) == -1) {
		printf("server: listen\n");
	}


	printf("Server: Waiting for connections...\n");
	struct node *list_root = list_create();
	int client_id = 0;
	char* client_list;
	char list_buff[MAXDATASIZE];
	int i;

	while(1){
		addr_size = sizeof their_addr;
		new_fd = accept(sockfd, (struct sockaddr *)&their_addr, &addr_size);
		if(new_fd == -1) {
			printf("server: accept\n");
			continue;
		}else{
			int* client_fd = malloc(sizeof(int));
			*client_fd = new_fd;
			client_id++;
			struct client_info *c_info = malloc(sizeof(struct client_info));
			c_info->id = client_id;
			c_info->socket = *client_fd;
			c_info->root = list_root;
			
			
							
			pthread_t *thr = malloc(sizeof(pthread_t));
			pthread_create(thr, NULL, client_thread, c_info);
			
			for(i = 0; i<MAXDATASIZE;i++){
  				list_buff[i] = '\0';
  			}
			
			client_list = list_listofclients(list_root);
			
			sprintf(list_buff, "connected client id's: %s \n", client_list);
			struct node *p;
			/*p = list_root;
			if (p==0) {
				return -1;
			}
			while ((p = p->next) != NULL) {
				send(p->socket, list_buff, strlen(list_buff)+1, MSG_NOSIGNAL);
			}*/
			
		}
		sleep(1);
	}
	return 0;
}//	End main

























