package repositories

import (
	"video-conf/core"
	"video-conf/database"
	"video-conf/database/models"

	"gorm.io/gorm"
)

type RoomRepositoryInterface interface {
	Connection() *gorm.DB
	Create(room models.Room, userIDs []string) (models.RoomUser, core.Error)
	FindRoomByID(id, ownerID string) models.Room
	FindRoomUserByID(id, userId string) models.RoomUser
	FindAll(userID string) []models.RoomUser
	Delete(roomId string)
}

type RoomRepository struct {
	DB *gorm.DB
}

func NewRoomRepository() *RoomRepository {
	return &RoomRepository{
		DB: database.Connection(),
	}
}

func (repo *RoomRepository) Connection() *gorm.DB {
	return repo.DB
}

func (repo *RoomRepository) Create(room models.Room, userIDs []string) (models.RoomUser, core.Error) {
	var roomUser models.RoomUser

	if len(userIDs) == 0 {
		return roomUser, core.Error{"reason": "You do not have access to do this"}
	}

	if room.Mode == models.ROOM_MODE_PRIVATE {
		var existRoomUser models.RoomUser

		repo.DB.Table("room_users").
			Joins("Room", repo.DB.Where(&models.Room{OwnerID: room.OwnerID, Mode: models.ROOM_MODE_PRIVATE})).
			Where("user_id = ? AND access = ?", userIDs[0], models.ROOM_ACCESS_MEMBER).
			First(&existRoomUser)

		return existRoomUser, nil
	}

	var newRoom models.Room

	repo.DB.Transaction(func(tx *gorm.DB) error {
		// create new room
		if err := tx.Create(&room).Scan(&newRoom).Error; err != nil {
			return err
		}

		// create owner room user
		ownerRoomUserErr := tx.Create(
			&models.RoomUser{
				RoomID: newRoom.ID,
				UserID: newRoom.OwnerID,
				Access: models.ROOM_ACCESS_OWNER,
			},
		).Scan(&roomUser).Error

		if ownerRoomUserErr != nil {
			return ownerRoomUserErr
		}

		// create members room user
		var roomUsers = []models.RoomUser{}

		for i := 0; i < len(userIDs); i++ {
			roomUsers = append(roomUsers, models.RoomUser{
				RoomID: newRoom.ID,
				UserID: userIDs[i],
				Access: models.ROOM_ACCESS_MEMBER,
			})
		}

		if err := tx.Create(&roomUsers).Error; err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	return roomUser, nil
}

func (repo *RoomRepository) FindRoomByID(id, ownerID string) models.Room {
	var room models.Room

	repo.DB.Table("rooms").First(&room, "id = ? AND owner_id = ?", id, ownerID)

	return room
}

func (repo *RoomRepository) FindRoomUserByID(id, userId string) models.RoomUser {
	var roomUser models.RoomUser

	repo.DB.Table("room_users").
		Where("room_id = ?", id).
		Where("user_id = ?", userId).
		Preload("Room.Users").
		First(&roomUser)

	return roomUser
}

func (repo *RoomRepository) FindAll(userID string) []models.RoomUser {
	var rooms []models.RoomUser

	repo.DB.Table("room_users").Where("user_id = ?", userID).Preload("Room").Find(&rooms)

	return rooms
}

func (repo *RoomRepository) Delete(roomId string) {
	repo.DB.
		Exec(
			"DELETE FROM rooms WHERE id = ?",
			roomId,
		)
}
