package models

type User struct {
  ID uint `json:"id" gorm: "primaryKey"`
  Names string `json:"names"`
  Email string `json:"email"`
}
