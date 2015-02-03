#include "direction.h"
#include "enums.h"


int get_direction(int order_table[N_FLOORS][N_BUTTONS], int state) {
	int current_floor = elev_get_floor_sensor_signal();
	static int last_floor;
	static int last_direction;
	
	if(current_floor != -1){
		last_floor = current_floor;
		if(state == GOING_UP){
			last_direction = GOING_UP;
		}else if(state == GOING_DOWN){
			last_direction = GOING_DOWN;
		}
	}

	//printf("\nLast Floor: ");
	//printf("%d", last_floor);
	if(current_floor == -1 && (state == GOING_UP || state == GOING_DOWN)) return -1;	
	if(current_floor == 0 && state == GOING_DOWN) return STOP;
	if(current_floor == N_FLOORS-1 && state == GOING_UP) return STOP;
	
	int next_event = STOP;
	int orders_up = count_orders_over_last_floor(order_table, last_floor);
	int orders_down = count_orders_under_last_floor(order_table, last_floor);

	switch (state) {
		case GOING_UP:
			if (order_table[current_floor][B_UP] == 1 || order_table[current_floor][B_ELEVATOR] == 1 || orders_up == 0) {
				next_event = STOP;
			}else{
				next_event = GO_UP;	
			}
			break;
		case GOING_DOWN:
			if (order_table[current_floor][B_DOWN] == 1 || order_table[current_floor][B_ELEVATOR] == 1 || orders_down == 0){
				next_event = STOP;
			}else{
				next_event = GO_DOWN;	
			}
			break;
		case STOP_GOING_UP:
			if(orders_up == 0){
				if(orders_down != 0){
					next_event = GO_DOWN;
				}else{
					next_event = STOP;
				}
			}else{
				next_event = GO_UP;
			}
			break;
		case STOP_GOING_DOWN:
			if(orders_down == 0){
				if(orders_up != 0){
					next_event = GO_UP;
				}else{
					next_event = STOP;
				}
			}else{
				next_event = GO_DOWN;
			}
			break;
		case IDLE:
			if (order_table[current_floor][0] == 1){
				next_event = GO_UP;
			}else if(order_table[current_floor][B_DOWN] == 1 || order_table[current_floor][B_ELEVATOR] == 1){
				next_event = GO_DOWN;
			}else if(orders_down != 0 && orders_down >= orders_up){
				next_event = GO_DOWN;
			}else if(orders_up != 0){
				next_event = GO_UP;
			}
			break;
		case EMERGENCY:
			if(orders_down == 0 && orders_up == 0){
				if(order_table[last_floor][B_ELEVATOR] == 1 && last_direction == GOING_UP){
					next_event = GO_DOWN;
				}else if(order_table[last_floor][B_ELEVATOR] == 1 && last_direction == GOING_DOWN){
					next_event = GO_UP;
				}
			}else if(orders_up > orders_down){
				next_event = GO_UP;
			}else{
				next_event = GO_DOWN;
			}
			break;
	}// End switch
	return next_event;
}
