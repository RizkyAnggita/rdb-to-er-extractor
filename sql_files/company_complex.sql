CREATE TABLE Person (
    SSN VARCHAR(9) NOT NULL PRIMARY KEY,
    Name VARCHAR(50) NOT NULL,
    Address VARCHAR(100) NOT NULL,
    Phone VARCHAR(15) NOT NULL,
    Email VARCHAR(50) NOT NULL
);

CREATE TABLE Employee (
    SSN VARCHAR(9) NOT NULL PRIMARY KEY,
    Hire_Date DATE NOT NULL,
    Salary DECIMAL(10,2) NOT NULL,
    Department VARCHAR(50) NOT NULL,
    FOREIGN KEY (SSN) REFERENCES Person(SSN)
);

CREATE TABLE Manager (
    SSN VARCHAR(9) NOT NULL PRIMARY KEY,
    Bonus DECIMAL(10,2) NOT NULL,
    FOREIGN KEY (SSN) REFERENCES Employee(SSN)
);

CREATE TABLE Dependent (
    ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(50) NOT NULL,
    Relationship VARCHAR(20) NOT NULL,
    Employee_SSN VARCHAR(9) NOT NULL,
    FOREIGN KEY (Employee_SSN) REFERENCES Employee(SSN)
);

CREATE TABLE Project (
    ID INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(50) NOT NULL,
    Start_Date DATE NOT NULL,
    End_Date DATE NOT NULL,
    Budget DECIMAL(10,2) NOT NULL
);

CREATE TABLE Works_On (
    Employee_SSN VARCHAR(9) NOT NULL,
    Project_ID INT NOT NULL,
    Hours DECIMAL(5,2) NOT NULL,
    PRIMARY KEY (Employee_SSN, Project_ID),
    FOREIGN KEY (Employee_SSN) REFERENCES Employee(SSN),
    FOREIGN KEY (Project_ID) REFERENCES Project(ID)
);

CREATE TABLE Dependents_of_Employees (
    Employee_SSN VARCHAR(9) NOT NULL,
    Dependent_ID INT NOT NULL,
    PRIMARY KEY (Employee_SSN, Dependent_ID),
    FOREIGN KEY (Employee_SSN) REFERENCES Employee(SSN),
    FOREIGN KEY (Dependent_ID) REFERENCES Dependent(ID)
);


INSERT INTO Person (SSN, Name, Address, Phone, Email)
VALUES 
    ('111111111', 'John Doe', '123 Main St.', '555-1234', 'johndoe@example.com'),
    ('222222222', 'Jane Smith', '456 Elm St.', '555-5678', 'janesmith@example.com'),
    ('333333333', 'Bob Johnson', '789 Oak St.', '555-9012', 'bobjohnson@example.com'),
    ('444444444', 'Sara Lee', '321 Pine St.', '555-3456', 'saralee@example.com');

INSERT INTO Employee (SSN, Hire_Date, Salary, Department)
VALUES 
    ('111111111', '2020-01-01', 50000.00, 'Sales'),
    ('222222222', '2018-06-01', 60000.00, 'Marketing'),
    ('333333333', '2019-09-01', 70000.00, 'Engineering'),
    ('444444444', '2021-03-01', 45000.00, 'Customer Support');

INSERT INTO Manager (SSN, Bonus)
VALUES 
    ('333333333', 5000.00),
    ('444444444', 2500.00);

INSERT INTO Dependent (Name, Relationship, Employee_SSN)
VALUES 
    ('Tom', 'Son', '111111111'),
    ('Mary', 'Daughter', '111111111'),
    ('Bob', 'Spouse', '222222222'),
    ('Joe', 'Son', '333333333'),
    ('Sue', 'Daughter', '333333333');

INSERT INTO Project (Name, Start_Date, End_Date, Budget)
VALUES 
    ('Project A', '2022-01-01', '2022-06-30', 100000.00),
    ('Project B', '2022-02-01', '2022-07-31', 50000.00),
    ('Project C', '2022-03-01', '2022-08-31', 75000.00);

INSERT INTO Works_On (Employee_SSN, Project_ID, Hours)
VALUES 
    ('111111111', 1, 80.00),
    ('111111111', 2, 40.00),
    ('222222222', 1, 60.00),
    ('333333333', 3, 100.00),
    ('444444444', 2, 20.00);

INSERT INTO Dependents_of_Employees (Employee_SSN, Dependent_ID)
VALUES 
    ('111111111', 1),
    ('111111111', 2),
    ('222222222', 3),
    ('333333333', 4),
    ('333333333', 5);
