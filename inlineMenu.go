package goBotUtils

import (
	"github.com/pepelazz/go-bot-telebot"
	"github.com/pkg/errors"
	"fmt"
)

type MenuItem interface {
	GetTitle() string
	GetId() string
}

// создание списка с pagination
func CreateInlineKb(list []interface{}, pageNum, perPage int) (kb [][]telebot.KeyboardButton, err error) {
	for i := perPage * pageNum; i < perPage * (pageNum + 1); i++ {
		if i > -1 && i < len(list) {
			v, ok := list[i].(MenuItem)
			if !ok {
				err = errors.New(fmt.Sprintf("Wrong type asserion in goBotUtils.CreateInlineKb %s", list[i]))
				return
			}
			kb = append(kb, []telebot.KeyboardButton{telebot.KeyboardButton{v.GetTitle(), "", v.GetId(), ""}})
		}
	}
	pages := []telebot.KeyboardButton{}
	if pageNum > 0 {
		pages = append(pages, telebot.KeyboardButton{"<<", "", "<<", ""})
	}
	if len(list) > (pageNum + 1) * perPage {
		pages = append(pages, telebot.KeyboardButton{">>", "", ">>", ""})

	}

	kb = append(kb, pages)
	return
}
