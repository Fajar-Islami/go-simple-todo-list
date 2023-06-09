CREATE TABLE todos (
  todo_id INT AUTO_INCREMENT PRIMARY KEY,
  activity_group_id INT,
   title VARCHAR(255) ,
  `priority` enum ('very-high','high','medium','low','very-low'),
   is_active BOOLEAN DEFAULT true,
  created_at DATETIME DEFAULT NOW(),
  updated_at DATETIME DEFAULT NOW(),
 FOREIGN KEY (activity_group_id) REFERENCES activities(activity_id) ON DELETE CASCADE
);
