package elasticsearch

import (
	"net/http"
	"github.com/gin-gonic/gin"
	es7 "github.com/elastic/go-elasticsearch/v7"
)

var cfg = es7.Config{
	Addresses: []string{
		"http://es:9200",
	},
}

func Client() gin.HandlerFunc {
	return func (c *gin.Context) {
		es, err := es7.NewClient(cfg)
		if err != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		c.Set("elastic", es)
	}	
}
