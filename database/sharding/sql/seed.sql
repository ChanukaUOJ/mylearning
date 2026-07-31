
ALTER TABLE users set UNLOGGED;

INSERT INTO users (first_name, middle_name, last_name, email, address, contact_no)
SELECT
    'User_first_' || i AS first_name,
    'User_middle_' || i AS middle_name,
    'User_last_' || i AS last_name,
    'user_' || i || '@example.com' AS email,
    'user_location_' || i AS address,
    'user_contact_' || i AS contact_no
FROM generate_series(1,10000000) AS i;

ALTER TABLE users set LOGGED;