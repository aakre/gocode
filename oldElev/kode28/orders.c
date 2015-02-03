#include "orders.h"
#include "enums.h"


void clear_order_table(int order_table[N_FLOORS][N_BUTTONS]) {
	int floor, button;
	for (floor=0; floor<N_FLOORS; floor++){
		for (button=0; button<N_BUTTONS;button++) {
		order_table[floor][button] = 0;
		}
	} 
}// End clear_order_table

void get_orders(int order_table[N_FLOORS][N_BUTTONS], int state) {
	int floor, button;
	for (floor=0; floor<N_FLOORS; floor++){
		for (button=0; button<N_BUTTONS;button++) {
			if(!(floor == 0 && button == 1) && !(floor == 3 && button == 0)){	// Checks if we're trying to check a non-existing button
				if(state != EMERGENCY || button == 2){				// Checks if we're in the state EMERGENCY (only allows orders from within the elevator if we are)
					if(elev_get_button_signal(button, floor) == 1){
					order_table[floor][button] = 1;
					}
				}
			}
		
		}
	}
}// End get_orders


void delete_orders_at_current_floor(int order_table[N_FLOORS][N_BUTTONS]) {
	int current_floor = elev_get_floor_sensor_signal();
	if(current_floor != -1){
		int button;
		for (button=0; button<N_BUTTONS; button++) {
			order_table[current_floor][button] = 0;
		}
	}
}//End delete orders


int count_orders_over_last_floor(int order_table[N_FLOORS][N_BUTTONS], int last_floor) {
	int floor, button, orders_up = 0;
	for (floor = last_floor + 1; floor < N_FLOORS; floor++){
		for (button=0; button<N_BUTTONS;button++) {
			if(!(floor == 0 && button == 1) && !(floor == 3 && button == 0)) {// Checks if we're trying to check a non-existing button
				if(order_table[floor][button] == 1){
					orders_up++;
				}			
			}
		}		
	}
    return orders_up;
}//End count orders over


int count_orders_under_last_floor(int order_table[N_FLOORS][N_BUTTONS], int last_floor) {
	int floor, button, orders_down = 0;
	for (floor = 0; floor < last_floor; floor++){
	    for (button=0; button<N_BUTTONS;button++) {
		    if(!(floor == 0 && button == 1) && !(floor == 3 && button == 0)) {// Checks if we're trying to check a non-existing button
		    	if(order_table[floor][button] == 1){
		    		orders_down++;
		    	}			
		    }
	    }		
    }
	return orders_down;
}//End count orders under	

