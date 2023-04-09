CREATE TABLE activities (
  activity_id INT AUTO_INCREMENT PRIMARY KEY,
   title VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci UNIQUE,
  email varchar(255) unique not null,
  created_at TIMESTAMP DEFAULT NOW()
);