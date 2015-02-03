#include "light_control.h"


void set_lights(int order_table[N_FLOORS][N_BUTTONS]){
	int floor, button;
	for (floor=0; floor<N_FLOORS; floor++){
		for (button=0; button<N_BUTTONS;button++) {
			if(!(floor == 0 && button == 1) && !(floor == 3 && button == 0)){ // Checks if we're trying to check a non-existing button
				elev_set_button_lamp(button, floor, order_table[floor][button]); // Sets button lights
			}
		}
	} 
	
	floor = elev_get_floor_sensor_signal();
	if (floor != -1) {
		elev_set_floor_indicator(floor); // Sets floor indicator lights
	}
}// End set_lights



void stop_lamp_on(){
	elev_set_stop_lamp(1);
}


void stop_lamp_off(){
	elev_set_stop_lamp(0);
}

