package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener los roles del contexto
		roles, exists := c.Get("roles")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Rol no encontrado en el contexto"})
			c.Abort()
			return
		}

		// Validar formato de roles
		rolesList, ok := roles.([]string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Formato de roles inválido"})
			c.Abort()
			return
		}

		// Definir roles válidos
		validRoles := map[string]struct{}{
			"ejecutivo_ventas":    {},
			"administrador":       {},
			"super_administrador": {},
		}

		// Verificar si el usuario tiene un rol válido
		isAdmin := false
		for _, role := range rolesList {
			if _, ok := validRoles[role]; ok {
				isAdmin = true
				break
			}
		}

		// Bloquear acceso si no tiene permisos
		if !isAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acceso denegado: Se requiere rol de administrador, ejecutivo_ventas o super_administrador"})
			c.Abort()
			return
		}

		// Continuar si tiene permisos
		c.Next()
	}
}
