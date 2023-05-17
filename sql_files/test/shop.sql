CREATE DATABASE shop;

USE shop;

CREATE TABLE offices (
  officeID INT PRIMARY KEY,
  officeName VARCHAR(50) NOT NULL,
  officePhone VARCHAR(150)
);

CREATE TABLE addresses (
  addressID INT,
  officeID INT,
  streetName VARCHAR(100),
  PRIMARY KEY (addressId, officeID),
  FOREIGN KEY (officeID) REFERENCES offices(officeID)
);

CREATE TABLE employees (
  employeeID INT PRIMARY KEY,
  officeID INT,
  employeeName VARCHAR(50),
  FOREIGN KEY (officeID) REFERENCES offices(officeID)
);

CREATE TABLE dependents (
  employeeID INT,
  dependentID INT,
  dependentName VARCHAR(50),
  dependentRelationship VARCHAR(50),
  PRIMARY KEY (employeeID, dependentID),
  FOREIGN KEY (employeeID) REFERENCES employees(employeeID)
);

CREATE TABLE family_members (
  employeeID INT,
  familyMemberName VARCHAR(50),
  PRIMARY KEY (employeeID, familyMemberName),
  FOREIGN KEY (employeeID) REFERENCES employees(employeeID)
);

INSERT INTO offices (officeID, officeName, officePhone)
VALUES
  (1, 'Burmott Office', '123-456-7890, 123-456-7892, 123-456-7893'),
  (2, 'Langston Office', '987-654-3210, 987-654-3211'),
  (3, 'Lancaster Office', '555-123-4567');

INSERT INTO addresses (addressId, officeID, streetName)
VALUES
  (1, 1, '123 Main Street'),
  (2, 2, '456 Langs Avenue'),
  (3, 3, '789 Oak Lane');

INSERT INTO employees (employeeID, officeID, employeeName)
VALUES
  (1, 1, 'John Doe'),
  (2, 1, 'Jane Smith'),
  (3, 2, 'Michael Johnson'),
  (4, 3, 'Emily Davis'),
  (5, 2, 'Rizky Anggita');

INSERT INTO Dependents (employeeID, dependentID, dependentName, dependentRelationship)
VALUES
  (1, 1, 'Mary Doe', 'Spouse'),
  (1, 2, 'Sara Doe', 'Child'),
  (2, 3, 'David Smith', 'Spouse'),
  (2, 4, 'Emily Smith', 'Child'),
  (3, 5, 'Amy Johnson', 'Spouse');

INSERT INTO family_members (employeeID, familyMemberName)
VALUES
  (1, 'Mrs. Jane Doe'),
  (1, 'Mr. Thomas Doe Jr.'),
  (2, 'Mrs. Smith'),
  (3, 'Mrs. Sarah Johnson'),
  (4, 'Mr. Albert Davis Jr.');