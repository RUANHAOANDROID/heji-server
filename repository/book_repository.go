package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"heji-server/domain"
	"heji-server/mongo"
)

// 账本存储库
type bookRepository struct {
	database   mongo.Database
	collection string
}

func (b bookRepository) FindOne(c context.Context, id primitive.ObjectID) (domain.Book, error) {
	coll := b.database.Collection(b.collection)
	filter := bson.M{"_id": id}
	var book domain.Book
	err := coll.FindOne(c, filter).Decode(&book)
	return book, err
}

func (b bookRepository) List(c context.Context, userId string) (*[]domain.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookRepository) Update(c context.Context, book *domain.Book) (*domain.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b bookRepository) CreateOne(c context.Context, book *domain.Book) error {
	coll := b.database.Collection(domain.CollBook)
	one, err := b.FindOne(c, book.ID)
	if err == nil && (book.ID == one.ID) {
		return errors.New("账本已存在！")
	}
	_, err = coll.InsertOne(c, book)
	if err != nil {
		return err
	}
	return err
}

func NewBookRepository(db mongo.Database, coll string) domain.BookRepository {
	return &bookRepository{
		database:   db,
		collection: coll,
	}
}
