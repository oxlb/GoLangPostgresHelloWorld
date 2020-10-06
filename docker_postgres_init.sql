CREATE TABLE students
(
    id bigint NOT NULL,
    name text COLLATE pg_catalog."default",
    CONSTRAINT student_pkey PRIMARY KEY (id)
    
);

INSERT INTO students(id, name) VALUES
 (1, 'A'),
 (2, 'B'),
 (3, 'C');