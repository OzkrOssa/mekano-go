package database

type envVars struct {
	DBHost     string
	DBName     string
	DBPort     string
	DBUser     string
	DBPassword string
}

type MekanoPayments struct {
	ID          uint   `gorm:"primaryKey"`
	Consecutive int    `gorm:"column:consecutive;unique"`
	CreateAt    string `gorm:"column:create_at"`
}

func (m *MekanoPayments) TableName() string {
	return "mekanopayments"
}
