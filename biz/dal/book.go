package dal

import (
	"hertz_demo/biz/dbmodel"
)

func CreateBook(books []*dbmodel.Book) error {
	return DB.Create(books).Error
}

func DeleteBook(bookID uint) error {
	return DB.Where("id = ?", bookID).Delete(&dbmodel.Book{}).Error
}

func UpdateBook(book *dbmodel.Book) error {
	return DB.Save(book).Error
}

func GetBookByID(bookID uint) (*dbmodel.Book, error) {
	var book dbmodel.Book
	err := DB.Where("id = ?", bookID).First(&book).Error
	return &book, err
}

func GetBookList(pageSize, offset int, title, author string) ([]*dbmodel.Book, error) {
	var books []*dbmodel.Book
	query := DB.Model(&dbmodel.Book{})

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if author != "" {
		query = query.Where("author LIKE ?", "%"+author+"%")
	}

	err := query.Limit(pageSize).Offset(offset).Find(&books).Error
	return books, err
}

func IsBookExists(title string) (bool, error) {
	var count int64
	err := DB.Model(&dbmodel.Book{}).Where("title = ?", title).Count(&count).Error
	return count > 0, err
}
