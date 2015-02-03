#include "door_engine_control.h"
#include "enums.h"

//Properly sets speed according to current speed.
void change_speed(int speed, int new_speed){
	switch(new_speed){
		case S_UP:
			elev_set_speed(VELOCITY);
			break;
		case S_DOWN:
			elev_set_speed(-VELOCITY);
			break;
		case S_STOP:
			elev_set_speed(0);
			if(speed == S_UP){
				elev_set_speed(-VELOCITY);
			}else if(speed == S_DOWN){
				elev_set_speed(VELOCITY);
			}
			usleep(STOP_LATENCY);
			elev_set_speed(0);
			break;	
	}
}// End change_speed

void open_door(){
	if(elev_get_floor_sensor_signal() != -1){
		elev_set_door_open_lamp(1);
	}
}// End open_door

void close_door(){
	elev_set_door_open_lamp(0);
}// End close_door
