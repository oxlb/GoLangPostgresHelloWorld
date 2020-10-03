CREATE TABLE students
(
    id bigint NOT NULL,
    name text COLLATE pg_catalog."default",
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp DEFAULT NULL,
    deleted_at timestamp DEFAULT NULL,
    CONSTRAINT student_pkey PRIMARY KEY (id)
    
);

INSERT INTO students(id, name) VALUES
 (1, 'A'),
 (2, 'B'),
 (3, 'C');