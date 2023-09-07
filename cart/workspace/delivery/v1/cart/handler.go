package cart

import (
	"net/http"

	"pensiel.com/domain/src/usecase/cart"
	"pensiel.com/material/src/contract"
	"pensiel.com/material/src/pensiel"
)

func (h *handler) GETPRODUCTCART(c pensiel.Context) *pensiel.Error {
	ctx := c.Context()
	invoice := c.Param("invoice")

	cart, err := h.uc.Cart.GetCartDetail(ctx, cart.GetCartDetailParam{
		Invoice: invoice,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, GetProductCartResponse{
		ResponseMeta: contract.ResponseMeta{
			Message: "product add to the cart",
		},
		Data: *cart,
	})
}

func (h *handler) ADDPRODUCTINTOCART(c pensiel.Context) *pensiel.Error {
	ctx := c.Context()
	invoice := c.Param("invoice")

	req := AddProductIntoCartRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	if err := h.uc.Cart.AddProductIntoCart(ctx, cart.AddProductIntoCartParam{
		Invoice:   invoice,
		ProductId: req.ProductId,
		Quantity:  req.Quantity,
	}); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, AddProductIntoCartResponse{
		ResponseMeta: contract.ResponseMeta{
			Message: "product add to the cart",
		},
	})
}

func (h *handler) APPLYPROMO(c pensiel.Context) *pensiel.Error {
	ctx := c.Context()
	invoice := c.Param("invoice")

	req := ApplyPromoRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	if err := h.uc.Cart.CartPromo(ctx, cart.CartPromoParam{
		Invoice:   invoice,
		PromoCode: req.PromoCode,
	}); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, AddProductIntoCartResponse{
		ResponseMeta: contract.ResponseMeta{
			Message: "Promotion successfully installed",
		},
	})
}

func (h *handler) UPDATEQUATITYPRODUCTCART(c pensiel.Context) *pensiel.Error {
	ctx := c.Context()
	invoice := c.Param("invoice")
	productId := c.Param("productId")

	req := UpdateQuantityProductCartRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	if err := h.uc.Cart.UpdateQuantityProductCart(ctx, cart.UpdateQuantityProductCartParam{
		Invoice:   invoice,
		ProductId: productId,
		Quantity:  req.Quantity,
	}); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, AddProductIntoCartResponse{
		ResponseMeta: contract.ResponseMeta{
			Message: "Product cart has been updated",
		},
	})

}

func (h *handler) REMOVEPRODUCTCART(c pensiel.Context) *pensiel.Error {
	ctx := c.Context()
	invoice := c.Param("invoice")
	productId := c.Param("productId")

	if err := h.uc.Cart.RemoveProductFromCart(ctx, cart.RemoveProductFormCartParam{
		Invoice:   invoice,
		ProductId: productId,
	}); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, AddProductIntoCartResponse{
		ResponseMeta: contract.ResponseMeta{
			Message: "Product cart has been remove",
		},
	})
}
