package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/minishop/internal/store"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

// Test func, return 200 and {"message" : "pong"}
func (s *App) Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		pong := store.Ping{Message: "pong"}
		c.JSON(http.StatusOK, pong)
	}
}

// Create new user if not exists with new wallet
func (s *App) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var u store.User
		Ok := store.Ping{
			Message: "user created",
		}

		if err := c.BindJSON(&u); err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, "user not created")
			return
		}

		account := u.Account
		password := u.Password
		role := u.Role

		if role == "" {
			u.Role = "user"
		}

		if account == "" || password == "" {
			s.Logger.Error("Some user fields is empty")
			respondWithError(c, 401, "Some user fields is empty")
			return
		}

		if err := s.Store.CreateUser(u); err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		c.JSON(http.StatusCreated, Ok)
	}
}

func (s *App) CreateItem() gin.HandlerFunc {
	return func(c *gin.Context) {

		var i store.Item
		Ok := store.Ping{
			Message: "item created",
		}

		if err := c.BindJSON(&i); err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		if i.Name == "" {
			s.Logger.Error("Some item fields is empty")
			respondWithError(c, 401, "Some item fields is empty")
			return
		}

		if err := s.Store.CreateItem(i); err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		c.JSON(http.StatusCreated, Ok)
	}
}

func (s *App) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		idStr := c.Query("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			s.Logger.Error(err)
			return
		}
		user, err := s.Store.GetUser(id)
		if err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		c.JSON(http.StatusCreated, user)
	}
}

func (s *App) GetItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Query("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			s.Logger.Error(err)
			return
		}
		item, err := s.Store.GetItem(id)
		if err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		c.JSON(http.StatusCreated, item)
	}
}

func (s *App) GetItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := s.Store.GetItems()
		if err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		c.JSON(http.StatusCreated, items)
	}
}

func (s *App) GetBasket() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Query("id_user")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			s.Logger.Error(err)
			return
		}
		basket, err := s.Store.GetBasket(id)
		if err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		c.JSON(http.StatusCreated, basket)
	}
}

func (s *App) GetBaskets() gin.HandlerFunc {
	return func(c *gin.Context) {
		baskets, err := s.Store.GetBaskets()
		if err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		c.JSON(http.StatusCreated, baskets)
	}
}

func (s *App) UpdateBasket() gin.HandlerFunc {
	return func(c *gin.Context) {
		var basket store.BasketUpdate

		Ok := store.Ping{
			Message: "basket updated",
		}

		if err := c.BindJSON(&basket); err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		if err := s.Store.UpdateBasket(basket); err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		c.JSON(http.StatusCreated, Ok)
	}
}
