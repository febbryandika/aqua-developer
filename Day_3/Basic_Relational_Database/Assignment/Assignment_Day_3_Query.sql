CREATE TABLE "customers" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "customer_name" char(50) NOT NULL
);

CREATE TABLE "products" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" char(50)
);

CREATE TABLE "orders" (
  "order_id" SERIAL PRIMARY KEY NOT NULL,
  "customer_id" int NOT NULL,
  "product_id" int NOT NULL,
  "order_date" date NOT NULL,
  "total" float NOT NULL
);

ALTER TABLE "orders"
ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "orders"
ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");


INSERT INTO "customers" ("customer_name")
VALUES ('Febbry');

INSERT INTO "customers" ("customer_name")
VALUES ('Andika');

INSERT INTO "customers" ("customer_name")
VALUES ('Ramadhan');


INSERT INTO "products" ("name")
VALUES ('Patin');

INSERT INTO "products" ("name")
VALUES ('Lele');

INSERT INTO "products" ("name")
VALUES ('Nila');


INSERT INTO "orders" ("customer_id", "product_id", "order_date", "total")
VALUES (1, 1, '2022-10-26', 10);

INSERT INTO "orders" ("customer_id", "product_id", "order_date", "total")
VALUES (1, 2, '2022-10-27', 3);

INSERT INTO "orders" ("customer_id", "product_id", "order_date", "total")
VALUES (1, 3, '2022-10-28', 7);

INSERT INTO "orders" ("customer_id", "product_id", "order_date", "total")
VALUES (2, 1, '2022-10-29', 5);

INSERT INTO "orders" ("customer_id", "product_id", "order_date", "total")
VALUES (2, 2, '2022-10-30', 1);

INSERT INTO "orders" ("customer_id", "product_id", "order_date", "total")
VALUES (2, 3, '2022-10-31', 2);

INSERT INTO "orders" ("customer_id", "product_id", "order_date", "total")
VALUES (3, 1, '2022-11-01', 9);

INSERT INTO "orders" ("customer_id", "product_id", "order_date", "total")
VALUES (3, 2, '2022-11-02', 4);

INSERT INTO "orders" ("customer_id", "product_id", "order_date", "total")
VALUES (3, 3, '2022-11-03', 6);


UPDATE "customers"
SET "customer_name" = 'Rama'
WHERE id = 3;

UPDATE "products"
SET "name" = 'Gurame'
WHERE id = 1;

UPDATE "orders"
SET "total" = '8'
WHERE order_id = 5;


DELETE FROM "orders" WHERE "customer_id" = 3 OR "product_id" = 3;
DELETE FROM "customers" WHERE "customer_name" = 'Rama';
DELETE FROM "products" WHERE "id" = 3;


SELECT customers."customer_name", products."name", orders."order_date", orders."total"
FROM "customers"
JOIN "orders" ON customers."id" = orders."customer_id"
JOIN "products" ON products."id" = orders."product_id";