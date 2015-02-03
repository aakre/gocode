--To make Ada compiler work in terminal window, enter: export PATH=/home/student/gnat/bin:$PATH

with Ada.Command_Line, Ada.Text_IO;
use Ada.Command_Line, Ada.Text_IO;


procedure Fibonacci is
	A : Natural := 0;
	B : Natural := 1;
	Sum : Natural := 0;
	N : Natural;
	
begin
	if Argument_Count < 1 then
		Put_Line("Error: No arguments");
		N := 0;
	else
		if Argument_Count > 1 then
			Put_Line("Warning: Too many arguments, using last one");
		end if;
		
		N := Integer'value(Argument(Argument_Count));
	end if;
	
	if N > 46 then
		Put_Line("Error: Integer overflow (wrap around). Fibonacci number to high (46 max)");
		return;
	end if;
	
	for I in 1..N loop
		Sum := A + B;
		A := B;
		B := Sum;
	end loop;
	
	Put("The Fibonacci number" & Integer'image(N) & " is:" & Integer'image(A));
	
end Fibonacci;
