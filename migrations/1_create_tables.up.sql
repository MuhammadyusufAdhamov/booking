CREATE TABLE IF NOT EXISTS "users"(
    "id" SERIAL PRIMARY KEY,
    "first_name" VARCHAR(255) NOT NULL,
    "last_name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "phone_number" VARCHAR(255) UNIQUE,
    "username" VARCHAR(255) UNIQUE,
    "password" VARCHAR(255) NOT NULL,
    "type" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "hotels"(
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER NOT NULL,
    "hotel_name" VARCHAR(255) NOT NULL,
    "hotel_location" VARCHAR(255) NOT NULL,
    "hotel_image_url" VARCHAR,
    "number_of_rooms" INTEGER NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "rooms"(
    "id" SERIAL PRIMARY KEY,
    "type" VARCHAR(255) NOT NULL,
    "number_of_room" INTEGER NOT NULL,
    "room_image_url" VARCHAR,
    "status" VARCHAR(255) NOT NULL,
    "hotel_id" INTEGER NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "bookings"(
    "id" SERIAL PRIMARY KEY,
    "room_id" INTEGER NOT NULL,
    "user_id" INTEGER NOT NULL,
    "hotel_id" INTEGER NOT NULL,
    "from_date" VARCHAR,
    "to_date" VARCHAR,
    "price" DECIMAL(8, 2) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE
    "hotels" ADD CONSTRAINT "hotels_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "rooms" ADD CONSTRAINT "rooms_hotels_id_foreign" FOREIGN KEY("hotel_id") REFERENCES "hotels"("id");
ALTER TABLE
    "bookings" ADD CONSTRAINT "booking_room_id_foreign" FOREIGN KEY("room_id") REFERENCES "rooms"("id");
ALTER TABLE
    "bookings" ADD CONSTRAINT "booking_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "bookings" ADD CONSTRAINT "booking_hotel_id_foreign" FOREIGN KEY("hotel_id") REFERENCES "hotels"("id");