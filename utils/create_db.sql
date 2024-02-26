CREATE ROLE dglabeluser WITH LOGIN PASSWORD 'dglabelpass';


CREATE DATABASE dglabeldb
    WITH
    OWNER = dglabeluser
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;