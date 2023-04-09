CREATE TABLE todos (
  todo_id INT AUTO_INCREMENT PRIMARY KEY,
  activity_group_id INT,
   title VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci UNIQUE,
  priority INT,
  created_at TIMESTAMP DEFAULT NOW(),
  FOREIGN KEY (activity_group_id) REFERENCES activities(activity_id)
);