package storage

import (
	"Archive-Adviser-Bot/lib/e"
	"crypto/sha1"
	"errors"
	"io"
	"fmt"
	"context"
)



type Storage interface{
	Save(ctx context.Context, p *Page) error
	PickRandom(ctx context.Context, UserName string) (*Page, error)
	IsExist(ctx context.Context, p *Page) (bool, error)
	Remove(ctx context.Context, p *Page) error
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
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}