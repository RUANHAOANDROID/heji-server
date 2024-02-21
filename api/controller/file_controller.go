package controller

import "heji-server/domain"

// FileController 用户API入口
type FileController struct {
	FilesUseCase domain.FilesUseCase //用户用例
}
