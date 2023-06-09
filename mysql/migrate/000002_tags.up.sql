CREATE TABLE TAG (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE TWEET_TAG (
   tweet_id INT NOT NULL,
   tag_id INT NOT NULL,
   PRIMARY KEY (tweet_id, tag_id),
   FOREIGN KEY (tweet_id) REFERENCES TWEET(id) ON DELETE CASCADE ON UPDATE CASCADE,
   FOREIGN KEY (tag_id) REFERENCES TAG(id) ON DELETE CASCADE ON UPDATE CASCADE
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci