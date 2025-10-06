module github/wrelin/ch05

go 1.25.1

replace github.com/wrelin/post05 => ../package

require (
	github.com/go-sql-driver/mysql v1.9.3
	github.com/wrelin/post05 v0.0.0-00010101000000-000000000000
)

require (
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/lib/pq v1.10.9 // indirect
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	golang.org/x/crypto v0.42.0 // indirect
	golang.org/x/text v0.29.0 // indirect
)
