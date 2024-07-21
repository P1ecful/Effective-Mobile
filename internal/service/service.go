package service

import (
	"em/internal/db/models"
	"log"
	"time"

	"gorm.io/gorm"
)

var (
	location, _ = time.LoadLocation("Europe/Moscow")
)

type TaskService struct {
	Log *log.Logger
	DB  *gorm.DB
}

func CreateNewService(log *log.Logger, db *gorm.DB) *TaskService {
	return &TaskService{
		Log: log,
		DB:  db,
	}
}

type IdRequest struct {
	ID int `json:"ID"`
}

type SummaryTimeResponse struct {
	UserID      int    `json:"User ID"`
	Task        string `json:"Task"`
	SummaryTime string `json:"Summary Time"`
}

// Methods for users
func (ts *TaskService) CreateUser(req *models.UserModel) error {
	result := ts.DB.Create(req)
	return result.Error
}

func (ts *TaskService) DeleteUser(id int) error {
	result := ts.DB.Delete(&models.UserModel{}, id)
	return result.Error
}

func (ts *TaskService) UpdateUser(req *models.UserModel) error {
	result := ts.DB.Model(&models.UserModel{}).Where("id = ?", req.ID).Updates(&models.UserModel{
		Name:           req.Name,
		Surname:        req.Surname,
		Address:        req.Address,
		PassportNumber: req.PassportNumber,
	})

	return result.Error
}

func (ts *TaskService) GetSummaryTime(userid int) *SummaryTimeResponse {
	var test *models.TaskModel
	result := ts.DB.First(&models.TaskModel{}, userid)
	result.Scan(&test)

	return &SummaryTimeResponse{
		UserID:      test.ID,
		Task:        test.Task,
		SummaryTime: time.Duration(int64(test.SummaryTime)).String(),
	}
}

func (ts *TaskService) GetUser(req *models.UserModel) []models.UserModel {
	var result []models.UserModel
	var response models.UserModel

	// get users
	ts.DB.Where(&models.UserModel{
		Name:           req.Name,
		ID:             req.ID,
		Surname:        req.Surname,
		Address:        req.Address,
		PassportNumber: req.PassportNumber,
	}).Find(&result, response)

	return result
}

// Methods for tasks
func (ts *TaskService) CreateTask(req *models.TaskModel) error {
	result := ts.DB.Create(req)
	return result.Error
}

func (ts *TaskService) StartTask(id int) error {
	result := ts.DB.Model(&models.TaskModel{}).Where("id = ?", id).Updates(&models.TaskModel{
		StartedAt: time.Now().UTC().In(location),
	})

	return result.Error
}

func (ts *TaskService) FinishTask(id int) error {
	var task models.TaskModel

	ts.DB.Model(&models.TaskModel{}).Where("id = ?", id).Updates(&models.TaskModel{
		FinishedAt: time.Now().UTC().In(location),
	})

	// get start and finish values
	prs := ts.DB.Where("ID = ?", id).Take(&models.TaskModel{})
	prs.Scan(&task)

	result := ts.DB.Model(&models.TaskModel{}).Where("id = ?", id).Updates(&models.TaskModel{
		SummaryTime: task.FinishedAt.Sub(task.StartedAt),
	})

	return result.Error
}
