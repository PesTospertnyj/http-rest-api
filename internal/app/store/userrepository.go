package store

import "github.com/PesTospertnyj/http-rest-api/internal/app/model"

//UserRepository ...
type UserRepository struct {
	store *Store
}

//Create ...
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	if err := r.store.db.
		QueryRow(
			`insert into users (email, encrypted_password) values ($1, $2) returning id`,
			u.Email,
			u.EncryptedPassword,
		).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

//FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	if err := r.store.db.QueryRow(
		"select id, email, encrypted_password from users where email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		return nil, err
	}

	return u, nil
}