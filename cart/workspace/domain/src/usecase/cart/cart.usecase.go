package cart

import (
	"context"
	"net/http"
	"sync"

	"pensiel.com/domain/src/data/cart"
	cartproduct "pensiel.com/domain/src/data/cart-product"
	"pensiel.com/domain/src/data/product"
	"pensiel.com/domain/src/data/promo"
	"pensiel.com/material/src/client"
	"pensiel.com/material/src/pensiel"
)

type Usecase interface {
	RemoveProductFromCart(ctx context.Context, data RemoveProductFormCartParam) *pensiel.Error
	CartPromo(ctx context.Context, data CartPromoParam) *pensiel.Error
	UpdateQuantityProductCart(ctx context.Context, data UpdateQuantityProductCartParam) *pensiel.Error
	AddProductIntoCart(ctx context.Context, data AddProductIntoCartParam) *pensiel.Error
	GetCartDetail(ctx context.Context, data GetCartDetailParam) (*GetCartDetailReturn, *pensiel.Error)
}

type usecase struct {
	*client.Client
	*Repository
}

type Repository struct {
	Product     product.Repository
	Cart        cart.Repository
	CartProduct cartproduct.Repository
	Promo       promo.Repository
}

func NewService(c *client.Client, r *Repository) Usecase {
	return &usecase{
		Client:     c,
		Repository: r,
	}
}

func (uc *usecase) AddProductIntoCart(ctx context.Context, data AddProductIntoCartParam) *pensiel.Error {
	err := uc.Dbi.Trx(ctx, func(tx context.Context) *pensiel.Error {
		var wg sync.WaitGroup
		var err *pensiel.Error
		var product *product.Response

		cart := &cart.EntityModel{
			Entity: cart.Entity{
				Invoice: data.Invoice,
			},
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			products, errs := uc.Repository.Product.GetProduct(tx, data.ProductId)

			if errs != nil {
				err = errs
				return
			}

			product = products
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			errs := uc.Repository.Cart.FindByInvoice(tx, cart)

			if errs != nil {
				err = errs
				return
			}
		}()

		wg.Wait()

		if err != nil {
			if err.Message != "cart not found" {
				return err
			}

			if err := uc.Cart.Insert(tx, cart); err != nil {
				return err
			}
		}

		if product.Data.Quatity <= 0 {
			return &pensiel.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "Quantity of the Product is not available",
			}
		}

		err = uc.CartProduct.InsertOrUpdateByCartIdAndProductId(tx, &cartproduct.EntityModel{
			Entity: cartproduct.Entity{
				CartId:    cart.SecondaryId,
				ProductId: data.ProductId,
				Quantity:  data.Quantity,
			},
		})

		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (uc *usecase) UpdateQuantityProductCart(ctx context.Context, data UpdateQuantityProductCartParam) *pensiel.Error {
	err := uc.Dbi.Trx(ctx, func(tx context.Context) *pensiel.Error {
		var wg sync.WaitGroup
		var err *pensiel.Error
		var product *product.Response

		cart := &cart.EntityModel{
			Entity: cart.Entity{
				Invoice: data.Invoice,
			},
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			products, errs := uc.Repository.Product.GetProduct(tx, data.ProductId)

			if errs != nil {
				err = errs
				return
			}

			product = products
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			errs := uc.Repository.Cart.FindByInvoice(tx, cart)

			if errs != nil {
				err = errs
				return
			}
		}()

		wg.Wait()

		if err != nil {
			return err
		}

		cartProduct := &cartproduct.EntityModel{
			Entity: cartproduct.Entity{
				CartId:    cart.SecondaryId,
				ProductId: data.ProductId,
			},
		}

		if err := uc.Repository.CartProduct.FindByIndex(tx, cartProduct); err != nil {
			return err
		}

		if product.Data.Quatity < data.Quantity {
			return &pensiel.Error{
				StatusCode: http.StatusBadRequest,
				Message:    "Quantity of the Product is not available",
			}
		}

		cartProduct.Quantity = data.Quantity

		if err := uc.CartProduct.Update(tx, cartProduct); err != nil {
			return err
		}

		return nil
	})

	return err
}

func (uc *usecase) CartPromo(ctx context.Context, data CartPromoParam) *pensiel.Error {
	err := uc.Dbi.Trx(ctx, func(tx context.Context) *pensiel.Error {
		var wg sync.WaitGroup
		var err *pensiel.Error
		var promo *promo.Response

		cart := &cart.EntityModel{
			Entity: cart.Entity{
				Invoice: data.Invoice,
			},
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			promos, errs := uc.Repository.Promo.GetPromo(ctx, data.PromoCode)

			if errs != nil {
				err = errs
				return
			}

			promo = promos
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			errs := uc.Repository.Cart.FindByInvoice(tx, cart)

			if errs != nil {
				err = errs
				return
			}
		}()

		wg.Wait()

		if err != nil {
			return err
		}

		cart.PromoCode = promo.Data.Code

		if err := uc.Cart.Update(tx, cart); err != nil {
			return err
		}

		return nil
	})

	return err
}

func (uc *usecase) RemoveProductFromCart(ctx context.Context, data RemoveProductFormCartParam) *pensiel.Error {
	err := uc.Dbi.Trx(ctx, func(tx context.Context) *pensiel.Error {
		cart := &cart.EntityModel{
			Entity: cart.Entity{
				Invoice: data.Invoice,
			},
		}

		if err := uc.Repository.Cart.FindByInvoice(tx, cart); err != nil {
			return err
		}

		cartProduct := &cartproduct.EntityModel{
			Entity: cartproduct.Entity{
				CartId:    cart.SecondaryId,
				ProductId: data.ProductId,
			},
		}

		if err := uc.Repository.CartProduct.FindByIndex(tx, cartProduct); err != nil {
			return err
		}

		if err := uc.Repository.CartProduct.Delete(tx, cartProduct); err != nil {
			return err
		}

		return nil
	})

	return err
}

func (uc *usecase) GetCartDetail(ctx context.Context, param GetCartDetailParam) (*GetCartDetailReturn, *pensiel.Error) {
	data := GetCartDetailSync{}
	var err *pensiel.Error

	dbx := uc.Dbi.Cnx(ctx)

	cart := &cart.EntityModel{
		Entity: cart.Entity{
			Invoice: param.Invoice,
		},
	}

	if err := uc.Repository.Cart.FindByInvoice(dbx, cart); err != nil {
		return nil, err
	}

	data.cart.MapFromEntityCart(cart)

	cartProduct := &cartproduct.EntityModel{
		Entity: cartproduct.Entity{
			CartId: cart.SecondaryId,
		},
	}

	cartProducts, err := uc.Repository.CartProduct.GetAllByIndex(dbx, cartProduct)

	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	for _, v := range cartProducts {
		wg.Add(1)
		go func(v cartproduct.EntityModel) {
			defer wg.Done()
			dataProduct := GetProductCartDataProduct{}
			products, errs := uc.Product.GetProduct(dbx, v.ProductId)

			data.mu.Lock()
			defer data.mu.Unlock()

			if errs != nil {
				err = errs
				return
			}

			dataProduct.MapFromEntityProduct(&products.Data)
			dataProduct.Quatity = v.Quantity
			dataProduct.TotalPrice = products.Data.Price * uint64(dataProduct.Quatity)

			data.cart.Price += dataProduct.TotalPrice

			data.cart.Product = append(data.cart.Product, dataProduct)
		}(v)
	}

	wg.Wait()

	if err != nil {
		return nil, err
	}

	if cart.PromoCode != "" {
		promo, err := uc.Promo.GetPromo(dbx, cart.PromoCode)

		if err != nil {
			return nil, err
		}

		data.cart.DiscountPrice = ((promo.Data.Discount * data.cart.Price) / 100)
	}

	if err != nil {
		return nil, err
	}

	data.cart.FinalPrice = data.cart.Price - data.cart.DiscountPrice

	return &data.cart, nil
}
