INSERT INTO roles (role_name) 
VALUES (
    ('teacher'),
    ('student')
);

INSERT INTO users (name, surname, username, email, role_id) 
VALUES (
    ('Melis', 'Verici', 'meliver', 'vericimels@hotmail.com', 1), 
    ('Abuzer', 'Kadayıf', 'abukada', 'abuzerkadayif@gmail.com', 2), 
    ('Meltem', 'Sever', 'melosev', 'melosev@gmail.com', 2), 
    ('Şakir', 'Çekiç', 'sakircekic', 'sakir.cekic@hotmail.com', 2), 
    ('Ebru', 'Cesur', 'cesurebru', 'ebcesr@yahoo.com', 2)
);