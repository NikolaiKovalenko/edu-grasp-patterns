//Explanation:
//Database Interface: Abstracts the operations that can be performed on a database, providing a stable interface that protects higher-level modules from changes in database implementations.
//Concrete Implementations (MySQLDatabase, PostgreSQLDatabase): Implement the Database interface, shielding the specific database operations from the rest of the system.
//DataManager Class: Uses the Database interface to perform operations, demonstrating that it can switch between different database implementations with minimal changes.
//Client Code: Demonstrates protection from variation by easily switching between different database implementations through the DataManager.

package main

import "fmt"

// Database interface abstracts the database operations
type Database interface {
	Connect() string
	Query(query string) string
}

// MySQLDatabase is a concrete implementation of Database interface
type MySQLDatabase struct{}

func (db *MySQLDatabase) Connect() string {
	return "Connected to MySQL database."
}

func (db *MySQLDatabase) Query(query string) string {
	return fmt.Sprintf("MySQL Executed query: %s", query)
}

// PostgreSQLDatabase is another concrete implementation of Database interface
type PostgreSQLDatabase struct{}

func (db *PostgreSQLDatabase) Connect() string {
	return "Connected to PostgreSQL database."
}

func (db *PostgreSQLDatabase) Query(query string) string {
	return fmt.Sprintf("PostgreSQL Executed query: %s", query)
}

// DataManager uses the Database interface, protecting it from changes in database implementations
type DataManager struct {
	database Database
}

// NewDataManager initializes a new DataManager with a specified Database
func NewDataManager(db Database) *DataManager {
	return &DataManager{database: db}
}

// PerformOperations connects to the database and performs a sample query
func (dm *DataManager) PerformOperations() {
	fmt.Println(dm.database.Connect())
	fmt.Println(dm.database.Query("SELECT * FROM users;"))
}

func main() {
	mysqlDB := &MySQLDatabase{}
	postgresDB := &PostgreSQLDatabase{}

	// Using DataManager with MySQL database
	manager := NewDataManager(mysqlDB)
	manager.PerformOperations()

	// Switching to PostgreSQL database
	manager = NewDataManager(postgresDB)
	manager.PerformOperations()
}

//OUTPUT
//
//Connected to MySQL database.
//MySQL Executed query: SELECT * FROM users;
//Connected to PostgreSQL database.
//PostgreSQL Executed query: SELECT * FROM users;
