package doctors

import (
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
	query := `"INSERT INTO doctors (firstname, lastname, email, phonenumber, department, password) VALUES (?, ?, ?, ?, ?, ?)`

	// Execute the query with the parameterized values
	_, err := s.db.ExecContext(ctx, query, doctor.FirstName, doctor.LastName, doctor.Email, doctor.PhoneNumber, doctor.Department, doctor.Password)
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
	query := `SELECT email, firstname FROM doctors WHERE email = ? AND password = ?`
	// Execute the query with the parameterized values
	row := s.db.QueryRowContext(ctx, query, email, password)
	var dbEmail, firstName string
	// Scan the result into the variables
	err := row.Scan(&dbEmail, &firstName)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("invalid email or password")
		}
		return fmt.Errorf("failed to query doctor %w", err)
	}
	if dbEmail != email {
		return fmt.Errorf("invalid email or password")
	}
	if firstName == "" {
		return fmt.Errorf("invalid email or password")
	}
	// If successful, return nil otherwise return an error
	return nil
}
