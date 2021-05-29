
-- +migrate Up
CREATE TABLE IF NOT EXISTS reservations(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY ,
    user_id INT NOT NULL ,
    check_in_at DATE NOT NULL ,
    check_out_at DATE NOT NULL ,
    consumption_point INT NOT NULL ,
    consumption_deposit INT NOT NULL ,
    detail VARCHAR(128) NOT NULL ,
    is_deleted BOOLEAN NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(id)
    );
-- +migrate Down
DROP TABLE IF EXISTS reservations;
