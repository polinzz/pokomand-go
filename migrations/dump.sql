CREATE TABLE `Restaurants`(
                              `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                              `name` TEXT NOT NULL,
                              `hub_id` INT UNSIGNED NOT NULL,
                              `foods` JSON NOT NULL,
                              `drinks` JSON NOT NULL
);

CREATE TABLE `User_Restaurants`(
                                   `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                                   `restaurant_id` INT UNSIGNED NOT NULL,
                                   `user_id` INT UNSIGNED NOT NULL
);
CREATE TABLE `Orders`
(
    `id`            INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `product`       JSON         NOT NULL,
    `restaurant_id` INT UNSIGNED NOT NULL,
    `price`         DOUBLE(8, 2) NOT NULL,
    `status`        TEXT         NOT NULL,
    `is_finish`     TINYINT(1)   NOT NULL DEFAULT 0,
    `retrieve_code` BIGINT       NOT NULL
);
CREATE TABLE `Users`(
                        `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                        `first_name` TEXT NOT NULL,
                        `last_name` TEXT NOT NULL,
                        `username` TEXT NOT NULL,
                        `password` TEXT NOT NULL,
                        `hub_id` INT NULL,
                        `restaurant_id` INT NULL,
                        `role` ENUM('hub','restaurant','costumers') NOT NULL DEFAULT 'hub'
);
CREATE TABLE `User_Hubs`(
                            `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                            `user_id` INT UNSIGNED NOT NULL,
                            `hub_id` INT UNSIGNED NOT NULL
);
CREATE TABLE `Hubs`(
                       `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                       `name` TEXT NOT NULL,
                       `user_id` INT NOT NULL
);
ALTER TABLE
    `User_Restaurants` ADD CONSTRAINT `user_restaurants_restaurant_id_foreign` FOREIGN KEY(`restaurant_id`) REFERENCES `Restaurants`(`id`);
ALTER TABLE
    `Orders` ADD CONSTRAINT `orders_restaurant_id_foreign` FOREIGN KEY(`restaurant_id`) REFERENCES `Restaurants`(`id`);
ALTER TABLE
    `User_Restaurants` ADD CONSTRAINT `user_restaurants_user_id_foreign` FOREIGN KEY(`user_id`) REFERENCES `Users`(`id`);
ALTER TABLE
    `User_Hubs` ADD CONSTRAINT `user_hubs_hub_id_foreign` FOREIGN KEY(`hub_id`) REFERENCES `Hubs`(`id`);
ALTER TABLE
    `Restaurants` ADD CONSTRAINT `restaurants_hub_id_foreign` FOREIGN KEY(`hub_id`) REFERENCES `Hubs`(`id`);
ALTER TABLE
    `User_Hubs` ADD CONSTRAINT `user_hubs_user_id_foreign` FOREIGN KEY(`user_id`) REFERENCES `Users`(`id`);