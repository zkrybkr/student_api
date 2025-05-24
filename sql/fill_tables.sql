INSERT INTO roles (role_name) 
VALUES
    ('teacher'),
    ('student');

INSERT INTO users (name, surname, username, email, roleID) 
VALUES 
    ('Melis', 'Verici', 'meliver', 'vericimels@hotmail.com', (SELECT ID FROM roles WHERE role_name = 'teacher')), 
    ('Ali', 'Kemal', 'ali.kemal', 'akemal01@gmail.com', (SELECT ID FROM roles WHERE role_name = 'teacher')),
    ('Abuzer', 'Kadayıf', 'abukada', 'abuzerkadayif@gmail.com', (SELECT ID FROM roles WHERE role_name = 'student')), 
    ('Meltem', 'Sever', 'melosev', 'melosev@gmail.com', (SELECT ID FROM roles WHERE role_name = 'student')), 
    ('Şakir', 'Çekiç', 'sakircekic', 'sakir.cekic@hotmail.com', (SELECT ID FROM roles WHERE role_name = 'student')), 
    ('Zehra', 'Gök', 'zhrgk', 'zgok@gmail.com', (SELECT ID FROM roles WHERE role_name = 'student')),
    ('Şakir', 'Baltacı', 'sakbal', 'baltas@hotmail.com', (SELECT ID FROM roles WHERE role_name = 'student')),
    ('Cevriye', 'Adalı', 'cev.ada', 'cevriye_adali@yandex.com', (SELECT ID FROM roles WHERE role_name = 'student')),
    ('Şaban', 'Uzan', 'suzan', 'suzan99@gmail.com', (SELECT ID FROM roles WHERE role_name = 'student')),
    ('Meltem', 'Doğru', 'mdogru', 'melo80@gmail.com', (SELECT ID FROM roles WHERE role_name = 'student')),
    ('Ebru', 'Cesur', 'cesurebru', 'ebcesr@yahoo.com', (SELECT ID FROM roles WHERE role_name = 'student'));