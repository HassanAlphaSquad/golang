CREATE TABLE users(
   ID  SERIAL PRIMARY KEY,
   NAME           TEXT      NOT NULL,
   AGE            INT       NOT NULL,
   ADDRESS        CHAR(50),
   SALARY         REAL
);

TRUNCATE TABLE users RESTART IDENTITY CASCADE;


INSERT INTO users (name, age, address, salary) VALUES 
('Zane Mitchell', 31, '900 Harbor Blvd', 68000),
('Lily Parker', 27, '105 Oakridge Dr', 54000),
('Ethan Cooper', 29, '210 Sunset Blvd', 59000),
('Sophia Reed', 34, '310 Hamilton St', 72000),
('Noah Bennett', 26, '415 Broadway Ave', 53000),
('Ava Foster', 40, '520 Lakeview Rd', 81000),
('Mason Carter', 33, '625 Elmwood Dr', 70000),
('Isabella Brooks', 25, '730 Pinecrest St', 52000),
('Logan Hayes', 38, '835 Ridgeway Ln', 76000),
('Chloe Richardson', 30, '940 West End Rd', 56000),
('Lucas Price', 35, '123 Sunset Way', 65000),
('Harper Watson', 28, '234 Cedar Creek Rd', 58000),
('Oliver Hughes', 41, '345 Birch Hollow Dr', 82000),
('Amelia Sanders', 32, '456 Meadowbrook Ln', 69000),
('William Murphy', 27, '567 Redwood Ct', 55000),
('Evelyn Flores', 39, '678 Skyline Blvd', 79000),
('James Stewart', 36, '789 Northview Dr', 73000),
('Charlotte Torres', 29, '890 Summit Ave', 60000),
('Benjamin Ward', 42, '901 Greenfield St', 85000),
('Scarlett Simmons', 31, '1022 Oceanfront Dr', 67000),
('Henry Collins', 37, '1123 Evergreen Pl', 74000),
('Victoria Powell', 26, '1224 Maple Ridge Rd', 52000),
('Daniel Scott', 40, '1325 Woodland Dr', 78000),
('Grace Peterson', 34, '1426 Riverbend Ln', 72000),
('Matthew Butler', 30, '1527 Lakeside Blvd', 63000),
('Avery Morris', 28, '1628 Hillside Ct', 57000),
('Sebastian Evans', 45, '1729 Parkwood Rd', 90000),
('Ella Jenkins', 33, '1830 Fairview St', 68000),
('Jack Russell', 39, '1931 Rolling Hills Dr', 77000),
('Mila Hughes', 25, '2032 Crestwood Ave', 51000),
('Carter Mitchell', 31, '2133 Highland Blvd', 66000),
('Luna Rivera', 37, '2234 Valley View Dr', 75000),
('Levi Gray', 29, '2335 Mountain Rd', 59000),
('Zoe Barnes', 28, '2436 Willow Ln', 55000),
('Hudson Brooks', 40, '2537 Maple Grove St', 82000),
('Penelope Cooper', 32, '2638 Pine Hill Rd', 70000),
('Alexander Wright', 35, '2739 Cherry Blossom Dr', 73000),
('Layla Stewart', 41, '2840 Oakview Ln', 85000),
('Julian Carter', 27, '2941 Sycamore Ave', 56000),
('Hannah Murphy', 30, '3042 Aspen Grove Blvd', 64000),
('Wyatt Adams', 33, '3143 Lakefront Rd', 71000),
('Aria Martin', 38, '3244 Golden Meadow Dr', 78000),
('Eli Ramirez', 42, '3345 Silverbrook Ln', 86000),
('Natalie Morgan', 29, '3446 Redwood Ln', 60000),
('Owen Bennett', 31, '3547 Sunset Hill Rd', 68000),
('Stella Hayes', 35, '3648 Blue Ridge St', 72000),
('Lincoln Rogers', 26, '3749 Whispering Pines Rd', 55000),
('Madison Bell', 39, '3850 Riverbend Blvd', 76000),
('Grayson Turner', 37, '3951 Evergreen Ct', 74000),
('Nathaniel Ford', 28, '4051 Elmwood Ave', 62000),
('Samantha Blake', 34, '4152 Pinehurst Dr', 71000),
('Joshua Bennett', 41, '4253 Sunset Ln', 80000),
('Lillian Hayes', 27, '4354 Meadowbrook St', 54000),
('Connor Stewart', 39, '4455 Maple Dr', 77000),
('Hannah Sullivan', 30, '4556 Birch St', 64000),
('Evan Peterson', 35, '4657 Valley Rd', 69000),
('Savannah Flores', 28, '4758 Ocean Ave', 61000),
('Owen Taylor', 31, '4859 Sunset Blvd', 66000),
('Natalie Ramirez', 36, '4960 Redwood St', 73000),
('Andrew Ward', 29, '5061 Crestwood Dr', 59000),
('Isla Hughes', 40, '5162 Oakwood Ln', 81000),
('Benjamin Adams', 33, '5263 Meadowbrook Blvd', 72000),
('Emily Cooper', 27, '5364 Golden Ave', 55000),
('Jacob Lewis', 38, '5465 Riverdale Rd', 75000),
('Avery Brooks', 25, '5566 Willow Way', 53000),
('Eli Scott', 45, '5667 Parkview St', 91000),
('Sophie Collins', 29, '5768 Silver Creek Rd', 60000),
('Jack Harris', 42, '5869 Maplewood Blvd', 86000),
('Victoria Jenkins', 37, '5970 Cedar Ln', 73000),
('Samuel King', 31, '6071 Northside Ave', 68000),
('Luna Murphy', 35, '6172 Sycamore Dr', 71000),
('Hudson Gray', 30, '6273 Pine Hollow Rd', 64000),
('Charlotte Turner', 39, '6374 Oak Valley Ln', 78000),
('Mason Reed', 34, '6475 Redwood Way', 70000),
('Lily Stewart', 26, '6576 Aspen Ridge Rd', 56000),
('Carter Mitchell', 32, '6677 Meadow View St', 69000),
('Emma Price', 41, '6778 Rolling Hills Blvd', 83000),
('Noah Barnes', 28, '6879 Whispering Pines Ave', 59000),
('Scarlett Hayes', 35, '6980 Lakeshore Rd', 72000),
('Julian Russell', 40, '7081 Riverbend Ct', 80000),
('Mia Foster', 33, '7182 Mountainview Rd', 71000),
('Grayson Powell', 27, '7283 Sunset Hollow Dr', 58000),
('Ella Simmons', 31, '7384 Oakwood Circle', 66000),
('Owen Bennett', 36, '7485 Mapleview St', 73000),
('Isabella Rogers', 25, '7586 Pine Haven Blvd', 54000),
('Levi Sanders', 42, '7687 Silver Lake Rd', 87000),
('Chloe Adams', 30, '7788 Birch Grove Ln', 62000),
('Henry Carter', 39, '7889 Cedarwood St', 76000),
('Natalie Hall', 34, '7990 Redwood Hills Dr', 72000),
('Andrew Morgan', 31, '8091 Skyline Ave', 69000),
('Sophia Turner', 28, '8192 Parkside Rd', 60000),
('Oliver Harris', 37, '8293 River Crest Blvd', 75000),
('Grace Taylor', 29, '8394 Sunset View Ln', 61000),
('Hudson Reed', 35, '8495 Golden Hollow St', 70000),
('Victoria Powell', 40, '8596 Maple Ridge Ave', 82000),
('Victoria Jenkins', 37, '5970 Cedar Ln', 73000),
('Samuel King', 31, '6071 Northside Ave', 68000),
('Luna Murphy', 35, '6172 Sycamore Dr', 71000),
('Hudson Gray', 30, '6273 Pine Hollow Rd', 64000),
('Charlotte Turner', 39, '6374 Oak Valley Ln', 78000),
('Mason Reed', 34, '6475 Redwood Way', 70000),
('Lily Stewart', 26, '6576 Aspen Ridge Rd', 56000),
('Carter Mitchell', 32, '6677 Meadow View St', 69000),
('Emma Price', 41, '6778 Rolling Hills Blvd', 83000),
('Noah Barnes', 28, '6879 Whispering Pines Ave', 59000),
('Scarlett Hayes', 35, '6980 Lakeshore Rd', 72000),
('Julian Russell', 40, '7081 Riverbend Ct', 80000),
('Mia Foster', 33, '7182 Mountainview Rd', 71000),
('Grayson Powell', 27, '7283 Sunset Hollow Dr', 58000),
('Ella Simmons', 31, '7384 Oakwood Circle', 66000),
('Owen Bennett', 36, '7485 Mapleview St', 73000),
('Isabella Rogers', 25, '7586 Pine Haven Blvd', 54000),
('Levi Sanders', 42, '7687 Silver Lake Rd', 87000),
('Chloe Adams', 30, '7788 Birch Grove Ln', 62000),
('Henry Carter', 39, '7889 Cedarwood St', 76000),
('Natalie Hall', 34, '7990 Redwood Hills Dr', 72000),
('Andrew Morgan', 31, '8091 Skyline Ave', 69000),
('Sophia Turner', 28, '8192 Parkside Rd', 60000),
('Oliver Harris', 37, '8293 River Crest Blvd', 75000),
('Grace Taylor', 29, '8394 Sunset View Ln', 61000),
('Hudson Reed', 35, '8495 Golden Hollow St', 70000),
('Lily Parker', 27, '105 Oakridge Dr', 54000),
('Ethan Cooper', 29, '210 Sunset Blvd', 59000),
('Sophia Reed', 34, '310 Hamilton St', 72000),
('Noah Bennett', 26, '415 Broadway Ave', 53000),
('Ava Foster', 40, '520 Lakeview Rd', 81000),
('Mason Carter', 33, '625 Elmwood Dr', 70000),
('Isabella Brooks', 25, '730 Pinecrest St', 52000),
('Logan Hayes', 38, '835 Ridgeway Ln', 76000),
('Chloe Richardson', 30, '940 West End Rd', 56000),
('Lucas Price', 35, '123 Sunset Way', 65000),
('Harper Watson', 28, '234 Cedar Creek Rd', 58000),
('Oliver Hughes', 41, '345 Birch Hollow Dr', 82000),
('Amelia Sanders', 32, '456 Meadowbrook Ln', 69000),
('William Murphy', 27, '567 Redwood Ct', 55000),
('Evelyn Flores', 39, '678 Skyline Blvd', 79000),
('James Stewart', 36, '789 Northview Dr', 73000),
('Charlotte Torres', 29, '890 Summit Ave', 60000),
('Benjamin Ward', 42, '901 Greenfield St', 85000),
('Hudson Brooks', 40, '2537 Maple Grove St', 82000),
('Penelope Cooper', 32, '2638 Pine Hill Rd', 70000),
('Alexander Wright', 35, '2739 Cherry Blossom Dr', 73000),
('Layla Stewart', 41, '2840 Oakview Ln', 85000),
('Julian Carter', 27, '2941 Sycamore Ave', 56000),
('Hannah Murphy', 30, '3042 Aspen Grove Blvd', 64000),
('Wyatt Adams', 33, '3143 Lakefront Rd', 71000),
('Aria Martin', 38, '3244 Golden Meadow Dr', 78000),
('Eli Ramirez', 42, '3345 Silverbrook Ln', 86000),
('Natalie Morgan', 29, '3446 Redwood Ln', 60000),
('Owen Bennett', 31, '3547 Sunset Hill Rd', 68000),
('Stella Hayes', 35, '3648 Blue Ridge St', 72000),
('Lincoln Rogers', 26, '3749 Whispering Pines Rd', 55000),
('Madison Bell', 39, '3850 Riverbend Blvd', 76000),
('Grayson Turner', 37, '3951 Evergreen Ct', 74000),
('Nathaniel Ford', 28, '4051 Elmwood Ave', 62000),
('Samantha Blake', 34, '4152 Pinehurst Dr', 71000),
('Joshua Bennett', 41, '4253 Sunset Ln', 80000),
('Lillian Hayes', 27, '4354 Meadowbrook St', 54000),
('Connor Stewart', 39, '4455 Maple Dr', 77000),
('Hannah Sullivan', 30, '4556 Birch St', 64000),
('Evan Peterson', 35, '4657 Valley Rd', 69000),
('Savannah Flores', 28, '4758 Ocean Ave', 61000),
('Owen Taylor', 31, '4859 Sunset Blvd', 66000),
('Natalie Ramirez', 36, '4960 Redwood St', 73000),
('Logan Hayes', 38, '835 Ridgeway Ln', 76000),
('Chloe Richardson', 30, '940 West End Rd', 56000),
('Lucas Price', 35, '123 Sunset Way', 65000),
('Harper Watson', 28, '234 Cedar Creek Rd', 58000),
('Oliver Hughes', 41, '345 Birch Hollow Dr', 82000),
('Amelia Sanders', 32, '456 Meadowbrook Ln', 69000),
('William Murphy', 27, '567 Redwood Ct', 55000),
('Evelyn Flores', 39, '678 Skyline Blvd', 79000),
('James Stewart', 36, '789 Northview Dr', 73000),
('Charlotte Torres', 29, '890 Summit Ave', 60000),
('Benjamin Ward', 42, '901 Greenfield St', 85000),
('Scarlett Simmons', 31, '1022 Oceanfront Dr', 67000),
('Henry Collins', 37, '1123 Evergreen Pl', 74000),
('Victoria Powell', 26, '1224 Maple Ridge Rd', 52000),
('Daniel Scott', 40, '1325 Woodland Dr', 78000),
('Grace Peterson', 34, '1426 Riverbend Ln', 72000),
('Matthew Butler', 30, '1527 Lakeside Blvd', 63000),
('Avery Morris', 28, '1628 Hillside Ct', 57000),
('Sebastian Evans', 45, '1729 Parkwood Rd', 90000),
('Ella Jenkins', 33, '1830 Fairview St', 68000),
('Jack Russell', 39, '1931 Rolling Hills Dr', 77000),
('Mila Hughes', 25, '2032 Crestwood Ave', 51000),
('Carter Mitchell', 31, '2133 Highland Blvd', 66000),
('Luna Rivera', 37, '2234 Valley View Dr', 75000),
('Levi Gray', 29, '2335 Mountain Rd', 59000),
('Zoe Barnes', 28, '2436 Willow Ln', 55000),
('Hudson Brooks', 40, '2537 Maple Grove St', 82000),
('Penelope Cooper', 32, '2638 Pine Hill Rd', 70000),
('Alexander Wright', 35, '2739 Cherry Blossom Dr', 73000),
('Layla Stewart', 41, '2840 Oakview Ln', 85000),
('Julian Carter', 27, '2941 Sycamore Ave', 56000),
('Hannah Murphy', 30, '3042 Aspen Grove Blvd', 64000),
('Wyatt Adams', 33, '3143 Lakefront Rd', 71000),
('Aria Martin', 38, '3244 Golden Meadow Dr', 78000),
('Eli Ramirez', 42, '3345 Silverbrook Ln', 86000),
('Natalie Morgan', 29, '3446 Redwood Ln', 60000),
('Owen Bennett', 31, '3547 Sunset Hill Rd', 68000),
('Stella Hayes', 35, '3648 Blue Ridge St', 72000),
('Lincoln Rogers', 26, '3749 Whispering Pines Rd', 55000),
('Madison Bell', 39, '3850 Riverbend Blvd', 76000),
('Grayson Turner', 37, '3951 Evergreen Ct', 74000),
('Nathaniel Ford', 28, '4051 Elmwood Ave', 62000),
('Samantha Blake', 34, '4152 Pinehurst Dr', 71000),
('Joshua Bennett', 41, '4253 Sunset Ln', 80000),
('Lillian Hayes', 27, '4354 Meadowbrook St', 54000),
('Connor Stewart', 39, '4455 Maple Dr', 77000),
('Hannah Sullivan', 30, '4556 Birch St', 64000),
('Evan Peterson', 35, '4657 Valley Rd', 69000),
('Savannah Flores', 28, '4758 Ocean Ave', 61000),
('Owen Taylor', 31, '4859 Sunset Blvd', 66000);