package models

import "time"

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type SignUp struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"  binding:"required,min=6"`
}

type CreateTaskRequest struct {
	TaskName string `json:"task_name" binding:"required"`
	Task     string `json:"task" binding:"required"`
}

type User struct {
	Id           uint           `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name" gorm:"not null"`
	Email        string         `json:"email" gorm:"unique;not null"`
	Password     string         `json:"-" gorm:"not null"`
	RefreshToken []RefreshToken `json:"-" gorm:"foreignKey:UserId"`
	Tasks        []Task         `json:"tasks,omitempty" gorm:"foreignKey:UserID"`
}

type RefreshToken struct {
	Id        uint      `gorm:"primaryKey"`
	UserId    uint      `gorm:"index"`
	TokenHash string    `gorm:"not null"`
	ExpireAt  time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}

type Task struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	TaskName  string    `json:"task_name"  gorm:"not null"`
	Task      string    `json:"task" gorm:"not null"`
	CreateAt  time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
	UserID    uint      `gorm:"foreignKey"`
}

// When you add an index to a column, the database maintains a separate,
// ordered lookup structure that stores the values of that column along with pointers to the actual rows,
//  so the database can find matching rows quickly without scanning the entire table.

//  where it belongs to ?
//  The index belongs to the database itself and is stored on disk alongside the table, inside the database’s own storage files.
