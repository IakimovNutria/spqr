SHOW key_ranges;
SHOW sharding_rules;
DROP KEY RANGE ALL;
-- DROP SHARDING RULE ALL; does not work
DROP SHARDING RULE rule1;
SHOW key_ranges;
SHOW sharding_rules;
ADD KEY RANGE krid1 FROM 1 TO 10 ROUTE TO sh1;
ADD KEY RANGE krid2 FROM 11 TO 20 ROUTE TO sh1;
ADD SHARDING RULE rule1 COLUMNS id;
SHOW key_ranges;
SHOW sharding_rules;
DROP SHARDING RULE rule1;
ADD SHARDING RULE r1 COLUMNS w_id;
ADD KEY RANGE krid1 FROM 1 TO 10 ROUTE TO sh1;
ADD KEY RANGE krid2 FROM 11 TO 20 ROUTE TO sh2;