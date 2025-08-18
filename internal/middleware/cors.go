package middleware

import "github.com/gin-contrib/cors"

func GetCorsConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	corsConfig.AllowHeaders = []string{"*"}

	return corsConfig
}
