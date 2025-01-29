package store

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
	"github.com/ybencab/todo-app/types"
	"github.com/ybencab/todo-app/utils"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(connStr string) *PostgresStore {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error: ", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("error: ", err.Error())
	}

	return &PostgresStore{db}
}

func (s *PostgresStore) Init() error {
	if err := s.CreateUsersTable(); err != nil {
		return err
	}
	if err := s.CreateTodosTable(); err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) CreateUsersTable() error {
	query := `create table if not exists user (
		id uuid primary key default gen_random_uuid(),
		username varchar(50) not null unique,
		email varchar(255) not null unique,
		created_at timestamptz default now(),
		hashed_password varchar(255) not null
	)`

	_, err := s.db.Query(query)
	return err
}

func (s *PostgresStore) CreateTodosTable() error {
	query := `create table if not exists todo (
		user_id uuid,
		title varchar(20) not null,
		description text not null,
		created_at timestamptz default now(),
		constraint fk_user foreign key (user_id) references user(id) on delete cascade
	)`

	_, err := s.db.Query(query)
	return err
}

func (s *PostgresStore) CreateTodo(todo *types.ToDo) error {
	query := `insert into todo
		(user_id, title, description, created_at)
		values ($1, $2, $3, $4)`
	
	_, err := s.db.Query(
		query,
		todo.UserID,
		todo.Title,
		todo.Description,
		todo.CreatedAt,
	)
	if err != nil {
		return err
	}
	
	return nil
}

func (s *PostgresStore) GetTodo(user_id, title string) (*types.ToDo, error) {
	rows, err := s.db.Query("select * from todo where id = $1 and user_id = $2", title, user_id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return utils.ScanRowIntoToDo(rows)
	}

	return nil, errors.New("todo not found")
}

func (s *PostgresStore) GetTodos() ([]*types.ToDo, error) {
	rows, err := s.db.Query("select * from todo")
	if err != nil {
		return nil, err
	}

	todos := []*types.ToDo{}
	for rows.Next() {
		todo, err := utils.ScanRowIntoToDo(rows)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (s *PostgresStore) CreateUser(req *types.User) error {
	query := `insert into user
		(id, username, email, created_at, hashed_password)
		values ($1, $2, $3, $4, $5)`

	_, err := s.db.Query(
		query,
		req.ID,
		req.Username,
		req.Email,
		req.CreatedAt,
		req.Password,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) GetUserByUsername(username string) (*types.User, error) {
	rows, err := s.db.Query("select * from user where username = $1", username)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return utils.ScanRowIntoUser(rows)
	}

	return nil, errors.New("user not found")	
}

func (s *PostgresStore) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("select * from user where email = $1", email)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return utils.ScanRowIntoUser(rows)
	}

	return nil, errors.New("user not found")
}

func (s *PostgresStore) GetAllUsers() ([]*types.User, error) {
	rows, err := s.db.Query("select * from user")
	if err != nil {
		return nil, err
	}

	users := []*types.User{}
	for rows.Next() {
		user, err := utils.ScanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil	
}
