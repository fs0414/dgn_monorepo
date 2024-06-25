package service

type UserServiceImpl interface {
	PublishToken(userId uint) (string, error)
}