package models

//Bu senaryo için sadece user Struct'ına ihtiyacımız var
type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email" gorm:"unique"` //gorm:"unique"` diyerek migrasyon sırasında e mail kolonunun tekrarlanamaz olmasını sağlıyoruz
	Message string `json:"message"`
}
