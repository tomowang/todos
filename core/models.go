package core

type User struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	Email     string `json:"email" gorm:"unique;not null;"`
	Password  string `json:"-"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime:milli"` // Use unix milli seconds as creating time
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:milli"` // Use unix milli seconds as updating time
	Todos     []Todo `json:"-"`
}

type Todo struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	UserID    uint   `json:"user_id"`
	Content   string `json:"content"`
	Status    uint   `json:"status"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime:milli"` // Use unix milli seconds as creating time
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:milli"` // Use unix milli seconds as updating time
}
