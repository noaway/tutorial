package service

type ItemServices interface {
	ShowName() string
}

type ItemService struct {
	Is    ItemServices
	Users Users
}

func (i *ItemService) Name() string {
	return i.Is.ShowName() + " ~ " + i.Users.GetUser().Name
}
