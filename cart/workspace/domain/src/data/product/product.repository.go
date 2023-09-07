package product

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"pensiel.com/material/src/pensiel"
	"pensiel.com/material/src/static"
)

type Repository interface {
	GetProduct(ctx context.Context, productId string) (*Response, *pensiel.Error)
}

type repository struct {
	host string
}

func NewRepository() Repository {
	return &repository{
		host: static.PRODUCT_SERVICE,
	}
}

func (r *repository) GetProduct(ctx context.Context, productId string) (*Response, *pensiel.Error) {
	url, err := url.ParseRequestURI(r.host)

	response := &Response{}

	if err != nil {
		return nil, &pensiel.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed parse request uri",
			Origin:     err,
		}
	}

	url.Path = fmt.Sprintf("/api/v1/product/%s", productId)

	resp, err := http.Get(url.String())

	if err != nil {
		return nil, &pensiel.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed request to product service",
			Origin:     err,
		}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, &pensiel.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed read response",
			Origin:     err,
		}
	}

	if err := json.Unmarshal(body, response); err != nil {
		return nil, &pensiel.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed unmarshal response",
			Origin:     err,
		}
	}

	if resp.StatusCode != 200 {
		return nil, &pensiel.Error{
			StatusCode: resp.StatusCode,
			Message:    response.Message,
		}
	}

	return response, nil
}
