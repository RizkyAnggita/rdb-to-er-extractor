CREATE KEYSPACE ecommerce
WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1};

USE ecommerce;

DROP TABLE IF EXISTS offices;
CREATE TABLE offices (
  phone TEXT,
  city TEXT,
  officeCode INT,
  PRIMARY KEY (officeCode)
);

DROP TABLE IF EXISTS employees;
CREATE TABLE employees (
  firstName TEXT,
  lastName TEXT,
  employeeNumber INT,
  reports_to INT,
  work_at_employees INT,
  PRIMARY KEY (employeeNumber)
);

DROP TABLE IF EXISTS reports_to;
CREATE TABLE reports_to (
  employeeNumber INT,
  employees INT,
  PRIMARY KEY (employeeNumber, employees)
);

DROP TABLE IF EXISTS work_at_employees;
CREATE TABLE work_at_employees (
  officeCode INT,
  employees INT,
  PRIMARY KEY (officeCode, employees)
);

DROP TABLE IF EXISTS staff;
CREATE TABLE staff (
  employeeNumber INT,
  shift TEXT,
  PRIMARY KEY (employeeNumber)
);

DROP TABLE IF EXISTS manager;
CREATE TABLE manager (
  employeeNumber INT,
  department TEXT,
  PRIMARY KEY (employeeNumber)
);

DROP TABLE IF EXISTS customers;
CREATE TABLE customers (
  customer_phones LIST<TEXT>,
  customerName TEXT,
  customerNumber INT,
  managed_customers INT,
  PRIMARY KEY (customerNumber)
);

DROP TABLE IF EXISTS managed_customers;
CREATE TABLE managed_customers (
  employeeNumber INT,
  customers INT,
  PRIMARY KEY (employeeNumber, customers)
);

DROP TABLE IF EXISTS orders;
CREATE TABLE orders (
  orderDate DATE,
  orderNumber INT,
  have_orders INT,
  PRIMARY KEY (orderNumber)
);

DROP TABLE IF EXISTS have_orders;
CREATE TABLE have_orders (
  customerNumber INT,
  orders INT,
  PRIMARY KEY (customerNumber, orders)
);

DROP TABLE IF EXISTS products;
CREATE TABLE products (
  quantityInStock INT,
  productDescription TEXT,
  productName TEXT,
  productCode TEXT,
  PRIMARY KEY (productCode)
);

DROP TABLE IF EXISTS orderdetails_orders;
CREATE TABLE orderdetails_orders (
  productCode TEXT,
  priceEach DECIMAL,
  quantityOrdered INT,
  orders INT,
  PRIMARY KEY (productCode, orders)
);

DROP TABLE IF EXISTS orderdetails_products;
CREATE TABLE orderdetails_products (
  orderNumber INT,
  priceEach DECIMAL,
  quantityOrdered INT,
  products TEXT,
  PRIMARY KEY (orderNumber, products)
);


DROP TABLE IF EXISTS payments;
CREATE TABLE payments (
  customerNumber INT,
  amount DECIMAL,
  paymentDate DATE,
  checkNumber TEXT,
  PRIMARY KEY (customerNumber, checkNumber)
);


INSERT INTO offices (officecode, city, phone) VALUES (1, 'San Francisco', '+1 650 219 4782');
INSERT INTO offices (officecode, city, phone) VALUES (2, 'Boston', '+1 617 555 1212');
INSERT INTO offices (officecode, city, phone) VALUES (3, 'Paris', '+33 14 723 4404');
INSERT INTO offices (officecode, city, phone) VALUES (4, 'Tokyo', '+81 33 224 5000');

INSERT INTO employees (employeenumber, lastname, firstname, work_at_employees, reports_to) VALUES (1002, 'Murphy', 'Diane', 1, 1501);
INSERT INTO employees (employeenumber, lastname, firstname, work_at_employees, reports_to) VALUES (1056, 'Patterson', 'Mary', 1, 1501);
INSERT INTO employees (employeenumber, lastname, firstname, work_at_employees, reports_to) VALUES (1076, 'Firrelli', 'Jeff', 1, 1501);
INSERT INTO employees (employeenumber, lastname, firstname, work_at_employees, reports_to) VALUES (1102, 'Rizky', 'Anggita', 1, 1501);
INSERT INTO employees (employeenumber, lastname, firstname, work_at_employees, reports_to) VALUES (1143, 'Bow', 'Anthony', 1, 1501);
INSERT INTO employees (employeenumber, lastname, firstname, work_at_employees, reports_to) VALUES (1165, 'Jennings', 'Leslie', 1, 1501);
INSERT INTO employees (employeenumber, lastname, firstname, work_at_employees, reports_to) VALUES (1166, 'Thompson', 'Steve', 2, 1056);
INSERT INTO employees (employeenumber, lastname, firstname, work_at_employees, reports_to) VALUES (1188, 'Firrelli', 'Julie', 2, 1143);
INSERT INTO employees (employeenumber, lastname, firstname, work_at_employees, reports_to) VALUES (1216, 'Patterson', 'Steve', 2, 1165);
INSERT INTO employees (employeenumber, lastname, firstname, work_at_employees, reports_to) VALUES (1286, 'Tseng', 'Foon Yue', 3, 1143);
INSERT INTO employees (employeenumber, lastname, firstname, work_at_employees, reports_to) VALUES (1323, 'Vanauf', 'George', 3, 1102);
INSERT INTO employees (employeenumber, lastname, firstname, work_at_employees, reports_to) VALUES (1337, 'Bondur', 'Loui', 4, 1102);
INSERT INTO employees (employeenumber, lastname, firstname, work_at_employees, reports_to) VALUES (1370, 'Hernandez', 'Gerard', 4, 1102);
INSERT INTO employees (employeenumber, lastname, firstname, work_at_employees, reports_to) VALUES (1401, 'Castillo', 'Pamela', 4, 1002);
INSERT INTO employees (employeenumber, lastname, firstname, work_at_employees, reports_to) VALUES (1501, 'Bott', 'Larry', 1, NULL);
INSERT INTO employees (employeenumber, lastname, firstname, work_at_employees, reports_to) VALUES (1504, 'Jones', 'Barry', 4, 1002);

INSERT INTO reports_to (employeenumber, employees) values(1401,	1002);
INSERT INTO reports_to (employeenumber, employees) values(1504,	1002);
INSERT INTO reports_to (employeenumber, employees) values(1166,	1056);
INSERT INTO reports_to (employeenumber, employees) values(1323,	1102);
INSERT INTO reports_to (employeenumber, employees) values(1337,	1102);
INSERT INTO reports_to (employeenumber, employees) values(1370,	1102);
INSERT INTO reports_to (employeenumber, employees) values(1188,	1143);
INSERT INTO reports_to (employeenumber, employees) values(1286,	1143);
INSERT INTO reports_to (employeenumber, employees) values(1216,	1165);
INSERT INTO reports_to (employeenumber, employees) values(1002,	1501);
INSERT INTO reports_to (employeenumber, employees) values(1056,	1501);
INSERT INTO reports_to (employeenumber, employees) values(1076,	1501);
INSERT INTO reports_to (employeenumber, employees) values(1102,	1501);
INSERT INTO reports_to (employeenumber, employees) values(1143,	1501);
INSERT INTO reports_to (employeenumber, employees) values(1165,	1501);

INSERT INTO work_at_employees (officecode, employees) VALUES(1,	1002);
INSERT INTO work_at_employees (officecode, employees) VALUES(1,	1056);
INSERT INTO work_at_employees (officecode, employees) VALUES(1,	1076);
INSERT INTO work_at_employees (officecode, employees) VALUES(1,	1102);
INSERT INTO work_at_employees (officecode, employees) VALUES(1,	1143);
INSERT INTO work_at_employees (officecode, employees) VALUES(1,	1165);
INSERT INTO work_at_employees (officecode, employees) VALUES(1,	1501);
INSERT INTO work_at_employees (officecode, employees) VALUES(2,	1166);
INSERT INTO work_at_employees (officecode, employees) VALUES(2,	1188);
INSERT INTO work_at_employees (officecode, employees) VALUES(2,	1216);
INSERT INTO work_at_employees (officecode, employees) VALUES(3,	1286);
INSERT INTO work_at_employees (officecode, employees) VALUES(3,	1323);
INSERT INTO work_at_employees (officecode, employees) VALUES(4,	1337);
INSERT INTO work_at_employees (officecode, employees) VALUES(4,	1370);
INSERT INTO work_at_employees (officecode, employees) VALUES(4,	1401);
INSERT INTO work_at_employees (officecode, employees) VALUES(4,	1504);

INSERT INTO manager (employeenumber, department) VALUES (1002, 'Sales');
INSERT INTO manager (employeenumber, department) VALUES (1056, 'Marketing');
INSERT INTO manager (employeenumber, department) VALUES (1076, 'Operations');
INSERT INTO manager (employeenumber, department) VALUES (1102, 'Human Resources');
INSERT INTO manager (employeenumber, department) VALUES (1143, 'IT');
INSERT INTO manager (employeenumber, department) VALUES (1165, 'Shipping');

INSERT INTO staff (employeenumber, shift) VALUES (1166,	'Morning');
INSERT INTO staff (employeenumber, shift) VALUES (1188,	'Evening');
INSERT INTO staff (employeenumber, shift) VALUES (1216,	'Morning');
INSERT INTO staff (employeenumber, shift) VALUES (1286,	'Afternoon');
INSERT INTO staff (employeenumber, shift) VALUES (1323,	'Evening');
INSERT INTO staff (employeenumber, shift) VALUES (1337,	'Morning');
INSERT INTO staff (employeenumber, shift) VALUES (1370,	'Afternoon');
INSERT INTO staff (employeenumber, shift) VALUES (1401,	'Evening');
INSERT INTO staff (employeenumber, shift) VALUES (1504,	'Morning');


INSERT INTO customers (customernumber, customername, customer_phones, managed_customers) VALUES (103, 'Atelier graphique', ['123-456-7890', '987-654-3210'], 1401);
INSERT INTO customers (customernumber, customername, customer_phones, managed_customers) VALUES (112, 'Signal Gift Stores', ['111-111-1111', '123-456-7890', '555-555-5555'], 1401);
INSERT INTO customers (customernumber, customername, customer_phones, managed_customers) VALUES (114, 'Australian Collectors Co.', ['555-555-5555'], 1401);
INSERT INTO customers (customernumber, customername, customer_phones, managed_customers) VALUES (119, 'La Rochelle Gifts', ['555-555-5555'], 1401);
INSERT INTO customers (customernumber, customername, customer_phones, managed_customers) VALUES (121, 'Baane Mini Imports', ['123-456-7890', '987-654-3210'], 1401);
INSERT INTO customers (customernumber, customername, customer_phones, managed_customers) VALUES (124, 'Mini Gifts Distributors Ltd.', ['11-111-1111'], 1504);
INSERT INTO customers (customernumber, customername, customer_phones, managed_customers) VALUES (125, 'Havel & Zbyszek Co', ['555-555-5555'], 1504);
INSERT INTO customers (customernumber, customername, customer_phones, managed_customers) VALUES (128, 'Blauer See Auto, Co.', ['555-555-5555'], 1504);
INSERT INTO customers (customernumber, customername, customer_phones, managed_customers) VALUES (129, 'Mini Wheels Co.', ['123-456-7890', '987-654-3210'], 1504);
INSERT INTO customers (customernumber, customername, customer_phones, managed_customers) VALUES (131, 'Land of Toys Inc.', ['111-111-1111', '123-456-7890'], 1504);

INSERT INTO managed_customers (employeenumber, customers) VALUES (1401, 103);
INSERT INTO managed_customers (employeenumber, customers) VALUES (1401, 112);
INSERT INTO managed_customers (employeenumber, customers) VALUES (1401, 114);
INSERT INTO managed_customers (employeenumber, customers) VALUES (1401, 119);
INSERT INTO managed_customers (employeenumber, customers) VALUES (1401, 121);
INSERT INTO managed_customers (employeenumber, customers) VALUES (1504, 124);
INSERT INTO managed_customers (employeenumber, customers) VALUES (1504, 125);
INSERT INTO managed_customers (employeenumber, customers) VALUES (1504, 128);
INSERT INTO managed_customers (employeenumber, customers) VALUES (1504, 129);
INSERT INTO managed_customers (employeenumber, customers) VALUES (1504, 131);

INSERT INTO orders (ordernumber, orderdate, have_orders) VALUES (10101, '2003-01-09', 103);
INSERT INTO orders (ordernumber, orderdate, have_orders) VALUES (10102, '2003-01-10', 112);
INSERT INTO orders (ordernumber, orderdate, have_orders) VALUES (10103, '2003-01-29', 112);
INSERT INTO orders (ordernumber, orderdate, have_orders) VALUES (10104, '2003-01-31', 114);
INSERT INTO orders (ordernumber, orderdate, have_orders) VALUES (10105, '2003-02-11', 114);
INSERT INTO orders (ordernumber, orderdate, have_orders) VALUES (10106, '2003-02-17', 121);
INSERT INTO orders (ordernumber, orderdate, have_orders) VALUES (10107, '2003-02-24', 121);
INSERT INTO orders (ordernumber, orderdate, have_orders) VALUES (10108, '2003-03-03', 103);
INSERT INTO orders (ordernumber, orderdate, have_orders) VALUES (10109, '2003-03-10', 119);
INSERT INTO orders (ordernumber, orderdate, have_orders) VALUES (10110, '2003-04-10', 119);

INSERT INTO have_orders (customernumber, orders) VALUES (103, 10101);
INSERT INTO have_orders (customernumber, orders) VALUES (103, 10108);
INSERT INTO have_orders (customernumber, orders) VALUES (112, 10102);
INSERT INTO have_orders (customernumber, orders) VALUES (112, 10103);
INSERT INTO have_orders (customernumber, orders) VALUES (114, 10104);
INSERT INTO have_orders (customernumber, orders) VALUES (114, 10105);
INSERT INTO have_orders (customernumber, orders) VALUES (119, 10109);
INSERT INTO have_orders (customernumber, orders) VALUES (119, 10110);
INSERT INTO have_orders (customernumber, orders) VALUES (121, 10106);
INSERT INTO have_orders (customernumber, orders) VALUES (121, 10107);

INSERT INTO products (productcode, productname, productdescription, quantityinstock) VALUES ('S10_1678', '1969 Harley Davidson Ultimate Chopper', 'This replica features working kickstand, front suspension, gear-shift lever', 7933);
INSERT INTO products (productcode, productname, productdescription, quantityinstock) VALUES ('S10_1949', '1952 Alpine Renault 1300', 'This 1:10 scale model replicates the car introduced in 1952.', 7305);
INSERT INTO products (productcode, productname, productdescription, quantityinstock) VALUES ('S10_2016', '1996 Moto Guzzi 1100i', 'Official Moto Guzzi logos and insignias, saddle bags located on side fork', 6625);
INSERT INTO products (productcode, productname, productdescription, quantityinstock) VALUES ('S10_4698', '2003 Harley-Davidson Eagle Drag Bike', 'Model features, official Harley Davidson logos and insignias, detachable rear wheelie bar.', 7209);
INSERT INTO products (productcode, productname, productdescription, quantityinstock) VALUES ('S10_4757', '1972 Alfa Romeo GTA', 'This 1:10 scale model features all opening doors, trunk, hood and detailed engine.', 3252);
INSERT INTO products (productcode, productname, productdescription, quantityinstock) VALUES ('S10_4962', '1962 LanciaA Delta 16V', 'Features opening hood, opening doors, opening trunk, wide fenders', 6791);
INSERT INTO products (productcode, productname, productdescription, quantityinstock) VALUES ('S12_1099', '1968 Ford Mustang', 'Hood, doors and trunk all open to reveal highly detailed engine and interior', 68);
INSERT INTO products (productcode, productname, productdescription, quantityinstock) VALUES ('S12_1108', '2001 Ferrari Enzo', 'Features opening hood, opening doors, opening trunk, detailed engine', 73);
INSERT INTO products (productcode, productname, productdescription, quantityinstock) VALUES ('S12_1666', '1958 Setra Bus', 'Model features 30 windows, skylights & glare resistant glass, working steering system', 1579);
INSERT INTO products (productcode, productname, productdescription, quantityinstock) VALUES ('S12_2823', '2002 Suzuki XREO', 'Features realistic head and tail lights, detailed interior, opening doors, and trunk', 9997);

INSERT INTO orderdetails_orders (orders, productcode, quantityordered, priceeach) VALUES (10101, 'S10_1678',20,	81.35);
INSERT INTO orderdetails_orders (orders, productcode, quantityordered, priceeach) VALUES (10101, 'S10_1949', 35, 76.56);
INSERT INTO orderdetails_orders (orders, productcode, quantityordered, priceeach) VALUES (10102, 'S10_2016', 20, 64.64);
INSERT INTO orderdetails_orders (orders, productcode, quantityordered, priceeach) VALUES (10102, 'S10_4698', 22, 95.15);
INSERT INTO orderdetails_orders (orders, productcode, quantityordered, priceeach) VALUES (10103, 'S10_1678', 30, 67.23);
INSERT INTO orderdetails_orders (orders, productcode, quantityordered, priceeach) VALUES (10103, 'S12_1099', 37, 50.52);
INSERT INTO orderdetails_orders (orders, productcode, quantityordered, priceeach) VALUES (10104, 'S10_4698', 42, 84.76);
INSERT INTO orderdetails_orders (orders, productcode, quantityordered, priceeach) VALUES (10104, 'S12_1108', 20, 98.04);
INSERT INTO orderdetails_orders (orders, productcode, quantityordered, priceeach) VALUES (10105, 'S10_4962', 38, 32.53);
INSERT INTO orderdetails_orders (orders, productcode, quantityordered, priceeach) VALUES (10105, 'S12_1108', 20, 138.51);
INSERT INTO orderdetails_orders (orders, productcode, quantityordered, priceeach) VALUES (10106, 'S10_1678', 22, 65.75);
INSERT INTO orderdetails_orders (orders, productcode, quantityordered, priceeach) VALUES (10106, 'S10_4962', 42, 44.22);
INSERT INTO orderdetails_orders (orders, productcode, quantityordered, priceeach) VALUES (10107, 'S10_1678', 23, 99.68);
INSERT INTO orderdetails_orders (orders, productcode, quantityordered, priceeach) VALUES (10107, 'S10_4962', 41, 77.05);
INSERT INTO orderdetails_orders (orders, productcode, quantityordered, priceeach) VALUES (10108, 'S12_1099', 42, 57.06);
INSERT INTO orderdetails_orders (orders, productcode, quantityordered, priceeach) VALUES (10108, 'S12_1666', 25, 92.03);
INSERT INTO orderdetails_orders (orders, productcode, quantityordered, priceeach) VALUES (10109, 'S12_1099', 20, 94.07);
INSERT INTO orderdetails_orders (orders, productcode, quantityordered, priceeach) VALUES (10109, 'S12_1666', 34, 62.24);
INSERT INTO orderdetails_orders (orders, productcode, quantityordered, priceeach) VALUES (10110, 'S10_4962', 33, 63.67);
INSERT INTO orderdetails_orders (orders, productcode, quantityordered, priceeach) VALUES (10110, 'S12_2823', 47, 53.32);

INSERT INTO orderdetails_products (ordernumber, products, quantityordered, priceeach) VALUES (10101, 'S10_1678',20,	81.35);
INSERT INTO orderdetails_products (ordernumber, products, quantityordered, priceeach) VALUES (10101, 'S10_1949', 35, 76.56);
INSERT INTO orderdetails_products (ordernumber, products, quantityordered, priceeach) VALUES (10102, 'S10_2016', 20, 64.64);
INSERT INTO orderdetails_products (ordernumber, products, quantityordered, priceeach) VALUES (10102, 'S10_4698', 22, 95.15);
INSERT INTO orderdetails_products (ordernumber, products, quantityordered, priceeach) VALUES (10103, 'S10_1678', 30, 67.23);
INSERT INTO orderdetails_products (ordernumber, products, quantityordered, priceeach) VALUES (10103, 'S12_1099', 37, 50.52);
INSERT INTO orderdetails_products (ordernumber, products, quantityordered, priceeach) VALUES (10104, 'S10_4698', 42, 84.76);
INSERT INTO orderdetails_products (ordernumber, products, quantityordered, priceeach) VALUES (10104, 'S12_1108', 20, 98.04);
INSERT INTO orderdetails_products (ordernumber, products, quantityordered, priceeach) VALUES (10105, 'S10_4962', 38, 32.53);
INSERT INTO orderdetails_products (ordernumber, products, quantityordered, priceeach) VALUES (10105, 'S12_1108', 20, 138.51);
INSERT INTO orderdetails_products (ordernumber, products, quantityordered, priceeach) VALUES (10106, 'S10_1678', 22, 65.75);
INSERT INTO orderdetails_products (ordernumber, products, quantityordered, priceeach) VALUES (10106, 'S10_4962', 42, 44.22);
INSERT INTO orderdetails_products (ordernumber, products, quantityordered, priceeach) VALUES (10107, 'S10_1678', 23, 99.68);
INSERT INTO orderdetails_products (ordernumber, products, quantityordered, priceeach) VALUES (10107, 'S10_4962', 41, 77.05);
INSERT INTO orderdetails_products (ordernumber, products, quantityordered, priceeach) VALUES (10108, 'S12_1099', 42, 57.06);
INSERT INTO orderdetails_products (ordernumber, products, quantityordered, priceeach) VALUES (10108, 'S12_1666', 25, 92.03);
INSERT INTO orderdetails_products (ordernumber, products, quantityordered, priceeach) VALUES (10109, 'S12_1099', 20, 94.07);
INSERT INTO orderdetails_products (ordernumber, products, quantityordered, priceeach) VALUES (10109, 'S12_1666', 34, 62.24);
INSERT INTO orderdetails_products (ordernumber, products, quantityordered, priceeach) VALUES (10110, 'S10_4962', 33, 63.67);
INSERT INTO orderdetails_products (ordernumber, products, quantityordered, priceeach) VALUES (10110, 'S12_2823', 47, 53.32);

INSERT INTO payments (customernumber, checknumber, paymentdate, amount) VALUES (103, 'HQ336336', '2004-10-19', 6066.78);
INSERT INTO payments (customernumber, checknumber, paymentdate, amount) VALUES (103, 'JM555205', '2003-06-05', 14571.44);
INSERT INTO payments (customernumber, checknumber, paymentdate, amount) VALUES (112, 'BO864823', '2004-12-18', 14191.12);
INSERT INTO payments (customernumber, checknumber, paymentdate, amount) VALUES (112, 'HQ55022', '2003-06-06', 32641.98);
INSERT INTO payments (customernumber, checknumber, paymentdate, amount) VALUES (114, 'GG31455', '2004-08-06', 45864.03);
INSERT INTO payments (customernumber, checknumber, paymentdate, amount) VALUES (114, 'MA765515', '2004-12-15', 82261.22);
INSERT INTO payments (customernumber, checknumber, paymentdate, amount) VALUES (119, 'DB933704', '2004-11-03', 7565.08);
INSERT INTO payments (customernumber, checknumber, paymentdate, amount) VALUES (119, 'LN373447', '2004-06-08', 7612.06);
INSERT INTO payments (customernumber, checknumber, paymentdate, amount) VALUES (121, 'JN558013', '2003-12-09', 11044.30);
INSERT INTO payments (customernumber, checknumber, paymentdate, amount) VALUES (121, 'OM314933', '2004-12-14', 16700.47);