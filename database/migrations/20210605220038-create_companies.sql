
-- +migrate Up
CREATE TABLE IF NOT EXISTS companies(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL ,
    status_id INT NOT NULL ,
    name VARCHAR(64) NOT NULL ,
    detail VARCHAR(128) NOT NULL ,
    color ENUM('RED', 'BLUE', 'YELLOW', 'GREEN', 'ORANGE', 'BLACK') NOT NULL DEFAULT 'BLACK',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (status_id) REFERENCES statuses(id)
    );

-- +migrate Down
