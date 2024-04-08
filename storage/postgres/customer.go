package postgres

import (
	"backend_course/rent_car/api/models"
	"backend_course/rent_car/pkg/logger"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type CustomerRepo struct {
	db     *pgxpool.Pool
	logger logger.ILogger
}

func NewCustomerRepo(db *pgxpool.Pool, log logger.ILogger) CustomerRepo {
	return CustomerRepo{
		db:     db,
		logger: log,
	}
}

func (c *CustomerRepo) Create(ctx context.Context, customer models.CreateCustomer) (string, error) {
	id := uuid.New().String()

	query := `INSERT INTO customers (
        id,
        first_name,
        last_name,
        email,
        phone,
		login,
		password,
        address,
        created_at,
        updated_at
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`

	_, err := c.db.Exec(ctx, query,
		id,
		customer.FirstName,
		customer.LastName,
		customer.Email,
		customer.Phone,
		customer.Login,
		customer.Password,
		customer.Address,
	)

	if err != nil {
		c.logger.Error("failed to create customer in database", logger.Error(err))
		return "", err
	}

	return id, nil
}

func (c *CustomerRepo) Update(ctx context.Context, customer models.UpdateCustomer, id string) (string, error) {
	query := `UPDATE customers SET
        first_name = $1,
        last_name = $2,
        email = $3,
        phone = $4,
        address = $5,
        updated_at = $6
    WHERE id = $7`

	_, err := c.db.Exec(ctx, query,
		customer.FirstName,
		customer.LastName,
		customer.Email,
		customer.Phone,
		customer.Address,
		time.Now(),
		id,
	)

	if err != nil {
		c.logger.Error("failed to update customer in database", logger.Error(err))
		return "", err
	}

	return id, nil
}

func (c *CustomerRepo) Login(ctx context.Context, login models.LoginCustomer) (string, error) {
	var hashedPass string

	query := `SELECT password
	FROM customers
	WHERE login = $1 AND deleted_at = 0`

	err := c.db.QueryRow(ctx, query,
		login.Phone,
	).Scan(&hashedPass)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("incorrect login")
		}
		c.logger.Error("failed to get customer password from database", logger.Error(err))
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(login.Password))
	if err != nil {
		return "", errors.New("password mismatch")
	}

	return "Logged in successfully", nil
}

func (c *CustomerRepo) ChangePassword(ctx context.Context, pass models.ChangePassword) (string, error) {
	var hashedPass string

	query := `SELECT password
	FROM customers
	WHERE login = $1 AND deleted_at = 0`

	err := c.db.QueryRow(ctx, query,
		pass.Login,
	).Scan(&hashedPass)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("incorrect login")
		}
		c.logger.Error("failed to get customer password from database", logger.Error(err))
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass.OldPassword))
	if err != nil {
		return "", errors.New("password mismatch")
	}

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.logger.Error("failed to generate customer new password", logger.Error(err))
		return "", err
	}

	query = `UPDATE customers SET 
		password = $1, 
		updated_at = CURRENT_TIMESTAMP 
	WHERE login = $2 AND deleted_at = 0`

	_, err = c.db.Exec(ctx, query, newHashedPassword, pass.Login)
	if err != nil {
		c.logger.Error("failed to change customer password in database", logger.Error(err))
		return "", err
	}

	return "Password changed successfully", nil
}

func (c *CustomerRepo) GetByID(ctx context.Context, id string) (models.Customer, error) {
	var (
		firstname       sql.NullString
		lastname        sql.NullString
		phone           sql.NullString
		email           sql.NullString
		address         sql.NullString
		createdat       sql.NullString
		updatedat       sql.NullString
		uniquecarscount sql.NullInt64
		orderscount     sql.NullInt64
	)

	query := `SELECT 
		id, 
		first_name, 
		last_name, 
		phone,
		email,
		address,
		created_at, 
		updated_at
		FROM customers WHERE id = $1 AND deleted_at = 0`

	row := c.db.QueryRow(ctx, query, id)

	customer := models.Customer{
		Orders: []models.Order{},
	}

	err := row.Scan(
		&customer.ID,
		&firstname,
		&lastname,
		&phone,
		&email,
		&address,
		&createdat,
		&updatedat,
	)

	if err != nil {
		c.logger.Error("failed to scan customer by ID from database", logger.Error(err))
		return models.Customer{}, err
	}

	customer.FirstName = firstname.String
	customer.LastName = lastname.String
	customer.Phone = phone.String
	customer.Email = email.String
	customer.Address = address.String
	customer.CreatedAt = createdat.String
	customer.UpdatedAt = updatedat.String

	orderQuery := `SELECT
		id,
		from_date,
		to_date,
		status,
		payment_status,
		created_at,
		updated_at
		FROM orders WHERE customer_id = $1`

	rows, err := c.db.Query(ctx, orderQuery, id)

	if err != nil {
		c.logger.Error("failed to get customer (orders) by ID from database", logger.Error(err))
		return models.Customer{}, err
	}
	defer rows.Close()

	var (
		order     models.Order
		fromdate  sql.NullString
		todate    sql.NullString
		status    sql.NullString
		paid      sql.NullBool
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	for rows.Next() {
		err := rows.Scan(
			&order.Id,
			&fromdate,
			&todate,
			&status,
			&paid,
			&createdAt,
			&updatedAt,
		)

		order.FromDate = fromdate.String
		order.ToDate = todate.String
		order.Status = status.String
		order.Paid = paid.Bool
		order.CreatedAt = createdAt.String
		order.UpdatedAt = updatedAt.String

		customer.Orders = append(customer.Orders, order)

		if err != nil {
			c.logger.Error("failed to scan customer orders from database", logger.Error(err))
			return models.Customer{}, err
		}
	}

	ordersCount := `SELECT
		COUNT(o.id)
		FROM orders o
		JOIN customers c ON c.id = o.customer_id
		WHERE c.id = $1`

	err = c.db.QueryRow(ctx, ordersCount, id).Scan(&orderscount)
	customer.OrdersCount = orderscount.Int64
	if err != nil {
		c.logger.Error("failed to get customer orders count from database", logger.Error(err))
		return models.Customer{}, err
	}

	uniqueCarsCount := `SELECT
		COUNT(DISTINCT o.car_id)
		FROM cars c
		JOIN orders o ON c.id = o.car_id
		WHERE o.customer_id = $1`

	err = c.db.QueryRow(ctx, uniqueCarsCount, id).Scan(&uniquecarscount)
	customer.UniqueCarsCount = uniquecarscount.Int64
	if err != nil {
		c.logger.Error("failed to get customer unique cars count from database", logger.Error(err))
		return models.Customer{}, err
	}

	return customer, nil
}

func (c *CustomerRepo) GetByLogin(ctx context.Context, login string) (models.Customer, error) {
	var (
		firstname sql.NullString
		lastname  sql.NullString
		phone     sql.NullString
		email     sql.NullString
		address   sql.NullString
		createdat sql.NullString
		updatedat sql.NullString
	)

	query := `SELECT 
		id, 
		first_name, 
		last_name, 
		phone,
		email,
		address,
		created_at, 
		updated_at,
		password
		FROM customers WHERE phone = $1 AND deleted_at = 0`

	row := c.db.QueryRow(ctx, query, login)

	customer := models.Customer{
		Orders: []models.Order{},
	}

	err := row.Scan(
		&customer.ID,
		&firstname,
		&lastname,
		&phone,
		&email,
		&address,
		&createdat,
		&updatedat,
		&customer.Password,
	)

	if err != nil {
		c.logger.Error("failed to scan customer by LOGIN from database", logger.Error(err))
		return models.Customer{}, err
	}

	customer.FirstName = firstname.String
	customer.LastName = lastname.String
	customer.Phone = phone.String
	customer.Email = email.String
	customer.Address = address.String
	customer.CreatedAt = createdat.String
	customer.UpdatedAt = updatedat.String

	return customer, nil
}

func (c *CustomerRepo) GetAll(ctx context.Context, req models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error) {
	var (
		resp      = models.GetAllCustomersResponse{}
		filter    string
		firstname sql.NullString
		lastname  sql.NullString
		phone     sql.NullString
		email     sql.NullString
		address   sql.NullString
		createdat sql.NullString
		updatedat sql.NullString
		count     sql.NullInt64
	)
	offset := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter = fmt.Sprintf(` AND (first_name ILIKE '%%%v%%' OR last_name ILIKE '%%%v%%' OR phone ILIKE '%%%v%%')`, req.Search, req.Search, req.Search)
	}

	filter += fmt.Sprintf(" OFFSET %v LIMIT %v", offset, req.Limit)

	query := `SELECT 
        id, 
        first_name, 
        last_name, 
        email,
        phone,
        address,
        created_at, 
        updated_at
        FROM customers WHERE deleted_at = 0` + filter

	rows, err := c.db.Query(ctx, query)
	if err != nil {
		c.logger.Error("failed to get all customers from database", logger.Error(err))
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var customer models.Customer

		err := rows.Scan(
			&customer.ID,
			&firstname,
			&lastname,
			&email,
			&phone,
			&address,
			&createdat,
			&updatedat,
		)
		if err != nil {
			c.logger.Error("failed to scan customers from database", logger.Error(err))
			return models.GetAllCustomersResponse{}, err
		}

		id := customer.ID

		customer.FirstName = firstname.String
		customer.LastName = lastname.String
		customer.Email = email.String
		customer.Phone = phone.String
		customer.Address = address.String
		customer.CreatedAt = createdat.String
		customer.UpdatedAt = updatedat.String

		customer.Orders = make([]models.Order, 0)

		orderQuery := `SELECT
            id,
            from_date,
            to_date,
            status,
            payment_status,
            created_at,
            updated_at
            FROM orders WHERE customer_id = $1`

		orderRows, err := c.db.Query(ctx, orderQuery, id)

		if err != nil {
			c.logger.Error("failed to get  customers orders from database", logger.Error(err))
			return models.GetAllCustomersResponse{}, err
		}
		defer orderRows.Close()

		var (
			orders          []models.Order
			fromdate        sql.NullString
			todate          sql.NullString
			status          sql.NullString
			paid            sql.NullBool
			createdAt       sql.NullString
			updatedAt       sql.NullString
			orderscount     sql.NullInt64
			uniquecarscount sql.NullInt64
		)

		for orderRows.Next() {
			var order models.Order
			err := orderRows.Scan(
				&order.Id,
				&fromdate,
				&todate,
				&status,
				&paid,
				&createdAt,
				&updatedAt,
			)

			order.FromDate = fromdate.String
			order.ToDate = todate.String
			order.Status = status.String
			order.Paid = paid.Bool
			order.CreatedAt = createdAt.String
			order.UpdatedAt = updatedAt.String

			if err != nil {
				c.logger.Error("failed to scan customers orders from database", logger.Error(err))
				return models.GetAllCustomersResponse{}, err
			}

			orders = append(orders, order)
		}

		customer.Orders = orders

		if customer.Orders != nil {
			ordersCountQuery := `SELECT COUNT(o.car_id) FROM orders o WHERE o.customer_id = $1`
			err = c.db.QueryRow(ctx, ordersCountQuery, id).Scan(&orderscount)
			customer.OrdersCount = orderscount.Int64
			if err != nil {
				c.logger.Error("failed to get customers car count from database", logger.Error(err))
				return models.GetAllCustomersResponse{}, err
			}
		}

		uniqueCarsCountQuery := `SELECT COUNT(DISTINCT o.car_id) FROM orders o WHERE o.customer_id = $1`
		err = c.db.QueryRow(ctx, uniqueCarsCountQuery, id).Scan(&uniquecarscount)
		customer.UniqueCarsCount = uniquecarscount.Int64
		if err != nil {
			c.logger.Error("failed to get customers car unique cars count from database", logger.Error(err))
			return resp, err
		}

		resp.Customers = append(resp.Customers, customer)
	}

	countQuery := `SELECT COUNT(id) FROM customers`
	err = c.db.QueryRow(ctx, countQuery).Scan(&count)
	resp.Count = count.Int64
	if err != nil {
		c.logger.Error("failed to get customers count from database", logger.Error(err))
		return models.GetAllCustomersResponse{}, err
	}

	return resp, nil
}

func (c *CustomerRepo) GetCustomerCars(ctx context.Context, name string, id string, boolean bool) (models.GetCustomerCarsResponse, error) {
	var (
		resp              models.GetCustomerCarsResponse
		customerCarsCount string
		query             string
		args              []interface{}
	)

	if boolean {
		query = `SELECT
            c.name,
            o.created_at,
            o.from_date,
            o.to_date,
            c.price
            FROM cars c
            INNER JOIN orders o ON c.id = o.car_id
            WHERE o.customer_id = $1`
		args = append(args, id)
	} else {
		query = `SELECT
            c.name,
            o.created_at,
            o.from_date,
            o.to_date,
            c.price
            FROM cars c
            INNER JOIN orders o ON c.id = o.car_id
            INNER JOIN customers cu ON cu.id = o.customer_id
            WHERE cu.id = $1 AND c.name = $2`
		args = append(args, id, name)
	}

	//fmt.Println(query)

	rows, err := c.db.Query(ctx, query, args...)
	if err != nil {
		c.logger.Error("failed to get customer cars from database", logger.Error(err))
		return models.GetCustomerCarsResponse{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			customerCar models.GetCustomerCars
			carName     sql.NullString
			createdAt   sql.NullString
			fromDate    sql.NullString
			toDate      sql.NullString
			price       sql.NullFloat64
		)

		err := rows.Scan(
			&carName,
			&createdAt,
			&fromDate,
			&toDate,
			&price,
		)

		if err != nil {
			c.logger.Error("failed to scan customer cars from database", logger.Error(err))
			return models.GetCustomerCarsResponse{}, err
		}

		customerCar.Duration, err = Duration(fromDate.String, toDate.String)
		if err != nil {
			c.logger.Error("failed to get duration", logger.Error(err))
			return models.GetCustomerCarsResponse{}, err
		}

		if price.Valid {
			customerCar.Price = price.Float64
		} else {
			customerCar.Price = 0
		}

		resp.CustomerCars = append(resp.CustomerCars, customerCar)
	}

	args = args[:0]

	if boolean {
		customerCarsCount = `SELECT
            COUNT(c.name)
            FROM cars c
            JOIN orders o ON c.id = o.car_id
            WHERE o.customer_id = $1`
		args = append(args, id)
	} else {
		customerCarsCount = `SELECT
            COUNT(c.name)
            FROM cars c
            JOIN orders o ON c.id = o.car_id
            JOIN customers cu ON cu.id = o.customer_id
            WHERE cu.id = $1 AND c.name = $2`
		args = append(args, id, name)
	}

	err = c.db.QueryRow(ctx, customerCarsCount, args...).Scan(&resp.Count)
	if err != nil {
		c.logger.Error("failed to get count of customer cars from database", logger.Error(err))
		return models.GetCustomerCarsResponse{}, err
	}

	return resp, nil
}

func (c *CustomerRepo) Delete(ctx context.Context, id string) error {
	query := `UPDATE customers SET deleted_at = date_part('epoch', CURRENT_TIMESTAMP)::int WHERE id = $1 AND deleted_at = 0`

	_, err := c.db.Exec(ctx, query, id)
	if err != nil {
		c.logger.Error("failed to delete customer from database", logger.Error(err))
		return err
	}

	return nil
}

func Duration(fD, tD string) (duration float64, err error) {
	fromDate, err := time.Parse(time.RFC3339, fD)
	if err != nil {
		return 0, err
	}

	toDate, err := time.Parse(time.RFC3339, tD)
	if err != nil {
		return 0, err
	}

	duration = float64(toDate.Sub(fromDate).Hours() / 24)
	return duration, nil
}

func (c *CustomerRepo) DeleteHard(ctx context.Context, id string) error {
	query := `DELETE FROM customers WHERE id = $1`

	_, err := c.db.Exec(ctx, query, id)
	if err != nil {
		c.logger.Error("failed to hard delete customer from database", logger.Error(err))
		return err
	}

	return nil
}

func (c *CustomerRepo) GetPassword(ctx context.Context, phone string) (string, error) {
	var hashedPass string

	query := `SELECT password
	FROM customers
	WHERE phone = $1 AND deleted_at = 0`

	err := c.db.QueryRow(ctx, query, phone).Scan(&hashedPass)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("incorrect phone")
		} else {
			c.logger.Error("failed to get customer password from database", logger.Error(err))
			return "", err
		}
	}

	return hashedPass, nil
}
