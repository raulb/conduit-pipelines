DROP TABLE IF EXISTS employees;
DROP SEQUENCE IF EXISTS employees_id_seq;
CREATE TABLE employees (
    id int NOT NULL,
    name varchar(255),
    full_time bool,
    -- updated_at timestamptz,
    PRIMARY KEY (id)
);
CREATE SEQUENCE employees_id_seq;
ALTER TABLE employees ALTER id SET DEFAULT NEXTVAL('employees_id_seq');
ALTER TABLE employees REPLICA IDENTITY FULL;

-- DROP TABLE IF EXISTS contractors;
-- DROP SEQUENCE IF EXISTS contractors_id_seq;
-- CREATE TABLE contractors (
--     id int NOT NULL,
--     name varchar(255),
--     full_time bool,
--     updated_at timestamptz,
--     PRIMARY KEY (id)
-- );
-- CREATE SEQUENCE contractors_id_seq;
-- ALTER TABLE contractors ALTER id SET DEFAULT NEXTVAL('contractors_id_seq');
-- ALTER TABLE contractors REPLICA IDENTITY FULL;
