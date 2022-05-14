package model

func GetAllShortCats() ([]Shortcat, error) {
	var result []Shortcat
	tx := db.Find(&result)
	return result, tx.Error
}

func GetShortCatById(id uint64) (Shortcat, error) {
	var result Shortcat
	tx := db.Where("id = ?", id).First((&result))
	return result, tx.Error
}

func GetShortCatByShortUrl(shortUrl string) (Shortcat, error) {
	var result Shortcat
	tx := db.Where("short_url = ?", shortUrl).First((&result))
	return result, tx.Error
}

func CreateShortCat(shortcat Shortcat) error {
	tx := db.Create(&shortcat)
	return tx.Error
}

func UpdateShortCat(shortcat Shortcat) error {
	tx := db.Save(&shortcat)
	return tx.Error
}

func DeleteShortCatById(id uint64) error {
	tx := db.Unscoped().Delete(&Shortcat{}, id)
	return tx.Error
}
