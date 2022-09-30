package model

type Category struct {
	//*gorm.Model
	ID       uint      `json:"id" gorm:"primary key"`
	Name     string    `json:"name" gorm:"type:varchar(50); not null;unique"`
	CreatedAt Time`json:"create_at"`
	UpdatedAt Time`json:"update_at"`
}



