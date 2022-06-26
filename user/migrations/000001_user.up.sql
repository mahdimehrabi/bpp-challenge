CREATE TABLE "users" (
    id      serial  PRIMARY KEY,
    vip BOOLEAN DEFAULT FALSE,
    name VARCHAR (50)
);
