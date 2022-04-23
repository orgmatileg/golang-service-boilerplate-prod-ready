package repository

// import (
// 	"context"
// 	"database/sql"
// 	"golang_service/model"
// 	"golang_service/pkg/logger"
// 	"time"
// )

// // IUserRepository ...
// type IUserRepository interface {
// 	Create(ctx context.Context, requestID string, reqData *model.User, tx *sql.Tx) (int64, error)
// 	FindByID(ctx context.Context, requestID string, id int64) (*model.User, error)
// 	FindByIDWithTx(ctx context.Context, requestID string, id int64, tx *sql.Tx) (*model.User, error)
// 	FindByEmail(ctx context.Context, requestID string, email string) (*model.User, error)
// 	FindByEmailWithStatus(ctx context.Context, requestID string, email string, status int) (*model.User, error)
// 	FindByPhoneNumber(ctx context.Context, requestID string, phoneNumber string) (*model.User, error)
// 	FindByPhoneNumberWithStatus(ctx context.Context, requestID string, phoneNumber string, status int) (*model.User, error)
// 	UpdateByID(ctx context.Context, requestID string, reqData *model.User, tx *sql.Tx) (int64, error)
// 	UpdatePINByID(ctx context.Context, requestID string, reqData *model.User, tx *sql.Tx) (int64, error)
// 	DeleteByID(ctx context.Context, requestID string, id int64, tx *sql.Tx) (int64, error)
// 	FindAll(ctx context.Context, requestID string, lastID, limit int64) ([]*model.User, error)
// 	UpdatePhoneByID(ctx context.Context, requestID string, reqData *model.User, tx *sql.Tx) (int64, error)
// 	DeleteAccount(ctx context.Context, requestID string, reqData *model.User, tx *sql.Tx) (int64, error)
// 	// Repo for cmd
// 	FindAllWithReferralNull(ctx context.Context, requestID string) ([]*model.User, error)
// }

// // UserRepository ...
// type UserRepository struct {
// 	db *sql.DB
// }

// // New ...
// func New(Conn *sql.DB) IUserRepository {
// 	return &UserRepository{Conn}
// }

// func (u *UserRepository) FindAll(ctx context.Context, requestID string, lastID, limit int64) ([]*model.User, error) {

// 	q := `SELECT id, phone_number, email, full_name, status, updated_at, created_at FROM users WHERE id > $1 LIMIT $2  `

// 	row, err := u.db.QueryContext(ctx, q, lastID, limit)
// 	if err != nil {
// 		logger.Errorf("FindAll | QueryContext | RequestID=%s | Error=%s", requestID, err.Error())
// 		return nil, err
// 	}

// 	defer func() {
// 		if e := row.Close(); e != nil {
// 			logger.Errorf("FindAll | row.Close() | RequestID=%s | Error=%s", requestID, err.Error())
// 		}
// 	}()

// 	lu := make([]*model.User, 0)

// 	for row.Next() {
// 		u := model.User{}
// 		err = row.Scan(
// 			&u.ID,
// 			&u.PhoneNumber,
// 			&u.Email,
// 			&u.FullName,
// 			&u.Status,
// 			&u.CreatedAt,
// 			&u.UpdatedAt,
// 		)
// 		if err != nil {
// 			logger.Errorf("FindAll | row.Scan | RequestID=%s | Error=%s", requestID, err.Error())
// 			return nil, err
// 		}
// 		lu = append(lu, &u)
// 	}
// 	return lu, nil
// }

// // UpdateByID ...
// func (u *UserRepository) UpdateByID(ctx context.Context, requestID string, reqData *model.User, tx *sql.Tx) (int64, error) {

// 	q := `
// 	UPDATE users
// 	SET full_name=$1, email=$2, status=$3, updated_at=$4
// 	WHERE id=$5`

// 	res, err := tx.ExecContext(ctx, q,
// 		reqData.FullName,
// 		reqData.Email,
// 		reqData.Status,
// 		time.Now(),
// 		reqData.ID,
// 	)
// 	if err != nil {
// 		logger.Errorf("UpdateByID | ExecContext | RequestID=%s | Error=%s", requestID, err.Error())
// 		return -1, err
// 	}
// 	rowUpdated, err := res.RowsAffected()
// 	if err != nil {
// 		logger.Errorf("UpdateByID | RowsAffected | RequestID=%s | Error=%s", requestID, err.Error())
// 		return -1, err
// 	}
// 	return rowUpdated, nil
// }

// // Create ...
// func (u *UserRepository) Create(ctx context.Context, requestID string, reqData *model.User, tx *sql.Tx) (int64, error) {

// 	q := `
// 	INSERT INTO users
// 	(email, phone_number, pin, full_name, status, updated_at, created_at)
// 	VALUES ($1, $2, $3, $4, $5, $6, $7)
// 	RETURNING id`

// 	now := time.Now()
// 	var insertedID int64

// 	err := tx.QueryRowContext(ctx, q, reqData.Email, reqData.PhoneNumber, reqData.PIN, reqData.FullName, reqData.Status, now, now).Scan(&insertedID)
// 	if err != nil {
// 		logger.Errorf("Create | ExecContext | RequestID=%s | Error=%s", requestID, err.Error())
// 		return -1, err
// 	}

// 	return insertedID, nil
// }

// // DeleteByID ...
// func (u *UserRepository) DeleteByID(ctx context.Context, requestID string, id int64, tx *sql.Tx) (int64, error) {

// 	q := "DELETE FROM users WHERE id = $1"

// 	res, err := tx.ExecContext(ctx, q, id)
// 	if err != nil {
// 		logger.Errorf("DeleteByID | ExecContext | RequestID=%s | Error=%s", requestID, err.Error())
// 		return -1, err
// 	}

// 	rowAffected, err := res.RowsAffected()
// 	if err != nil {
// 		logger.Errorf("DeleteByID | RowsAffected | Error=%s", err.Error())
// 		return -1, err
// 	}

// 	return rowAffected, nil
// }

// // FindByID ..
// func (u *UserRepository) FindByID(ctx context.Context, requestID string, id int64) (*model.User, error) {
// 	q := `SELECT id,
// 		 phone_number,
// 		 email,
// 		 full_name,
// 		 status,
// 		 updated_at,
// 		 created_at,
// 		 pin,
// 		 is_delete
// 		FROM users WHERE id = $1`
// 	m := model.User{}

// 	err := u.db.QueryRowContext(ctx, q, id).Scan(
// 		&m.ID,
// 		&m.PhoneNumber,
// 		&m.Email,
// 		&m.FullName,
// 		&m.Status,
// 		&m.UpdatedAt,
// 		&m.CreatedAt,
// 		&m.PIN,
// 		&m.IsDelete,
// 	)
// 	if err != nil {
// 		logger.Errorf("FindByID | QueryRowContext | RequestID=%s | Error=%s", requestID, err.Error())
// 		return nil, err
// 	}

// 	return &m, nil
// }

// // FindByEmail ..
// func (u *UserRepository) FindByEmail(ctx context.Context, requestID string, email string) (*model.User, error) {
// 	q := "SELECT id, phone_number, email, full_name, status, updated_at, created_at FROM users WHERE email = $1"
// 	m := model.User{}

// 	err := u.db.QueryRowContext(ctx, q, email).Scan(
// 		&m.ID,
// 		&m.PhoneNumber,
// 		&m.Email,
// 		&m.FullName,
// 		&m.Status,
// 		&m.UpdatedAt,
// 		&m.CreatedAt,
// 	)
// 	if err != nil {
// 		logger.Errorf("FindByEmail | QueryRowContext | RequestID=%s | Error=%s", requestID, err.Error())
// 		return nil, err
// 	}

// 	return &m, nil
// }

// // FindByPhoneNumber ...
// func (u *UserRepository) FindByPhoneNumber(ctx context.Context, requestID string, phoneNumber string) (*model.User, error) {
// 	q := "SELECT id, phone_number, email, pin, full_name, status, updated_at, created_at FROM users WHERE phone_number = $1"
// 	m := model.User{}

// 	err := u.db.QueryRowContext(ctx, q, phoneNumber).Scan(
// 		&m.ID,
// 		&m.PhoneNumber,
// 		&m.Email,
// 		&m.PIN,
// 		&m.FullName,
// 		&m.Status,
// 		&m.UpdatedAt,
// 		&m.CreatedAt,
// 	)
// 	if err != nil {
// 		logger.Errorf("FindByPhoneNumber | QueryRowContext | RequestID=%s | Error=%s", requestID, err.Error())
// 		return nil, err
// 	}

// 	return &m, nil
// }

// // FindByPhoneNumberWithStatus ...
// func (u *UserRepository) FindByPhoneNumberWithStatus(ctx context.Context, requestID string, phoneNumber string, status int) (*model.User, error) {
// 	q := "SELECT id, phone_number, pin, email, full_name, status, updated_at, created_at FROM users WHERE phone_number = $1 AND status = $2"
// 	m := model.User{}

// 	err := u.db.QueryRowContext(ctx, q, phoneNumber, status).Scan(
// 		&m.ID,
// 		&m.PhoneNumber,
// 		&m.PIN,
// 		&m.Email,
// 		&m.FullName,
// 		&m.Status,
// 		&m.UpdatedAt,
// 		&m.CreatedAt,
// 	)
// 	if err != nil {
// 		logger.Errorf("FindByPhoneNumber | QueryRowContext | RequestID=%s | Error=%s", requestID, err.Error())
// 		return nil, err
// 	}

// 	return &m, nil
// }

// // FindByEmailWithStatus ...
// func (u *UserRepository) FindByEmailWithStatus(ctx context.Context, requestID string, email string, status int) (*model.User, error) {
// 	q := "SELECT id, phone_number, email, full_name, status, updated_at, created_at FROM users WHERE email = $1 AND status = $2"
// 	m := model.User{}

// 	err := u.db.QueryRowContext(ctx, q, email, status).Scan(
// 		&m.ID,
// 		&m.PhoneNumber,
// 		&m.Email,
// 		&m.FullName,
// 		&m.Status,
// 		&m.UpdatedAt,
// 		&m.CreatedAt,
// 	)
// 	if err != nil {
// 		logger.Errorf("FindByPhoneNumber | QueryRowContext | RequestID=%s | Error=%s", requestID, err.Error())
// 		return nil, err
// 	}

// 	return &m, nil
// }

// // FindByID ..
// func (u *UserRepository) FindByIDWithTx(ctx context.Context, requestID string, id int64, tx *sql.Tx) (*model.User, error) {
// 	q := "SELECT id, phone_number, email, full_name, status, updated_at, created_at FROM users WHERE id = $1"
// 	m := model.User{}

// 	err := tx.QueryRowContext(ctx, q, id).Scan(
// 		&m.ID,
// 		&m.PhoneNumber,
// 		&m.Email,
// 		&m.FullName,
// 		&m.Status,
// 		&m.UpdatedAt,
// 		&m.CreatedAt,
// 	)
// 	if err != nil {
// 		logger.Errorf("FindByID | QueryRowContext | RequestID=%s | Error=%s", requestID, err.Error())
// 		return nil, err
// 	}

// 	return &m, nil
// }

// // UpdatePINByID ...
// func (u *UserRepository) UpdatePINByID(ctx context.Context, requestID string, reqData *model.User, tx *sql.Tx) (int64, error) {
// 	q := `UPDATE users
// 	SET pin=$1,
// 		updated_at=$2
// 	WHERE id=$3`
// 	res, err := tx.ExecContext(ctx, q,
// 		reqData.PIN,
// 		time.Now(),
// 		reqData.ID,
// 	)
// 	if err != nil {
// 		logger.Errorf("UpdatePINByID | ExecContext | RequestID=%s | Error=%s", requestID, err.Error())
// 		return -1, err
// 	}
// 	rowUpdated, err := res.RowsAffected()
// 	if err != nil {
// 		logger.Errorf("UpdatePINByID | RowsAffected | RequestID=%s | Error=%s", requestID, err.Error())
// 		return -1, err
// 	}
// 	return rowUpdated, nil
// }

// // UpdatePhoneByID ...
// func (u *UserRepository) UpdatePhoneByID(ctx context.Context, requestID string, reqData *model.User, tx *sql.Tx) (int64, error) {

// 	q := `
// 	UPDATE users
// 	SET phone_number=$1, updated_at=$2
// 	WHERE id=$3`

// 	res, err := tx.ExecContext(ctx, q,
// 		reqData.PhoneNumber,
// 		time.Now(),
// 		reqData.ID,
// 	)
// 	if err != nil {
// 		logger.Errorf("UpdatePhoneByID | ExecContext | RequestID=%s | Error=%s", requestID, err.Error())
// 		return -1, err
// 	}
// 	rowUpdated, err := res.RowsAffected()
// 	if err != nil {
// 		logger.Errorf("UpdatePhoneByID | RowsAffected | RequestID=%s | Error=%s", requestID, err.Error())
// 		return -1, err
// 	}
// 	return rowUpdated, nil
// }

// // DeleteAccount ...
// func (u *UserRepository) DeleteAccount(ctx context.Context, requestID string, reqData *model.User, tx *sql.Tx) (int64, error) {

// 	q := `
// 	UPDATE users
// 	SET phone_number=$1,
// 		email=$2,
// 		is_delete=$3,
// 		updated_at=$4
// 	WHERE id=$5`

// 	res, err := tx.ExecContext(ctx, q,
// 		reqData.PhoneNumber,
// 		reqData.Email,
// 		true,
// 		time.Now(),
// 		reqData.ID,
// 	)
// 	if err != nil {
// 		logger.Errorf("DeleteAccount | ExecContext | RequestID=%s | Error=%s", requestID, err.Error())
// 		return -1, err
// 	}
// 	rowUpdated, err := res.RowsAffected()
// 	if err != nil {
// 		logger.Errorf("DeleteAccount | RowsAffected | RequestID=%s | Error=%s", requestID, err.Error())
// 		return -1, err
// 	}
// 	return rowUpdated, nil
// }

// // FindAllWithReferralNull
// func (u *UserRepository) FindAllWithReferralNull(ctx context.Context, requestID string) ([]*model.User, error) {

// 	q := `select
// 		u.id,
// 		u.phone_number,
// 		u.email,
// 		u.full_name,
// 		u.status,
// 		u.updated_at,
// 		u.created_at
// 	from
// 		users u
// 	left join partner_coupons pc on
// 		pc.user_id = u.id
// 	where
// 		pc is null`

// 	row, err := u.db.QueryContext(ctx, q)
// 	if err != nil {
// 		logger.Errorf("FindAllWithReferralNull | QueryContext | RequestID=%s | Error=%s", requestID, err.Error())
// 		return nil, err
// 	}

// 	defer func() {
// 		if e := row.Close(); e != nil {
// 			logger.Errorf("FindAllWithReferralNull | row.Close() | RequestID=%s | Error=%s", requestID, err.Error())
// 		}
// 	}()

// 	lu := make([]*model.User, 0)

// 	for row.Next() {
// 		u := model.User{}
// 		err = row.Scan(
// 			&u.ID,
// 			&u.PhoneNumber,
// 			&u.Email,
// 			&u.FullName,
// 			&u.Status,
// 			&u.CreatedAt,
// 			&u.UpdatedAt,
// 		)
// 		if err != nil {
// 			logger.Errorf("FindAllWithReferralNull | row.Scan | RequestID=%s | Error=%s", requestID, err.Error())
// 			return nil, err
// 		}
// 		lu = append(lu, &u)
// 	}
// 	return lu, nil
// }
