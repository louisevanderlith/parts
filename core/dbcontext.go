package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
)

type PartsContext interface {
	GetSpare(key hsk.Key) (Spare, error)
	FindLatestSpares(page, size int) (records.Page, error)
	CreateSpare(obj Spare) (hsk.Key, error)
	UpdateSpare(key hsk.Key, obj Spare) error
}

func CreateContext() PartsContext {
	ctx = context{
		Spares: husk.NewTable(Spare{}),
	}

	return ctx
}

func Shutdown() {
	ctx.Spares.Save()
}

func Context() PartsContext {
	return ctx
}

type context struct {
	Spares husk.Table
}

var ctx context

func (c context) GetSpare(key hsk.Key) (Spare, error) {
	rec, err := c.Spares.FindByKey(key)

	if err != nil {
		return Spare{}, err
	}

	return rec.GetValue().(Spare), nil
}

func (c context) FindLatestSpares(page, size int) (records.Page, error) {
	return c.Spares.Find(page, size, op.Everything())
}

func (c context) CreateSpare(obj Spare) (hsk.Key, error) {
	return c.Spares.Create(obj)
}

func (c context) UpdateSpare(key hsk.Key, obj Spare) error {
	return c.Spares.Update(key, obj)
}
