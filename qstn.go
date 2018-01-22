package goBotUtils

import (
	"strconv"
	"github.com/pkg/errors"
	"math/rand"
)

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

type SimpleFactForQstn struct {
	Id    string      `json:"id"`
	Title string      `json:"title"`
	Image string      `json:"image"`
	Info  interface{} `json:"info"`
}

/* реализация интерфейса Qstn - простой вопрос с методами проверки и формирования вопроса */
type SimpleQstn struct {
	Id          string   `json:"id"`
	Answers     []string `json:"answers"`
	TrueAnswers []string `json:"true_answer"`
	Image       string   `json:"image"`
}

func (q *SimpleQstn) Check(userAnswer []string) (bool, interface{}) {
	//находим правильные ответы, которые пользователь не отметил и неправильные ответы, которые он отметил.
	missedTrueAnswers := CompareStringArray(q.TrueAnswers, userAnswer)
	wrongUserAnswers := CompareStringArray(userAnswer, q.TrueAnswers)
	// если оба массива пустые, то ответ правильный
	if len(missedTrueAnswers) == 0 && len(wrongUserAnswers) == 0 {
		return true, nil
	}
	return false, nil
}

func (q *SimpleQstn) GetId() string {
	return q.Id
}

func (q *SimpleQstn) GetAnswers() []string {
	return q.Answers
}

func (q *SimpleQstn) GetTrueAnswers() []string {
	return q.TrueAnswers
}

func (q *SimpleQstn) GetImage() string {
	return q.Image
}

/* реализация интерфейса QstnProvider - простой провайдер */
type SimpleQstnProvider struct {
	Data map[string]*SimpleFactForQstn
}

func (qc *SimpleQstnProvider) GetQstnIds(excludeIds []string, count int) ([]string, error) {
	// проверяем что если количество вопросов - excludeIds > count, то переопределяем count в сторону уменьшения
	if len(qc.Data)-len(excludeIds) < count {
		count = len(qc.Data) - len(excludeIds) - 1
	}
	if count < 1 {
		return nil, errors.New("GetQstnIds parameter 'count' must be positive")
	}
	res := []string{}
	for len(res) <= count {
		id := strconv.Itoa(rand.Intn(len(qc.Data)) + 1)
		isExist := false
		for _, v := range excludeIds {
			if v == id {
				isExist = true
				break
			}
		}
		if !isExist {
			res = append(res, id)
			excludeIds = append(excludeIds, id)
		}
	}
	return res, nil
}

func (qc *SimpleQstnProvider) GetQstn(id string, excludeIds []string) (Qstn, error) {
	factForQstn := qc.Data[id]
	if factForQstn == nil {
		return nil, errors.Errorf("GetQstn factForQstn not found id:%v", id)

	}

	q := SimpleQstn{}
	q.Id = factForQstn.Id
	q.Answers = []string{factForQstn.Title}
	q.TrueAnswers = []string{factForQstn.Title}

	q.Image = factForQstn.Image

	for len(q.Answers) < 4 {
		id := strconv.Itoa(rand.Intn(len(qc.Data)) + 1)
		isExist := false
		for _, v := range excludeIds {
			if v == id || q.Id == id {
				isExist = true
				break
			}
		}
		if !isExist {
			q.Answers = append(q.Answers, qc.Data[id].Title)
			excludeIds = append(excludeIds, id)
		}
	}

	// перемешиваем ответы
	for i := range q.Answers {
		j := rand.Intn(i + 1)
		q.Answers[i], q.Answers[j] = q.Answers[j], q.Answers[i]
	}

	return &q, nil
}