CREATE DISTRIBUTION ds1 COLUMN TYPES integer;
CREATE SHARDING RULE hash1 DISTRIBUTED KEY col1 HASH FUNCTION IDENT FOR DISTRIBUTION ds1;
CREATE SHARDING RULE hash2 DISTRIBUTED KEY col2 HASH FUNCTION IDENTITY FOR DISTRIBUTION ds1;
CREATE SHARDING RULE hash3 DISTRIBUTED KEY col3 HASH FUNCTION MURMUR HASH FOR DISTRIBUTION ds1;
CREATE SHARDING RULE hash4 DISTRIBUTED KEY col4 HASH FUNCTION CITY HASH FOR DISTRIBUTION ds1;

SHOW sharding_rules;

DROP DISTRIBUTION ALL CASCADE;
DROP SHARDING RULE ALL;
DROP KEY RANGE ALL;