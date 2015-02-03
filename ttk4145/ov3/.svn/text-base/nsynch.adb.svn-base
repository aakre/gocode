with Ada.Text_IO;
use Ada.Text_IO;

procedure NSynch is

   N : constant := 10;
   Q : Natural := 0;

   protected Manager is
      entry Synchronize;
   private
      entry Wait;
   end Manager;

   protected body Manager is
      entry Synchronize when Q < N is
      begin
         Q := Q + 1;
         requeue Wait;
      end Synchronize;
      
      entry Wait when Q = N is
      begin
         if Wait'Count = 0 then
            Q := 0;
			end if;
		end Wait;
   end Manager;

   task type Worker;

   task body Worker is
   begin
      for i in 0..10 loop
         Manager.Synchronize;
         Put("!");
         Manager.Synchronize;
         Put(".");
		end loop;  
   end Worker;

   Workers : array (1 .. N) of Worker;

begin
   null;
end NSynch;






