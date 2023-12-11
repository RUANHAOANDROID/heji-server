package controller

import "heji-server/domain"

type BookController struct {
	BookUseCase domain.BookRepository
}
