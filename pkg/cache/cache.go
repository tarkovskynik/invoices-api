package cache

import (
	"errors"
	"invoices"
	"sync"
)

type Cache struct {
	data map[int]invoices.Invoice
	sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[int]invoices.Invoice),
	}
}

func (c *Cache) CreateCache(inv *invoices.Invoice) {
	c.Lock()
	c.data[inv.Id] = *inv
	c.Unlock()
}

func (c *Cache) GetCache(id int) (invoices.Invoice, error) {
	c.Lock()
	defer c.Unlock()

	invoice, ok := c.data[id]

	if !ok {
		return invoice, errors.New("invoice not found")
	}

	return invoice, nil
}

func (c *Cache) UpdateCache(id int, inv invoices.Invoice) {
	c.Lock()
	inv.Id = id
	c.data[id] = inv
	c.Unlock()
}

func (c *Cache) DeleteCache(id int) {
	c.Lock()
	defer c.Unlock()
	delete(c.data, id)
}
