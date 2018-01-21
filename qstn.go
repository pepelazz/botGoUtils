package goBotUtils

type QstnProvider interface {
	GetQstn(id string, excludeIds []string) (Qstn, error)
	GetQstnIds(excludeIds []string, count int) ([]string, error) // метод получения порции id'шников вопросов, которые исключают переданные excludeIds
}

type Qstn interface {
	GetId() string
	GetAnswers() []string
	GetImage() string
	Check(userAnswer []string) (bool, interface{})
	GetTrueAnswers() []string
}
