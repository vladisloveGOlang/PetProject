CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    Email VARCHAR(255) NOT NULL,
    Password VARCHAR(255) NOT NULL,
    DeletedAt TIMESTAMP WITH TIME ZONE,
    CreatedAt TIMESTAMP WITH TIME ZONE,
    UpdatedAt TIMESTAMP WITH TIME ZONE
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
   NEW.UpdatedAt = NOW();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_Updated_at
BEFORE UPDATE 
ON users
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE OR REPLACE FUNCTION create_timestamp()
RETURNS TRIGGER AS $$
BEGIN
   NEW.CreatedAt = NOW();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER create_CreatedAt
BEFORE INSERT 
ON users
FOR EACH ROW
EXECUTE FUNCTION create_timestamp();

CREATE TRIGGER create_Updated_at
BEFORE INSERT 
ON users
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();