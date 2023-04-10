CREATE TABLE activities (
  activity_id INT AUTO_INCREMENT PRIMARY KEY,
   title VARCHAR(255),
  email varchar(255) unique not null,
  created_at DATETIME DEFAULT NOW(),
  updated_at DATETIME DEFAULT NOW()
);