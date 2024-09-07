create table if not exists Professor
(
    prof_id    int primary key auto_increment,
    prof_lname varchar(50),
    prof_fname varchar(50)
    );

create table if not exists Course
(
    course_id   int primary key auto_increment,
    course_name varchar(255)
    );


create table if not exists Room
(
    room_id  int primary key auto_increment,
    room_loc varchar(50),
    room_cap varchar(50),
    class_id int
    );

create table if not exists Class
(
    class_id   int primary key auto_increment,
    class_name varchar(255),
    prof_id    int,
    course_id  int,
    room_id    int,
    FOREIGN KEY (prof_id) REFERENCES Professor (prof_id),
    FOREIGN KEY (course_id) REFERENCES Course (course_id),
    FOREIGN KEY (room_id) REFERENCES Room (room_id)
    );

alter table Room
    add FOREIGN KEY (class_id) REFERENCES Class (class_id);


create table if not exists Student
(
    stud_id     int primary key auto_increment,
    stud_fname  varchar(50),
    stud_lname  varchar(50),
    stud_street varchar(255),
    stud_city   varchar(50),
    stud_zip    varchar(10)
    );

create table if not exists Enroll
(
    stud_id  int,
    class_id int,
    grade    varchar(3),
    primary key (stud_id, class_id),
    FOREIGN KEY (stud_id) REFERENCES Student (stud_id),
    FOREIGN KEY (class_id) REFERENCES Class (class_id)
    );