package cache

import (
	"errors"
	"invoices"
	"sync"
)

type Cache struct {
	//db *sql.DB
	counter int
	data map[int]invoices.Invoice
	sync.Mutex

}

func NewCache() *Cache {
	return &Cache{
		data: make(map[int]invoices.Invoice),
		counter: 1,
	}
}

func (c *Cache) Create(inv *invoices.Invoice){
	//row := c.db.QueryRow()
	c.Lock()
	inv.Id = c.counter
	c.data[inv.Id] = *inv
	c.counter ++
	c.Unlock()
}

func (c *Cache) Get(id int) (invoices.Invoice, error){
	c.Lock()
	defer c.Unlock()

	invoice,ok := c.data[id]

	if !ok {
		return invoice, errors.New("invoice not found")
	}

	return invoice,nil
}

func (c *Cache) GetAll() ([]invoices.Invoice, error) {
	var result []invoices.Invoice
	c.Lock()
	defer c.Unlock()
	inv := c.data

	for _,value := range inv{
		result = append(result,value)
	}
	return result,nil
}

func (c *Cache) Update(id int, inv invoices.Invoice){
	c.Lock()
	c.data[id] = inv
	c.Unlock()
}

func (c *Cache) Delete(id int){
	c.Lock()
	defer c.Unlock()
	delete(c.data, id)
}