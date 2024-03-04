package storage

import (
	"Archive-Adviser-Bot/lib/e"
	"crypto/sha1"
	"errors"
	"io"
)



type Storage interface{
	Save(p *Page) error
	PickRandom(UserName string) (*Page, error)
	IsExist(p *Page) (bool, error)
	Remove(p *Page) error
}


var ErrNoSavedPages = errors.New("no saved page")

type Page struct{
	URL string
	UserName string
}

func (p Page) Hash() (string, error){
	h := sha1.New()
	if _, err := io.WriteString(h, p.URL); err != nil{
		return "", e.Wrap("cant calculate hash", err)
	}
	if _, err := io.WriteString(h, p.UserName); err != nil{
		return "", e.Wrap("cant calculate hash", err)
	}
	return string(h.Sum(nil)), nil
}