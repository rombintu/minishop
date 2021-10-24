package app

import "github.com/gin-gonic/gin"

func (s *App) VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			respondWithError(c, 401, "API token required")
			return
		}

		maker, err := NewJWTMaker(s.Config.Server.Secret)
		if err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		if _, err := maker.VerifyToken(token); err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		c.Next()
	}
}

func (s *App) VerifyTokenManager() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			respondWithError(c, 401, "API token required")
			return
		}

		maker, err := NewJWTMaker(s.Config.Server.Secret)
		if err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		payload, err := maker.VerifyToken(token)
		if err != nil {
			s.Logger.Error(err)
			respondWithError(c, 401, err.Error())
			return
		}

		if payload.Username != "manager" {
			respondWithError(c, 401, "you not administrator")
			return
		}
		c.Next()
	}
}
