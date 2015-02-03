#include "signals.h"
#include "enums.h"
	

void activate_emergency(int order_table[N_FLOORS][N_BUTTONS]){
	if(elev_get_floor_sensor_signal() != -1){
		open_door();
		timer_set();	
	}
	clear_order_table(order_table);
	stop_lamp_on();
	printf("\nEMERGENCY ACTIVATED\n");
} // End activate_emergency

void deactivate_emergency(){
	close_door();
	stop_lamp_off();
	printf("\nEMERGENCY DEACTIVATED\n");
} // End deactivate_emergency

int get_obstruction(int order_table[N_FLOORS][N_BUTTONS], int state){ //Recognizes orders at the current floor as an obstruction (when the door is already open)
	int obstruction = 0;
	if(state == STOP_GOING_UP || state == STOP_GOING_DOWN){
		int button;
		if(elev_get_floor_sensor_signal() != -1){
			int current_floor = elev_get_floor_sensor_signal();
			for (button=0; button<N_BUTTONS; button++) {
				obstruction += order_table[current_floor][button];
			}
		}
		if(obstruction){
			delete_orders_at_current_floor(order_table);
		}
	}
	obstruction += elev_get_obstruction_signal();
	return obstruction;
} // End get_obstruction

int get_emergency(int state){
	int emergency = 0;
	emergency += elev_get_stop_signal();
	if((state == IDLE || state == STOP_GOING_UP || state == STOP_GOING_DOWN) && elev_get_floor_sensor_signal() == -1){
		emergency ++;
		printf("\nWARNING: Elevator slipped past floor; check if VELOCITY is too high/STOP_LATENCY too low\n");
    printf("(Try setting VELOCITY to 200 and STOP_LATENCY to 10000 on elevator with sloppy engine)\n");
	}
	return emergency;
}  // End get_emergency
