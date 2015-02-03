--To make Ada compiler work in terminal window, enter: export PATH=/home/student/gnat/bin:$PATH

with Ada.Text_IO;
with Ada.Strings.Unbounded; use Ada.Strings.Unbounded;
with Ada.Strings.Unbounded.Text_IO; use Ada.Strings.Unbounded.Text_IO;

procedure Hello_World is
	S : Unbounded_String;
	
begin
	Ada.Text_IO.Put_Line("Hello, world!");
	Ada.Text_IO.Put_Line("What's your name?");
	
	while S = "" loop
		S := Get_line;
		if S = "" then
			Ada.Text_IO.Put_Line("I didn't quite catch that, what was your name again?");
		end if;
	end loop;
	
	Put("Welcome, " & S & "!");
	
end Hello_World;

