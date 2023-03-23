CREATE DATABASE my_school;

USE my_school;

CREATE TABLE Student (
  StudentID INT NOT NULL AUTO_INCREMENT,
  Name VARCHAR(50) NOT NULL,
  Email VARCHAR(50) NOT NULL,
  PRIMARY KEY (StudentID)
);

CREATE TABLE Course (
  CourseID INT NOT NULL AUTO_INCREMENT,
  Title VARCHAR(50) NOT NULL,
  Description TEXT,
  PRIMARY KEY (CourseID)
);

CREATE TABLE Teacher (
  TeacherID INT NOT NULL AUTO_INCREMENT,
  Name VARCHAR(50) NOT NULL,
  Email VARCHAR(50) NOT NULL,
  PRIMARY KEY (TeacherID)
);

CREATE TABLE Enrollment (
  StudentID INT NOT NULL,
  CourseID INT NOT NULL,
  PRIMARY KEY (StudentID, CourseID),
  FOREIGN KEY (StudentID) REFERENCES Student(StudentID),
  FOREIGN KEY (CourseID) REFERENCES Course(CourseID)
);

CREATE TABLE Teach (
  StudentID INT NOT NULL,
  CourseID INT NOT NULL,
  TeacherID INT NOT NULL,
  PRIMARY KEY (StudentID, CourseID, TeacherID),
  FOREIGN KEY (StudentID, CourseID) REFERENCES Enrollment(StudentID, CourseID),
  FOREIGN KEY (TeacherID) REFERENCES Teacher(TeacherID)
);


-- Insert sample students
INSERT INTO Student (Name, Email) VALUES
('John Doe', 'johndoe@example.com'),
('Jane Smith', 'janesmith@example.com'),
('Bob Johnson', 'bobjohnson@example.com');

-- Insert sample courses
INSERT INTO Course (Title, Description) VALUES
('Introduction to Computer Science', 'An introductory course on computer science.'),
('Data Structures and Algorithms', 'A course on data structures and algorithms.'),
('Database Systems', 'A course on database design and management.');

-- Insert sample teachers
INSERT INTO Teacher (Name, Email) VALUES
('Prof. Smith', 'smith@example.com'),
('Prof. Lee', 'lee@example.com'),
('Prof. Kim', 'kim@example.com');

-- Insert sample enrollments
INSERT INTO Enrollment (StudentID, CourseID) VALUES
(1, 1),
(1, 2),
(2, 1),
(2, 3),
(3, 2),
(3, 3);

-- Insert sample teach assignments
INSERT INTO Teach (StudentID, CourseID, TeacherID) VALUES
(1, 1, 1),
(1, 2, 2),
(2, 1, 1),
(2, 3, 3),
(3, 2, 2),
(3, 3, 3);
