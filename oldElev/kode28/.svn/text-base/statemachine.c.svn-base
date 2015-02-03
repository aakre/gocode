//File includes statemachine and statemachine case functions.

#include "statemachine.h"
#include "enums.h"


//LEAN, MEAN STATEMACHINE
int handle_event(int state, int event, int order_table[N_FLOORS][N_BUTTONS]) {
	int next_state = state;
	switch (state) {
	case GOING_UP:
		next_state = handle_event_going_up(event, order_table);
		break;
	
	case GOING_DOWN:
		next_state = handle_event_going_down(event, order_table);	
		break;
		
	case STOP_GOING_UP:
		next_state = handle_event_stop_going_up(event, order_table);
		break;
	
	case STOP_GOING_DOWN:
		next_state = handle_event_stop_going_down(event, order_table);	
		break;
			
	case IDLE:
		next_state = handle_event_idle(event, order_table);
		break;

	case EMERGENCY:
		next_state = handle_event_emergency(event, order_table);
		break;

	case BOOT:
		next_state = handle_event_boot();
		break;
		
	} // End statemachine
	return next_state;	
}// End handle_event



//CASE-FUNKSJONER
int handle_event_going_up(int event, int order_table[N_FLOORS][N_BUTTONS]) {
	int next_state = GOING_UP;
	switch (event) {

			case EMERGENCY_BUTTON:
				change_speed(S_UP, S_STOP);
				activate_emergency(order_table);
				next_state = EMERGENCY;
				break;

			case GO_UP:
				next_state = GOING_UP;
				break;
	
			case GO_DOWN:
				change_speed(S_UP, S_DOWN);
				next_state = GOING_DOWN;
				break;
			
			case STOP:
				change_speed(S_UP, S_STOP);
				next_state = STOP_GOING_UP;
				delete_orders_at_current_floor(order_table);
				open_door();
				timer_set();
				break;
			
			case OBSTRUCTION:
				change_speed(S_UP, S_STOP);
				activate_emergency(order_table);
				next_state = EMERGENCY;
				break;
		}
	return next_state;
}





int handle_event_going_down(int event, int order_table[N_FLOORS][N_BUTTONS]) {
	int next_state = GOING_DOWN;
	switch (event) {

			case EMERGENCY_BUTTON:
				change_speed(S_DOWN, S_STOP);
				activate_emergency(order_table);
				next_state = EMERGENCY;
				break;

			case GO_UP:
				change_speed(S_DOWN, S_UP);
				next_state = GOING_UP;
				break;

			case GO_DOWN:
				next_state = GOING_DOWN;
				break;
			
			case STOP:
				change_speed(S_DOWN, S_STOP);
				next_state = STOP_GOING_DOWN;
				delete_orders_at_current_floor(order_table);
				open_door();
				timer_set();
				break;
			
			case OBSTRUCTION:
				change_speed(S_DOWN, S_STOP);
				activate_emergency(order_table);
				next_state = EMERGENCY;
				break;
		}
	return next_state;
}

int handle_event_stop_going_up(int event, int order_table[N_FLOORS][N_BUTTONS]) {
	int next_state = STOP_GOING_UP;	
	switch (event) {
	
			case EMERGENCY_BUTTON:
				activate_emergency(order_table);
				next_state = EMERGENCY;
				break;

			case GO_UP:
				if(timer_elapsed() > DOOR_DELAY){
					close_door();
					change_speed(S_STOP, S_UP);
					next_state = GOING_UP;
				}else{
					next_state = STOP_GOING_UP;
				}
				break;

			case GO_DOWN:
				if(timer_elapsed() > DOOR_DELAY){
					close_door();
					change_speed(S_STOP, S_DOWN);
					next_state = GOING_DOWN;
				}else{
					next_state = STOP_GOING_DOWN;
				}
				break;
			
			case STOP:
				if(timer_elapsed() > DOOR_DELAY){
					close_door();
					next_state = IDLE;
				}else{
					next_state = STOP_GOING_UP;
				}
				break;
			
			case OBSTRUCTION:
				timer_set();
				next_state = STOP_GOING_UP;
				break;
			
		}
		return next_state;
}

int handle_event_stop_going_down(int event, int order_table[N_FLOORS][N_BUTTONS]) {
	int next_state = STOP_GOING_DOWN;	
	switch (event) {

			case EMERGENCY_BUTTON:
				activate_emergency(order_table);
				next_state = EMERGENCY;
				break;

			case GO_UP:
				if(timer_elapsed() > DOOR_DELAY){
					close_door();
					change_speed(S_STOP, S_UP);
					next_state = GOING_UP;
				}else{
					next_state = STOP_GOING_UP;
				}
				break;

			case GO_DOWN:
				if(timer_elapsed() > DOOR_DELAY){
					close_door();
					change_speed(S_STOP, S_DOWN);
					next_state = GOING_DOWN;
				}else{
					next_state = STOP_GOING_DOWN;
				}
				break;
			
			case STOP:
				if(timer_elapsed() > DOOR_DELAY){
					close_door();
					next_state = IDLE;
				}else{
					next_state = STOP_GOING_DOWN;
				}
				break;
			
			case OBSTRUCTION:
				timer_set();
				next_state = STOP_GOING_DOWN;
				break;
			
		}
		return next_state;
}

int handle_event_idle(int event, int order_table[N_FLOORS][N_BUTTONS]) {
	int next_state = IDLE;
	switch (event) {

			case EMERGENCY_BUTTON:
				activate_emergency(order_table);
				next_state = EMERGENCY;
				break;

			case GO_UP:
				close_door();
				change_speed(S_STOP, S_UP);
				next_state = GOING_UP;
				break;

			case GO_DOWN:
				close_door();
				change_speed(S_STOP, S_DOWN);
				next_state = GOING_DOWN;
				break;
			
			case STOP:
				next_state = IDLE;
				break;
			
			case OBSTRUCTION:
				activate_emergency(order_table);
				next_state = EMERGENCY;
				break;
			
		}
	return next_state;
}


int handle_event_emergency(int event, int order_table[N_FLOORS][N_BUTTONS]) {
	int next_state = EMERGENCY;	
		switch (event) {

			case EMERGENCY_BUTTON:
				activate_emergency(order_table);
				next_state = EMERGENCY;
				break;

			case GO_UP:
				if(elev_get_floor_sensor_signal() != -1 && timer_elapsed() < DOOR_DELAY){	//Elevator will wait for # seconds defined by DOOR_DELAY before continuing when at a floor
					return EMERGENCY;
				}
				deactivate_emergency();
				change_speed(S_STOP, S_UP);
				next_state = GOING_UP;
				break;

			case GO_DOWN:
				if(elev_get_floor_sensor_signal() != -1 && timer_elapsed() < DOOR_DELAY){	//Elevator will wait for # seconds defined by DOOR_DELAY before continuing when at a floor
					return EMERGENCY;
				}
				deactivate_emergency();
				change_speed(S_STOP, S_DOWN);
				next_state = GOING_DOWN;
				break;
			
			case STOP:
				next_state = EMERGENCY;
				break;
			
			case OBSTRUCTION:
				timer_set();
				next_state = EMERGENCY;
				break;
			
		}
		return next_state;
}


int handle_event_boot() {
	int next_state = IDLE;
	int floor = elev_get_floor_sensor_signal();
	printf("\nBOOT IN PROGRESS; PLEASE WAIT...\n");
	while(floor == -1){
		change_speed(STOP, S_DOWN);
		floor = elev_get_floor_sensor_signal();
	}
	elev_set_floor_indicator(floor);
	change_speed(S_DOWN, STOP);
	printf("\nELEVATOR READY\n");
	return next_state;
}
	

	





















