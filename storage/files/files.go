package files

import (
	"Archive-Adviser-Bot/lib/e"
	"Archive-Adviser-Bot/storage"
	"encoding/gob"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)


type files struct{
	BasePath string
}

func New(basepath string) files{
	return files{
		BasePath: basepath,
	}
}

func (f *files)Save(p storage.Page) (err error){
	defer func() { err = e.WrapIfErr("cant save", err)}()

	fpath := filepath.Join(f.BasePath, p.UserName)

	if err := os.Mkdir(fpath, 0774); err != nil{
		return err
	}
	fname, err := p.Hash()
	if err != nil{
		return err
	}
	file, err := os.Create(fname)
	if err != nil{
		return err
	}
	defer func(){ _ = file.Close()}()
	err = gob.NewEncoder(file).Encode(p)
	if err != nil{
		return err
	}
	return nil
}

func (f *files)PickRandom(UserName string) (p *storage.Page, err error){
	defer func() { err = e.WrapIfErr("cant pick random item", err)}()
	fpath := filepath.Join(f.BasePath, UserName)
	allfiles, err := os.ReadDir(fpath)
	if err != nil{
		return nil, err
	}
	if len(allfiles) == 0{
		return nil, errors.New("no saved page")
	}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(allfiles))
	file := allfiles[n]
	return f.ReadAndDecode(filepath.Join(fpath, file.Name())) 
}


func (f files)Remove(p *storage.Page) error{
	h, err := p.Hash()
	if err != nil{
		return e.Wrap("cant remove page", err)
	}
	path := filepath.Join(f.BasePath, p.UserName, h)
	err = os.Remove(path)
	if err != nil{
		msg := fmt.Sprintf("cant remove file %s", path)
		return e.Wrap(msg, err) 
	}
	return nil
}

func (f files)IsExist(p *storage.Page) (bool, error){
	h, err := p.Hash()
	if err != nil{
		return false, e.Wrap("cant remove page", err)
	}
	path := filepath.Join(f.BasePath, p.UserName, h)
	_, err = os.Stat(path)
	switch {
	case errors.Is(err, os.ErrNotExist) :
		return false, nil
	case err == nil:
		return true, nil
	default:
		return false, e.Wrap("cant remove page", err)
	}

}

func (file files)ReadAndDecode(filepath string) (*storage.Page, error){
	f, err := os.Open(filepath)
	if err != nil{
		return nil, e.Wrap("Cant open file", err)
	}
	var pa storage.Page
	if err = gob.NewDecoder(f).Decode(&pa); err != nil{
		return nil, e.Wrap("cant Decode file", err)
	}
	return &pa, nil
}