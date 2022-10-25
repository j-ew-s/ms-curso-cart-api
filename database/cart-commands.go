package database

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/j-ew-s/ms-curso-cart-api/cart-services/models"
)

const UserCartKey = "UC"

var ctx = context.Background()

func GetCart(c *redis.Client, userid string) *models.Cart {
	userKey := UserCartKey + userid
	cart := &models.Cart{}
	val, err := c.Get(userKey).Result()
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal([]byte(val), &cart)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(fmt.Printf(":: RESULT :: %v  :::", cart))

	return cart
}

func RemoveFromCart() {
	fmt.Println("RemoveFromCart")

}

func ClearCart() {
	fmt.Println("ClearCart")

}

func AddToCart(c *redis.Client, cart models.Cart, userid string) error {

	userKey := UserCartKey + userid

	cartJson, err := json.Marshal(cart)
	if err != nil {
		fmt.Println(err)
	}

	err = c.Set(userKey, cartJson, 24*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}
