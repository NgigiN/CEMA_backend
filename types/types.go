// This file contains the types and interfaces used in the application.
// It acts as a contract between the different components of the application.
// It defines the data structures and methods that are used to interact with the database.

// structs define the data structures used in the application
// interfaces define the methods in each service that are used to interact with the database
package types

type DoctorStore interface {
	RegisterDoctors(doctor DoctorRegistration) error
	LoginDoctor(email, password string) error
}

type DoctorRegistration struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
	Department  string `json:"department"`
	Password    string `json:"password"`
}

type DocLogInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type ClientStore interface {
	RegisterClients(client Client) error
	EnrollClient(phoneNumber string, programName string) error
	SearchClient(phonenumber string) (ClientResponse, error)
	GetAllClients() ([]Client, error)
	UpdateClient(client Client) error
	DeleteClient(phonenumber string) error
	CreatePrescription(prescription Prescription) error
	UpdatePrescription(prescription Prescription) error
	GetPrescriptionsByClient(clientID int) ([]Prescription, error)
}
type Client struct {
	ID               int     `json:"id"`
	FirstName        string  `json:"firstname"`
	LastName         string  `json:"lastname"`
	PhoneNumber      string  `json:"phonenumber"`
	Age              int     `json:"age"`
	Height           float32 `json:"height"`
	Weight           float32 `json:"weight"`
	EmergencyContact string  `json:"emergency_contact"`
	EmergencyNumber  string  `json:"emergency_number"`
}

type ClientResponse struct {
	ID               int            `json:"id"`
	FirstName        string         `json:"firstname"`
	LastName         string         `json:"lastname"`
	PhoneNumber      string         `json:"phonenumber"`
	Height           float64        `json:"height"`
	Weight           float64        `json:"weight"`
	Age              int            `json:"age"`
	EmergencyContact string         `json:"emergency_contact"`
	EmergencyNumber  string         `json:"emergency_number"`
	Programs         []Programs     `json:"programs"`
	Prescriptions    []Prescription `json:"prescriptions"`
}

type ProgramsStore interface {
	RegisterPrograms(programs Programs) error
	GetPrograms() ([]Programs, error)
}

type Programs struct {
	Name     string `json:"name"`
	Symptoms string `json:"symptoms"`
}

type ProgramEnrollment struct {
	ProgramID int `json:"program_id"`
	ClientID  int `json:"client_id"`
}

type Prescription struct {
	ID         int      `json:"id"`
	ClientPhone string `json:"client_phone"`
	DoctorID   int      `json:"doctor_id"`
	Medicines  []string `json:"medicines"`
	DateIssued string   `json:"date_issued"`
}
