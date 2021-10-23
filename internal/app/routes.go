package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rombintu/minishop/internal/store"
)

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
		notOk := store.Ping{
			Message: "user not created",
		}
		if err := c.BindJSON(&u); err != nil {
			s.Logger.Error(err)
			c.JSON(http.StatusConflict, notOk)
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
			c.JSON(http.StatusConflict, notOk)
			return
		}

		if err := s.Store.CreateUser(u); err != nil {
			c.JSON(http.StatusConflict, notOk)
			s.Logger.Error(err)
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
		notOk := store.Ping{
			Message: "item not created",
		}
		if err := c.BindJSON(&i); err != nil {
			s.Logger.Error(err)
			c.JSON(http.StatusConflict, notOk)
			return
		}

		if i.Name == "" {
			s.Logger.Error("Some item fields is empty")
			c.JSON(http.StatusConflict, notOk)
			return
		}

		if err := s.Store.CreateItem(i); err != nil {
			c.JSON(http.StatusConflict, notOk)
			s.Logger.Error(err)
			return
		}

		c.JSON(http.StatusCreated, Ok)
	}
}

func (s *App) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Query("id")
		notOk := store.Ping{
			Message: "user not found",
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			s.Logger.Error(err)
			return
		}
		user, err := s.Store.GetUser(id)
		if err != nil {
			c.JSON(http.StatusConflict, notOk)
			s.Logger.Error(err)
			return
		}

		c.JSON(http.StatusCreated, user)
	}
}

func (s *App) GetItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Query("id")
		notOk := store.Ping{
			Message: "item not found",
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			s.Logger.Error(err)
			return
		}
		item, err := s.Store.GetItem(id)
		if err != nil {
			c.JSON(http.StatusConflict, notOk)
			s.Logger.Error(err)
			return
		}

		c.JSON(http.StatusCreated, item)
	}
}

func (s *App) GetItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		notOk := store.Ping{
			Message: "items not found",
		}
		items, err := s.Store.GetItems()
		if err != nil {
			c.JSON(http.StatusConflict, notOk)
			s.Logger.Error(err)
			return
		}

		c.JSON(http.StatusCreated, items)
	}
}

func (s *App) GetBasket() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Query("id_user")
		notOk := store.Ping{
			Message: "basket not found",
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			s.Logger.Error(err)
			return
		}
		basket, err := s.Store.GetBasket(id)
		if err != nil {
			c.JSON(http.StatusConflict, notOk)
			s.Logger.Error(err)
			return
		}

		c.JSON(http.StatusCreated, basket)
	}
}

func (s *App) GetBaskets() gin.HandlerFunc {
	return func(c *gin.Context) {
		notOk := store.Ping{
			Message: "basket not found",
		}
		baskets, err := s.Store.GetBaskets()
		if err != nil {
			c.JSON(http.StatusConflict, notOk)
			s.Logger.Error(err)
			return
		}

		c.JSON(http.StatusCreated, baskets)
	}
}
