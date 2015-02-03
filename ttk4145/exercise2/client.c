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
#include <fcntl.h>

#define PORT "9001"
#define HOST "localhost"
#define MAXDATASIZE 100
#define MAXID 10

//PROTIP: cat /etc/hosts

void client_connect(int* sockfd, char host[MAXDATASIZE]){
	struct addrinfo hints, *res, *p;
  
	memset(&hints, 0, sizeof hints);
	hints.ai_family = AF_INET;
	hints.ai_socktype = SOCK_STREAM;
  
	if((getaddrinfo("localhost", PORT, &hints, &res)) != 0) {
		fprintf(stderr, "client: getaddrinfo\n");
	}
	
	*sockfd = 0;
	while(*sockfd <= 0){
		printf("client: attempting connection...\n");
		for(p=res; p!=NULL; p->ai_next) {
			if ((*sockfd = socket(p->ai_family, p->ai_socktype, p->ai_protocol)) == -1) {
				fprintf(stderr, "client: socket\n");
				sleep(1);
				continue;
			}
			if (connect(*sockfd, p->ai_addr, p->ai_addrlen) == -1) {
				close(*sockfd);
				fprintf(stderr, "client: connect\n");
				sleep(1);
				continue;
			}
			break;
		}
		if(p==NULL) {
		fprintf(stderr, "client: failed to connect server at %s \n", HOST);
		}
	}
	printf("client: connected to server at %s \n", HOST);
	freeaddrinfo(res);
}

int main(int argc, char* argv[]){
	int numbytes, i, sockfd, flags;
	char buf_out[MAXDATASIZE], buf_in[MAXDATASIZE], *r_id, s_id[MAXID], *r_msg, s_msg[MAXDATASIZE-MAXID];
	//char host[MAXDATASIZE] = HOST;
	
	if (-1 == (flags = fcntl(sockfd, F_GETFL, 0)))
		{
		flags = 0;
	}
	fcntl(sockfd, F_SETFL, flags | O_NONBLOCK);

	client_connect(&sockfd, HOST);
	
	while(1){
		for(i = 0; i<MAXDATASIZE;i++){
			buf_in[i] = '\0';
			buf_out[i] = '\0';
		}
		r_id = "2";
		r_msg = "Hello";
		sleep(1);
		sprintf(buf_out, "to:%s;msg:%s;", r_id, r_msg);
	
		if (send(sockfd, buf_out, strlen(buf_out)+1, MSG_NOSIGNAL) == -1) {
			fprintf(stderr, "client: send\n");
		}else{
			printf("client: msg sent: %s \n", r_msg);
		}
	
		if ((numbytes = recv(sockfd, buf_in, MAXDATASIZE-1, 0)) == -1) {
			printf("client: recv\n");
		}else if(numbytes == 0){
			printf("client: server connection lost\n");
			printf("client: reconnecting...\n");
			client_connect(&sockfd, HOST);
		}else{
			if(strstr(buf_in, "connected client id's:")){
				printf("%s \n", buf_in);
			}else{
				sscanf(buf_in, "to:%s;msg:%s;", s_id, s_msg);
				printf("Message from client id %s : %s \n", s_id, s_msg);
			}
		}
 
	}

	close(sockfd);

	return 0;
};
