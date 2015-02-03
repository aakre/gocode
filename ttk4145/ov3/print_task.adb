with Ada.Text_IO, Ada.Real_Time;
use Ada.Text_IO, Ada.Real_Time;

procedure Print_Task is
	
	task Hello;
	task World;
	
	task body Hello is
		Next : Time := Clock;
	begin
		loop
			Put_Line("Hello");			
			delay 1.0;
		end loop;
	end Hello;
	
	task body World is
		Next : Time := Clock;
	begin
		loop
			Put_Line("World");
			Delay 2.0;
		end loop;
	end World;

begin
	null;
end Print_Task;
