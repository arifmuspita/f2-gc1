package repository

import (
	"database/sql"
	"f2_gc1/model"
)

type ICustomerRepository interface {
	Create(customer *model.Customer) error
	Update(customer *model.Customer) error
	Delete(id int) error
	GetAll() ([]*model.Customer, error)
	GetByID(id int) (*model.Customer, error)
}

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) ICustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (c *customerRepository) Create(customer *model.Customer) error {
	_, err := c.db.Exec("INSERT INTO customers (name, email, phone) values (?,?,?)", customer.Name, customer.Email, customer.Phone)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) Update(customer *model.Customer) error {
	_, err := c.db.Exec("UPDATE customers SET name=?, email=?, phone=? WHERE id=?", customer.Name, customer.Email, customer.Phone, customer.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) Delete(id int) error {
	_, err := c.db.Exec("DELETE FROM customers WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) GetAll() ([]*model.Customer, error) {
	rows, err := c.db.Query("SELECT id, name, email, phone FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []*model.Customer
	for rows.Next() {
		var customer model.Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Phone)
		if err != nil {
			return nil, err
		}
		customers = append(customers, &customer)
	}
	return customers, nil
}

func (c *customerRepository) GetByID(id int) (*model.Customer, error) {
	var customer model.Customer
	err := c.db.QueryRow("SELECT id, name, email, phone FROM customers WHERE id = ?", id).
		Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Phone)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
