#include "run_elevator.h"
#include "enums.h"

int main(){
	printf("\nInitialization starting...\n");
	if(!elev_init()){
		printf("\nFailed to initialize elevator...\n");
		return 0;
	}
	printf("\nInitialization complete\n");

	int state = BOOT;
	int event = STOP;
	int order_table[N_FLOORS][N_BUTTONS];
	clear_order_table(order_table);
	
	while(1){
		//printf("\nCurrent state: %d, Current event: %d", state, event);

		state = handle_event(state, event, order_table);
		get_orders(order_table, state);
		set_lights(order_table);

		int temp_event = get_direction(order_table, state);
		if(get_emergency(state)){
			event = EMERGENCY_BUTTON;
		}else if(get_obstruction(order_table, state)){
			event = OBSTRUCTION;
		}else if(temp_event != -1){
			event = temp_event;
		}
	}
	return 0;
} // End main
