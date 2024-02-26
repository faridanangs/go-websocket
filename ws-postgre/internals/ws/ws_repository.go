package ws

import (
	"context"
	"database/sql"
	"errors"
	"ws_postgre/helper"
)

type RoomRepositoryIPLM struct{}

func NewRoomRepositoryIPLM() *RoomRepositoryIPLM {
	return &RoomRepositoryIPLM{}
}

func (repo *RoomRepositoryIPLM) Save(ctx context.Context, tx *sql.Tx, req Room) Room {
	query := "INSERT INTO rooms (id, name, user_id, created_at) values($1,$2,$3,$4)"
	_, err := tx.ExecContext(ctx, query, req.ID, req.Name, req.UserID, req.CreatedAt)

	helper.HelperError(err, "error creating room repository")

	row := tx.QueryRowContext(ctx, `
	select r.id, r.name, r.user_id, r.created_at, u.id, u.name, u.email, u.created_at
	from rooms as r
	join users as u on u.id = r.user_id
	where r.id = $1
	`, req.ID)
	room := Room{}

	err = row.Scan(
		&room.ID, &room.Name, &room.UserID, &room.CreatedAt,
		&room.User.ID, &room.User.Name, &room.User.Email, &room.User.CreatedAt,
	)
	helper.HelperError(err, "error scanning getById room repository")

	return room

}
func (repo *RoomRepositoryIPLM) GetByID(ctx context.Context, tx *sql.Tx, roomID string) (Room, error) {
	row := tx.QueryRowContext(ctx, `
	select r.id, r.name, r.user_id, r.created_at, u.id, u.name, u.email, u.created_at
	from rooms as r
	join users as u on u.id = r.user_id
	where r.id = $1
	`, roomID)

	// userScan := internaluser.UserResponse{}

	room := Room{}

	err := row.Scan(
		&room.ID, &room.Name, &room.UserID, &room.CreatedAt,
		// &userScan.ID, &userScan.Name, &userScan.Email, &userScan.CreatedAt
		&room.User.ID, &room.User.Name, &room.User.Email, &room.User.CreatedAt,
	)
	if err != nil {
		return Room{}, errors.New("room not found")
	}

	// room.User = append(room.User, userScan)

	return room, nil
}
func (repo *RoomRepositoryIPLM) GetAll(ctx context.Context, tx *sql.Tx) []Room {
	rows, err := tx.QueryContext(ctx,
		`
	select r.id, r.name, r.user_id, r.created_at, u.id, u.name, u.email, u.created_at
	from rooms as r
	join users as u on u.id = r.user_id
	`)
	helper.HelperError(err, "error querying getAll room repository")

	defer rows.Close()
	rooms := []Room{}
	// userScan := internaluser.UserResponse{}

	for rows.Next() {
		var room Room
		err := rows.Scan(&room.ID, &room.Name, &room.UserID, &room.CreatedAt,
			// &userScan.ID, &userScan.Name, &userScan.Email, &userScan.CreatedAt
			&room.User.ID, &room.User.Name, &room.User.Email, &room.User.CreatedAt,
		)
		helper.HelperError(err, "error scanning getAll room repository")

		// room.User = append(room.User, userScan)

		rooms = append(rooms, room)
	}

	return rooms

}
func (repo *RoomRepositoryIPLM) SaveMessage(ctx context.Context, tx *sql.Tx, req Message) {
	query := "INSERT INTO messages (user_id, username, contents, room_id, created_at) values($1,$2,$3,$4, $5)"
	_, err := tx.ExecContext(ctx, query, req.ID, req.Username, req.Content, req.RoomID, req.CreatedAt)

	helper.HelperError(err, "error creating message repository")

}
