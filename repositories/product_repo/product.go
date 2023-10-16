package product_repo

import (
	"database/sql"
	"fmt"
	"go-learn/entities"
	"go-learn/library/meta"
	"strings"

	"github.com/google/uuid"
)

func (c *_ProductRepoImp) Create(pr *entities.Product) error {

	queryInsert := `INSERT INTO products (id, product_name, product_price, stock, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := c.conn.Exec(queryInsert, pr.ID, pr.ProductName, pr.ProductPrice, pr.Stock, pr.CreatedAt, pr.UpdatedAt)
	if err != nil {
		err = fmt.Errorf("executing query: %w", err)
		return err
	}
	return nil
}

func (c *_ProductRepoImp) Detail(id uuid.UUID) (*entities.Product, error) {
	query := `
		SELECT 
			id, 
			product_name, 
			product_price, 
			stock, 
			created_at, 
			updated_at 
		FROM 
			products
		WHERE 
			id = $1`

	var object entities.Product

	err := c.conn.QueryRow(query, id).Scan(
		&object.ID,
		&object.ProductName,
		&object.ProductPrice,
		&object.Stock,
		&object.CreatedAt,
		&object.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		err = fmt.Errorf("scanning activity objects: %w", err)
		return nil, err
	}

	return &object, nil
}

func (c *_ProductRepoImp) GetAll(m *meta.Metadata) ([]entities.Product, error) {
	q, err := meta.ParseMetaData(m, c)
	if err != nil {
		return nil, err
	}
	stmt := `SELECT id, product_name, product_price, stock, created_at, updated_at FROM products
	`
	queries := QueryStatement(stmt)
	var (
		searchBy = q.SearchBy
		order    = q.OrderBy
	)
	if len(q.Search) > 2 {
		if len(q.SearchBy) != 0 {
			queries = queries.Where(searchBy, like, q.Search)
		}
	}
	if q.DateEnd.Valid && q.DateFrom.Valid {
		queries = queries.Where(order, between, q.DateFrom.Time.Local(), q.DateEnd.Time.Local())
	}

	query, _, args := queries.Order(order, direction(strings.ToUpper(q.OrderDirection))).
		Offset(q.Offset).
		Limit(q.Limit).Build()

	rows, err := c.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	collections := make([]entities.Product, 0)
	for rows.Next() {
		var f entities.Product
		if err := rows.Scan(
			&f.ID,
			&f.ProductName,
			&f.ProductPrice,
			&f.Stock,
			&f.CreatedAt,
			&f.UpdatedAt,
		); err != nil {
			return nil, err
		}

		collections = append(collections, f)
	}

	return collections, nil
}

func (c *_ProductRepoImp) Sortable(field string) bool {
	switch field {
	case "created_at", "updated_at":
		return true
	default:
		return false
	}
}

func (c *_ProductRepoImp) CreateSales(payload entities.SalesPayload, reqQuantity int) error {
	tx, err := c.conn.Begin()
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	insertQuery := `
		INSERT INTO sales (
			id,
			product_id,
			customer_id,
			quantity,
			created_at,
			updated_at
		)VALUES(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		)
	`

	_, err = tx.Exec(insertQuery, payload.ID, payload.ProductsID, payload.CustomerID, payload.Quantity, payload.CreatedAt, payload.UpdatedAt)
	if err != nil {
		err = fmt.Errorf("executing query insert: %w", err)
		return err
	}

	updateQuery := `
		UPDATE 
			products
		SET
			stock = $2,
			updated_at = $3
		WHERE
			id = $1
	`

	_, err = tx.Exec(updateQuery, payload.ProductsID, reqQuantity, payload.UpdatedAt)
	if err != nil {
		err = fmt.Errorf("executing query insert: %w", err)
		return err
	}
	return nil
}

func (c *_ProductRepoImp) GetSales(m *meta.Metadata) ([]entities.SalesResponse, error) {
	q, err := meta.ParseMetaData(m, c)
	if err != nil {
		return nil, err
	}
	stmt := `SELECT 
				s.id, 
				p.product_name,
				c.customer_name,
				s.quantity,
				s.created_at, 
				s.updated_at 
			FROM 
				sales s
			JOIN
				products p
			ON 
				p.id = s.product_id
			JOIN 
				customers c 
			ON 
				c.id = s.customer_id
	`
	queries := QueryStatement(stmt)
	var (
		searchBy = q.SearchBy
		order    = q.OrderBy
	)
	if len(q.Search) > 2 {
		if len(q.SearchBy) != 0 {
			queries = queries.Where(searchBy, like, q.Search)
		}
	}
	if q.DateEnd.Valid && q.DateFrom.Valid {
		queries = queries.Where(order, between, q.DateFrom.Time.Local(), q.DateEnd.Time.Local())
	}

	query, _, args := queries.Order(order, direction(strings.ToUpper(q.OrderDirection))).
		Offset(q.Offset).
		Limit(q.Limit).Build()

	rows, err := c.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	collections := make([]entities.SalesResponse, 0)
	for rows.Next() {
		var f entities.SalesResponse
		if err := rows.Scan(
			&f.ID,
			&f.ProductName,
			&f.CustomerName,
			&f.Quantity,
			&f.CreatedAt,
			&f.UpdatedAt,
		); err != nil {
			return nil, err
		}

		collections = append(collections, f)
	}

	return collections, nil
}
