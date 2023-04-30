CREATE DATABASE ecommerce;

USE ecommerce;

-- Creating the offices table
CREATE TABLE offices (
  officeCode INT(11) NOT NULL,
  city VARCHAR(50) NOT NULL,
  phone VARCHAR(50) NOT NULL,
  PRIMARY KEY (officeCode)
);

-- Creating the employees table
CREATE TABLE employees (
  employeeNumber INT(11) NOT NULL,
  lastName VARCHAR(50) NOT NULL,
  firstName VARCHAR(50) NOT NULL,
  officeCode INT(11) NOT NULL,
  reportsTo INT(11),
  PRIMARY KEY (employeeNumber),
  FOREIGN KEY (officeCode) REFERENCES offices(officeCode),
  FOREIGN KEY (reportsTo) REFERENCES employees(employeeNumber)
);

-- Creating the customers table
CREATE TABLE customers (
  customerNumber INT(11) NOT NULL,
  customerName VARCHAR(50) NOT NULL,
  phone VARCHAR(50) NOT NULL,
  salesRepEmployeeNumber INT(11),
  PRIMARY KEY (customerNumber),
  FOREIGN KEY (salesRepEmployeeNumber) REFERENCES employees(employeeNumber)
);

-- Creating the payments table
CREATE TABLE payments (
  customerNumber INT(11) NOT NULL,
  checkNumber VARCHAR(50) NOT NULL,
  paymentDate DATE NOT NULL,
  amount DECIMAL(10,2) NOT NULL,
  PRIMARY KEY (customerNumber, checkNumber),
  FOREIGN KEY (customerNumber) REFERENCES customers(customerNumber)
);

-- Creating the orders table
CREATE TABLE orders (
  orderNumber INT(11) NOT NULL,
  orderDate DATE NOT NULL,
  customerNumber INT(11) NOT NULL,
  PRIMARY KEY (orderNumber),
  FOREIGN KEY (customerNumber) REFERENCES customers(customerNumber)
);

-- Creating the products table
CREATE TABLE products (
  productCode VARCHAR(15) NOT NULL,
  productName VARCHAR(70) NOT NULL,
  productDescription TEXT NOT NULL,
  quantityInStock SMALLINT(6) NOT NULL,
  PRIMARY KEY (productCode)
);

-- Creating the orderdetails table
CREATE TABLE orderdetails (
  orderNumber INT(11) NOT NULL,
  productCode VARCHAR(15) NOT NULL,
  quantityOrdered INT(11) NOT NULL,
  priceEach DECIMAL(10,2) NOT NULL,
  PRIMARY KEY (orderNumber, productCode),
  FOREIGN KEY (orderNumber) REFERENCES orders(orderNumber),
  FOREIGN KEY (productCode) REFERENCES products(productCode)
);


CREATE TABLE `staff` (
  employeeNumber INT(11) NOT NULL,
  shift VARCHAR(50) NOT NULL,
  PRIMARY KEY (employeeNumber),
  FOREIGN KEY (employeeNumber) REFERENCES employees(employeeNumber)
);

CREATE TABLE `manager` (
  employeeNumber INT(11) NOT NULL,
  department VARCHAR(50) NOT NULL,
  PRIMARY KEY (employeeNumber),
  FOREIGN KEY (employeeNumber) REFERENCES employees(employeeNumber)
);

INSERT INTO offices (officeCode, city, phone) VALUES
(1, 'San Francisco', '+1 650 219 4782'),
(2, 'Boston', '+1 617 555 1212'),
(3, 'Paris', '+33 14 723 4404'),
(4, 'Tokyo', '+81 33 224 5000');


INSERT INTO employees (employeeNumber, lastName, firstName, officeCode, reportsTo) VALUES
(1501, 'Bott', 'Larry', 1, NULL),
(1002, 'Murphy', 'Diane', 1, 1501),
(1056, 'Patterson', 'Mary', 1, 1501),
(1076, 'Firrelli', 'Jeff', 1, 1501),
(1102, 'Rizky', 'Anggita', 1, 1501),
(1143, 'Bow', 'Anthony', 1, 1501),
(1165, 'Jennings', 'Leslie', 1, 1501),
(1166, 'Thompson', 'Steve', 2, 1056),
(1188, 'Firrelli', 'Julie', 2, 1143),
(1216, 'Patterson', 'Steve', 2, 1165),
(1286, 'Tseng', 'Foon Yue', 3, 1143),
(1323, 'Vanauf', 'George', 3, 1102),
(1337, 'Bondur', 'Loui', 4, 1102),
(1370, 'Hernandez', 'Gerard', 4, 1102),
(1401, 'Castillo', 'Pamela', 4, 1002),
(1504, 'Jones', 'Barry', 4, 1002);

INSERT INTO `manager` (employeeNumber, department) VALUES
(1002, 'Sales'),
(1056, 'Marketing'),
(1076, 'Operations'),
(1102, 'Human Resources'),
(1143, 'IT'),
(1165, 'Shipping');


INSERT INTO staff (employeeNumber, shift) VALUES
  (1166, 'Morning'),
  (1188, 'Evening'),
  (1216, 'Morning'),
  (1286, 'Afternoon'),
  (1323, 'Evening'),
  (1337, 'Morning'),
  (1370, 'Afternoon'),
  (1401, 'Evening'),
  (1504, 'Morning');


INSERT INTO customers (customerNumber, customerName, phone, salesRepEmployeeNumber) VALUES
(103, 'Atelier graphique', '40.32.2555', 1401),
(112, 'Signal Gift Stores', '7025551838', 1401),
(114, 'Australian Collectors, Co.', '03 9520 4555', 1401),
(119, 'La Rochelle Gifts', '40.67.8555', 1401),
(121, 'Baane Mini Imports', '07-98 9555', 1401),
(124, 'Mini Gifts Distributors Ltd.', '4155551450', 1504),
(125, 'Havel & Zbyszek Co', '+48 22 555 55 55', 1504),
(128, 'Blauer See Auto, Co.', '6155558276', 1504),
(129, 'Mini Wheels Co.', '6505555787', 1504),
(131, 'Land of Toys Inc.', '2125557818', 1504);

-- Populating the payments table
INSERT INTO payments (customerNumber, checkNumber, paymentDate, amount)
VALUES (103, 'HQ336336', '2004-10-19', 6066.78),
(103, 'JM555205', '2003-06-05', 14571.44),
(112, 'BO864823', '2004-12-18', 14191.12),
(112, 'HQ55022', '2003-06-06', 32641.98),
(114, 'GG31455', '2004-08-06', 45864.03),
(114, 'MA765515', '2004-12-15', 82261.22),
(119, 'DB933704', '2004-11-03', 7565.08),
(119, 'LN373447', '2004-06-08', 7612.06),
(121, 'JN558013', '2003-12-09', 11044.3),
(121, 'OM314933', '2004-12-14', 16700.47);

-- Populating the orders table
INSERT INTO orders (orderNumber, orderDate, customerNumber)
VALUES
(10101, '2003-01-09', 103),
(10102, '2003-01-10', 112),
(10103, '2003-01-29', 112),
(10104, '2003-01-31', 114),
(10105, '2003-02-11', 114),
(10106, '2003-02-17', 121),
(10107, '2003-02-24', 121),
(10108, '2003-03-03', 103),
(10109, '2003-03-10', 119),
(10110, '2003-04-10', 119);

INSERT INTO products (productCode, productName, productDescription, quantityInStock)
VALUES
('S10_1678', '1969 Harley Davidson Ultimate Chopper', 'This replica features working kickstand, front suspension, gear-shift lever', 7933),
('S10_1949', '1952 Alpine Renault 1300', 'This 1:10 scale model replicates the car introduced in 1952.', 7305),
('S10_2016', '1996 Moto Guzzi 1100i', 'Official Moto Guzzi logos and insignias, saddle bags located on side fork', 6625),
('S10_4698', '2003 Harley-Davidson Eagle Drag Bike', 'Model features, official Harley Davidson logos and insignias, detachable rear wheelie bar.', 7209),
('S10_4757', '1972 Alfa Romeo GTA', 'This 1:10 scale model features all opening doors, trunk, hood and detailed engine.', 3252),
('S10_4962', '1962 LanciaA Delta 16V', 'Features opening hood, opening doors, opening trunk, wide fenders', 6791),
('S12_1099', '1968 Ford Mustang', 'Hood, doors and trunk all open to reveal highly detailed engine and interior', 68),
('S12_1108', '2001 Ferrari Enzo', 'Features opening hood, opening doors, opening trunk, detailed engine', 73),
('S12_1666', '1958 Setra Bus', 'Model features 30 windows, skylights & glare resistant glass, working steering system', 1579),
('S12_2823', '2002 Suzuki XREO', 'Features realistic head and tail lights, detailed interior, opening doors, and trunk', 9997);

INSERT INTO orderdetails (orderNumber, productCode, quantityOrdered, priceEach)
VALUES
(10101, 'S10_1678', 20, 81.35),
(10101, 'S10_1949', 35, 76.56),
(10102, 'S10_2016', 20, 64.64),
(10102, 'S10_4698', 22, 95.15),
(10103, 'S12_1099', 37, 50.52),
(10103, 'S10_1678', 30, 67.23),
(10104, 'S10_4698', 42, 84.76),
(10104, 'S12_1108', 20, 98.04),
(10105, 'S12_1108', 20, 138.51),
(10105, 'S10_4962', 38, 32.53),
(10106, 'S10_4962', 42, 44.22),
(10106, 'S10_1678', 22, 65.75),
(10107, 'S10_4962', 41, 77.05),
(10107, 'S10_1678', 23, 99.68),
(10108, 'S12_1099', 42, 57.06),
(10108, 'S12_1666', 25, 92.03),
(10109, 'S12_1666', 34, 62.24),
(10109, 'S12_1099', 20, 94.07),
(10110, 'S12_2823', 47, 53.32),
(10110, 'S10_4962', 33, 63.67);
