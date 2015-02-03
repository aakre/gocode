#include <unistd.h>

#include "elev.h"

#define VELOCITY 200  //Sets the speed of the elevator (50 - 300)
#define STOP_LATENCY 100  //Decides for how long [µs] the elevater will reverse its direction before stopping. Useful on sloppy motors (100 - 100000).
#define DOOR_DELAY 3  //# of seconds the elevator waits before closing the doors

void change_speed(int speed, int new_speed);
void open_door();
void close_door();
