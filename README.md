# Simple Login & Products API

Developer: Bhenedictho Adriel Natalawa Pampang

## Tech Specifications
### Libraries/Frameworks Used

- Go (version 1.19)
- Gorilla Mux
- PostgreSQL

### Architecture/Modularity

The project follows a clean architecture pattern, separating concerns into different layers:

- **Presentation Layer**: Contains the API handlers and controllers.
- **Service Layer**: Contains the business logic.
- **Repository Layer**: Deals with data storage and retrieval.
- **Database Layer**: Connects to the database.

## Quick Start
### Entity Relationship Diagram
```bash
![DB DESIGN](https://github.com/DitoAdriel99/test-backend-mkp/blob/main/db-design.png)
```
Pada rancangan di atas, setiap produk yang dijual dan penjualan yang terjadi direkam menggunakan FK yang menghubungkan entitas Product dan Sales dengan entitas lainnya. Informasi penjualan mencakup detail penjualan seperti produk yang dijual, customer yang melakukan pembelian, tanggal penjualan, dan jumlah produk yang terjual. Stok produk juga terkait dengan entitas Produk untuk memastikan kontrol stok yang tepat.

Hubungan (Relasi):
Sales-Product:

Relasi antara Sales dan Product dengan kunci asing ProductID.
Sales-Customer:

Relasi antara Sales dan Customer dengan kunci asing CustomerID.

Hubungan antara entitas-entitas diwakili dengan panah yang menunjukkan kunci asing yang menghubungkan entitas-entitas tersebut. Misalnya, Penjualan memiliki kunci asing ProductID dan CustomerID yang menghubungkannya dengan Produk dan Customer.

Hubungan "Sales-Product" menggambarkan bahwa setiap penjualan terkait dengan satu produk, dan hubungan "Sales-Customers" menggambarkan bahwa setiap penjualan terkait dengan satu pelanggan. Semua penjualan mencatat ID product, ID cusomer, tanggal penjualan, dan jumlah produk yang dijual. Stok produk terkait dengan Produk untuk mengelola persediaan.
### API Documentation
#### Login
```bash
curl --location 'localhost:9999/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "Admin",
    "password": "admin"
}'
```
#### Register
```bash
curl --location 'localhost:9999/register' \
--header 'Content-Type: application/json' \
--data '{
    "username": "Admin",
    "password": "admin"
}
'
```
#### Create Product
```bash
curl --location 'localhost:9999/product' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkFkbWluIiwiaXNzIjoiQWRtaW4iLCJleHAiOjE2OTc0NzQ4NDYsImlhdCI6MTY5NzQ2NzY0Nn0.47ELKYZ3ADINMl4Of1yqSCY4GUqnnYqRHDyl8ZeUDLw' \
--data '{
    "product_name": "test1",
    "product_price": 32000,
    "stock": 100
}'
```
#### Get Product
```bash
curl --location 'localhost:9999/product' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkRpdG8iLCJpc3MiOiJEaXRvIiwiZXhwIjoxNjk3NDcxOTY1LCJpYXQiOjE2OTc0NjQ3NjV9.7simuJwrnqjOc-zlugciqh9A1O4cncStcm4CuGzyQeo'
```
#### Create Sales
```bash
curl --location 'localhost:9999/sales' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkRpdG8iLCJpc3MiOiJEaXRvIiwiZXhwIjoxNjk3NDcxOTY1LCJpYXQiOjE2OTc0NjQ3NjV9.7simuJwrnqjOc-zlugciqh9A1O4cncStcm4CuGzyQeo' \
--data '{
    "product_id": "3d000293-6c2c-11ee-b3d9-00ffb41493a0",
    "customer_id": "a242231e-fd14-4de0-82c1-f80a0a5d689b",
    "quantity": 1
}'
```
#### Get Sales
```bash
curl --location 'localhost:9999/sales?search_by=customer_name&search=Michael%20Johnson' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IkRpdG8iLCJpc3MiOiJEaXRvIiwiZXhwIjoxNjk3NDcxOTY1LCJpYXQiOjE2OTc0NjQ3NjV9.7simuJwrnqjOc-zlugciqh9A1O4cncStcm4CuGzyQeo'
```
### Installation guide
#### 1. install go version 1.19
```bash
# Please Read This Link Installation Guide of GO

# Link Download -> https://go.dev/dl/
# Link Install -> https://go.dev/doc/install

```

#### 2. Run the application
```bash
# run command :
git clone https://github.com/DitoAdriel99/test-backend-mkp.git

# install dependency
go mod tidy

# setup env
ENV=

DB_DRIVER=postgres
DB_USERNAME=        #change to your db username
DB_PASSWORD=        #change to your db password
DB_HOST=            #change to your db host
DB_PORT=            #change to your db port 
DB_DATABASE=        #change to your db name 
DB_URL=             #postgres://{DB_USERNAME}:{DB_PASSWORD}@{DB_HOST}:{DB_PORT}/{DB_DATABASE}?sslmode=disable

KEY=                #change to your key
EXPIRED=            #change to your expiration time

# Run App
make start

# Migrate db
make migrate-up //this for up migrations
make migrate-down //this for down migrations
```