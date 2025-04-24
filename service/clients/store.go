// This file hanldes the data access layer for the clients service.
package clients

import (
	"cema_backend/types"
	"context"
	"database/sql"
	"fmt"
)

// struct that declares the database connection
type Store struct {
	db *sql.DB
}

// NewStore initializes a new Store with the given database connection
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// RegisterClients saves a new client in the database
func (s *Store) RegisterClients(client types.Client) error {
	// context is used to manage the lifetime of the request
	ctx := context.Background()
	// Insert queries are seperated to prevent SQL injection
	query := `INSERT INTO clients (firstname, lastname, phonenumber, height, weight, age, emergencycontact, emergencynumber) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	// Execute the query with the parametized values
	_, err := s.db.ExecContext(ctx, query, client.FirstName, client.LastName, client.PhoneNumber, client.Height, client.Weight, client.Age, client.EmergencyContact, client.EmergencyNumber)
	if err != nil {
		return fmt.Errorf("failed to save client in DB %w", err)
	}
	return nil
}

// EnrollClient enrolls a client in a program
func (s *Store) EnrollClient(email string, programID string) error {
	ctx := context.Background()

	var clientID int
	err := s.db.QueryRowContext(ctx, "SELECT id FROM clients WHERE email = ?", email).Scan(&clientID)
	if err != nil {
		return fmt.Errorf("could not find client by email: %w", err)
	}

	query := `INSERT INTO enrollments (program_id, client_id) VALUES (?, ?)`
	_, err = s.db.ExecContext(ctx, query, programID, clientID)
	if err != nil {
		return fmt.Errorf("failed to enroll client in program %w", err)
	}
	return nil
}

// SearchClient retrieves a client by their phone number(which in this case I assume is unique)
// and returns their details along with the programs they are enrolled in
func (s *Store) SearchClient(phonenumber string) (types.ClientResponse, error) {
	ctx := context.Background()

	var client types.ClientResponse

	// Get client data
	clientQuery := `SELECT id, firstname, lastname, phonenumber, height, weight, age, emergencycontact, emergencynumber FROM clients WHERE phonenumber = ?`
	err := s.db.QueryRowContext(ctx, clientQuery, phonenumber).Scan(
		&client.ID, &client.FirstName, &client.LastName,
		&client.PhoneNumber, &client.Height, &client.Weight,
		&client.Age, &client.EmergencyContact, &client.EmergencyNumber,
	)
	// if the client is not found, return an empty client
	if err != nil {
		return client, fmt.Errorf("failed to retrieve client: %w", err)
	}

	// Get program related to the client
	programQuery := `
		SELECT p.name, p.symptoms, p.severity
		FROM enrollments e
		JOIN programs p ON e.program_id = p.id
		WHERE e.client_id = ?
	`
	rows, err := s.db.QueryContext(ctx, programQuery, client.ID)
	if err != nil {
		return client, fmt.Errorf("failed to retrieve programs: %w", err)
	}
	defer rows.Close()

	// Scan the rows and loop through them appending them to the client's programs
	for rows.Next() {
		var program types.Programs
		if err := rows.Scan(&program.Name, &program.Symptoms, &program.Severity); err != nil {
			return client, err
		}
		client.Programs = append(client.Programs, program)
	}
	return client, nil
}

// GetAllClients retrieves all clients from the database
// and returns them as a slice of Client structs
func (s *Store) GetAllClients() ([]types.Client, error) {
	// context is used to manage the lifetime of the request
	ctx := context.Background()
	query := `SELECT id, firstname, lastname, phonenumber, height, weight, age, emergencycontact, emergencynumber FROM clients`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve clients: %w", err)
	}
	defer rows.Close()

	// Scan the rows and loop through them appending them to the clients slice
	var clients []types.Client
	for rows.Next() {
		var client types.Client
		if err := rows.Scan(&client.ID, &client.FirstName, &client.LastName,
			&client.PhoneNumber, &client.Height, &client.Weight,
			&client.Age, &client.EmergencyContact, &client.EmergencyNumber); err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	// Check for any errors encountered during iteration if any
	// and return the clients slice otherwise, we will have looped through all the rows
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return clients, nil
}

// UpdateClient updates the details of a client in the database
func (s *Store) UpdateClient(client types.Client) error {
	// context is used to manage the lifetime of the request
	ctx := context.Background()
	query := `UPDATE clients SET firstname = ?, lastname = ?, phonenumber = ?, height = ?, weight = ?, age = ?, emergencycontact = ?, emergencynumber = ? WHERE id = ?`

	_, err := s.db.ExecContext(ctx, query, client.FirstName, client.LastName, client.PhoneNumber, client.Height, client.Weight, client.Age, client.EmergencyContact, client.EmergencyNumber)
	if err != nil {
		return fmt.Errorf("failed to update client %w", err)
	}
	return nil
}

// DeleteClient deletes a client from the database
func (s *Store) DeleteClient(phonenumber string) error {
	// context is used to manage the lifetime of the request
	ctx := context.Background()
	query := `DELETE FROM clients WHERE phonenumber = ?`
	_, err := s.db.ExecContext(ctx, query, phonenumber)
	if err != nil {
		return fmt.Errorf("failed to delete client %w", err)
	}
	return nil
}
