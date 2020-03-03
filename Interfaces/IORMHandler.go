package Interfaces

type IORMHandler interface {
	InitialMigration(connectionString string) error
}
