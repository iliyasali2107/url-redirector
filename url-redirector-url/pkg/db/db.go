package db

import (
	"context"
	"fmt"
	"log"

	"url-redirecter-url/pkg/models"

	"github.com/jackc/pgx/v5"
)

type Storage interface {
	InsertUrl(models.Url) (models.Url, error)
	GetActiveUrl(int64) (models.Url, error)
	Activate(int64) (int64, error)
	Deactivate(int64) (int64, error)
	GetUrl(int64) (models.Url, error)
	GetUserUrls(userID int64) ([]models.Url, error)
}

type storage struct {
	DB *pgx.Conn
}

func Init(url string) Storage {
	ctx := context.Background()

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	query := `CREATE TABLE IF NOT EXISTS urls (
		id SERIAL PRIMARY KEY,
		user_id SMALLINT,
		url VARCHAR(255),
		active BOOLEAN
	);`

	_, err = conn.Exec(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	return &storage{conn}
}

func (s *storage) InsertUrl(url models.Url) (models.Url, error) {
	query := `INSERT INTO urls(user_id, url, active) VALUES($1, $2, $3) RETURNING id, user_id, url, active`
	var res models.Url
	err := s.DB.QueryRow(context.Background(), query, url.UserID, url.Url, url.Active).Scan(&res.ID, &res.UserID, &res.Url, &res.Active)
	if err != nil {
		return models.Url{}, fmt.Errorf("failed to insert url %w", err)
	}

	return res, nil
}

func (s *storage) GetActiveUrl(userID int64) (models.Url, error) {
	query := `SELECT * FROM urls WHERE user_id = $1 AND active = true`
	var url models.Url
	err := s.DB.QueryRow(context.Background(), query, userID).Scan(&url.ID, &url.UserID, &url.Url, &url.Active)
	if err != nil {
		return models.Url{}, err
	}

	return url, nil
}

func (s *storage) Activate(urlID int64) (int64, error) {
	query := `UPDATE urls SET active = true WHERE id = $1 RETURNING id;`
	var id int64
	err := s.DB.QueryRow(context.Background(), query, urlID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *storage) GetUrl(id int64) (models.Url, error) {
	query := `SELECT * FROM urls WHERE id = $1`
	var url models.Url
	err := s.DB.QueryRow(context.Background(), query, id).Scan(&url.ID, &url.UserID, &url.Url, &url.Active)
	if err != nil {
		return models.Url{}, err
	}

	return url, nil
}

func (s *storage) Deactivate(urlID int64) (int64, error) {
	query := `UPDATE urls SET active = false WHERE id = $1 AND active = true RETURNING id;`
	var id int64
	err := s.DB.QueryRow(context.Background(), query, urlID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *storage) GetUserUrls(userID int64) ([]models.Url, error) {
	query := `SELECT * FROM urls WHERE user_id = $1`

	var urls []models.Url

	rows, err := s.DB.Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}

	hasRows := false

	for rows.Next() {
		var url models.Url
		err := rows.Scan(&url.ID, &url.UserID, &url.Url, &url.Active)
		if err != nil {
			return nil, err
		}

		urls = append(urls, url)
		hasRows = true
	}

	if !hasRows {
		return nil, pgx.ErrNoRows
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return urls, err
}
