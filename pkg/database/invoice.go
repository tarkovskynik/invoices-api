package database

import (
	"errors"
	"invoices"
	"sync"
)

type InvoiceRepository struct {
	//db *sql.DB
	counter int
	data map[int]invoices.Invoice
	sync.Mutex

}

func NewInvoiceRepository() *InvoiceRepository {
	return &InvoiceRepository{
		data: make(map[int]invoices.Invoice),
		counter: 1,
	}
}

func (r *InvoiceRepository) Create(inv *invoices.Invoice){
	//row := r.db.QueryRow()
	r.Lock()
	inv.Id = r.counter
	r.data[inv.Id] = *inv
	r.counter ++
	r.Unlock()
}

func (r *InvoiceRepository) Get(id int) (invoices.Invoice, error){
	r.Lock()
	defer r.Unlock()

	invoice,ok := r.data[id]

	if !ok {
		return invoice, errors.New("invoice not found")
	}

	return invoice,nil
}

func (r *InvoiceRepository) GetAll() ([]invoices.Invoice, error) {
	var result []invoices.Invoice
	r.Lock()
	defer r.Unlock()
	inv := r.data

	for _,value := range inv{
		result = append(result,value)
	}
	return result,nil
}

func (r *InvoiceRepository) Update(id int, inv invoices.Invoice){
	r.Lock()
	r.data[id] = inv
	r.Unlock()
}

func (r *InvoiceRepository) Delete(id int){
	r.Lock()
	defer r.Unlock()
	delete(r.data, id)
}