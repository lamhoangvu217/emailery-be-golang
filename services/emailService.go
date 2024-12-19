package services

import (
	"github.com/lamhoangvu217/emailery-be-golang/database"
	"github.com/lamhoangvu217/emailery-be-golang/models"
)

func GetAllEmailsService() ([]models.Email, error) {
	var emails []models.Email
	if err := database.DB.Find(&emails).Error; err != nil {
		return nil, err
	}
	return emails, nil
}

func DeleteTempEmail(tempEmail *models.Email, tempEmailId string) error {
	if err := database.DB.Where("id = ?", tempEmailId).Delete(tempEmail).Error; err != nil {
		return err
	}
	return nil
}
