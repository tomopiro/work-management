USE work;
CREATE TABLE IF NOT EXISTS `employees`
(
 `id`               VARCHAR(4),
 `name`             VARCHAR(20) NOT NULL,
 `created_at`       Datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);
