CREATE OR REPLACE FUNCTION seed_foos_and_bars()
  RETURNS integer AS
$BODY$
DECLARE
   f_id integer;
BEGIN
   INSERT INTO foos(name)
   VALUES ('first')
   RETURNING id INTO f_id;

   -- Inserts the new client and references the inserted person
   INSERT INTO bars(foo_id, value) VALUES (f_id, 10);
   INSERT INTO bars(foo_id, value) VALUES (f_id, 11);

    RETURN f_id;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE;

SELECT * FROM seed_foos_and_bars();
