package models

import "gorm.io/gorm"

//This is a one to many table releationship with the departments belonging to many employees but one employee belonging to one department

type Department struct {
	gorm.Model          `gorm:"primarykey" json:"id"`
	Name string         `json:"name"`
	Region string       `json:"region"`
	Employee []Employee `gorm:"foreignKey:DepartmentID" json:"employee"` // A slice of employees because a department can belong to many employees
}

type Employee struct {
	gorm.Model             `gorm:"primarykey" json:"id"`
	Name string           `json:"name"`
	DepartmentID uint     `json:"department_id"`
	Department            `gorm:"foreignkey:DepartmentID" json:"department"`
}

