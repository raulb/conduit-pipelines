-- Create the employees table
CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    department VARCHAR(50) NOT NULL,
    salary NUMERIC(10, 2) NOT NULL,
    hire_date DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_employees_department ON employees(department);
CREATE INDEX IF NOT EXISTS idx_employees_hire_date ON employees(hire_date);

-- Grant necessary privileges to the postgres user
ALTER USER postgres WITH REPLICATION;

-- Create a publication for logical replication
-- This allows the connector to listen for changes on the employees table
CREATE PUBLICATION conduit_publication FOR TABLE employees;

-- We can't create the replication slot here, as it requires superuser privileges
-- and is typically created dynamically by the connector at runtime.
-- However, we can try to create it if it doesn't exist:

DO $$
BEGIN
    -- Check if the slot already exists
    IF NOT EXISTS (SELECT 1 FROM pg_replication_slots WHERE slot_name = 'conduit_test') THEN
        -- Try to create the replication slot
        BEGIN
            PERFORM pg_create_logical_replication_slot('conduit_test', 'pgoutput');
        EXCEPTION WHEN OTHERS THEN
            RAISE NOTICE 'Could not create replication slot automatically. The connector will create it at runtime if it has sufficient privileges.';
        END;
    END IF;
END $$;


