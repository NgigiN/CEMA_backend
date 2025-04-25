package programs

import (
	"cema_backend/types"
	"database/sql"
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

// RegisterPrograms saves a new program's details in the database.
// It takes a Programs struct as input and returns an error if the operation fails.
func (s *Store) RegisterPrograms(programs types.Programs) error {
	_, err := s.db.Exec("INSERT INTO programs (name, symptoms) VALUES (?, ?)",
		programs.Name, programs.Symptoms)
	if err != nil {
		return err
	}
	return nil
}

// GetPrograms retrieves all programs from the database.
// It returns a slice of Programs structs and an error if the operation fails.
func (s *Store) GetPrograms() ([]types.Programs, error) {
	rows, err := s.db.Query("SELECT name, symptoms FROM programs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var programs []types.Programs
	for rows.Next() {
		var program types.Programs
		if err := rows.Scan(&program.Name, &program.Symptoms); err != nil {
			return nil, err
		}
		programs = append(programs, program)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return programs, nil
}
