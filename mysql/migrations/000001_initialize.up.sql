CREATE TABLE TWEETS (
    id INT NOT NULL AUTO_INCREMENT,
    category TINYINT NOT NULL DEFAULT 0 COMMENT '0: own, 1: like',
    add_date DATETIME NOT NULL,
    url VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL,
    PRIMARY KEY (id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
