#include "timer.h"

static double timer_start = 0;

void timer_set(){
	timer_start = clock();	
}// End timer_set

double timer_elapsed(){
	double time_elapsed = ((double)(clock() - timer_start))/CLOCKS_PER_SEC;
	return time_elapsed;
}// End timer_elapsed
