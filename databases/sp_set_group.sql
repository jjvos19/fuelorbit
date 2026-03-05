--- 
 -- Descripcion: SP para agrupar registros.
 -- Fecha: 25/02/2026

create or replace procedure sp_set_group()
language plpgsql 
as $$
declare 
   wn_group_id bigint := 0;
   wc_hash_text varchar(64) := null;
   wn_start bigint := 0;
   wn_finish bigint := 0;
 begin
	select min(t.id), max(t.id) into wn_start, wn_finish
	  from (
	select td.id 
	  from tkr_data td  
	 where td.group_blck = 0
	 order by td.id asc
	 limit 10) as t;
 	select 
	       trim(md5(string_agg(t.cadena, '%%'))) into wc_hash_text
	  from (
			select td.id || '*'|| td.tanker_id || '*' || td.gps_coordinate || '*' || td.volume || '*' || td.state_motor || '*' || td.hash_device || '*' || td.send_date as cadena, td.group_blck, td.id 
			  from tkr_data td  
			 where td.group_blck = 0
			   and td.id between  wn_start and wn_finish
			 order by td.id asc
			 ) as t
	 group by t.group_blck;
	insert into public.tkr_group(hash_group, data_id_start, data_id_finish, created_at)
	       values(wc_hash_text, wn_start, wn_finish, now()) returning id into wn_group_id;
	update public.tkr_data 
	   set group_blck = wn_group_id
	 where id between wn_start and wn_finish;
 end;
$$ ;
 