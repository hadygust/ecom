package orders

import (
	"context"
	"errors"
	"fmt"

	repo "github.com/hadygust/ecom/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductNoStock  = errors.New("product have not enough stock")
	ErrItemNoQuantity  = errors.New("quantity must be at least 1 for each choosen item")
)

type Service interface {
	PlaceOrder(ctx context.Context, order createOrderParams) (repo.Order, error)
}

type svc struct {
	repo *repo.Queries
	db   *pgx.Conn
}

func NewService(repo *repo.Queries, db *pgx.Conn) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) PlaceOrder(ctx context.Context, order createOrderParams) (repo.Order, error) {

	if order.CustomerId == 0 {
		return repo.Order{}, fmt.Errorf("customer ID is required")
	}

	if len(order.OrderItems) <= 0 {
		return repo.Order{}, fmt.Errorf("Must include at least 1 order items")
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Order{}, err
	}
	defer tx.Rollback(ctx)

	qtx := s.repo.WithTx(tx)

	in_order, err := s.repo.CreateOrder(ctx, order.CustomerId)
	if err != nil {
		return repo.Order{}, err
	}
	defer tx.Rollback(ctx)

	for _, item := range order.OrderItems {
		product, err := s.repo.FindProductByID(ctx, item.ProductId)
		if err != nil {
			return repo.Order{}, ErrProductNotFound
		}
		if item.Quantity <= 0 {
			return repo.Order{}, ErrItemNoQuantity
		}
		if product.Quantity-item.Quantity < 0 {
			return repo.Order{}, ErrProductNoStock
		}
		defer tx.Rollback(ctx)

		_, err = qtx.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID:     in_order.ID,
			ProductID:   product.ID,
			Quantity:    item.Quantity,
			PriceInCent: product.PriceInCent * item.Quantity,
		})
		if err != nil {
			return repo.Order{}, err
		}
		defer tx.Rollback(ctx)

		_, err = qtx.UpdateProductStock(ctx, repo.UpdateProductStockParams{
			Quantity: product.Quantity - item.Quantity,
			ID:       product.ID,
		})
		if err != nil {
			return repo.Order{}, err
		}
		defer tx.Rollback(ctx)
	}

	tx.Commit(ctx)
	return in_order, nil
}
