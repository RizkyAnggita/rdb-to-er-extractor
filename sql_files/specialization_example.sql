-- Create the database
CREATE DATABASE company;

-- Use the company database
USE company;

-- Create the parent table
CREATE TABLE Employee (
  SSN INT PRIMARY KEY,
  Name VARCHAR(255),
  Birthdate DATE,
  Gender ENUM('Male', 'Female'),
  Hiredate DATE,
  JobTitle VARCHAR(255),
  Salary DECIMAL(10, 2)
);

-- Create the child tables for specialization
CREATE TABLE Manager (
  SSN INT PRIMARY KEY,
  NumEmployees INT,
  FOREIGN KEY (SSN) REFERENCES Employee(SSN)
) ENGINE=InnoDB;

CREATE TABLE Programmer (
  SSN INT PRIMARY KEY,
  NumProjects INT,
  FOREIGN KEY (SSN) REFERENCES Employee(SSN)
) ENGINE=InnoDB;


INSERT INTO Employee (SSN, Name, Birthdate, Gender, Hiredate, JobTitle, Salary) 
VALUES 
  (111223333, 'John Smith', '1990-01-01', 'Male', '2018-01-01', 'Software Engineer', 80000),
  (222334444, 'Jane Doe', '1995-05-05', 'Female', '2020-01-01', 'Manager', 120000),
  (333445555, 'Bob Johnson', '1985-10-10', 'Male', '2015-01-01', 'Programmer', 90000);

-- Insert data into the "Manager" table
INSERT INTO Manager (SSN, NumEmployees)
VALUES
  (222334444, 10);

-- Insert data into the "Programmer" table
INSERT INTO Programmer (SSN, NumProjects)
VALUES
  (333445555, 5);