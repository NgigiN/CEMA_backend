CREATE TABLE IF NOT EXISTS prescriptions (
    id SERIAL PRIMARY KEY,
    client_phone VARCHAR(15) NOT NULL,
    doctor_id INT NOT NULL,
    medicines TEXT NOT NULL,
    date_issued DATE NOT NULL,
    FOREIGN KEY (doctor_id) REFERENCES doctors (id) ON DELETE CASCADE
);