CREATE TABLE enrollments (
  id INT AUTO_INCREMENT PRIMARY KEY,
  client_id INT NOT NULL,
  program_id INT NOT NULL,
  enrolled_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE,
  FOREIGN KEY (program_id) REFERENCES programs(id) ON DELETE CASCADE,
  UNIQUE KEY unique_enrollment (client_id, program_id)
);
