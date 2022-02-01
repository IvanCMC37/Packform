from asyncio.windows_events import NULL
import psycopg2

#Create database
dbname = 'postgres_testdb'
username = 'postgres'
pwd = '529740123'
hostname = 'localhost'
port_id = '5432'
conn = None
cursor = None

try:
   #establishing the connection
   conn = psycopg2.connect(
      database="postgres", user= username , password=pwd , host= hostname , port= port_id
   )
   conn.autocommit = True

   #Creating a cursor object using the cursor() method
   cursor = conn.cursor()

   #Drop if a database exist then create a new one
   cursor.execute(f'DROP DATABASE IF EXISTS {dbname}')
   print("Database dropped successfully........")
   cursor.execute(f'CREATE database {dbname}')
   print("Database created successfully........")

   #Create table
   conn = psycopg2.connect(
      database=dbname, user= username , password=pwd , host= hostname , port= port_id
   )
   conn.autocommit = True
   cursor = conn.cursor()

   cursor.execute("DROP TABLE IF EXISTS COMPANIES CASCADE")
   sql ='''CREATE TABLE COMPANIES(
      COMPANY_ID INT PRIMARY KEY NOT NULL,
      COMPANY_NAME TEXT NOT NULL
   )'''
   cursor.execute(sql)
   print("Table COMPANIES created successfully........")

   cursor.execute("DROP TABLE IF EXISTS CUSTOMERS CASCADE")
   sql ='''CREATE TABLE CUSTOMERS(
      USER_ID TEXT PRIMARY KEY NOT NULL UNIQUE,
      LOGIN TEXT NOT NULL UNIQUE,
      PASSWORD TEXT NOT NULL,
      NAME TEXT NOT NULL,
      COMPANY_ID INT REFERENCES COMPANIES(COMPANY_ID) NOT NULL,
      CREDIT_CARDS TEXT NOT NULL
   )'''
   cursor.execute(sql)
   print("Table CUSTOMERS created successfully........")

   cursor.execute("DROP TABLE IF EXISTS ORDERS CASCADE")
   sql ='''CREATE TABLE ORDERS(
      ID INT PRIMARY KEY NOT NULL UNIQUE,
      ORDER_DATE TEXT NOT NULL,
      ORDER_NAME TEXT NOT NULL,
      CUSTOMER_ID TEXT NOT NULL REFERENCES CUSTOMERS(USER_ID)
   )'''
   cursor.execute(sql)
   print("Table ORDERS created successfully........")

   cursor.execute("DROP TABLE IF EXISTS ORDER_ITEMS CASCADE")
   sql ='''CREATE TABLE ORDER_ITEMS(
      ID INT PRIMARY KEY NOT NULL UNIQUE,
      ORDER_ID INT NOT NULL REFERENCES ORDERS(ID),
      PRICE_PER_UNIT NUMERIC DEFAULT 0.00,
      QUANTITY INT NOT NULL,
      PRODUCT TEXT NOT NULL
   )'''
   cursor.execute(sql)
   print("Table ORDER_ITEMS created successfully........")

   cursor.execute("DROP TABLE IF EXISTS DELIVERIES CASCADE")
   sql ='''CREATE TABLE DELIVERIES(
      ID INT PRIMARY KEY NOT NULL UNIQUE,
      ORDER_ITEM_ID INT NOT NULL REFERENCES ORDER_ITEMS(ID),
      DELIVERED_QUANTITY INT NOT NULL
   )'''
   cursor.execute(sql)
   print("Table DELIVERIES created successfully........")

   #FILL IN DATA
   with open('Data\Test task - Postgres - customer_companies.csv', 'r') as f:
      next(f) # Skip the header row.
      cursor.copy_expert("copy companies from stdin (format csv)", f)
      
   with open('Data\Test task - Postgres - customers.csv', 'r') as f:
      next(f) # Skip the header row.
      cursor.copy_expert("copy customers from stdin (format csv)", f)

   with open('Data\Test task - Postgres - orders.csv', 'r') as f:
      next(f) # Skip the header row.
      cursor.copy_expert("copy orders from stdin (format csv)", f)

   with open('Data\Test task - Postgres - order_items.csv', 'r') as f:
      next(f) # Skip the header row.
      cursor.copy_expert("copy order_items from stdin (format csv)", f)
      cursor.execute('UPDATE order_items SET price_per_unit = 0.0 WHERE price_per_unit is NULL')

   with open('Data\Test task - Postgres - deliveries.csv', 'r') as f:
      next(f) # Skip the header row.
      cursor.copy_expert("copy deliveries from stdin (format csv)", f)

except Exception as error:
   print(error)
finally:
   if cursor is not None:
      cursor.close()
   if conn is not None:
      conn.close()