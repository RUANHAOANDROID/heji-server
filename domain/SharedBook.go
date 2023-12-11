package domain

const (
	CollShared = "shared_books"
)

type SharedBook struct {
	//账本
	BookID string `bson:"book_id"`
	//邀请码
	Code string `bson:"code"`
	//到期时间
	ExpiredTime string `bson:"expired_time"`
}
