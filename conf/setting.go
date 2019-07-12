package conf

import "time"

const (
	// debug or release
	RUN_MODE = "debug"

	// [app]
	PAGE_SIZE  = 20
	JWT_SECRET = "CardTask"

	// [server]
	HTTP_PORT     = 8080
	READ_TIMEOUT  = time.Duration(60) * time.Second
	WRITE_TIMEOUT = time.Duration(60) * time.Second

	// [database]
	DB_TYPE     = "postgres"
	DB_USER     = "ki"
	DB_PASSWORD = "1234go"
	// 127.0.0.1:3306
	DB_HOST         = "127.0.0.1"
	DB_PORT         = "5432"
	DB_NAME         = "Card"
	DB_TABLE_PREFIX = "Card_"

	//CORS

)
