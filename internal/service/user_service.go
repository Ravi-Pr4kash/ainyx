package service

import (
    "context"
    "time"

    "ainyx/db/sqlc"
    "github.com/jackc/pgx/v5/pgtype"
)

type UserService struct {
    db *sqlc.Queries
}

func NewUserService(db *sqlc.Queries) *UserService {
    return &UserService{db: db}
}

func convertTimeToPgDate(t time.Time) pgtype.Date {
    var d pgtype.Date
    d.Time = t
    d.Valid = true
    return d
}

//
// CREATE USER
//
func (s *UserService) CreateUser(ctx context.Context, name string, dob time.Time) (sqlc.User, error) {
    params := sqlc.CreateUserParams{
        Name: name,
        Dob:  convertTimeToPgDate(dob),
    }
    return s.db.CreateUser(ctx, params)
}

//
// GET USER BY ID
//
func (s *UserService) GetUserByID(ctx context.Context, id int32) (sqlc.User, error) {
    return s.db.GetUserByID(ctx, id)
}

//
// LIST USERS
//
func (s *UserService) ListUsers(ctx context.Context) ([]sqlc.User, error) {
    return s.db.ListUsers(ctx)
}

//
// UPDATE USER
//
func (s *UserService) UpdateUser(ctx context.Context, id int32, name string, dob time.Time) (sqlc.User, error) {
    params := sqlc.UpdateUserParams{
        Name: name,
        Dob:  convertTimeToPgDate(dob),
        ID:   id,
    }
    return s.db.UpdateUser(ctx, params)
}

//
// DELETE USER
//
func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
    return s.db.DeleteUser(ctx, id)
}
