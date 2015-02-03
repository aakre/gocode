#include "door_engine_control.h"
#include "signals.h"
#include "orders.h"
#include "timer.h"
#include "elev.h"

int handle_event(int state, int event, int order_table[N_FLOORS][N_BUTTONS]);
int handle_event_going_up(int event, int order_table[N_FLOORS][N_BUTTONS]);
int handle_event_going_down(int event, int order_table[N_FLOORS][N_BUTTONS]);
int handle_event_stop_going_up(int event, int order_table[N_FLOORS][N_BUTTONS]);
int handle_event_stop_going_down(int event, int order_table[N_FLOORS][N_BUTTONS]);
int handle_event_idle(int event, int order_table[N_FLOORS][N_BUTTONS]);
int handle_event_emergency(int event, int order_table[N_FLOORS][N_BUTTONS]);
int handle_event_boot();