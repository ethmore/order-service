package controllers

import (
	"fmt"
	"net/http"
	"order/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type OrderInfo struct {
	Token              string
	AddressID          string
	CardLastFourDigits string
}

func Test() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var orderInfo OrderInfo
		if err := ctx.ShouldBindBodyWith(&orderInfo, binding.JSON); err != nil {
			fmt.Println("body: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		address, err := services.GetDetailedAddressByID(orderInfo.Token, orderInfo.AddressID)
		if err != nil {
			fmt.Println("get address: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		cart, err := services.GetCartInfo(orderInfo.Token)
		if err != nil {
			fmt.Println("get cart: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		cartProducts, err := services.GetCartProducts(orderInfo.Token)
		if err != nil {
			fmt.Println("get cart products: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		var productIDs []string
		for _, product := range cartProducts {
			productIDs = append(productIDs, product.Id)
		}

		productsAndSellers, err := services.GetProductsSellers(orderInfo.Token, productIDs)
		if err != nil {
			fmt.Println("GetProductsSellers: ", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		var order services.Order
		var product services.Product_

		order.Token = orderInfo.Token
		for i, j := range cart.Items {
			product.Title = j.Title
			product.Qty = cartProducts[i].Qty
			product.Price = cart.Items[i].TotalPrice

			if productsAndSellers[i].ProductID == j.Id {
				product.SellerName = productsAndSellers[i].Seller
			} else {
				for _, j := range productsAndSellers {
					if j.ProductID == cart.Items[i].Id {
						product.SellerName = j.Seller
					}
				}
			}
			order.Products = append(order.Products, product)
		}

		order.TotalPrice = cart.TotalCartPrice
		order.ShipmentAddressID = address.Id

		order.CardLastFourDigits = orderInfo.CardLastFourDigits
		order.PaymentStatus = "Success"
		order.OrderStatus = "Preparing"

		currentTime := time.Now().Format("02-Jan-2006 15:04:05")

		order.OrderTime = currentTime

		_, ordErr := services.InsertOrder(order)
		if ordErr != nil {
			fmt.Println("inser Order: ", ordErr)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}
