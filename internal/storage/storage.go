package storage

type Storage interface {
	CreateEmployee(name string, email string, age int) (int64, error)
}
