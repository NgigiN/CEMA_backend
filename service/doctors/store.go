package doctors

import (
	"cema_backend/auth"
	"cema_backend/types"
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

// NewStore initializes a new Store instance with the given database connection.
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// RegisterDoctors saves a new doctor's details in the database.
func (s *Store) RegisterDoctors(doctor types.DoctorRegistration) error {
	// context is used to manage the lifetime of the request
	ctx := context.Background()

	// Hash the password
	hashedPassword, err := auth.HashPassword(doctor.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	// Update the query to use the hashed password
	query := `INSERT INTO doctors (firstname, lastname, email, phonenumber, department, password) VALUES (?, ?, ?, ?, ?, ?)`

	// Execute the query with the hashed password
	_, err = s.db.ExecContext(ctx, query, doctor.FirstName, doctor.LastName, doctor.Email, doctor.PhoneNumber, doctor.Department, hashedPassword)
	if err != nil {
		return fmt.Errorf("failed to save doctor in DB %w", err)
	}
	// If successful, return nil otherwise return an error
	return nil
}

// LoginDoctor verifies a doctor's credentials in the database.
func (s *Store) LoginDoctor(email, password string) error {
	// context is used to manage the lifetime of the request
	ctx := context.Background()

	// Retrieve the hashed password from the database
	query := `SELECT password FROM doctors WHERE email = ?`
	var storedHashedPassword string
	err := s.db.QueryRowContext(ctx, query, email).Scan(&storedHashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("invalid email or password")
		}
		return fmt.Errorf("failed to query doctor: %w", err)
	}

	// Verify the provided password against the stored hash
	if !auth.CheckPasswordHash(password, storedHashedPassword) {
		return fmt.Errorf("invalid email or password")
	}

	// If successful, return nil
	return nil
}
