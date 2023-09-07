package static

import (
	"os"
	"strings"
)

var (
	PORT         = strings.TrimSpace(os.Getenv("PORT"))
	MODE         = strings.TrimSpace(os.Getenv("MODE"))
	ENVIRONTMENT = strings.TrimSpace(os.Getenv("ENVIRONTMENT"))

	COOKIE_SECURE = strings.TrimSpace(os.Getenv("COOKIE_SECURE")) == "true"

	DATABASE_HOST     = strings.TrimSpace(os.Getenv("DATABASE_HOST"))
	DATABASE_PORT     = strings.TrimSpace(os.Getenv("DATABASE_PORT"))
	DATABASE_USER     = strings.TrimSpace(os.Getenv("DATABASE_USER"))
	DATABASE_PASSWORD = strings.TrimSpace(os.Getenv("DATABASE_PASSWORD"))
	DATABASE_NAME     = strings.TrimSpace(os.Getenv("DATABASE_NAME"))

	PRODUCT_SERVICE = strings.TrimSpace(os.Getenv("PRODUCT_SERVICE"))
	PROMO_SERVICE   = strings.TrimSpace(os.Getenv("PROMO_SERVICE"))

	HASH_SALT = strings.TrimSpace(os.Getenv("HASH_SALT"))
	HASH_COST = strings.TrimSpace(os.Getenv("HASH_COST"))
)
