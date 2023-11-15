package simple_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var GinCmd = &cobra.Command{
	Use:   "gin",
	Short: "gin server",
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.Default()

		r.GET("/", func(c *gin.Context) {
			c.String(200, "hello from tom")
		})

		r.Run(":9999")
	},
}
