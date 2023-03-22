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
