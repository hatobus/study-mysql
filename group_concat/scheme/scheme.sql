DROP TABLE IF EXISTS user;
CREATE TABLE user (
    id int(11) NOT NULL AUTO_INCREMENT,
    name varchar(16) NOT NULL,
    PRIMARY KEY (id)
);

DROP TABLE IF EXISTS attachment;
CREATE TABLE attachment (
    user_id int(11) NOT NULL,
    attachment_name varchar(32) NOT NULL
);
