package postgres

import (
	"github.com/g0shi4ek/store/internal/store/domain"
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
)

type ItemsRepository struct{
	pg *pgxpool.Pool
}

func NewItemsRepository(pgx * pgxpool.Pool) domain.ItemRepository{
	return &ItemsRepository{pg: pgx}
}


func (r * ItemsRepository) Create(ctx context.Context, item *domain.Item) error{
	query := "INSERT INTO items (name, amount, provider, price) VALUES ($1, $2, $3, $4) RETURNING id"
	return r.pg.QueryRow(ctx, query, item.Name, item.Amount, item.Provider, item.Price).Scan(&item.Id)
}

func (r * ItemsRepository) Update(ctx context.Context, item *domain.Item) error{
	query := "UPDATE items SET name = $1, amount = $2, provider = $3, price = $4 WHERE id = $5"
	_, err := r.pg.Exec(ctx, query, item.Name, item.Amount, item.Provider, item.Price, item.Id)
	return err
}

func (r * ItemsRepository) Delete(ctx context.Context, id int) error{
	query := "DELETE FROM items WHERE id = $1"
	_, err := r.pg.Exec(ctx, query, id)
	return err
}

func (r * ItemsRepository) GetById(ctx context.Context, id int) (*domain.Item, error){
	var newItem domain.Item
	query := "SELECT id, name, amount, provider, price FROM items WHERE id = $1"
	err := r.pg.QueryRow(ctx, query, id).Scan(&newItem.Id, &newItem.Name, &newItem.Amount, &newItem.Provider, &newItem.Price)
	if err != nil{
		return nil, err
	}
	return &newItem, nil
}

func (r * ItemsRepository) GetByName(ctx context.Context, name string) (*domain.Item, error){
	var newItem domain.Item
	query := "SELECT id, name, amount, provider, price FROM items WHERE name = $1"
	err := r.pg.QueryRow(ctx, query, name).Scan(&newItem.Id, &newItem.Name, &newItem.Amount, &newItem.Provider, &newItem.Price)
	if err != nil{
		return nil, err
	}
	return &newItem, nil
}

func (r *ItemsRepository) GetAll(ctx context.Context) ([]*domain.Item, error) {
	query := "SELECT id, name, amount, provider, price FROM items"
	rows, err := r.pg.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*domain.Item
	for rows.Next() {
		var item domain.Item
		err := rows.Scan(&item.Id, &item.Name, &item.Amount, &item.Provider, &item.Price)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}