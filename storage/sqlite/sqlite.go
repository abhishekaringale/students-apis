package sqlite

import (
	"database/sql"

	"github.com/abhishekaringale/students-api/internal/config"
	"github.com/abhishekaringale/students-api/internal/types"
	_ "github.com/glebarez/go-sqlite"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite", cfg.StoragePath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT UNIQUE,
		age INTEGER
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{Db: db}, nil

}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO students (name, email,age) VALUES (?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	var student types.Student

	err := s.Db.QueryRow("SELECT id, name, email, age FROM students WHERE id = ?", id).Scan(
		&student.Id,
		&student.Name,
		&student.Email,
		&student.Age,
	)
	if err != nil {
		return types.Student{}, err
	}

	return student, nil
}

func (s *Sqlite) GetStudents() ([]types.Student, error) {
	rows, err := s.Db.Query("SELECT id, name, email, age FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []types.Student
	for rows.Next() {
		var student types.Student
		err := rows.Scan(
			&student.Id,
			&student.Name,
			&student.Email,
			&student.Age,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}
