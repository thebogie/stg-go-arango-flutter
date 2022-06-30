package model

type EntityVenue struct {
	Key     string  `json:"_key;"`
	Address string  `json:"address"`
	Lat     float64 `json:"lat"`
	Lng     string  `json:"lng"`
}

/*
func (entity *EntityStudent) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *EntityStudent) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}


type EntityPlayer struct {
	Key       string `json:"_key"`
	Firstname string `json:"firstname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
*/
